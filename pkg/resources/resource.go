// Copyright 2020 Yunion
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resources

import (
	"context"

	onecloudv1 "yunion.io/x/onecloud-service-operator/api/v1"
)

type OCResource interface {
	GetResourceName() Resource
	GetIResource() onecloudv1.IResource
	Create(ctx context.Context, params interface{}) (onecloudv1.ExternalInfoBase, error)
	Delete(ctx context.Context) (onecloudv1.ExternalInfoBase, error)
	GetStatus(ctx context.Context) (onecloudv1.IResourceStatus, error)
}
