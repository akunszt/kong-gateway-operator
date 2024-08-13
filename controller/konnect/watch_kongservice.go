package konnect

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	ctrllog "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	configurationv1alpha1 "github.com/kong/kubernetes-configuration/api/configuration/v1alpha1"
	konnectv1alpha1 "github.com/kong/kubernetes-configuration/api/konnect/v1alpha1"
)

// TODO(pmalek): this can be extracted and used in reconciler.go
// as every Konnect entity will have a reference to the KonnectAPIAuthConfiguration.
// This would require:
// - mapping function from non List types to List types
// - a function on each Konnect entity type to get the API Auth configuration
//   reference from the object
// - lists have their items stored in Items field, not returned via a method

// KongServiceReconciliationWatchOptions returns the watch options for
// the KongService.
func KongServiceReconciliationWatchOptions(
	cl client.Client,
) []func(*ctrl.Builder) *ctrl.Builder {
	return []func(*ctrl.Builder) *ctrl.Builder{
		func(b *ctrl.Builder) *ctrl.Builder {
			return b.Watches(
				&konnectv1alpha1.KonnectAPIAuthConfiguration{},
				handler.EnqueueRequestsFromMapFunc(
					enqueueKongServiceForKonnectAPIAuthConfiguration(cl),
				),
			)
		},
		func(b *ctrl.Builder) *ctrl.Builder {
			return b.Watches(
				&konnectv1alpha1.KonnectControlPlane{},
				handler.EnqueueRequestsFromMapFunc(
					enqueueKongServiceForKonnectControlPlane(cl),
				),
			)
		},
	}
}

func enqueueKongServiceForKonnectAPIAuthConfiguration(
	cl client.Client,
) func(ctx context.Context, obj client.Object) []reconcile.Request {
	return func(ctx context.Context, obj client.Object) []reconcile.Request {
		auth, ok := obj.(*konnectv1alpha1.KonnectAPIAuthConfiguration)
		if !ok {
			return nil
		}
		var l configurationv1alpha1.KongServiceList
		if err := cl.List(ctx, &l, &client.ListOptions{
			// TODO: change this when cross namespace refs are allowed.
			Namespace: auth.GetNamespace(),
		}); err != nil {
			return nil
		}

		var ret []reconcile.Request
		for _, svc := range l.Items {
			switch svc.Spec.ControlPlaneRef.Type {
			case configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef:
				nn := types.NamespacedName{
					Name:      svc.Spec.ControlPlaneRef.KonnectNamespacedRef.Name,
					Namespace: svc.Spec.ControlPlaneRef.KonnectNamespacedRef.Namespace,
				}
				// TODO: change this when cross namespace refs are allowed.
				if nn.Namespace != auth.Namespace {
					continue
				}
				var cp konnectv1alpha1.KonnectControlPlane
				if err := cl.Get(ctx, nn, &cp); err != nil {
					ctrllog.FromContext(ctx).Error(
						err,
						"failed to get KonnectControlPlane",
						"KonnectControlPlane", nn,
					)
					continue
				}

				// TODO: change this when cross namespace refs are allowed.
				if cp.GetKonnectAPIAuthConfigurationRef().Name != auth.Name {
					continue
				}

				ret = append(ret, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: svc.Namespace,
						Name:      svc.Name,
					},
				})
			default:
				ctrllog.FromContext(ctx).Error(
					fmt.Errorf("unsupported ControlPlaneRef type %q", svc.Spec.ControlPlaneRef.Type),
					"unsupported ControlPlaneRef for KongService",
					"KongService", svc,
				)
				continue
			}
		}
		return ret
	}
}

func enqueueKongServiceForKonnectControlPlane(
	cl client.Client,
) func(ctx context.Context, obj client.Object) []reconcile.Request {
	return func(ctx context.Context, obj client.Object) []reconcile.Request {
		cp, ok := obj.(*konnectv1alpha1.KonnectControlPlane)
		if !ok {
			return nil
		}
		var l configurationv1alpha1.KongServiceList
		if err := cl.List(ctx, &l, &client.ListOptions{
			// TODO: change this when cross namespace refs are allowed.
			Namespace: cp.GetNamespace(),
		}); err != nil {
			return nil
		}

		var ret []reconcile.Request
		for _, svc := range l.Items {
			switch svc.Spec.ControlPlaneRef.Type {
			case configurationv1alpha1.ControlPlaneRefKonnectNamespacedRef:
				// TODO: change this when cross namespace refs are allowed.
				if svc.Spec.ControlPlaneRef.KonnectNamespacedRef.Name != cp.Name {
					continue
				}

				ret = append(ret, reconcile.Request{
					NamespacedName: types.NamespacedName{
						Namespace: svc.Namespace,
						Name:      svc.Name,
					},
				})

			default:
				ctrllog.FromContext(ctx).Error(
					fmt.Errorf("unsupported ControlPlaneRef type %q", svc.Spec.ControlPlaneRef.Type),
					"unsupported ControlPlaneRef for KongService",
					"KongService", svc,
				)
				continue
			}
		}
		return ret
	}
}