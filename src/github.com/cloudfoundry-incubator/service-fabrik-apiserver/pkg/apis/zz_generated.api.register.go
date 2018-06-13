/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package apis

import (
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup"
	backupv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup/v1alpha1"
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment"
	deploymentv1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment/v1alpha1"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
)

// GetAllApiBuilders returns all known APIGroupBuilders
// so they can be registered with the apiserver
func GetAllApiBuilders() []*builders.APIGroupBuilder {
	return []*builders.APIGroupBuilder{
		GetBackupAPIBuilder(),
		GetDeploymentAPIBuilder(),
	}
}

var backupApiGroup = builders.NewApiGroupBuilder(
	"backup.servicefabrik.io",
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/backup").
	WithUnVersionedApi(backup.ApiVersion).
	WithVersionedApis(
		backupv1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetBackupAPIBuilder() *builders.APIGroupBuilder {
	return backupApiGroup
}

var deploymentApiGroup = builders.NewApiGroupBuilder(
	"deployment.servicefabrik.io",
	"github.com/cloudfoundry-incubator/service-fabrik-apiserver/pkg/apis/deployment").
	WithUnVersionedApi(deployment.ApiVersion).
	WithVersionedApis(
		deploymentv1alpha1.ApiVersion,
	).
	WithRootScopedKinds()

func GetDeploymentAPIBuilder() *builders.APIGroupBuilder {
	return deploymentApiGroup
}
