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
	"context"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/atconf"
	"github.com/vertica/vertica-kubernetes/pkg/cmds"
	"github.com/vertica/vertica-kubernetes/pkg/names"
	"github.com/vertica/vertica-kubernetes/pkg/paths"
	"github.com/vertica/vertica-kubernetes/pkg/test"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ctrl "sigs.k8s.io/controller-runtime"
)

var _ = Describe("k8s/install_reconcile_test", func() {
	ctx := context.Background()

	It("should detect no install is needed", func() {
		vdb := vapi.MakeVDB()
		test.CreatePods(ctx, k8sClient, vdb, true)
		defer test.DeletePods(ctx, k8sClient, vdb)

		sc := &vdb.Spec.Subclusters[0]
		fpr := &cmds.FakePodRunner{}
		pfact := createPodFactsDefault(fpr)
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, pfact)
		drecon := actor.(*InstallReconciler)
		Expect(drecon.Reconcile(ctx, &ctrl.Request{})).Should(Equal(ctrl.Result{}))
		for i := int32(0); i < 3; i++ {
			Expect(drecon.PFacts.Detail[names.GenPodName(vdb, sc, i)].isInstalled).Should(BeTrue(), fmt.Sprintf("Pod index %d", i))
		}
	})

	It("should try install if a pod has not run the installer yet", func() {
		vdb := vapi.MakeVDB()
		test.CreatePods(ctx, k8sClient, vdb, test.AllPodsRunning)
		defer test.DeletePods(ctx, k8sClient, vdb)

		sc := &vdb.Spec.Subclusters[0]
		fpr := &cmds.FakePodRunner{}
		pfact := createPodFactsDefault(fpr)
		Expect(pfact.Collect(ctx, vdb)).Should(Succeed())
		pfact.Detail[names.GenPodName(vdb, sc, 1)].dbExists = false
		pfact.Detail[names.GenPodName(vdb, sc, 1)].isInstalled = false
		pfact.Detail[names.GenPodName(vdb, sc, 2)].dbExists = false
		pfact.Detail[names.GenPodName(vdb, sc, 2)].isInstalled = false
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, pfact)
		drecon := actor.(*InstallReconciler)
		drecon.ATWriter = &atconf.FakeWriter{}
		Expect(drecon.Reconcile(ctx, &ctrl.Request{})).Should(Equal(ctrl.Result{}))
		cmdHist := fpr.FindCommands(fmt.Sprintf("cat > %s", paths.AdminToolsConf))
		Expect(len(cmdHist)).Should(Equal(3))
		// We should see two instances of creating the install indicator -- one at each host that we install at
		cmdHist = fpr.FindCommands(vdb.GenInstallerIndicatorFileName())
		Expect(len(cmdHist)).Should(Equal(2))
	})

	It("should skip call exec on a pod if is not yet running", func() {
		vdb := vapi.MakeVDB()
		const ScSize = 2
		vdb.Spec.Subclusters[0].Size = ScSize
		vdb.Status.Subclusters = []vapi.SubclusterStatus{
			{Name: vdb.Spec.Subclusters[0].Name, InstallCount: ScSize - 1, Detail: []vapi.VerticaDBPodStatus{}},
		}
		test.CreatePods(ctx, k8sClient, vdb, test.AllPodsNotRunning)
		defer test.DeletePods(ctx, k8sClient, vdb)

		fpr := &cmds.FakePodRunner{Results: cmds.CmdResults{}}
		pfact := MakePodFacts(vdbRec, fpr)
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, &pfact)
		drecon := actor.(*InstallReconciler)
		drecon.ATWriter = &atconf.FakeWriter{}
		res, err := drecon.Reconcile(ctx, &ctrl.Request{})
		Expect(err).Should(Succeed())
		Expect(res.Requeue).Should(BeTrue())
		Expect(len(fpr.Histories)).Should(Equal(0))
	})

	It("try install when not all pods are running", func() {
		vdb := vapi.MakeVDB()
		const ScIndex = 0
		sc := &vdb.Spec.Subclusters[ScIndex]
		sc.Size = 2
		vdb.Status.Subclusters = []vapi.SubclusterStatus{
			{Name: vdb.Spec.Subclusters[0].Name, InstallCount: sc.Size - 1, Detail: []vapi.VerticaDBPodStatus{}},
		}
		test.CreatePods(ctx, k8sClient, vdb, test.AllPodsNotRunning)
		defer test.DeletePods(ctx, k8sClient, vdb)
		// Make only pod -1 runable.
		const PodIndex = 1
		test.SetPodStatus(ctx, k8sClient, 1 /* funcOffset */, names.GenPodName(vdb, sc, 1), ScIndex, PodIndex, test.AllPodsRunning)

		fpr := &cmds.FakePodRunner{}
		pfact := MakePodFacts(vdbRec, fpr)
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, &pfact)
		drecon := actor.(*InstallReconciler)
		res, err := drecon.Reconcile(ctx, &ctrl.Request{})
		Expect(err).Should(Succeed())
		Expect(res.Requeue).Should(BeTrue())
	})

	It("install should accept eula", func() {
		vdb := vapi.MakeVDB()
		const ScIndex = 0
		sc := &vdb.Spec.Subclusters[ScIndex]
		sc.Size = 2
		test.CreatePods(ctx, k8sClient, vdb, test.AllPodsRunning)
		defer test.DeletePods(ctx, k8sClient, vdb)

		fpr := &cmds.FakePodRunner{}
		pfact := createPodFactsWithInstallNeeded(ctx, vdb, fpr)
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, pfact)
		drecon := actor.(*InstallReconciler)
		err := drecon.acceptEulaIfMissing(ctx)
		Expect(err).Should(Succeed())
		cmds := fpr.FindCommands(paths.EulaAcceptanceScript)
		Expect(len(cmds)).Should(Equal(4)) // 2 for each pod; 1 to copy and 1 to execute the script
	})

	It("should install pods in pod-index order", func() {
		vdb := vapi.MakeVDB()
		const ScIndex = 0
		sc := &vdb.Spec.Subclusters[ScIndex]
		sc.Size = 3
		test.CreatePods(ctx, k8sClient, vdb, test.AllPodsRunning)
		defer test.DeletePods(ctx, k8sClient, vdb)

		fpr := &cmds.FakePodRunner{}
		pfact := createPodFactsWithInstallNeeded(ctx, vdb, fpr)
		// Make pod-1 not running.  This will prevent install of pod-1 and pod-2
		pn := names.GenPodName(vdb, sc, 1)
		pfact.Detail[pn].isPodRunning = false
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, pfact)
		drecon := actor.(*InstallReconciler)
		podList, err := drecon.getInstallTargets(ctx)
		Expect(err).Should(Succeed())
		Expect(len(podList)).Should(Equal(1))
		Expect(podList[0].name).Should(Equal(names.GenPodName(vdb, sc, 0)))
	})

	It("should generate certs only on supported vertica versions", func() {
		vdb := vapi.MakeVDB()
		secret := corev1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "tls-secret",
				Namespace: vdb.Namespace,
			},
			Data: map[string][]byte{
				corev1.TLSPrivateKeyKey:   []byte("pk"),
				corev1.TLSCertKey:         []byte("cert"),
				paths.HTTPServerCACrtName: []byte("ca"),
			},
		}
		Expect(k8sClient.Create(ctx, &secret)).Should(Succeed())
		defer func() { Expect(k8sClient.Delete(ctx, &secret)) }()
		vdb.Spec.HTTPServerMode = vapi.HTTPServerModeEnabled
		vdb.Spec.HTTPServerTLSSecret = secret.Name
		vdb.Annotations[vapi.VersionAnnotation] = "v12.0.0"

		fpr := &cmds.FakePodRunner{}
		pfact := createPodFactsWithInstallNeeded(ctx, vdb, fpr)
		actor := MakeInstallReconciler(vdbRec, logger, vdb, fpr, pfact)
		drecon := actor.(*InstallReconciler)
		for _, val := range pfact.Detail {
			Expect(drecon.genCreateConfigDirsScript(val)).ShouldNot(ContainSubstring(paths.HTTPTLSConfDir))
		}
		err := drecon.generateHTTPCerts(ctx)
		Expect(err).Should(Succeed())
		cmds := fpr.FindCommands(paths.HTTPTLSConfFileName)
		Expect(len(cmds)).Should(Equal(0))

		vdb.Annotations[vapi.VersionAnnotation] = vapi.HTTPServerMinVersion
		for _, val := range pfact.Detail {
			Expect(drecon.genCreateConfigDirsScript(val)).Should(ContainSubstring(paths.HTTPTLSConfDir))
		}
		err = drecon.generateHTTPCerts(ctx)
		Expect(err).Should(Succeed())
		cmds = fpr.FindCommands(paths.HTTPTLSConfFileName)
		Expect(len(cmds)).Should(Equal(int(vdb.Spec.Subclusters[0].Size)))
	})
})
