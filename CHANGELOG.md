# Changelog
All notable changes to this project will be documented in this file.

The file is generated by [Changie](https://github.com/miniscruff/changie).


## 1.10.2 - 2023-04-11
### Changed
* [#367](https://github.com/vertica/vertica-kubernetes/issues/367) Use 12.0.4 as default vertica server image
* [#365](https://github.com/vertica/vertica-kubernetes/issues/365) Moved to operator-sdk v1.28.0
### Fixed
* [#369](https://github.com/vertica/vertica-kubernetes/issues/369) Fix helm install without cluster admin priv
* [#362](https://github.com/vertica/vertica-kubernetes/issues/362) Support subcluster names with underscores, such as default_subcluster.
* [#360](https://github.com/vertica/vertica-kubernetes/issues/360) Run rebalance shards on new subcluster created in a v11 database that was migrated from enterprise
* [#353](https://github.com/vertica/vertica-kubernetes/issues/353) Setup keys for client side agent access

## 1.10.1 - 2023-03-13
### Added
* [#349](https://github.com/vertica/vertica-kubernetes/issues/349) Backdoor to run the Vertica agent. This is to be used for development purposes only.
### Changed
* [#342](https://github.com/vertica/vertica-kubernetes/issues/342) The default value for spec.httpServerMode is to enable the http server in server versions 12.0.4 or newer.
* [#343](https://github.com/vertica/vertica-kubernetes/issues/343) Remove keys from the vertica-k8s container. This will be available in the first server version after 12.0.4.
### Fixed
* [#345](https://github.com/vertica/vertica-kubernetes/issues/345) Regression in 1.10.0 that prevents the operator from restarting vertica if the pod has sidecars.

## 1.10.0 - 2023-02-26
### Added
* [#337](https://github.com/vertica/vertica-kubernetes/issues/337) Add config knob for pod-level securityContext of vertica pod's
* [#328](https://github.com/vertica/vertica-kubernetes/issues/328) Allow scheduling rules for operator pod
* [#325](https://github.com/vertica/vertica-kubernetes/issues/325) Add startupProbe and livenessProbe for the server
* [#320](https://github.com/vertica/vertica-kubernetes/issues/320) Add an init program to the vertica-k8s container to reap zombies. This will be available in server versions 12.0.4 and higher.
### Changed
* [#332](https://github.com/vertica/vertica-kubernetes/issues/332) Allow revive when local paths aren't known
* [#323](https://github.com/vertica/vertica-kubernetes/issues/323) Use 12.0.3 as default vertica server image
* [#320](https://github.com/vertica/vertica-kubernetes/issues/320) Use fsGroup security policy so that mounted PVs have write access for dbadmin
### Removed
* [#320](https://github.com/vertica/vertica-kubernetes/issues/320) Support for Vertica server 11.0.0. New minimum version it supports is 11.0.1.
### Fixed
* [#336](https://github.com/vertica/vertica-kubernetes/issues/336) Allow revive if some data paths differ
* [#330](https://github.com/vertica/vertica-kubernetes/issues/330) Make disk full errors more prominent

## 1.9.0 - 2023-01-11
### Added
* [#309](https://github.com/vertica/vertica-kubernetes/issues/309) Allow the readinessProbe to be configured
* [#308](https://github.com/vertica/vertica-kubernetes/issues/308) Allow posix path as communal path
* [#300](https://github.com/vertica/vertica-kubernetes/issues/300) Include a label in the operator's Prometheus metrics to identify the database uniquely
* [#290](https://github.com/vertica/vertica-kubernetes/issues/290) Exposed the http port in the service object
* [#287](https://github.com/vertica/vertica-kubernetes/issues/287) Allow authorization to /metrics endpoint with TLS certificates
### Changed
* [#304](https://github.com/vertica/vertica-kubernetes/issues/304) Prometheus metrics for subcluster to include label for subcluster oid rather than subcluster name
* [#296](https://github.com/vertica/vertica-kubernetes/issues/296) Moved to operator-sdk v1.25.2
* [#290](https://github.com/vertica/vertica-kubernetes/issues/290) Renamed spec.httpServerSecret in VerticaDB to spec.httpServerTLSSecret
* [#287](https://github.com/vertica/vertica-kubernetes/issues/287) Default value for prometheus.createRBACProxy helm chart parameter is now true
### Deprecated
* [#287](https://github.com/vertica/vertica-kubernetes/issues/287) prometheus.createServiceMonitor helm chart parameter
### Fixed
* [#301](https://github.com/vertica/vertica-kubernetes/issues/301) Don't start the metric endpoint if metrics are disabled
* [#299](https://github.com/vertica/vertica-kubernetes/issues/299) Remove metrics for subclusters when VerticaDB is deleted
* [#292](https://github.com/vertica/vertica-kubernetes/issues/292) Extend the internal timeout for admintools to allow a slow revive to succeed
* [#291](https://github.com/vertica/vertica-kubernetes/issues/291) vdb-gen to handle db's that don't have authentication parms for communal storage

## 1.8.0 - 2022-11-18
### Added
* [#257](https://github.com/vertica/vertica-kubernetes/issues/257) Run the operator with readOnlyRootFilesystem set to true
* [#265](https://github.com/vertica/vertica-kubernetes/issues/265) Allow IAM authentication to communal storage
* [#274](https://github.com/vertica/vertica-kubernetes/issues/274) Allow catalog path to be specified in VerticaDB
* [#282](https://github.com/vertica/vertica-kubernetes/issues/282) Ability to skip package install during create db
### Changed
* [#254](https://github.com/vertica/vertica-kubernetes/issues/254) Moved to operator-sdk v1.23.0
* [#266](https://github.com/vertica/vertica-kubernetes/issues/266) Helm install with serviceAccountNameOverride will add roles/rolebindings
* [#268](https://github.com/vertica/vertica-kubernetes/issues/268) Default TLS cert for webhook is now generated internally rather than through cert-manager.
* [#273](https://github.com/vertica/vertica-kubernetes/issues/273) Allow webhook CA bundle to be taken from secret instead of helm chart parameter
### Fixed
* [#258](https://github.com/vertica/vertica-kubernetes/issues/258) Don't interrupt a slow Vertica startup
* [#259](https://github.com/vertica/vertica-kubernetes/issues/259) Hide communal credentials from the operator log
* [#262](https://github.com/vertica/vertica-kubernetes/issues/262) The vdbgen tool should be able to set ksafety, image and requestSize, when needed, to appropriate values taken from the database
* [#264](https://github.com/vertica/vertica-kubernetes/issues/264) Allow environment variables to flow down to Vertica process
* [#271](https://github.com/vertica/vertica-kubernetes/issues/271) Some pods may fail to run for a server upgrade change
* [#270](https://github.com/vertica/vertica-kubernetes/issues/270) Upgrade operator and server together may cause admintools to fail in the container due to lack of EULA acceptance
* [#275](https://github.com/vertica/vertica-kubernetes/issues/275) Allow local paths to share the same mount point
* [#280](https://github.com/vertica/vertica-kubernetes/issues/280) Operator pod readiness probe to wait for webhook
* [#283](https://github.com/vertica/vertica-kubernetes/issues/283) Improve the stability of the operator in big clusters

## 1.7.0 - 2022-08-26
### Added
* [#230](https://github.com/vertica/vertica-kubernetes/issues/230) Allow vstack and cores to be taken in the container
* [#232](https://github.com/vertica/vertica-kubernetes/issues/232) Ability to override the names of k8s objects in helm chart
* [#244](https://github.com/vertica/vertica-kubernetes/issues/244) Automated resize of the PV
* [#246](https://github.com/vertica/vertica-kubernetes/issues/246) Add feature gate to try out the experimental http server
* [#248](https://github.com/vertica/vertica-kubernetes/issues/248) Support for Java UDx's in the full Vertica image
* [#250](https://github.com/vertica/vertica-kubernetes/issues/250) Added e2e-udx testsuite to the CI
### Changed
* [#238](https://github.com/vertica/vertica-kubernetes/issues/238) Moved to operator-sdk v1.22.2
* [#239](https://github.com/vertica/vertica-kubernetes/issues/239) GitHub CI overhaul
* [#245](https://github.com/vertica/vertica-kubernetes/issues/245) Update server container base image to Ubuntu focal-20220801
### Fixed
* [#233](https://github.com/vertica/vertica-kubernetes/issues/233) Allow Vertica upgrade from 11.x to 12.x.
* [#234](https://github.com/vertica/vertica-kubernetes/issues/234) Update app.kubernetes.io/version in all objects when upgrading the operator
* [#234](https://github.com/vertica/vertica-kubernetes/issues/234) Prevent the need to restart the pods when the operator is upgraded
* [#234](https://github.com/vertica/vertica-kubernetes/issues/234) Allow operator upgrade from <= 1.1.0
* [#235](https://github.com/vertica/vertica-kubernetes/issues/235) Helm chart parm 'prometheus.createProxyRBAC' missed a required manifest
* [#247](https://github.com/vertica/vertica-kubernetes/issues/247) Add webhook rule to prevent use of restricted paths for local paths (data or depot)

## 1.6.0 - 2022-06-24
### Added
* [#224](https://github.com/vertica/vertica-kubernetes/issues/224) Allow spread communication encryption to be set in the VerticaDB CR
* [#227](https://github.com/vertica/vertica-kubernetes/issues/227) Warning message if v12.0.0 server and cgroups v2
### Changed
* [#218](https://github.com/vertica/vertica-kubernetes/issues/218) Use limits for pod when running admintools
* [#219](https://github.com/vertica/vertica-kubernetes/issues/219) Include zlib dev package in vertica-k8s image
* [#223](https://github.com/vertica/vertica-kubernetes/issues/223) Renamed Prometheus metrics exposed through the operator

## 1.5.0 - 2022-06-03
### Added
* [#206](https://github.com/vertica/vertica-kubernetes/issues/206) Push down more state into /etc/podinfo
* [#202](https://github.com/vertica/vertica-kubernetes/issues/202) Log events when shard/node ratio is not optimal
* [#199](https://github.com/vertica/vertica-kubernetes/issues/199) Add new prometheus metrics for the operator
* [#198](https://github.com/vertica/vertica-kubernetes/issues/198) Expose prometheus service for operator
* [#195](https://github.com/vertica/vertica-kubernetes/issues/195) Integrate autoscaler with VerticaDB
### Changed
* [#214](https://github.com/vertica/vertica-kubernetes/issues/214) Move to operator-sdk v1.21.0
### Fixed
* [#204](https://github.com/vertica/vertica-kubernetes/issues/204) Prevent requeueTime/upgradeRequeueTime from being negative in the webhook
* [#203](https://github.com/vertica/vertica-kubernetes/issues/203) Don't clear out installed/dbadded state for pods when they are pending
* [#202](https://github.com/vertica/vertica-kubernetes/issues/202) When creating the db, we should also choose the first primary subcluster
* [#201](https://github.com/vertica/vertica-kubernetes/issues/201) Improved handling for pending pods

## 1.4.0 - 2022-05-03
### Added
* [#189](https://github.com/vertica/vertica-kubernetes/issues/189) Additional subcluster options to better customize network load balancers
* [#170](https://github.com/vertica/vertica-kubernetes/issues/170) Helm parameters to allow deployment of the operator from private registries
* [#183](https://github.com/vertica/vertica-kubernetes/issues/183) Scale down will drain active connections before removing pod
* [#171](https://github.com/vertica/vertica-kubernetes/issues/171) Allow existing serviceaccount to be used
* [#168](https://github.com/vertica/vertica-kubernetes/issues/168) Added ability to configure RequeueAfter for upgrade reconciles. This delay can be specified through '.spec.upgradeRequeueTime' parameter. Prior to this, an online upgrade could wait upto 20 minutes before retrying.
### Changed
* [#187](https://github.com/vertica/vertica-kubernetes/issues/187) Change server container base image to ubuntu
* [#188](https://github.com/vertica/vertica-kubernetes/issues/188) Set the minimum TLS version of the webhook to TLS 1.3
* [#166](https://github.com/vertica/vertica-kubernetes/issues/166) Batch 'admintools -t db_add_node' for faster scale up
* [#165](https://github.com/vertica/vertica-kubernetes/issues/165) Move to operator-sdk v1.18.0
### Fixed
* [#191](https://github.com/vertica/vertica-kubernetes/issues/191) Allow database names with uppercase characters
* [#186](https://github.com/vertica/vertica-kubernetes/issues/186) Handle the scenario when restart is needed because the StatefulSets were deleted.  We ensure the necessary k8s objects are created before driving restart.
* [#178](https://github.com/vertica/vertica-kubernetes/issues/178) Avoid a second cluster restart after offline upgrade has completed successfully.
* [#176](https://github.com/vertica/vertica-kubernetes/issues/176) Upgrade path detection should allow skipping service packs


## 1.3.1 - 2022-03-02
### Fixed
* [#164](https://github.com/vertica/vertica-kubernetes/issues/164) Order the environment variables that were converted from annotations.  Prior to this fix, it was quite easy to get the statefulset controller to go into a repeated rolling upgrade.  The order ensures the statefulset doesn't appear to change between reconcile cycles.
* [#161](https://github.com/vertica/vertica-kubernetes/issues/161) Tolerate slashes being at the end of the communal endpoint url

## 1.3.0 - 2022-02-18
### Added
* [#146](https://github.com/vertica/vertica-kubernetes/issues/146) All annotations in the CR will be converted to environment variables in the containers.
* [#150](https://github.com/vertica/vertica-kubernetes/issues/150) Allow multiple subclusters to share the same Service object
* [#150](https://github.com/vertica/vertica-kubernetes/issues/150) Support for an online upgrade policy
* [#143](https://github.com/vertica/vertica-kubernetes/issues/143) New helm parameters to control the logging level and log path location for the operator pod
* [#81](https://github.com/vertica/vertica-kubernetes/issues/81) Support for RedHat OpenShift 4.8+
### Fixed
* [#151](https://github.com/vertica/vertica-kubernetes/issues/151) Subcluster names with hyphens were prevented from being the default subcluster.  This caused issues when creating the database and removal of subclusters.

## 1.2.0 - 2021-12-21
### Added
* [#87](https://github.com/vertica/vertica-kubernetes/issues/87) Support for Azure Blob Storage (azb://) as a communal endpoint.
* [#87](https://github.com/vertica/vertica-kubernetes/issues/87) Support for Google Cloud Storage (gs://) as a communal endpoint.
* [#87](https://github.com/vertica/vertica-kubernetes/issues/87) Support for HDFS (webhdfs://) as a communal endpoint.
* [#88](https://github.com/vertica/vertica-kubernetes/issues/88) Support for HDFS (swebhdfs://) as a communal endpoint.
* [#89](https://github.com/vertica/vertica-kubernetes/issues/89) Added the ability to specify custom volume mounts for use within the Vertica container.
* [#91](https://github.com/vertica/vertica-kubernetes/issues/91) Support for Kerberos authentication
* [#94](https://github.com/vertica/vertica-kubernetes/issues/94) Ability to specify custom ssh keys
* [#59](https://github.com/vertica/vertica-kubernetes/issues/59) New initPolicy called ScheduleOnly.  Use this policy when you have a vertica cluster running outside of Kubernetes and you want to provision new nodes to run inside Kubernetes.  Most of the automation is disabled when running in this mode.
### Removed
* [#88](https://github.com/vertica/vertica-kubernetes/issues/88) Removed support for Vertica 10.1.1.  The operator only supports Vertica 11.0.0 or higher.
### Fixed
* [#90](https://github.com/vertica/vertica-kubernetes/issues/90) Timing with scale down that can cause corruption in admintools.conf
* [#99](https://github.com/vertica/vertica-kubernetes/issues/99) The RollingUpdate strategy can kick-in after an image change causing pods in the cluster to restart again.
* [#101](https://github.com/vertica/vertica-kubernetes/issues/101) The image change can be marked complete before we finish the restart of the pods.
* [#113](https://github.com/vertica/vertica-kubernetes/issues/113) Restart of a cluster that has nodes in read-only state.  This is needed to run the operator with Vertica version 11.0.2 or newer.


## 1.1.0 - 2021-10-24
### Added
* [#42](https://github.com/vertica/vertica-kubernetes/issues/42) Added the ability to specify custom volumes for use within sidecars.
* [#57](https://github.com/vertica/vertica-kubernetes/issues/57) Added the ability to specify a custom CA file to authenticate s3 communal storage over https.  Previously https was only allowed for AWS.
* [#54](https://github.com/vertica/vertica-kubernetes/issues/54) Added the ability to mount additional certs in the Vertica container.  These certs can be specified through the new '.spec.certSecrets' parameter.
### Changed
* [#39](https://github.com/vertica/vertica-kubernetes/issues/39) Calls to update_vertica are removed.  The operator will modify admintools.conf for install/uninstall now.  This speeds up the time it takes to scale out.
* [#43](https://github.com/vertica/vertica-kubernetes/issues/43) Start the admission controller webhook as part of the operator pod.  This removes the helm chart and container for the webhook.  To order to use the webhook with the namespace scoped operator, the NamespaceDefaultLabelName feature gate must be enabled (on by default in 1.21+) or the namespace must have the label 'kubernetes.io/metadata.name=<nsName>' set.
* [#46](https://github.com/vertica/vertica-kubernetes/issues/46) Relax the dependency that the webhook requires cert-manager.  The default behaviour is to continue to depend on cert-manager.  But we now allow custom certs to be added through new helm chart parameters.
* [#51](https://github.com/vertica/vertica-kubernetes/issues/51) The operator automatically follows the upgrade procedure when the '.spec.image' is changed.  This removes the upgrade-vertica.sh script that previously handled this outside of the operator.
### Fixed
* [#47](https://github.com/vertica/vertica-kubernetes/issues/47) Communal storage on AWS s3.  The timeouts the operator had set were too low preventing a create DB from succeeding.
* [#58](https://github.com/vertica/vertica-kubernetes/issues/58) Increased the memory limit for the operator pod and made it configurable as a helm parameter.
* [#61](https://github.com/vertica/vertica-kubernetes/issues/61) Allow the AWS region to be specified in the CR.

## 1.0.0 - 2021-08-16

### Added
* Kubernetes operator (verticadb-operator) added to manage the lifecycle of a Vertica cluster
* helm chart (verticadb-operator) added to install the operator
* helm chart (verticadb-webhook) added to install the admission controller webhook
* Standalone tool (vdb-gen) that can be used to create a CR from a database for the purpose of migrating it to Kubernetes.

### Removed
* helm chart (vertica) was removed as it was made obsolete by the verticadb-operator

## 0.1.0 - 2021-04-30

### Added
* Helm chart (vertica) for statefulset deployment
