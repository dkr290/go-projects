/*
Copyright 2025.

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

package controller

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsbankingcirclenetv1alpha1 "api-operator/api/v1alpha1"
)

const (
	appLabel        = "my-api"
	versionLabel    = "version"
	ingressName     = "my-api-ingress"
	imagePullsecret = "regcred"
)

// AppVersionReconciler reconciles a AppVersion object
type AppVersionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=apps.bankingcircle.net.bankingcircle.net,resources=appversions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps.bankingcircle.net.bankingcircle.net,resources=appversions/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps.bankingcircle.net.bankingcircle.net,resources=appversions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the AppVersion object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *AppVersionReconciler) Reconcile(
	ctx context.Context,
	req ctrl.Request,
) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	// trying to fetch appversion CR instance
	var appVersion appsbankingcirclenetv1alpha1.AppVersion
	if err := r.Get(ctx, req.NamespacedName, &appVersion); err != nil {
		if errors.IsNotFound(err) {
			logger.Info("AppVersion resource not found. Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		logger.Error(err, "Failed to get AppVersion")
		return ctrl.Result{}, err
	}
	// List existing Deployments for the API using the label "app=my-api from the contants it is subject to change"
	var deploymentList appsv1.DeploymentList
	if err := r.List(ctx, &deploymentList, client.MatchingLabels{"app": appLabel}); err != nil {
		logger.Error(err, "Failed to list Deployments")
		return ctrl.Result{}, err
	}
	// get the active current version from the Deployment byt version= key
	existingVersions := map[string]*appsv1.Deployment{}
	for k, v := range deploymentList.Items {
		if ver, exists := v.Labels[versionLabel]; exists {
			existingVersions[ver] = &deploymentList.Items[k]
		}
	}
	desiredVersion := appVersion.Spec.Version
	// If the desired version is not deployed, create new Deployment and respective Service. We just do ingress mapping after
	if _, exists := existingVersions[desiredVersion]; !exists {
		logger.Info("Creating new deployment and service for version", "version", desiredVersion)
		// adding new deployment
		deployment := r.constructDeployment(appVersion)
		// setting appVersion as owner for garbage collection best practices
		if err := controllerutil.SetControllerReference(&appVersion, deployment, r.Scheme); err != nil {
			return ctrl.Result{}, err
		}
		if err := r.Create(ctx, deployment); err != nil {
			logger.Error(err, "Failed to create Deployment", "Deployment", deployment)
			return ctrl.Result{}, err
		}
		// creating the connecting service
		// creating the service --------------------------------

		service := r.constructService(appVersion)
		if err := controllerutil.SetControllerReference(&appVersion, service, r.Scheme); err != nil {
			return ctrl.Result{}, err
		}
		// adding desired version ,adding to our collection of existing versions
		existingVersions[desiredVersion] = deployment

	}
	// Sort the versions (assumes versions in this application are only in the form of like "v21", "v22", etc.)
	sortedVersions := sortVersions(existingVersions)
	if len(sortedVersions) > 2 {
		toDelete := sortedVersions[:len(sortedVersions)-2]
		for _, ver := range toDelete {
			logger.Info("Deleting old deployment and service for version", "version", ver)
			// Delete the old deployment
			if dep, ok := existingVersions[ver]; ok {
				if err := r.Delete(ctx, dep); err != nil {
					logger.Error(err, "Failed to delete Deployment", "version", ver)
				}
			}

			// Delete the old service which should be associated with the old deployment

			svcName := serviceName(ver)
			var svc corev1.Service
			if err := r.Get(ctx, client.ObjectKey{Namespace: appVersion.Namespace, Name: svcName}, &svc); err == nil {
				if err := r.Delete(ctx, &svc); err != nil {
					logger.Error(
						err,
						"Failed to delete Service",
						"Service",
						svcName,
						"version",
						ver,
					)
					return ctrl.Result{}, err
				}
			}
		}
		// Keep only the newest two versions.
		sortedVersions = sortedVersions[len(sortedVersions)-2:]
	}
	if err := r.reconcileIngress(ctx, sortedVersions, req.Namespace); err != nil {
		logger.Error(err, "Failed to reconcile Ingress")
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

func (r *AppVersionReconciler) constructDeployment(
	appVersion appsbankingcirclenetv1alpha1.AppVersion,
) *appsv1.Deployment {
	labels := map[string]string{
		"app":     appLabel,
		"version": appVersion.Spec.Version,
	}

	replicas := int32(1)
	if appVersion.Spec.Replicas != nil {
		replicas = *appVersion.Spec.Replicas
	}

	objectMetaData := metav1.ObjectMeta{
		Name:      "my-api-" + strings.ToLower(appVersion.Spec.Version),
		Namespace: appVersion.Namespace,
		Labels:    labels,
	}

	specData := appsv1.DeploymentSpec{
		Replicas: &replicas,
		Selector: &metav1.LabelSelector{
			MatchLabels: labels,
		},
		Template: corev1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: labels,
			},
			Spec: corev1.PodSpec{
				Containers: []corev1.Container{
					{
						Name:  "api",
						Image: appVersion.Spec.Image,
						Ports: []corev1.ContainerPort{
							{ContainerPort: appVersion.Spec.Port},
						},
					},
				},
				ImagePullSecrets: []corev1.LocalObjectReference{
					{
						Name: imagePullsecret,
					},
				},
			},
		},
	}

	deploy := &appsv1.Deployment{
		ObjectMeta: objectMetaData,
		Spec:       specData,
	}

	return deploy
}

func (r *AppVersionReconciler) constructService(
	appVersion appsbankingcirclenetv1alpha1.AppVersion,
) *corev1.Service {
	metadata := metav1.ObjectMeta{
		Name:      serviceName(appVersion.Spec.Version),
		Namespace: appVersion.Namespace,
		Labels: map[string]string{
			"app":     appLabel,
			"version": appVersion.Spec.Version,
		},
	}
	spec := corev1.ServiceSpec{
		Selector: map[string]string{
			"app":     appLabel,
			"version": appVersion.Spec.Version,
		},
		Ports: []corev1.ServicePort{
			{
				Port:       appVersion.Spec.Port,
				TargetPort: intstr.FromInt(int(appVersion.Spec.Port)),
				Protocol:   corev1.ProtocolTCP,
			},
		},
	}

	svc := &corev1.Service{
		ObjectMeta: metadata,
		Spec:       spec,
	}

	return svc
}

func (r *AppVersionReconciler) reconcileIngress(
	ctx context.Context,
	versions []string,
	namespace string,
) error {
	ingress := &networkingv1.Ingress{}

	// Check if the Ingress already exists
	err := r.Get(ctx, client.ObjectKey{Namespace: namespace, Name: ingressName}, ingress)
	// Check if the Ingress already existso
	if err != nil && errors.IsNotFound(err) {
		// ingress does not exists and creating new one
		ingress = r.constructIngress(versions, namespace)
		return r.Create(ctx, ingress)
	} else if err != nil {
		return err
	}
	newIngress := r.constructIngress(versions, namespace)
	ingress.Spec = newIngress.Spec
	return r.Update(ctx, ingress)
}

func (r *AppVersionReconciler) constructIngress(
	versions []string,
	namespace string,
) *networkingv1.Ingress {
	var paths []networkingv1.HTTPIngressPath
	for _, ver := range versions {
		path := networkingv1.HTTPIngressPath{
			Path: "/" + ver,
			PathType: func() *networkingv1.PathType {
				pt := networkingv1.PathTypePrefix
				return &pt
			}(),
			Backend: networkingv1.IngressBackend{
				Service: &networkingv1.IngressServiceBackend{
					Name: serviceName(ver),
					Port: networkingv1.ServiceBackendPort{
						Number: 80,
					},
				},
			},
		}
		paths = append(paths, path)
	}
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ingressName,
			Namespace: namespace,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: "my-api.bankingcircle.net",
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: paths,
						},
					},
				},
			},
			TLS: []networkingv1.IngressTLS{
				{
					Hosts:      []string{"my-api.bankingcircle.net"},
					SecretName: "my-api-tls-secret",
				},
			},
		},
	}
	return ingress
}

// returns a standardized Service name based on the version.
func serviceName(version string) string {
	return fmt.Sprintf("my-api-%s-service", strings.ToLower(version))
}

// sortVersions sorts the version strings (assuming formats like "v21", "v22") Only thoose are accepted in this application
func sortVersions(versions map[string]*appsv1.Deployment) []string {
	var vers []string
	for ver := range versions {
		vers = append(vers, ver)
	}
	sort.Slice(vers, func(i, j int) bool {
		return parseVersion(vers[i]) < parseVersion(vers[j])
	})
	return vers
}

// parseVersion converts a version string (e.g., "v22") by stripping the "v" and returning its integer part.
func parseVersion(ver string) int {
	trimmed := strings.TrimPrefix(ver, "v")
	num, err := strconv.Atoi(trimmed)
	if err != nil {
		return 0
	}
	return num
}

// SetupWithManager sets up the controller with the Manager.
// SetupWithManager sets up the controller with the Manager.
func (r *AppVersionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsbankingcirclenetv1alpha1.AppVersion{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&networkingv1.Ingress{}).
		Complete(r)
}
