VERTICA_RPM?=vertica-x86_64.RHEL6.latest.rpm
BUILDER_OS_VERSION?=7.9.2009
BASE_OS_VERSION?=focal-20220316
VERTICA_IMG?=vertica-k8s
MINIMAL_VERTICA_IMG?=
NO_KEYS?=
VERTICA_VERSION?=$(shell rpm --nosignature -qp --queryformat '%{VERSION}-%{RELEASE}' packages/$(VERTICA_RPM))

all: docker-build-vertica

.PHONY: docker-build-vertica
docker-build-vertica: Dockerfile packages/package-checksum-patcher.py
	docker pull ubuntu:focal ## make sure we use the latest focal image
	docker build \
		--platform linux/amd64 \
		-f Dockerfile \
		--label minimal=${MINIMAL_VERTICA_IMG} \
		--label os-version=${BASE_OS_VERSION} \
		--label vertica-version=${VERTICA_VERSION} \
		--build-arg MINIMAL=${MINIMAL_VERTICA_IMG} \
		--build-arg VERTICA_RPM=${VERTICA_RPM} \
		--build-arg NO_KEYS=${NO_KEYS} \
		--build-arg BASE_OS_VERSION=${BASE_OS_VERSION} \
		--build-arg BUILDER_OS_VERSION=${BUILDER_OS_VERSION} \
		-t ${VERTICA_IMG} .
