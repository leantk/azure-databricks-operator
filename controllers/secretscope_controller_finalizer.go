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

package controllers

import (
	"context"
	"fmt"

	databricksv1beta1 "github.com/microsoft/azure-databricks-operator/api/v1beta1"
)

func (r *SecretScopeReconciler) addFinalizer(instance *databricksv1beta1.SecretScope) error {
	instance.AddFinalizer(databricksv1beta1.SecretScopeFinalizerName)
	err := r.Update(context.Background(), instance)
	if err != nil {
		return fmt.Errorf("failed to update secret scope finalizer: %v", err)
	}
	return nil
}

func (r *SecretScopeReconciler) handleFinalizer(instance *databricksv1beta1.SecretScope) error {
	if instance.HasFinalizer(databricksv1beta1.SecretScopeFinalizerName) {
		if err := r.delete(instance); err != nil {
			return err
		}

		instance.RemoveFinalizer(databricksv1beta1.SecretScopeFinalizerName)
		if err := r.Update(context.Background(), instance); err != nil {
			return err
		}
	}
	// Our finalizer has finished, so the reconciler can do nothing.
	return nil
}
