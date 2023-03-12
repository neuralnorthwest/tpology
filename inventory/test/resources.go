// Copyright 2023 Scott M. Long
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import "github.com/neuralnorthwest/tpology/resource"

// Example of an application resource:
//
// name: some-app
// description: Some application
// owner: some-team
// application:
//   git:
//     repo: some-repo
//     branch: some-branch
//     path: some-path
//   deployments:
//     - environment: some-environment
//
// Example of an environment resource:
//
// name: some-environment
// description: Some environment
// owner: some-team
// environment:
//   - cluster: some-cluster
//     namespace: some-namespace
//
// Example of a cluster resource:
//
// name: some-cluster
// description: Some cluster
// owner: some-team
// cluster:
//   provider: some-provider
//   project: some-project
//   region: some-region
//   zone: some-zone
//
// Example of a provider resource:
//
// name: some-provider
// description: Some provider
// owner: some-team
// provider:
//   type: some-type
//   account: some-account

// makeTestResources returns a plausible set of resources for testing.
func MakeTestResources() []*resource.Resource {
	resources := []*resource.Resource{}
	// create some providers - gcp and aws
	provider_gcp := &resource.Resource{
		Name:        "gcp",
		Description: "Google Cloud Platform",
		Owner:       "gcp-team",
		Kind:        "provider",
		Data: map[string]interface{}{
			"type":    "gcp",
			"account": "gcp-account",
		},
	}
	provider_aws := &resource.Resource{
		Name:        "aws",
		Description: "Amazon Web Services",
		Owner:       "aws-team",
		Kind:        "provider",
		Data: map[string]interface{}{
			"type":    "aws",
			"account": "aws-account",
		},
	}

	// create some clusters - gcp-dev, gcp-prod, aws-dev, aws-prod
	cluster_gcp_dev := &resource.Resource{
		Name:        "gcp-dev",
		Description: "Google Cloud Platform Development Cluster",
		Owner:       "gcp-team",
		Kind:        "cluster",
		Data: map[string]interface{}{
			"provider": "gcp",
			"project":  "gcp-dev-project",
			"region":   "us-east1",
			"zone":     "us-east1-b",
		},
	}
	cluster_gcp_prod := &resource.Resource{
		Name:        "gcp-prod",
		Description: "Google Cloud Platform Production Cluster",
		Owner:       "gcp-team",
		Kind:        "cluster",
		Data: map[string]interface{}{
			"provider": "gcp",
			"project":  "gcp-prod-project",
			"region":   "us-east1",
			"zone":     "us-east1-b",
		},
	}
	cluster_aws_dev := &resource.Resource{
		Name:        "aws-dev",
		Description: "Amazon Web Services Development Cluster",
		Owner:       "aws-team",
		Kind:        "cluster",
		Data: map[string]interface{}{
			"provider": "aws",
			"project":  "aws-dev-project",
			"region":   "us-east-1",
			"zone":     "us-east-1b",
		},
	}
	cluster_aws_prod := &resource.Resource{
		Name:        "aws-prod",
		Description: "Amazon Web Services Production Cluster",
		Owner:       "aws-team",
		Kind:        "cluster",
		Data: map[string]interface{}{
			"provider": "aws",
			"project":  "aws-prod-project",
			"region":   "us-east-1",
			"zone":     "us-east-1b",
		},
	}

	// create some environments - dev and prod
	environment_dev := &resource.Resource{
		Name:        "dev",
		Description: "Development Environment",
		Owner:       "dev-team",
		Kind:        "environment",
		Data: map[string]interface{}{
			"clusters": []interface{}{
				map[string]interface{}{
					"cluster":   "gcp-dev",
					"namespace": "dev",
				},
				map[string]interface{}{
					"cluster":   "aws-dev",
					"namespace": "dev",
				},
			},
		},
	}
	environment_prod := &resource.Resource{
		Name:        "prod",
		Description: "Production Environment",
		Owner:       "prod-team",
		Kind:        "environment",
		Data: map[string]interface{}{
			"clusters": []interface{}{
				map[string]interface{}{
					"cluster":   "gcp-prod",
					"namespace": "prod",
				},
				map[string]interface{}{
					"cluster":   "aws-prod",
					"namespace": "prod",
				},
			},
		},
	}

	// create some applications - app1 and app2
	application_app1 := &resource.Resource{
		Name:        "app1",
		Description: "Application 1",
		Owner:       "app1-team",
		Kind:        "application",
		Data: map[string]interface{}{
			"git": map[string]interface{}{
				"repo":   "app1-repo",
				"branch": "app1-branch",
				"path":   "app1-path",
			},
			"deployments": []interface{}{
				map[string]interface{}{
					"environment": "dev",
				},
				map[string]interface{}{
					"environment": "prod",
				},
			},
		},
	}
	application_app2 := &resource.Resource{
		Name:        "app2",
		Description: "Application 2",
		Owner:       "app2-team",
		Kind:        "application",
		Data: map[string]interface{}{
			"git": map[string]interface{}{
				"repo":   "app2-repo",
				"branch": "app2-branch",
				"path":   "app2-path",
			},
			"deployments": []interface{}{
				map[string]interface{}{
					"environment": "dev",
				},
				map[string]interface{}{
					"environment": "prod",
				},
			},
		},
	}

	resources = append(resources, provider_gcp)
	resources = append(resources, provider_aws)
	resources = append(resources, cluster_gcp_dev)
	resources = append(resources, cluster_gcp_prod)
	resources = append(resources, cluster_aws_dev)
	resources = append(resources, cluster_aws_prod)
	resources = append(resources, environment_dev)
	resources = append(resources, environment_prod)
	resources = append(resources, application_app1)
	resources = append(resources, application_app2)

	return resources
}
