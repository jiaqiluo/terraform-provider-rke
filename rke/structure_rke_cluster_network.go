package rke

import (
	rancher "github.com/rancher/rke/types"
)

// Flatteners

func flattenRKEClusterNetworkCalico(in *rancher.CalicoNetworkProvider) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}

	if len(in.CloudProvider) > 0 {
		obj["cloud_provider"] = in.CloudProvider
	}

	return []interface{}{obj}
}

func flattenRKEClusterNetworkCanal(in *rancher.CanalNetworkProvider) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}

	if len(in.FlannelNetworkProvider.Iface) > 0 {
		obj["iface"] = in.FlannelNetworkProvider.Iface
	}

	return []interface{}{obj}
}

func flattenRKEClusterNetworkFlannel(in *rancher.FlannelNetworkProvider) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}

	if len(in.Iface) > 0 {
		obj["iface"] = in.Iface
	}

	return []interface{}{obj}
}

func flattenRKEClusterNetworkWeave(in *rancher.WeaveNetworkProvider) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}

	if len(in.Password) > 0 {
		obj["password"] = in.Password
	}

	return []interface{}{obj}
}

func flattenRKEClusterNetworkAci(in *rancher.AciNetworkProvider) []interface{} {
	obj := make(map[string]interface{})
	if in == nil {
		return []interface{}{}
	}
	if len(in.SystemIdentifier) > 0 {
		obj["system_id"] = in.SystemIdentifier
	}
	if in.ApicHosts != nil {
		obj["apic_hosts"] = toArrayInterface(in.ApicHosts)
	}
	if len(in.Token) > 0 {
		obj["token"] = in.Token
	}
	if len(in.ApicUserName) > 0 {
		obj["apic_user_name"] = in.ApicUserName
	}
	if len(in.ApicUserKey) > 0 {
		obj["apic_user_key"] = in.ApicUserKey
	}
	if len(in.ApicUserCrt) > 0 {
		obj["apic_user_crt"] = in.ApicUserCrt
	}
	if len(in.ApicRefreshTime) > 0 {
		obj["apic_refresh_time"] = in.ApicRefreshTime
	}
	if len(in.VmmDomain) > 0 {
		obj["vmm_domain"] = in.VmmDomain
	}
	if len(in.VmmController) > 0 {
		obj["vmm_controller"] = in.VmmController
	}
	if len(in.EncapType) > 0 {
		obj["encap_type"] = in.EncapType
	}
	if len(in.NodeSubnet) > 0 {
		obj["node_subnet"] = in.NodeSubnet
	}
	if len(in.McastRangeStart) > 0 {
		obj["mcast_range_start"] = in.McastRangeStart
	}
	if len(in.McastRangeEnd) > 0 {
		obj["mcast_range_end"] = in.McastRangeEnd
	}
	if len(in.AEP) > 0 {
		obj["aep"] = in.AEP
	}
	if len(in.VRFName) > 0 {
		obj["vrf_name"] = in.VRFName
	}
	if len(in.VRFTenant) > 0 {
		obj["vrf_tenant"] = in.VRFTenant
	}
	if len(in.L3Out) > 0 {
		obj["l3out"] = in.L3Out
	}
	if in.L3OutExternalNetworks != nil {
		obj["l3out_external_networks"] = toArrayInterface(in.L3OutExternalNetworks)
	}
	if len(in.DynamicExternalSubnet) > 0 {
		obj["extern_dynamic"] = in.DynamicExternalSubnet
	}
	if len(in.StaticExternalSubnet) > 0 {
		obj["extern_static"] = in.StaticExternalSubnet
	}
	if len(in.ServiceGraphSubnet) > 0 {
		obj["node_svc_subnet"] = in.ServiceGraphSubnet
	}
	if len(in.KubeAPIVlan) > 0 {
		obj["kube_api_vlan"] = in.KubeAPIVlan
	}
	if len(in.ServiceVlan) > 0 {
		obj["service_vlan"] = in.ServiceVlan
	}
	if len(in.InfraVlan) > 0 {
		obj["infra_vlan"] = in.InfraVlan
	}
	if len(in.Tenant) > 0 {
		obj["tenant"] = in.Tenant
	}
	if len(in.OVSMemoryLimit) > 0 {
		obj["ovs_memory_limit"] = in.OVSMemoryLimit
	}
	if len(in.OVSMemoryRequest) > 0 {
		obj["ovs_memory_request"] = in.OVSMemoryRequest
	}
	if len(in.ImagePullPolicy) > 0 {
		obj["image_pull_policy"] = in.ImagePullPolicy
	}
	if len(in.ImagePullSecret) > 0 {
		obj["image_pull_secret"] = in.ImagePullSecret
	}
	if len(in.ServiceMonitorInterval) > 0 {
		obj["service_monitor_interval"] = in.ServiceMonitorInterval
	}
	if len(in.PBRTrackingNonSnat) > 0 {
		obj["pbr_tracking_non_snat"] = in.PBRTrackingNonSnat
	}
	if len(in.InstallIstio) > 0 {
		obj["install_istio"] = in.InstallIstio
	}
	if len(in.IstioProfile) > 0 {
		obj["istio_profile"] = in.IstioProfile
	}
	if len(in.DropLogEnable) > 0 {
		obj["drop_log_enable"] = in.DropLogEnable
	}
	if len(in.ControllerLogLevel) > 0 {
		obj["controller_log_level"] = in.ControllerLogLevel
	}
	if len(in.HostAgentLogLevel) > 0 {
		obj["host_agent_log_level"] = in.HostAgentLogLevel
	}
	if len(in.OpflexAgentLogLevel) > 0 {
		obj["opflex_log_level"] = in.OpflexAgentLogLevel
	}
	if len(in.UseAciCniPriorityClass) > 0 {
		obj["use_aci_cni_priority_class"] = in.UseAciCniPriorityClass
	}
	if len(in.NoPriorityClass) > 0 {
		obj["no_priority_class"] = in.NoPriorityClass
	}
	if len(in.MaxNodesSvcGraph) > 0 {
		obj["max_nodes_svc_graph"] = in.MaxNodesSvcGraph
	}
	if len(in.SnatContractScope) > 0 {
		obj["snat_contract_scope"] = in.SnatContractScope
	}
	if len(in.PodSubnetChunkSize) > 0 {
		obj["pod_subnet_chunk_size"] = in.PodSubnetChunkSize
	}
	if len(in.EnableEndpointSlice) > 0 {
		obj["enable_endpoint_slice"] = in.EnableEndpointSlice
	}
	if len(in.SnatNamespace) > 0 {
		obj["snat_namespace"] = in.SnatNamespace
	}
	if len(in.EpRegistry) > 0 {
		obj["ep_registry"] = in.EpRegistry
	}
	if len(in.OpflexMode) > 0 {
		obj["opflex_mode"] = in.OpflexMode
	}
	if len(in.SnatPortRangeStart) > 0 {
		obj["snat_port_range_start"] = in.SnatPortRangeStart
	}
	if len(in.SnatPortRangeEnd) > 0 {
		obj["snat_port_range_end"] = in.SnatPortRangeEnd
	}
	if len(in.SnatPortsPerNode) > 0 {
		obj["snat_ports_per_node"] = in.SnatPortsPerNode
	}
	if len(in.OpflexClientSSL) > 0 {
		obj["opflex_client_ssl"] = in.OpflexClientSSL
	}
	if len(in.UsePrivilegedContainer) > 0 {
		obj["use_privileged_container"] = in.UsePrivilegedContainer
	}
	if len(in.UseHostNetnsVolume) > 0 {
		obj["use_host_netns_volume"] = in.UseHostNetnsVolume
	}
	if len(in.UseOpflexServerVolume) > 0 {
		obj["use_opflex_server_volume"] = in.UseOpflexServerVolume
	}
	if len(in.SubnetDomainName) > 0 {
		obj["subnet_domain_name"] = in.SubnetDomainName
	}
	if in.KafkaBrokers != nil {
		obj["kafka_brokers"] = toArrayInterface(in.KafkaBrokers)
	}
	if len(in.KafkaClientCrt) > 0 {
		obj["kafka_client_crt"] = in.KafkaClientCrt
	}
	if len(in.KafkaClientKey) > 0 {
		obj["kafka_client_key"] = in.KafkaClientKey
	}
	if len(in.CApic) > 0 {
		obj["capic"] = in.CApic
	}
	if len(in.UseAciAnywhereCRD) > 0 {
		obj["use_aci_anywhere_crd"] = in.UseAciAnywhereCRD
	}
	if len(in.OverlayVRFName) > 0 {
		obj["overlay_vrf_name"] = in.OverlayVRFName
	}
	if len(in.GbpPodSubnet) > 0 {
		obj["gbp_pod_subnet"] = in.GbpPodSubnet
	}
	if len(in.RunGbpContainer) > 0 {
		obj["run_gbp_container"] = in.RunGbpContainer
	}
	if len(in.RunOpflexServerContainer) > 0 {
		obj["run_opflex_server_container"] = in.RunOpflexServerContainer
	}
	if len(in.OpflexServerPort) > 0 {
		obj["opflex_server_port"] = in.OpflexServerPort
	}
	if len(in.DurationWaitForNetwork) > 0 {
		obj["duration_wait_for_network"] = in.DurationWaitForNetwork
	}
	if len(in.DisableWaitForNetwork) > 0 {
		obj["disable_wait_for_network"] = in.DisableWaitForNetwork
	}
	if len(in.ApicSubscriptionDelay) > 0 {
		obj["apic_subscription_delay"] = in.ApicSubscriptionDelay
	}
	if len(in.ApicRefreshTickerAdjust) > 0 {
		obj["apic_refresh_ticker_adjust"] = in.ApicRefreshTickerAdjust
	}
	if len(in.DisablePeriodicSnatGlobalInfoSync) > 0 {
		obj["disable_periodic_snat_global_info_sync"] = in.DisablePeriodicSnatGlobalInfoSync
	}
	if len(in.OpflexDeviceDeleteTimeout) > 0 {
		obj["opflex_device_delete_timeout"] = in.OpflexDeviceDeleteTimeout
	}
	if len(in.MTUHeadRoom) > 0 {
		obj["mtu_head_room"] = in.MTUHeadRoom
	}
	if len(in.NodePodIfEnable) > 0 {
		obj["node_pod_if_enable"] = in.NodePodIfEnable
	}
	if len(in.SriovEnable) > 0 {
		obj["sriov_enable"] = in.SriovEnable
	}
	if len(in.MultusDisable) > 0 {
		obj["multus_disable"] = in.MultusDisable
	}
	if len(in.UseClusterRole) > 0 {
		obj["use_cluster_role"] = in.UseClusterRole
	}
	if len(in.NoWaitForServiceEpReadiness) > 0 {
		obj["no_wait_for_service_ep_readiness"] = in.NoWaitForServiceEpReadiness
	}
	if len(in.AddExternalSubnetsToRdconfig) > 0 {
		obj["add_external_subnets_to_rdconfig"] = in.AddExternalSubnetsToRdconfig
	}
	if len(in.ServiceGraphEndpointAddDelay) > 0 {
		obj["service_graph_endpoint_add_delay"] = in.ServiceGraphEndpointAddDelay
	}
	if len(in.ServiceGraphEndpointAddServices) > 0 {
		obj["service_graph_endpoint_add_services"] = toArrayMapInterface(in.ServiceGraphEndpointAddServices)
	}
	if len(in.HppOptimization) > 0 {
		obj["hpp_optimization"] = in.HppOptimization
	}
	if len(in.SleepTimeSnatGlobalInfoSync) > 0 {
		obj["sleep_time_snat_global_info_sync"] = in.SleepTimeSnatGlobalInfoSync
	}
	if len(in.OpflexAgentOpflexAsyncjsonEnabled) > 0 {
		obj["opflex_agent_opflex_asyncjson_enabled"] = in.OpflexAgentOpflexAsyncjsonEnabled
	}
	if len(in.OpflexAgentOvsAsyncjsonEnabled) > 0 {
		obj["opflex_agent_ovs_asyncjson_enabled"] = in.OpflexAgentOvsAsyncjsonEnabled
	}
	if len(in.OpflexAgentPolicyRetryDelayTimer) > 0 {
		obj["opflex_agent_policy_retry_delay_timer"] = in.OpflexAgentPolicyRetryDelayTimer
	}
	if len(in.AciMultipod) > 0 {
		obj["aci_multipod"] = in.AciMultipod
	}
	if len(in.OpflexDeviceReconnectWaitTimeout) > 0 {
		obj["opflex_device_reconnect_wait_timeout"] = in.OpflexDeviceReconnectWaitTimeout
	}
	if len(in.AciMultipodUbuntu) > 0 {
		obj["aci_multipod_ubuntu"] = in.AciMultipodUbuntu
	}
	if len(in.DhcpRenewMaxRetryCount) > 0 {
		obj["dhcp_renew_max_retry_count"] = in.DhcpRenewMaxRetryCount
	}
	if len(in.DhcpDelay) > 0 {
		obj["dhcp_delay"] = in.DhcpDelay
	}
	if len(in.UseSystemNodePriorityClass) > 0 {
		obj["use_system_node_priority_class"] = in.UseSystemNodePriorityClass
	}
	if len(in.AciContainersControllerMemoryRequest) > 0 {
		obj["aci_containers_controller_memory_request"] = in.AciContainersControllerMemoryRequest
	}
	if len(in.AciContainersControllerMemoryLimit) > 0 {
		obj["aci_containers_controller_memory_limit"] = in.AciContainersControllerMemoryLimit
	}
	if len(in.AciContainersHostMemoryRequest) > 0 {
		obj["aci_containers_host_memory_request"] = in.AciContainersHostMemoryRequest
	}
	if len(in.AciContainersHostMemoryLimit) > 0 {
		obj["aci_containers_host_memory_limit"] = in.AciContainersHostMemoryLimit
	}
	if len(in.McastDaemonMemoryRequest) > 0 {
		obj["mcast_daemon_memory_request"] = in.McastDaemonMemoryRequest
	}
	if len(in.McastDaemonMemoryLimit) > 0 {
		obj["mcast_daemon_memory_limit"] = in.McastDaemonMemoryLimit
	}
	if len(in.OpflexAgentMemoryRequest) > 0 {
		obj["opflex_agent_memory_request"] = in.OpflexAgentMemoryRequest
	}
	if len(in.OpflexAgentMemoryLimit) > 0 {
		obj["opflex_agent_memory_limit"] = in.OpflexAgentMemoryLimit
	}
	if len(in.AciContainersMemoryRequest) > 0 {
		obj["aci_containers_memory_request"] = in.AciContainersMemoryRequest
	}
	if len(in.AciContainersMemoryLimit) > 0 {
		obj["aci_containers_memory_limit"] = in.AciContainersMemoryLimit
	}
	if len(in.OpflexAgentStatistics) > 0 {
		obj["opflex_agent_statistics"] = in.OpflexAgentStatistics
	}
	if len(in.AddExternalContractToDefaultEpg) > 0 {
		obj["add_external_contract_to_default_epg"] = in.AddExternalContractToDefaultEpg
	}
	if len(in.EnableOpflexAgentReconnect) > 0 {
		obj["enable_opflex_agent_reconnect"] = in.EnableOpflexAgentReconnect
	}
	if len(in.OpflexOpensslCompat) > 0 {
		obj["opflex_openssl_compat"] = in.OpflexOpensslCompat
	}
	if in.NodeSnatRedirectExclude != nil {
		obj["opflex_openssl_compat"] = toArrayMapInterface(in.NodeSnatRedirectExclude)
	}
	if len(in.TolerationSeconds) > 0 {
		obj["toleration_seconds"] = in.TolerationSeconds
	}
	if len(in.DisableHppRendering) > 0 {
		obj["disable_hpp_rendering"] = in.DisableHppRendering
	}
	if len(in.ApicConnectionRetryLimit) > 0 {
		obj["apic_connection_retry_limit"] = in.ApicConnectionRetryLimit
	}
	if len(in.TaintNotReadyNode) > 0 {
		obj["taint_not_ready_node"] = in.TaintNotReadyNode
	}
	if len(in.DropLogDisableEvents) > 0 {
		obj["drop_log_disable_events"] = in.DropLogDisableEvents
	}
	return []interface{}{obj}
}

