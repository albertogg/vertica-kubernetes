/*
 (c) Copyright [2021-2023] Open Text.
 Licensed under the Apache License, Version 2.0 (the "License");
 You may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package vdb

import (
	"bytes"
	"context"
	"fmt"

	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/cmds"
	"github.com/vertica/vertica-kubernetes/pkg/names"
	"github.com/vertica/vertica-kubernetes/pkg/paths"
	"k8s.io/apimachinery/pkg/types"
)

// prepLocalData Prepare for the add node or create_db by removing any local
// data/depot dirs and ensuring proper ownership.
// This step is necessary because of a lack of cleanup in admintools if any of
// these commands fail.
func prepLocalData(ctx context.Context, vdb *vapi.VerticaDB, prunner cmds.PodRunner, podName types.NamespacedName) error {
	locPaths := []string{vdb.GetDBDataPath(), vdb.GetDBDepotPath(), vdb.GetDBCatalogPath()}
	var rmCmds bytes.Buffer
	rmCmds.WriteString("set -o errexit\n")
	for _, path := range locPaths {
		rmCmds.WriteString(fmt.Sprintf("[[ -d %s ]] && rm -rf %s || true\n", path, path))
	}
	// We also need to ensure the dbadmin owns the depot directory.  When the
	// directory are first mounted they are owned by root.  Vertica handles changing
	// the ownership of the config, log and data directory.  This function exists to
	// handle the depot directory. This can be skipped if the depotPath is
	// shared with one of the data or catalog paths.
	if vdb.Spec.Local.DepotPath != vdb.Spec.Local.DataPath && vdb.Spec.Local.DepotPath != vdb.Spec.Local.GetCatalogPath() {
		rmCmds.WriteString(fmt.Sprintf("sudo chown dbadmin:verticadba -R %s/%s", paths.LocalDataPath, vdb.GetPVSubPath("depot")))
	}

	cmd := []string{"bash", "-c", fmt.Sprintf("cat > %s<<< '%s'; bash %s",
		paths.PrepScript, rmCmds.String(), paths.PrepScript)}
	if _, _, err := prunner.ExecInPod(ctx, podName, names.ServerContainer, cmd...); err != nil {
		return err
	}
	return nil
}

// debugDumpAdmintoolsConf will dump specific info from admintools.conf for logging purposes
// +nolint
func debugDumpAdmintoolsConf(ctx context.Context, prunner cmds.PodRunner, atPod types.NamespacedName) {
	// Dump out vital informating from admintools.conf for logging purposes. We
	// rely on the logging that is done inside ExecInPod.
	cmd := []string{
		"bash", "-c",
		fmt.Sprintf(`ls -l %s && grep '^node\|^v_\|^host' %s`, paths.AdminToolsConf, paths.AdminToolsConf),
	}
	// Since this is for debugging purposes all errors are ignored
	prunner.ExecInPod(ctx, atPod, names.ServerContainer, cmd...) //nolint:errcheck
}

// debugDumpAdmintoolsConfForPods will dump debug information for admintools.conf for a list of pods
func debugDumpAdmintoolsConfForPods(ctx context.Context, prunner cmds.PodRunner, pods []*PodFact) {
	for _, pod := range pods {
		debugDumpAdmintoolsConf(ctx, prunner, pod.name)
	}
}
