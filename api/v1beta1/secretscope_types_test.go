/*
Copyright 2019 microsoft.

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

package v1beta1

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// These tests are written in BDD-style using Ginkgo framework. Refer to
// http://onsi.github.io/ginkgo to learn more.

var _ = Describe("SecretScope", func() {
	var (
		key              types.NamespacedName
		created, fetched *SecretScope
	)

	BeforeEach(func() {
		// Add any setup steps that needs to be executed before each test
	})

	AfterEach(func() {
		// Add any teardown steps that needs to be executed after each test
	})

	// Add Tests for OpenAPI validation (or additonal CRD features) specified in
	// your API definition.
	// Avoid adding tests for vanilla CRUD operations because they would
	// test Kubernetes API server, which isn't the goal here.
	Context("Create API", func() {

		It("should create an object successfully", func() {

			key = types.NamespacedName{
				Name:      "foo",
				Namespace: "default",
			}
			created = &SecretScope{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "foo",
					Namespace: "default",
				}}

			By("creating an API obj")
			Expect(k8sClient.Create(context.Background(), created)).To(Succeed())

			fetched = &SecretScope{}
			Expect(k8sClient.Get(context.Background(), key, fetched)).To(Succeed())
			Expect(fetched).To(Equal(created))

			By("deleting the created object")
			Expect(k8sClient.Delete(context.Background(), created)).To(Succeed())
			Expect(k8sClient.Get(context.Background(), key, created)).ToNot(Succeed())
		})

		It("should correctly handle isSubmitted", func() {
			secretScope := &SecretScope{}
			Expect(secretScope.IsSubmitted()).To(BeFalse())
		})

		It("should correctly handle finalizers", func() {
			secretScope := &SecretScope{
				ObjectMeta: metav1.ObjectMeta{
					DeletionTimestamp: &metav1.Time{
						Time: time.Now(),
					},
				},
			}
			Expect(secretScope.IsBeingDeleted()).To(BeTrue())

			secretScope.AddFinalizer(SecretScopeFinalizerName)
			Expect(len(secretScope.GetFinalizers())).To(Equal(1))
			Expect(secretScope.HasFinalizer(SecretScopeFinalizerName)).To(BeTrue())

			secretScope.RemoveFinalizer(SecretScopeFinalizerName)
			Expect(len(secretScope.GetFinalizers())).To(Equal(0))
			Expect(secretScope.HasFinalizer(SecretScopeFinalizerName)).To(BeFalse())
		})
	})

})
