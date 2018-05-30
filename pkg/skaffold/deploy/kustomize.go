/*
Copyright 2018 Google LLC

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

package deploy

import (
	"context"
	"io"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/build"
	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/schema/v1alpha2"
)

type KustomizeDeployer struct {
	*v1alpha2.DeployConfig
	kubeContext string
}

// NewKustomizeDeployer returns a new KustomizeDeployer for a DeployConfig filled
// with the needed configuration for `helm`
func NewKustomizeDeployer(cfg *v1alpha2.DeployConfig, kubeContext string) *KustomizeDeployer {
	return &KustomizeDeployer{
		DeployConfig: cfg,
		kubeContext:  kubeContext,
	}
}

func (h *KustomizeDeployer) Deploy(ctx context.Context, out io.Writer, builds []build.Build) error {

	return nil
}

// Not implemented
func (k *KustomizeDeployer) Dependencies() ([]string, error) {
	return nil, nil
}

// Cleanup deletes what was deployed by calling Deploy.
func (h *KustomizeDeployer) Cleanup(ctx context.Context, out io.Writer) error {

	return nil
}
