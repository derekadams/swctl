/**
 * Copyright © 2014-2020 The SiteWhere Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package infra

import (
	"fmt"
	"net/http"

	apiextensionsclientset "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	errors "k8s.io/apimachinery/pkg/api/errors"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"

	"github.com/sitewhere/swctl/pkg/resources"
)

// infraTemplateResource template for resources files
type infraTemplateResource struct {
	FileTemplate string
	FileCount    int
	Enabled      bool
}

var infraTemplate = []infraTemplateResource{
	infraTemplateResource{
		FileTemplate: "/infra/01-mosquitto/infra-mosquitto-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/02-redis/infra-redis-%02d.yaml",
		FileCount:    9,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/03-zookeeper/infra-zookeeper-%02d.yaml",
		FileCount:    5,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/04-kafka/infra-kafka-%02d.yaml",
		FileCount:    1,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/05-postgresql/infra-postgresql-%02d.yaml",
		FileCount:    3,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/06-syncope/infra-syncope-%02d.yaml",
		FileCount:    8,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/07-warp10/infra-warp10-%02d.yaml",
		FileCount:    4,
		Enabled:      false,
	},
	infraTemplateResource{
		FileTemplate: "/infra/08-influxdb/infra-influxdb-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/09-nifi/infra-nifi-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra/10-zoo-entrance/infra-zoo-entrance-%02d.yaml",
		FileCount:    3,
		Enabled:      true,
	},
}

var infraTemplateMin = []infraTemplateResource{
	infraTemplateResource{
		FileTemplate: "/infra-min/01-mosquitto/infra-mosquitto-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/02-redis/infra-redis-%02d.yaml",
		FileCount:    7,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/03-zookeeper/infra-zookeeper-%02d.yaml",
		FileCount:    5,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/04-kafka/infra-kafka-%02d.yaml",
		FileCount:    1,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/05-postgresql/infra-postgresql-%02d.yaml",
		FileCount:    3,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/06-syncope/infra-syncope-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/07-warp10/infra-warp10-%02d.yaml",
		FileCount:    4,
		Enabled:      false,
	},
	infraTemplateResource{
		FileTemplate: "/infra-min/08-influxdb/infra-influxdb-%02d.yaml",
		FileCount:    4,
		Enabled:      true,
	},
}

// InstallSiteWhereInfrastructureDefault installs SiteWhere Infrastructure components
func InstallSiteWhereInfrastructureDefault(waitReady bool,
	statikFS http.FileSystem,
	clientset kubernetes.Interface,
	apiextensionsClientset apiextensionsclientset.Interface,
	config *rest.Config) error {
	var err error

	for _, tpl := range infraTemplate {
		if tpl.Enabled {
			for i := 1; i <= tpl.FileCount; i++ {
				var infraResource = fmt.Sprintf(tpl.FileTemplate, i)
				err = resources.InstallResourceFromFile(infraResource, statikFS, clientset, apiextensionsClientset, config)
				if err != nil && !errors.IsAlreadyExists(err) {
					return err
				}
			}
		}
	}

	// if waitReady {
	// 	err = waitForDeploymentAvailable(clientset, "sitewhere-infrastructure-mosquitto", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Deployment sitewhere-infrastructure-mosquitto: ")
	// 	// 	color.Info.Println("Available")
	// 	// }

	// 	err = waitForDeploymentAvailable(clientset, "sitewhere-syncope", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Deployment sitewhere-syncope: ")
	// 	// 	color.Info.Println("Available")
	// 	// }

	// 	for i := 0; i < 3; i++ {
	// 		podName := fmt.Sprintf("sitewhere-kafka-zookeeper-%d", i)
	// 		err = waitForPodContainersRunning(clientset, podName, resources.SitewhereSystemNamespace() )
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// if config.IsVerbose() {
	// 		// 	fmt.Print(fmt.Sprintf("Pod %s: ", podName))
	// 		// 	color.Info.Println("Ready")
	// 		// }
	// 	}

	// 	for i := 0; i < 3; i++ {
	// 		podName := fmt.Sprintf("sitewhere-kafka-kafka-%d", i)
	// 		err = waitForPodContainersRunning(clientset, podName, resources.SitewhereSystemNamespace() )
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// if config.IsVerbose() {
	// 		// 	fmt.Print(fmt.Sprintf("Pod %s: ", podName))
	// 		// 	color.Info.Println("Ready")
	// 		// }
	// 	}

	// 	err = waitForPodContainersRunning(clientset, "sitewhere-postgresql-0", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-postgresql-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }

	// 	for i := 0; i < 3; i++ {
	// 		podName := fmt.Sprintf("sitewhere-infrastructure-redis-ha-server-%d", i)
	// 		err = waitForPodContainersRunning(clientset, podName, resources.SitewhereSystemNamespace() )
	// 		if err != nil {
	// 			return err
	// 		}
	// 		// if config.IsVerbose() {
	// 		// 	fmt.Print(fmt.Sprintf("Pod %s: ", podName))
	// 		// 	color.Info.Println("Ready")
	// 		// }
	// 	}

	// 	// err = waitForPodContainersRunning(clientset, "sitewhere-infrastructure-warp10-0", resources.SitewhereSystemNamespace() )
	// 	// if err != nil {
	// 	// 	return err
	// 	// }
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-infrastructure-warp10-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }
	// }
	return err
}

// InstallSiteWhereInfrastructureMinimal install SiteWhere Infrastructure, minimal profile.
// Minimal profile includes essential components only.
func InstallSiteWhereInfrastructureMinimal(waitReady bool,
	statikFS http.FileSystem,
	clientset kubernetes.Interface,
	apiextensionsClientset apiextensionsclientset.Interface,
	config *rest.Config) error {
	var err error

	for _, tpl := range infraTemplateMin {
		if tpl.Enabled {
			for i := 1; i <= tpl.FileCount; i++ {
				var infraResource = fmt.Sprintf(tpl.FileTemplate, i)
				err = resources.InstallResourceFromFile(infraResource, statikFS, clientset, apiextensionsClientset, config)
				if err != nil && !errors.IsAlreadyExists(err) {
					return err
				}
			}
		}
	}

	// if waitReady {
	// 	err = waitForDeploymentAvailable(clientset, "sitewhere-infrastructure-mosquitto", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Deployment sitewhere-infrastructure-mosquitto: ")
	// 	// 	color.Info.Println("Available")
	// 	// }

	// 	err = waitForDeploymentAvailable(clientset, "sitewhere-kafka-entity-operator", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Deployment sitewhere-kafka-entity-operator: ")
	// 	// 	color.Info.Println("Available")
	// 	// }

	// 	err = waitForDeploymentAvailable(clientset, "sitewhere-syncope", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Deployment sitewhere-syncope: ")
	// 	// 	color.Info.Println("Available")
	// 	// }

	// 	err = waitForPodContainersRunning(clientset, "sitewhere-infrastructure-zookeeper-0", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-infrastructure-zookeeper-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }
	// 	// TODO if not minimal, wait for -1 and -2

	// 	err = waitForPodContainersRunning(clientset, "sitewhere-kafka-kafka-0", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-kafka-kafka-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }

	// 	err = waitForPodContainersRunning(clientset, "sitewhere-postgresql-0", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-postgresql-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }

	// 	err = waitForPodContainersRunning(clientset, "sitewhere-infrastructure-redis-ha-server-0", resources.SitewhereSystemNamespace() )
	// 	if err != nil {
	// 		return err
	// 	}
	// 	// if config.IsVerbose() {
	// 	// 	fmt.Print("Pod sitewhere-infrastructure-redis-ha-server-0: ")
	// 	// 	color.Info.Println("Ready")
	// 	// }
	// }
	return nil
}

// UninstallSiteWhereInfrastructureMinimal uninstall SiteWhere infrastructure components
func UninstallSiteWhereInfrastructureMinimal(statikFS http.FileSystem,
	clientset kubernetes.Interface,
	apiextensionsClientset apiextensionsclientset.Interface,
	config *rest.Config) error {
	var err error

	for _, tpl := range infraTemplateMin {
		if tpl.Enabled {
			for i := 1; i <= tpl.FileCount; i++ {
				var infraResource = fmt.Sprintf(tpl.FileTemplate, i)
				err = resources.UninstallResourceFromFile(infraResource, statikFS, clientset, apiextensionsClientset, config)
				if err != nil && !errors.IsNotFound(err) {
					return err
				}
			}
		}
	}
	return nil
}

// UninstallSiteWhereInfrastructureDefault uninstall SiteWhere minimal infrastructure components
func UninstallSiteWhereInfrastructureDefault(statikFS http.FileSystem,
	clientset kubernetes.Interface,
	apiextensionsClientset apiextensionsclientset.Interface,
	config *rest.Config) error {
	var err error

	for _, tpl := range infraTemplate {
		if tpl.Enabled {
			for i := 1; i <= tpl.FileCount; i++ {
				var infraResource = fmt.Sprintf(tpl.FileTemplate, i)
				err = resources.UninstallResourceFromFile(infraResource, statikFS, clientset, apiextensionsClientset, config)
				if err != nil && !errors.IsNotFound(err) {
					return err
				}
			}
		}
	}
	return nil
}