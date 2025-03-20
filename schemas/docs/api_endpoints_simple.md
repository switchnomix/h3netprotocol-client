# H3Net Protocol API Endpoints (Simplified)

Generated on: Thu Mar 20 23:29:22 UTC 2025

| Category | Endpoint | Methods |
| -------- | -------- | ------- |
| acl | /config/acls | GET, POST |
| | /config/acls/device/{deviceId} | GET, POST |
| | /config/acls/{aclId} | GET, PUT |
| | | |
| config | /config/acls/{deviceId}/{aclName} | GET, PUT, DELETE |
| | /config/acls/{deviceId}/{aclName}/rules | POST |
| | /config/acls/{deviceId}/{aclName}/rules/{ruleId} | PUT, DELETE |
| | /config/devices | GET, POST |
| | /config/devices/info/bulk | POST |
| | /config/devices/{deviceId} | GET, PUT |
| | /config/devices/{deviceId}/info | GET, POST |
| | /config/devices/{deviceId}/interfaces | GET |
| | /config/devices/{deviceId}/interfaces/{interfaceId} | PUT |
| | /config/devices/{deviceId}/system | GET, PUT |
| | /config/routing | GET, PUT |
| | /config/switch | GET, POST |
| | /config/vlan | GET, PUT, POST |
| | | |
| core | /devices | GET, POST |
| | /devices/{deviceId} | GET |
| | /health | GET |
| | /hello | POST |
| | | |
| discovery | /discovery/capabilities | GET, POST |
| | /discovery/capabilities/bulk | POST |
| | /discovery/network | GET, POST |
| | /discovery/network/neighbors/{deviceId} | GET |
| | /discovery/network/options | GET, PUT |
| | /discovery/network/topology | GET |
| | | |
| events | /events | POST |
| | | |
| operations | /operations/state | GET, POST |
| | /operations/state/{deviceId} | GET, POST |
| | /operations/state/{deviceId}/interfaces | GET |
| | /operations/state/{deviceId}/reboot | POST |
| | /operations/state/{deviceId}/tables | GET |
| | | |
| pipeline | /config/pipeline | GET, POST |
| | /config/pipeline/flows | GET, POST |
| | /config/pipeline/tables | GET, POST |
| | /pipeline/flows | GET, POST |
| | /tables | GET, POST |
| | /tables/{tableId} | GET, PUT |
| | /tables/{tableId}/pipeline | POST |
| | | |
| routing | /config/routing/all | GET |
| | /config/routing/routes/device/{routeId} | GET, PUT, DELETE |
| | /config/routing/{deviceId} | GET, PUT |
| | /config/routing/{deviceId}/bgp | GET, PUT |
| | /config/routing/{deviceId}/bgp/neighbors | GET, POST |
| | /config/routing/{deviceId}/ospf | GET, PUT |
| | /config/routing/{deviceId}/static | GET, POST |
| | /config/routing/{deviceId}/validate | POST |
| | | |
| security | /security/devices/{deviceId}/roles | GET, PUT |
| | /security/roles | GET, POST |
| | /security/roles/{roleName} | GET, PUT, DELETE |
| | | |
| switch | /config/switches | GET |
| | /config/switches/{deviceId} | GET, PUT |
| | /config/switches/{deviceId}/features | GET, PUT |
| | /config/switches/{deviceId}/ports | GET |
| | /config/switches/{deviceId}/ports/{portId} | GET, PUT |
| | | |
| telemetry | /telemetry/interface_stats | POST |
| | /telemetry/meters | GET |
| | /telemetry/metrics | GET, POST |
| | /telemetry/metrics/bulk | POST |
| | /telemetry/metrics/{deviceId} | GET |
| | /telemetry/metrics/{deviceId}/environmental | GET |
| | /telemetry/metrics/{deviceId}/resources | GET |
| | /telemetry/metrics/{deviceId}/thresholds | GET, PUT |
| | /telemetry/stats | GET |
| | /telemetry/stats/bulk | POST |
| | /telemetry/stats/{deviceId} | GET |
| | /telemetry/stats/{deviceId}/interfaces | GET |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/counters | GET, POST, DELETE |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/errors | GET |
| | /telemetry/stats/{deviceId}/interfaces/{interfaceId}/qos | GET |
| | | |
| vlan | /config/vlans | GET |
| | /config/vlans/id/{vlanId} | GET, PUT |
| | /config/vlans/{deviceId} | GET, POST |
| | /config/vlans/{deviceId}/{vlanId} | GET, PUT, DELETE |
| | /config/vlans/{deviceId}/{vlanId}/ports | PUT |
| | /config/vlans/{deviceId}/{vlanId}/properties | PUT |
