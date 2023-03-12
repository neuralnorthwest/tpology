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

package inventory

import (
	"testing"

	"github.com/neuralnorthwest/tpology/inventory/test"
	"github.com/stretchr/testify/assert"
)

// Test_BuildGraph tests the BuildGraph function.
func Test_Inventory_BuildGraph(t *testing.T) {
	resources := test.MakeTestResources()
	inv := New(resources...)
	g, err := inv.BuildGraph()
	if err != nil {
		t.Errorf("Error building graph: %v", err)
	}
	assert.Len(t, g.Nodes, 4)
	assert.Len(t, g.Nodes["application"], 2)
	assert.Len(t, g.Nodes["environment"], 2)
	assert.Len(t, g.Nodes["cluster"], 4)
	assert.Len(t, g.Nodes["provider"], 2)

	if resource, ok := g.Nodes["application"]["app1"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if deployments, ok := data["deployments"]; ok {
			assert.Len(t, deployments.([]interface{}), 2)
			expectedEnv := []string{"dev", "prod"}
			for i, env := range deployments.([]interface{}) {
				envMap := env.(map[string]interface{})
				if envResource, ok := envMap["environment"]; ok {
					if envNode, ok := envResource.(*Node); ok {
						assert.Equal(t, expectedEnv[i], envNode.Name)
					} else {
						t.Errorf("Error getting environment resource")
					}
				} else {
					t.Errorf("Error getting environment")
				}
			}
		} else {
			t.Errorf("Error getting deployments")
		}
	} else {
		t.Errorf("Error getting app1")
	}

	if resource, ok := g.Nodes["application"]["app2"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if deployments, ok := data["deployments"]; ok {
			assert.Len(t, deployments.([]interface{}), 2)
			expectedEnv := []string{"dev", "prod"}
			for i, env := range deployments.([]interface{}) {
				envMap := env.(map[string]interface{})
				if envResource, ok := envMap["environment"]; ok {
					if envNode, ok := envResource.(*Node); ok {
						assert.Equal(t, expectedEnv[i], envNode.Name)
					} else {
						t.Errorf("Error getting environment resource")
					}
				} else {
					t.Errorf("Error getting environment")
				}
			}
		} else {
			t.Errorf("Error getting deployments")
		}
	} else {
		t.Errorf("Error getting app2")
	}

	if resource, ok := g.Nodes["environment"]["dev"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if clusters, ok := data["clusters"]; ok {
			assert.Len(t, clusters.([]interface{}), 2)
			expectedCluster := []string{"gcp-dev", "aws-dev"}
			for i, cluster := range clusters.([]interface{}) {
				clusterMap := cluster.(map[string]interface{})
				if clusterResource, ok := clusterMap["cluster"]; ok {
					if clusterNode, ok := clusterResource.(*Node); ok {
						assert.Equal(t, expectedCluster[i], clusterNode.Name)
					} else {
						t.Errorf("Error getting cluster resource")
					}
				} else {
					t.Errorf("Error getting cluster")
				}
			}
		} else {
			t.Errorf("Error getting clusters")
		}
	} else {
		t.Errorf("Error getting dev")
	}

	if resource, ok := g.Nodes["environment"]["prod"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if clusters, ok := data["clusters"]; ok {
			assert.Len(t, clusters.([]interface{}), 2)
			expectedCluster := []string{"gcp-prod", "aws-prod"}
			for i, cluster := range clusters.([]interface{}) {
				clusterMap := cluster.(map[string]interface{})
				if clusterResource, ok := clusterMap["cluster"]; ok {
					if clusterNode, ok := clusterResource.(*Node); ok {
						assert.Equal(t, expectedCluster[i], clusterNode.Name)
					} else {
						t.Errorf("Error getting cluster resource")
					}
				} else {
					t.Errorf("Error getting cluster")
				}
			}
		} else {
			t.Errorf("Error getting clusters")
		}
	} else {
		t.Errorf("Error getting prod")
	}

	if resource, ok := g.Nodes["cluster"]["gcp-dev"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if provider, ok := data["provider"]; ok {
			if providerNode, ok := provider.(*Node); ok {
				assert.Equal(t, "gcp", providerNode.Name)
			} else {
				t.Errorf("Error getting provider resource")
			}
		} else {
			t.Errorf("Error getting provider")
		}
	} else {
		t.Errorf("Error getting gcp-dev")
	}

	if resource, ok := g.Nodes["cluster"]["gcp-prod"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if provider, ok := data["provider"]; ok {
			if providerNode, ok := provider.(*Node); ok {
				assert.Equal(t, "gcp", providerNode.Name)
			} else {
				t.Errorf("Error getting provider resource")
			}
		} else {
			t.Errorf("Error getting provider")
		}
	} else {
		t.Errorf("Error getting gcp-prod")
	}

	if resource, ok := g.Nodes["cluster"]["aws-dev"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if provider, ok := data["provider"]; ok {
			if providerNode, ok := provider.(*Node); ok {
				assert.Equal(t, "aws", providerNode.Name)
			} else {
				t.Errorf("Error getting provider resource")
			}
		} else {
			t.Errorf("Error getting provider")
		}
	} else {
		t.Errorf("Error getting aws-dev")
	}

	if resource, ok := g.Nodes["cluster"]["aws-prod"]; ok {
		data := resource.GraphData.(map[string]interface{})
		if provider, ok := data["provider"]; ok {
			if providerNode, ok := provider.(*Node); ok {
				assert.Equal(t, "aws", providerNode.Name)
			} else {
				t.Errorf("Error getting provider resource")
			}
		} else {
			t.Errorf("Error getting provider")
		}
	} else {
		t.Errorf("Error getting aws-prod")
	}
}
