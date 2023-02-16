/*
Copyright 2023.

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

	recommenderv1beta1 "github.com/Ishujeet/maverick/api/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// RecommenderObjectReconciler reconciles a RecommenderObject object
type RecommenderObjectReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=recommender.maverick.com,resources=recommenderobjects,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=recommender.maverick.com,resources=recommenderobjects/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=recommender.maverick.com,resources=recommenderobjects/finalizers,verbs=update
//+kubebuilder:rbac:groups=appsV1,resources=deployment,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=appsV1,resources=deployment/status,verbs=get

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the RecommenderObject object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *RecommenderObjectReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var recommenderObject recommenderv1beta1.RecommenderObject
	if err := r.Get(ctx, req.NamespacedName, &recommenderObject); err != nil {
		log.Error(err, "unable to fetch RecommenderObject")
		// we'll ignore not-found errors, since they can't be fixed by an immediate
		// requeue (we'll need to wait for a new notification), and we can get them
		// on deleted requests.
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	deployment := &appsv1.Deployment{}
	if err := r.Get(ctx, types.NamespacedName{Name: recommenderObject.Spec.TargetRef.Name, Namespace: recommenderObject.Namespace}, deployment); err != nil && errors.IsNotFound(err) {
		log.Error(err, "Deployment not found")
		//If deployment not exists Recommender Object won't work
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *RecommenderObjectReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&recommenderv1beta1.RecommenderObject{}).
		Complete(r)
}
