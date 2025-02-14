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

package v1beta1

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("verticadb_types", func() {
	const FakeUID = "abcdef"

	It("should include UID in path if IncludeUIDInPath is set", func() {
		vdb := MakeVDB()
		vdb.ObjectMeta.UID = FakeUID
		vdb.Spec.Communal.IncludeUIDInPath = true
		Expect(vdb.GetCommunalPath()).Should(ContainSubstring(string(vdb.ObjectMeta.UID)))
	})

	It("should not include UID in path if IncludeUIDInPath is not set", func() {
		vdb := MakeVDB()
		vdb.ObjectMeta.UID = FakeUID
		vdb.Spec.Communal.IncludeUIDInPath = false
		Expect(vdb.GetCommunalPath()).ShouldNot(ContainSubstring(string(vdb.ObjectMeta.UID)))
	})

	It("should require a transient subcluster", func() {
		vdb := MakeVDB()
		vdb.Spec.Subclusters = []Subcluster{
			{Name: "sc1"},
			{Name: "sc2"},
		}
		// Transient is only required if specified
		Expect(vdb.RequiresTransientSubcluster()).Should(BeFalse())
		vdb.Spec.TemporarySubclusterRouting.Names = []string{"sc1"}
		Expect(vdb.RequiresTransientSubcluster()).Should(BeFalse())
		vdb.Spec.TemporarySubclusterRouting.Template = Subcluster{
			Name:      "the-transient-sc-name",
			Size:      1,
			IsPrimary: false,
		}
		Expect(vdb.RequiresTransientSubcluster()).Should(BeTrue())
	})

	It("should say that the httpServer is not enabled if value is blank", func() {
		vdb := MakeVDB()
		vdb.Spec.HTTPServerMode = ""
		Expect(vdb.IsHTTPServerEnabled()).Should(BeFalse())
		vdb.Spec.HTTPServerMode = HTTPServerModeEnabled
		Expect(vdb.IsHTTPServerEnabled()).Should(BeTrue())
	})

	It("should enable httpServer if auto is set and vertica version is new enough", func() {
		vdb := MakeVDB()
		vdb.Spec.HTTPServerMode = HTTPServerModeAuto
		vdb.ObjectMeta.Annotations[VersionAnnotation] = "v11.0.2"
		Expect(vdb.IsHTTPServerEnabled()).Should(BeFalse())
		vdb.ObjectMeta.Annotations[VersionAnnotation] = HTTPServerMinVersion
		Expect(vdb.IsHTTPServerEnabled()).Should(BeFalse())
		vdb.ObjectMeta.Annotations[VersionAnnotation] = HTTPServerAutoMinVersion
		Expect(vdb.IsHTTPServerEnabled()).Should(BeTrue())
	})
})
