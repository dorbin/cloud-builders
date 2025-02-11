/*
Copyright 2019 Google, Inc. All rights reserved.
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
package testservices

import (
	"context"
)

// TestGcloud implements the GcloudService interface.
type TestGcloud struct {
	ContainerImagesDescribeResp string
	ContainerImagesDescribeErr  error

	ContainerClustersGetCredentialsErr error

	ConfigGetValueResp string
	ConfigGetValueErr  error
}

// ContainerImageDescribe calls `gcloud container images describe <image> --format=<format>` and returns stdout.
func (g *TestGcloud) ContainerImagesDescribe(ctx context.Context, image, format string) (string, error) {
	return g.ContainerImagesDescribeResp, g.ContainerImagesDescribeErr
}

// ContainerClustersGetCredentials calls `gcloud container clusters get-credentials <clusterName> --zone=<clusterLocation> --project=<clusterProject>`.
func (g *TestGcloud) ContainerClustersGetCredentials(ctx context.Context, clusterName, clusterLocation, clusterProject string) error {
	return g.ContainerClustersGetCredentialsErr
}

// ConfigGetValue calls `gcloud config get-value <property>` and returns stdout.
func (g *TestGcloud) ConfigGetValue(ctx context.Context, property string) (string, error) {
	return g.ConfigGetValueResp, g.ConfigGetValueErr
}
