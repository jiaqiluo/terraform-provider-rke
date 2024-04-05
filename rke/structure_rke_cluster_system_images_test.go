package rke

import (
	"reflect"
	"testing"

	rancher "github.com/rancher/rke/types"
)

var (
	testRKEClusterSystemImagesConf      rancher.RKESystemImages
	testRKEClusterSystemImagesInterface []interface{}
)

func init() {
	testRKEClusterSystemImagesConf = rancher.RKESystemImages{
		Etcd:                      "etcd",
		Alpine:                    "alpine",
		NginxProxy:                "nginx_proxy",
		CertDownloader:            "cert_downloader",
		KubernetesServicesSidecar: "kubernetes_services_sidecar",
		KubeDNS:                   "kube_dns",
		DNSmasq:                   "dnsmasq",
		KubeDNSSidecar:            "kube_dns_sidecar",
		KubeDNSAutoscaler:         "kube_dns_autoscaler",
		CoreDNS:                   "coredns",
		CoreDNSAutoscaler:         "coredns_autoscaler",
		Nodelocal:                 "nodelocal",
		Kubernetes:                "kubernetes",
		Flannel:                   "flannel",
		FlannelCNI:                "flannel_cni",
		CalicoNode:                "calico_node",
		CalicoCNI:                 "calico_cni",
		CalicoControllers:         "calico_controllers",
		CalicoCtl:                 "calico_ctl",
		CalicoFlexVol:             "calico_flex_vol",
		CanalNode:                 "canal_node",
		CanalCNI:                  "canal_cni",
		CanalControllers:          "canal_controllers",
		CanalFlannel:              "canal_flannel",
		CanalFlexVol:              "canal_flex_vol",
		WeaveNode:                 "weave_node",
		WeaveCNI:                  "weave_cni",
		PodInfraContainer:         "pod_infra_container",
		Ingress:                   "ingress",
		IngressBackend:            "ingress_backend",
		IngressWebhook:            "ingress_webhook",
		MetricsServer:             "metrics_server",
		WindowsPodInfraContainer:  "windows_pod_infra_container",
		AciCniDeployContainer:     "aci_cni_deploy_container",
		AciHostContainer:          "aci_host_container",
		AciOpflexContainer:        "aci_opflex_container",
		AciMcastContainer:         "aci_mcast_container",
		AciOpenvSwitchContainer:   "aci_ovs_container",
		AciControllerContainer:    "aci_controller_container",
		AciGbpServerContainer:     "aci_gbp_server_container",
		AciOpflexServerContainer:  "aci_opflex_server_container",
	}
	testRKEClusterSystemImagesInterface = []interface{}{
		map[string]interface{}{
			"etcd":                        "etcd",
			"alpine":                      "alpine",
			"nginx_proxy":                 "nginx_proxy",
			"cert_downloader":             "cert_downloader",
			"kubernetes_services_sidecar": "kubernetes_services_sidecar",
			"kube_dns":                    "kube_dns",
			"dnsmasq":                     "dnsmasq",
			"kube_dns_sidecar":            "kube_dns_sidecar",
			"kube_dns_autoscaler":         "kube_dns_autoscaler",
			"coredns":                     "coredns",
			"coredns_autoscaler":          "coredns_autoscaler",
			"nodelocal":                   "nodelocal",
			"kubernetes":                  "kubernetes",
			"flannel":                     "flannel",
			"flannel_cni":                 "flannel_cni",
			"calico_node":                 "calico_node",
			"calico_cni":                  "calico_cni",
			"calico_controllers":          "calico_controllers",
			"calico_ctl":                  "calico_ctl",
			"calico_flex_vol":             "calico_flex_vol",
			"canal_node":                  "canal_node",
			"canal_cni":                   "canal_cni",
			"canal_controllers":           "canal_controllers",
			"canal_flannel":               "canal_flannel",
			"canal_flex_vol":              "canal_flex_vol",
			"weave_node":                  "weave_node",
			"weave_cni":                   "weave_cni",
			"pod_infra_container":         "pod_infra_container",
			"ingress":                     "ingress",
			"ingress_backend":             "ingress_backend",
			"ingress_webhook":             "ingress_webhook",
			"metrics_server":              "metrics_server",
			"windows_pod_infra_container": "windows_pod_infra_container",
			"aci_cni_deploy_container":    "aci_cni_deploy_container",
			"aci_host_container":          "aci_host_container",
			"aci_opflex_container":        "aci_opflex_container",
			"aci_mcast_container":         "aci_mcast_container",
			"aci_ovs_container":           "aci_ovs_container",
			"aci_controller_container":    "aci_controller_container",
			"aci_gbp_server_container":    "aci_gbp_server_container",
			"aci_opflex_server_container": "aci_opflex_server_container",
		},
	}
}

func TestFlattenRKEClusterSystemImages(t *testing.T) {

	cases := []struct {
		Input          rancher.RKESystemImages
		ExpectedOutput []interface{}
	}{
		{
			testRKEClusterSystemImagesConf,
			testRKEClusterSystemImagesInterface,
		},
	}

	for _, tc := range cases {
		output := flattenRKEClusterSystemImages(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from flattener.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}

func TestExpandRKEClusterSystemImages(t *testing.T) {

	cases := []struct {
		Input          []interface{}
		ExpectedOutput rancher.RKESystemImages
	}{
		{
			testRKEClusterSystemImagesInterface,
			testRKEClusterSystemImagesConf,
		},
	}

	for _, tc := range cases {
		output := expandRKEClusterSystemImages(tc.Input)
		if !reflect.DeepEqual(output, tc.ExpectedOutput) {
			t.Fatalf("Unexpected output from expander.\nExpected: %#v\nGiven:    %#v",
				tc.ExpectedOutput, output)
		}
	}
}
