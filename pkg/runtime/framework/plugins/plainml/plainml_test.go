/*
Copyright 2025 The Kubeflow Authors.

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

package plainml

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1ac "k8s.io/client-go/applyconfigurations/core/v1"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/ptr"

	trainer "github.com/kubeflow/trainer/pkg/apis/trainer/v1alpha1"
	"github.com/kubeflow/trainer/pkg/constants"
	"github.com/kubeflow/trainer/pkg/runtime"
	"github.com/kubeflow/trainer/pkg/runtime/framework"
	utiltesting "github.com/kubeflow/trainer/pkg/util/testing"
)

func TestPlainML(t *testing.T) {
	cases := map[string]struct {
		trainJob  *trainer.TrainJob
		info      *runtime.Info
		wantError error
		wantInfo  *runtime.Info
	}{
		"no action when info is null": {},
		"no action when mlPolicy is null": {
			info: &runtime.Info{
				Labels: map[string]string{"key": "value"},
			},
			wantInfo: &runtime.Info{
				Labels: map[string]string{"key": "value"},
			},
		},
		"no action when mlPolicy torch is not null": {
			info: &runtime.Info{
				Labels: map[string]string{"key": "value"},
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						TorchPolicy("auto", nil).
						Obj(),
				},
			},
			wantInfo: &runtime.Info{
				Labels: map[string]string{"key": "value"},
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						TorchPolicy("auto", nil).
						Obj(),
				},
			},
		},
		"no action when mlPolicy mpi is not null": {
			info: &runtime.Info{
				Labels: map[string]string{"key": "value"},
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						MPIPolicy(ptr.To[int32](1), ptr.To(trainer.MPIImplementationOpenMPI), nil, ptr.To(false)).
						Obj(),
				},
			},
			wantInfo: &runtime.Info{
				Labels: map[string]string{"key": "value"},
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						MPIPolicy(ptr.To[int32](1), ptr.To(trainer.MPIImplementationOpenMPI), nil, ptr.To(false)).
						Obj(),
				},
			},
		},
		"trainJob trainer numNodes are respected rather than runtimeInfo": {
			trainJob: utiltesting.MakeTrainJobWrapper(metav1.NamespaceDefault, "test").
				Trainer(
					utiltesting.MakeTrainJobTrainerWrapper().NumNodes(200).Obj(),
				).
				Obj(),
			info: runtime.NewInfo(
				runtime.WithMLPolicy(
					utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						Obj(),
				),
				runtime.WithPodSpecReplicas(constants.JobTrainerNode, 100, nil, corev1ac.PodSpec().
					WithContainers(corev1ac.Container().WithName(constants.ContainerTrainer)),
				),
			),
			wantInfo: &runtime.Info{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().WithNumNodes(200).Obj(),
				},
				TemplateSpec: runtime.TemplateSpec{
					PodSets: []runtime.PodSet{{
						Name: constants.JobTrainerNode,
						Containers: []runtime.Container{{
							Name: constants.ContainerTrainer,
						}},
					}},
				},
				Scheduler: &runtime.Scheduler{TotalRequests: map[string]runtime.TotalResourceRequest{
					constants.JobTrainerNode: {Replicas: 200},
				}},
			},
		},
		"trainJob trainer env are respected rather than runtimeInfo": {
			trainJob: utiltesting.MakeTrainJobWrapper(metav1.NamespaceDefault, "test").
				Trainer(
					utiltesting.MakeTrainJobTrainerWrapper().
						ContainerEnv(corev1.EnvVar{
							Name:  "CONFLICT",
							Value: "FROM_TRAINER",
						},
						).
						NumNodes(1).
						Obj(),
				).
				Obj(),
			info: runtime.NewInfo(
				runtime.WithMLPolicy(
					utiltesting.MakeMLPolicyWrapper().Obj(),
				),
				runtime.WithPodSpecReplicas(constants.JobTrainerNode, 100, nil, corev1ac.PodSpec().
					WithContainers(corev1ac.Container().
						WithName(constants.ContainerTrainer).
						WithEnv(corev1ac.EnvVar().
							WithName("CONFLICT").
							WithValue("FROM_TRAINER"),
						),
					),
				),
			),
			wantInfo: &runtime.Info{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(1).
						Obj(),
				},
				TemplateSpec: runtime.TemplateSpec{
					PodSets: []runtime.PodSet{{
						Name: constants.JobTrainerNode,
						Containers: []runtime.Container{{
							Name: constants.ContainerTrainer,
							Env: []corev1ac.EnvVarApplyConfiguration{
								*corev1ac.EnvVar().WithName("CONFLICT").WithValue("FROM_TRAINER"),
							},
						}},
					}},
				},
				Scheduler: &runtime.Scheduler{TotalRequests: map[string]runtime.TotalResourceRequest{
					constants.JobTrainerNode: {Replicas: 1},
				}},
			},
		},
		"override trainer numNodes to runtimeInfo scheduler totalRequests": {
			trainJob: utiltesting.MakeTrainJobWrapper(metav1.NamespaceDefault, "test").
				Trainer(
					utiltesting.MakeTrainJobTrainerWrapper().NumNodes(200).Obj(),
				).
				Obj(),
			info: runtime.NewInfo(
				runtime.WithMLPolicy(
					utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(100).
						Obj(),
				),
				runtime.WithPodSpecReplicas(constants.JobTrainerNode, 100, corev1.ResourceList{
					corev1.ResourceCPU: resource.MustParse("200m"),
				}, corev1ac.PodSpec().
					WithContainers(corev1ac.Container().WithName(constants.ContainerTrainer)),
				),
			),
			wantInfo: &runtime.Info{
				Labels:      make(map[string]string),
				Annotations: make(map[string]string),
				RuntimePolicy: runtime.RuntimePolicy{
					MLPolicy: utiltesting.MakeMLPolicyWrapper().
						WithNumNodes(200).
						Obj(),
				},
				TemplateSpec: runtime.TemplateSpec{
					PodSets: []runtime.PodSet{{
						Name: constants.JobTrainerNode,
						Containers: []runtime.Container{{
							Name: constants.ContainerTrainer,
						}},
					}},
				},
				Scheduler: &runtime.Scheduler{TotalRequests: map[string]runtime.TotalResourceRequest{
					constants.JobTrainerNode: {
						Replicas:    200,
						PodRequests: corev1.ResourceList{corev1.ResourceCPU: resource.MustParse("200m")},
					}},
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, ctx := ktesting.NewTestContext(t)
			var cancel func()
			ctx, cancel = context.WithCancel(ctx)
			t.Cleanup(cancel)
			cliBuilder := utiltesting.NewClientBuilder()
			p, err := New(ctx, cliBuilder.Build(), nil)
			if err != nil {
				t.Fatalf("Failed to initialize PlainML plugin: %v", err)
			}
			err = p.(framework.EnforceMLPolicyPlugin).EnforceMLPolicy(tc.info, tc.trainJob)
			if diff := cmp.Diff(tc.wantError, err, cmpopts.EquateErrors()); len(diff) != 0 {
				t.Errorf("Unexpected error from EnforceMLPolicy (-want,+got):\n%s", diff)
			}
			if diff := cmp.Diff(tc.wantInfo, tc.info,
				cmpopts.SortSlices(func(a, b string) bool { return a < b }),
				cmpopts.SortMaps(func(a, b string) bool { return a < b }),
			); len(diff) != 0 {
				t.Errorf("Unexpected RuntimeInfo (-want,+got):\n%s", diff)
			}
		})
	}
}
