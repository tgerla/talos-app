module github.com/talos-systems/ui

go 1.13

require (
	github.com/Equanox/gotron v0.2.23
	github.com/otiai10/curr v0.0.0-20190513014714-f5a3d24e5776 // indirect
	github.com/pkg/errors v0.8.1
	github.com/talos-systems/talos v0.1.0
	k8s.io/apimachinery v0.0.0
	k8s.io/client-go v0.0.0
)

replace (
	github.com/docker/distribution v2.7.1+incompatible => github.com/docker/distribution v2.7.1-0.20190205005809-0d3efadf0154+incompatible
	github.com/opencontainers/runtime-spec v1.0.1 => github.com/opencontainers/runtime-spec v0.1.2-0.20180301181910-fa4b36aa9c99
	k8s.io/api => k8s.io/api v0.0.0-20190822053644-6185379c914a
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190822063004-0670dc4fec4e
	k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20190820074809-31b1e1ea64dc
	k8s.io/apiserver => k8s.io/apiserver v0.0.0-20190822060508-785eacbd19ae
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190822063658-442a64f3fed7
	k8s.io/client-go => k8s.io/client-go v0.0.0-20190822054823-0a74433fb222
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190822065847-2058b41dfbb6
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.0.0-20190822065536-566e5fc137f7
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190820100630-060a3d12ce80
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190822055535-1f6a258f5d89
	k8s.io/cri-api => k8s.io/cri-api v0.0.0-20190820110325-95eec93e2395
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.0.0-20190822070154-f51cd605b3ee
	k8s.io/klog => k8s.io/klog v0.3.1
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190822061015-a4f93a8219ed
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.0.0-20190822065235-826221481525
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.0.0-20190822064323-7e0495d8a3ff
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.0.0-20190822064931-4470440ed041
	k8s.io/kubectl => k8s.io/kubectl v0.0.0-20190822071625-14af4a63a1e1
	k8s.io/kubelet => k8s.io/kubelet v0.0.0-20190822064626-fa8f3d935631
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.0.0-20190822070624-3a30a18bba71
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190822063337-6c03eb8600ee
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.0.0-20190822061642-ab22eab63834
	k8s.io/utils => k8s.io/utils v0.0.0-20190801114015-581e00157fb1
)

replace github.com/talos-systems/talos => ../talos
