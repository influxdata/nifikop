// Copyright © 2019 Banzai Cloud
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

package common

import (
	"fmt"
	"github.com/go-logr/logr"
	"github.com/erdrix/nifikop/pkg/apis/nifi/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// clusterRefLabel is the label key used for referencing NifiUsers/NifiDataflow
// to a NifiCluster
var ClusterRefLabel = "nifiCluster"


// requeueWithError is a convenience wrapper around logging an error message
// separate from the stacktrace and then passing the error through to the controller
// manager
func RequeueWithError(logger logr.Logger, msg string, err error) (reconcile.Result, error) {
	// Info log the error message and then let the reconciler dump the stacktrace
	logger.Info(msg)
	return reconcile.Result{}, err
}

// reconciled returns an empty result with nil error to signal a successful reconcile
// to the controller manager
func Reconciled() (reconcile.Result, error) {
	return reconcile.Result{}, nil
}


// clusterLabelString returns the label value for a cluster reference
func ClusterLabelString(cluster *v1alpha1.NifiCluster) string {
	return fmt.Sprintf("%s.%s", cluster.Name, cluster.Namespace)
}

// applyClusterRefLabel ensures a map of labels contains a reference to a parent nifi cluster
func ApplyClusterRefLabel(cluster *v1alpha1.NifiCluster, labels map[string]string) map[string]string {
	labelValue := ClusterLabelString(cluster)
	if labels == nil {
		labels = make(map[string]string, 0)
	}
	if label, ok := labels[ClusterRefLabel]; ok {
		if label != labelValue {
			labels[ClusterRefLabel] = labelValue
		}
	} else {
		labels[ClusterRefLabel] = labelValue
	}
	return labels
}