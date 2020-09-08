/*
Copyright (c) SiteWhere, LLC. All rights reserved. http://www.sitewhere.com

The software in this package is published under the terms of the CPAL v1.0
license, a copy of which has been included with this distribution in the
LICENSE file.
*/

// Package internal Implements swctl internal use only functions
package internal

import (
	"log"
	"os"
	"path/filepath"
	"runtime"

	apixv1beta1client "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset/typed/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//var clientset *kubernetes.Clientset
//var apixClient *apixv1beta1client.ApiextensionsV1beta1Client

// GetKubeConfigFromKubeconfig Buid a Kubernetes Config from ~/.kube/config
func GetKubeConfigFromKubeconfig() (*rest.Config, error) {
	kubeconfig := filepath.Join(userHomeDir(), ".kube", "config")
	return GetKubeConfig(kubeconfig)
}

// GetKubeConfig Buid a Kubernetes Config from a filepath
func GetKubeConfig(pathToCfg string) (*rest.Config, error) {
	if pathToCfg == "" {
		// in cluster access
		return rest.InClusterConfig()
	}

	cfg, err := clientcmd.BuildConfigFromFlags("", pathToCfg)

	if err != nil {
		log.Println("Using in cluster config")
		return GetKubeConfig("")
	}
	return cfg, nil
}

// GetKubernetesClient Returns Kubernetes Client
func GetKubernetesClient(pathToCfg string) (*kubernetes.Clientset, error) {
	config, err := GetKubeConfig(pathToCfg)

	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

// GetKubernetesClientV1Beta1 Returns Kubernetes Client v1 beta1 interface
func GetKubernetesClientV1Beta1(pathToCfg string) (*apixv1beta1client.ApiextensionsV1beta1Client, error) {
	config, err := GetKubeConfig(pathToCfg)

	if err != nil {
		return nil, err
	}

	return apixv1beta1client.NewForConfig(config)
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return home
		}
	}
	return os.Getenv("HOME")
}