func flattenRKEClusterNetwork(in rancher.NetworkConfig) []interface{} {
	obj := make(map[string]interface{})

	if in.CalicoNetworkProvider != nil {
		obj["calico_network_provider"] = flattenRKEClusterNetworkCalico(in.CalicoNetworkProvider)
	}

	if in.CanalNetworkProvider != nil {
		obj["canal_network_provider"] = flattenRKEClusterNetworkCanal(in.CanalNetworkProvider)
	}

	if in.FlannelNetworkProvider != nil {
		obj["flannel_network_provider"] = flattenRKEClusterNetworkFlannel(in.FlannelNetworkProvider)
	}

	if in.WeaveNetworkProvider != nil {
		obj["weave_network_provider"] = flattenRKEClusterNetworkWeave(in.WeaveNetworkProvider)
	}

	if in.AciNetworkProvider != nil {
		obj["aci_network_provider"] = flattenRKEClusterNetworkAci(in.AciNetworkProvider)
	}

	if in.MTU > 0 {
		obj["mtu"] = in.MTU
	}

	if len(in.Options) > 0 {
		obj["options"] = toMapInterface(in.Options)
	}

	if len(in.Plugin) > 0 {
		obj["plugin"] = in.Plugin
	}

	return []interface{}{obj}
}

