// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package kubernetes

import (
	"k8s.io/client-go/kubernetes"

	"github.com/tensorchord/envd-server/pkg/runtime"
)

type generalProvisioner struct {
	client kubernetes.Interface

	namespace string
}

func NewProvisioner(client kubernetes.Interface) runtime.Provisioner {
	return &generalProvisioner{
		client:    client,
		namespace: "default",
	}
}