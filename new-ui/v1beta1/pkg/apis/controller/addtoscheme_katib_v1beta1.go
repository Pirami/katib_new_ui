/*
Copyright 2022 The Kubeflow Authors.

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

package apis

import (
	experiments "github.com/kubeflow/katib/pkg/apis/controller/experiments/v1beta1"
	suggestions "github.com/kubeflow/katib/pkg/apis/controller/suggestions/v1beta1"
	trials "github.com/kubeflow/katib/pkg/apis/controller/trials/v1beta1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		experiments.SchemeBuilder.AddToScheme,
		trials.SchemeBuilder.AddToScheme,
		suggestions.SchemeBuilder.AddToScheme,
	)
}