// Expanders

func expandRKEClusterNetworkCalico(p []interface{}) *rancher.CalicoNetworkProvider {
	obj := &rancher.CalicoNetworkProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["cloud_provider"].(string); ok && len(v) > 0 {
		obj.CloudProvider = v
	}

	return obj
}

func expandRKEClusterNetworkCanal(p []interface{}) *rancher.CanalNetworkProvider {
	obj := &rancher.CanalNetworkProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["iface"].(string); ok && len(v) > 0 {
		obj.FlannelNetworkProvider.Iface = v
	}

	return obj
}

func expandRKEClusterNetworkFlannel(p []interface{}) *rancher.FlannelNetworkProvider {
	obj := &rancher.FlannelNetworkProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["iface"].(string); ok && len(v) > 0 {
		obj.Iface = v
	}

	return obj
}

func expandRKEClusterNetworkWeave(p []interface{}) *rancher.WeaveNetworkProvider {
	obj := &rancher.WeaveNetworkProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["password"].(string); ok && len(v) > 0 {
		obj.Password = v
	}

	return obj
}

func expandRKEClusterNetworkAci(p []interface{}) *rancher.AciNetworkProvider {
	obj := &rancher.AciNetworkProvider{}
	if len(p) == 0 || p[0] == nil {
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["system_id"].(string); ok && len(v) > 0 {
		obj.SystemIdentifier = v
	}
	if v, ok := in["apic_hosts"].([]interface{}); ok && len(v) > 0 {
		obj.ApicHosts = toArrayString(v)
	}
	if v, ok := in["token"].(string); ok && len(v) > 0 {
		obj.Token = v
	}
	if v, ok := in["apic_user_name"].(string); ok && len(v) > 0 {
		obj.ApicUserName = v
	}
	if v, ok := in["apic_user_key"].(string); ok && len(v) > 0 {
		obj.ApicUserKey = v
	}
	if v, ok := in["apic_user_crt"].(string); ok && len(v) > 0 {
		obj.ApicUserCrt = v
	}
	if v, ok := in["encap_type"].(string); ok && len(v) > 0 {
		obj.EncapType = v
	}
	if v, ok := in["mcast_range_start"].(string); ok && len(v) > 0 {
		obj.McastRangeStart = v
	}
	if v, ok := in["mcast_range_end"].(string); ok && len(v) > 0 {
		obj.McastRangeEnd = v
	}
	if v, ok := in["aep"].(string); ok && len(v) > 0 {
		obj.AEP = v
	}
	if v, ok := in["vrf_name"].(string); ok && len(v) > 0 {
		obj.VRFName = v
	}
	if v, ok := in["vrf_tenant"].(string); ok && len(v) > 0 {
		obj.VRFTenant = v
	}
	if v, ok := in["l3out"].(string); ok && len(v) > 0 {
		obj.L3Out = v
	}
	if v, ok := in["node_subnet"].(string); ok && len(v) > 0 {
		obj.NodeSubnet = v
	}
	if v, ok := in["l3out_external_networks"].([]interface{}); ok && len(v) > 0 {
		obj.L3OutExternalNetworks = toArrayString(v)
	}
	if v, ok := in["extern_dynamic"].(string); ok && len(v) > 0 {
		obj.DynamicExternalSubnet = v
	}
	if v, ok := in["extern_static"].(string); ok && len(v) > 0 {
		obj.StaticExternalSubnet = v
	}
	if v, ok := in["node_svc_subnet"].(string); ok && len(v) > 0 {
		obj.ServiceGraphSubnet = v
	}
	if v, ok := in["kube_api_vlan"].(string); ok && len(v) > 0 {
		obj.KubeAPIVlan = v
	}
	if v, ok := in["service_vlan"].(string); ok && len(v) > 0 {
		obj.ServiceVlan = v
	}
	if v, ok := in["infra_vlan"].(string); ok && len(v) > 0 {
		obj.InfraVlan = v
	}
	if v, ok := in["snat_port_range_start"].(string); ok && len(v) > 0 {
		obj.SnatPortRangeStart = v
	}
	if v, ok := in["snat_port_range_end"].(string); ok && len(v) > 0 {
		obj.SnatPortRangeEnd = v
	}
	if v, ok := in["snat_ports_per_node"].(string); ok && len(v) > 0 {
		obj.SnatPortsPerNode = v
	}

	return obj
}

func expandRKEClusterNetwork(p []interface{}) rancher.NetworkConfig {
	obj := rancher.NetworkConfig{}
	if len(p) == 0 || p[0] == nil {
		obj.Plugin = rkeClusterNetworkPluginDefault
		return obj
	}
	in := p[0].(map[string]interface{})

	if v, ok := in["calico_network_provider"].([]interface{}); ok && len(v) > 0 {
		obj.CalicoNetworkProvider = expandRKEClusterNetworkCalico(v)
	}

	if v, ok := in["canal_network_provider"].([]interface{}); ok && len(v) > 0 {
		obj.CanalNetworkProvider = expandRKEClusterNetworkCanal(v)
	}

	if v, ok := in["flannel_network_provider"].([]interface{}); ok && len(v) > 0 {
		obj.FlannelNetworkProvider = expandRKEClusterNetworkFlannel(v)
	}

	if v, ok := in["weave_network_provider"].([]interface{}); ok && len(v) > 0 {
		obj.WeaveNetworkProvider = expandRKEClusterNetworkWeave(v)
	}

	if v, ok := in["aci_network_provider"].([]interface{}); ok && len(v) > 0 {
		obj.AciNetworkProvider = expandRKEClusterNetworkAci(v)
	}

	if v, ok := in["mtu"].(int); ok && v > 0 {
		obj.MTU = v
	}

	if v, ok := in["options"].(map[string]interface{}); ok && len(v) > 0 {
		obj.Options = toMapString(v)
	}

	if v, ok := in["plugin"].(string); ok && len(v) > 0 {
		obj.Plugin = v
	}

	return obj
}
