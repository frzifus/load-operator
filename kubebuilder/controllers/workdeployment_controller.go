/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	workv1 "load-operator/api/v1"
)

// WorkDeploymentReconciler reconciles a WorkDeployment object
type WorkDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=work.klimlive.de,resources=workdeployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=work.klimlive.de,resources=workdeployments/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=work.klimlive.de,resources=workdeployments/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(frzifus): Modify the Reconcile function to compare the state specified by
// the WorkDeployment object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *WorkDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info(fmt.Sprintf("work-reconcile: request: %+v ", req))

	// TODO(frzifus): clean-up
	custom := &workv1.WorkDeployment{}
	key := types.NamespacedName{Name: req.Name, Namespace: req.Namespace}
	if err := r.Get(ctx, key, custom); err != nil && !apierrors.IsNotFound(err) {
		logger.Info(fmt.Sprintf("work-reconcile: crd not found! request: %+v ", req))
		return ctrl.Result{}, err
	}

	deploy := &appsv1.Deployment{}
	if err := r.Get(ctx, key, deploy); apierrors.IsNotFound(err) {
		logger.Info("Creating worker...")
		d := r.mkDeploy(req.Name, req.Namespace, custom.Spec.TargetLoad, custom.Spec.TargetMemory)
		if err := r.Create(ctx, d); err != nil {
			return ctrl.Result{}, err
		}
		custom.Status.Name = custom.Spec.Name
		custom.Status.TargetLoad = custom.Spec.TargetLoad
		custom.Status.TargetMemory = custom.Spec.TargetMemory
		if err := r.Update(ctx, custom); err != nil {
			return ctrl.Result{}, err
		}
	} else if err != nil {
		return ctrl.Result{}, err
	}

	if custom.Spec.TargetLoad == 0 && custom.Spec.TargetMemory == 0 {
		logger.Info("Deleting worker...")
		return ctrl.Result{}, r.Delete(ctx, deploy)
	}

	if custom.Status.Name != custom.Spec.Name ||
		custom.Status.TargetLoad != custom.Spec.TargetLoad ||
		custom.Status.TargetMemory != custom.Spec.TargetMemory {
		updateDeploy(deploy, custom.Spec.TargetLoad, custom.Spec.TargetMemory)
		logger.Info("Updating worker...")
		if err := r.Update(ctx, deploy); err != nil {
			return ctrl.Result{}, err
		}
	}

	custom.Status.Name = custom.Spec.Name
	custom.Status.TargetLoad = custom.Spec.TargetLoad
	custom.Status.TargetMemory = custom.Spec.TargetMemory
	if err := r.Update(ctx, custom); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func updateDeploy(deploy *appsv1.Deployment, load uint8, mem uint16) {
	// TODO(frzifus): update args
}

func (r *WorkDeploymentReconciler) mkDeploy(
	name string,
	namespace string,
	load uint8,
	mem uint16,
) *appsv1.Deployment {
	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: func(v int32) *int32 { return &v }(1), // TODO
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app.klimlive.de": name,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app.klimlive.de": name,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name: "worker",
							// TODO(frzifus): replace with worker img
							Image: "nginx:1.19",
							// TODO(frzifus): use load & mem as args
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
						},
					},
				},
			},
		},
	}
}

// SetupWithManager sets up the controller with the Manager.
func (r *WorkDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&workv1.WorkDeployment{}).
		Complete(r)
}
