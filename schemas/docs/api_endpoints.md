# H3Net Protocol API Endpoints

Generated on: Thu Mar 20 23:47:37 UTC 2025

| Category | Endpoint | GET | PUT | POST | DELETE |
| -------- | -------- | --- | --- | ---- | ------ |
| acl | /config/acls | getAllAcls | ❌ | createAcl | ❌ |
| | /config/acls/device/{deviceId} | getDeviceAcls | ❌ | createDeviceAcl | ❌ |
| | /config/acls/{aclId} | getAclConfig | updateAclConfig | ❌ | ❌ |
| | | | | | |
| config | /config/acls/{deviceId}/{aclName} | getACL | updateACL | ❌ | deleteACL |
| | /config/acls/{deviceId}/{aclName}/rules | ❌ | ❌ | addACLRule | ❌ |
| | /config/acls/{deviceId}/{aclName}/rules/{ruleId} | ❌ | updateACLRule | ❌ | deleteACLRule |
| | /config/devices | getAllDeviceConfigs | ❌ | createOrUpdateDevice | ❌ |
| | /config/devices/info/bulk | ❌ | ❌ | requestBulkDeviceInfo | ❌ |
| | /config/devices/{deviceId} | getDeviceConfig | updateDeviceConfig | ❌ | ❌ |
| | /config/devices/{deviceId}/info | getDeviceInfo | ❌ | requestDeviceInfo | ❌ |
| | /config/devices/{deviceId}/interfaces | getDeviceInterfaces | ❌ | ❌ | ❌ |
| | /config/devices/{deviceId}/interfaces/{interfaceId} | ❌ | updateInterfaceConfig | ❌ | ❌ |
| | /config/devices/{deviceId}/system | getSystemSettings | updateSystemSettings | ❌ | ❌ |
| | /config/routing | getRoutingConfig | putRoutingConfig | ❌ | ❌ |
| | /config/switch | getSwitchConfig | ❌ | updateSwitchConfig | ❌ |
| | /config/vlan | getVlanConfig | updateVlanConfig | createVlanConfig | ❌ |
| | | | | | |
| core | /devices | getDevices | ❌ | registerDevice | ❌ |
| | /devices/{deviceId} | getDevice | ❌ | ❌ | ❌ |
| | /health | core.checkHealth | ❌ | ❌ | ❌ |
| | /hello | ❌ | ❌ | core.sendHello | ❌ |
| | | | | | |
| discovery | /discovery/capabilities | getDeviceCapabilities | ❌ | queryDeviceCapabilities | ❌ |
| | /discovery/capabilities/bulk | ❌ | ❌ | queryBulkCapabilities | ❌ |
| | /discovery/network | getNetworkDiscoveryStatus | ❌ | performNetworkDiscovery | ❌ |
| | /discovery/network/neighbors/{deviceId} | getDeviceNeighbors | ❌ | ❌ | ❌ |
| | /discovery/network/options | getNetworkDiscoveryOptions | updateNetworkDiscoveryOptions | ❌ | ❌ |
| | /discovery/network/topology | getNetworkTopology | ❌ | ❌ | ❌ |
| | | | | | |
| events | /events | ❌ | ❌ | events.sendEventNotification | ❌ |
| | | | | | |
| operations | /operations/state | getOperationalState | ❌ | queryDeviceState | ❌ |
| | /operations/state/{deviceId} | getDeviceOperationalState | ❌ | querySpecificDeviceState | ❌ |
| | /operations/state/{deviceId}/interfaces | getDeviceInterfaceStates | ❌ | ❌ | ❌ |
| | /operations/state/{deviceId}/reboot | ❌ | ❌ | rebootDevice | ❌ |
| | /operations/state/{deviceId}/tables | getTableEntries | ❌ | ❌ | ❌ |
| | | | | | |
| pipeline | /config/pipeline | getPipelineConfig | ❌ | updatePipelineConfig | ❌ |
| | /config/pipeline/flows | getFlowEntries | ❌ | updateFlowEntries | ❌ |
| | /config/pipeline/tables | getMatchActionTables | ❌ | createMatchActionTable | ❌ |
| | /pipeline/flows | getFlows | ❌ | createFlow | ❌ |
| | /tables | getPipelineTables | ❌ | createTable | ❌ |
| | /tables/{tableId} | getPipelineTable | updateTable | ❌ | ❌ |
| | /tables/{tableId}/pipeline | ❌ | ❌ | updateTablePipelineConfig | ❌ |
| | | | | | |
| routing | /config/routing/all | getAllRoutingConfigs | ❌ | ❌ | ❌ |
| | /config/routing/routes/device/{routeId} | getRouteConfig | updateRouteConfig | ❌ | deleteRouteConfig |
| | /config/routing/{deviceId} | getDeviceRoutingConfig | updateDeviceRoutingConfig | ❌ | ❌ |
| | /config/routing/{deviceId}/bgp | getBGPConfig | updateBGPConfig | ❌ | ❌ |
| | /config/routing/{deviceId}/bgp/neighbors | getBGPNeighbors | ❌ | addBGPNeighbor | ❌ |
| | /config/routing/{deviceId}/ospf | getOSPFConfig | updateOSPFConfig | ❌ | ❌ |
| | /config/routing/{deviceId}/static | getStaticRoutes | ❌ | addStaticRoute | ❌ |
| | /config/routing/{deviceId}/validate | ❌ | ❌ | validateRoutingConfig | ❌ |
| | | | | | |
| security | /security/devices/{deviceId}/roles | getDeviceRoles | updateDeviceRoles | ❌ | ❌ |
| | /security/roles | getRoles | ❌ | createRole | ❌ |
| | /security/roles/{roleName} | getRole | updateRole | ❌ | deleteRole |
| | | | | | |
| switch | /config/switches | getAllSwitches | ❌ | ❌ | ❌ |
| | /config/switches/{deviceId} | getSwitchConfigById | updateSwitchConfigById | ❌ | ❌ |
| | /config/switches/{deviceId}/features | getSwitchFeatures | updateSwitchFeatures | ❌ | ❌ |
| | /config/switches/{deviceId}/ports | getAllPorts | ❌ | ❌ | ❌ |
| | /config/switches/{deviceId}/ports/{portId} | getPortConfig | updatePortConfig | ❌ | ❌ |
| | | | | | |
| telemetry | /telemetry/interface_stats | ❌ | ❌ | reportInterfaceStats | ❌ |
| | /telemetry/meters | getMeterStats | ❌ | ❌ | ❌ |
| | /telemetry/metrics | getTelemetryMetrics | ❌ | reportTelemetryMetrics | ❌ |
| | /telemetry/metrics/bulk | ❌ | ❌ | getBulkMetrics | ❌ |
| | /telemetry/metrics/{deviceId} | getDeviceMetrics | ❌ | ❌ | ❌ |
| | /telemetry/metrics/{deviceId}/environmental | getEnvironmentalMetrics | ❌ | ❌ | ❌ |
| | /telemetry/metrics/{deviceId}/resources | getResourceMetrics | ❌ | ❌ | ❌ |
| | /telemetry/metrics/{deviceId}/thresholds | getMetricThresholds | updateMetricThresholds | ❌ | ❌ |
| | /telemetry/stats | getInterfaceStats | ❌ | ❌ | ❌ |
| | /telemetry/stats/bulk | ❌ | ❌ | getBulkStats | ❌ |
| | /telemetry/stats/{deviceId} | getDeviceStats | ❌ | ❌ | ❌ |
| | /telemetry/stats/{deviceId}/interfaces | getDeviceInterfaceStats | ❌ | ❌ | ❌ |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/counters | getInterfaceCounters | ❌ | updateInterfaceCounters | clearInterfaceCounters |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/errors | getErrorStats | ❌ | ❌ | ❌ |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/qos | getQoSStats | ❌ | ❌ | ❌ |
| | | | | | |
| vlan | /config/vlans | getAllVlanConfigs | ❌ | ❌ | ❌ |
| | /config/vlans/id/{vlanId} | getVlanConfigById | updateVlanConfigById | ❌ | ❌ |
| | /config/vlans/{deviceId} | getDeviceVLANs | ❌ | createVLAN | ❌ |
| | /config/vlans/{deviceId}/{vlanId} | getVLAN | updateVLAN | ❌ | deleteVLAN |
| | /config/vlans/{deviceId}/{vlanId}/ports | ❌ | updateVLANPorts | ❌ | ❌ |
| | /config/vlans/{deviceId}/{vlanId}/properties | ❌ | updateVLANProperties | ❌ | ❌ |
