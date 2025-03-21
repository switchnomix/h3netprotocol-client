/*
 * H3Net Protocol API
 *
 * H3Net Protocol schema with various messages and configurations.
 *
 * The version of the OpenAPI document: 1.0.0
 * Contact: h3netprotocol@googlegroups.com
 * Generated by: https://openapi-generator.tech
 */

use crate::models;
use serde::{Deserialize, Serialize};

#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct NetworkTopologyLinksInner {
    /// Source device ID
    #[serde(rename = "source")]
    pub source: String,
    /// Target device ID
    #[serde(rename = "target")]
    pub target: String,
    /// Source port name
    #[serde(rename = "sourcePort", skip_serializing_if = "Option::is_none")]
    pub source_port: Option<String>,
    /// Target port name
    #[serde(rename = "targetPort", skip_serializing_if = "Option::is_none")]
    pub target_port: Option<String>,
    /// Type of link
    #[serde(rename = "linkType", skip_serializing_if = "Option::is_none")]
    pub link_type: Option<String>,
    /// Protocol used to discover this link
    #[serde(rename = "discoveryProtocol", skip_serializing_if = "Option::is_none")]
    pub discovery_protocol: Option<DiscoveryProtocol>,
}

impl NetworkTopologyLinksInner {
    pub fn new(source: String, target: String) -> NetworkTopologyLinksInner {
        NetworkTopologyLinksInner {
            source,
            target,
            source_port: None,
            target_port: None,
            link_type: None,
            discovery_protocol: None,
        }
    }
}
/// Protocol used to discover this link
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum DiscoveryProtocol {
    #[serde(rename = "lldp")]
    Lldp,
    #[serde(rename = "cdp")]
    Cdp,
    #[serde(rename = "fdb")]
    Fdb,
    #[serde(rename = "bgp_neighbors")]
    BgpNeighbors,
}

impl Default for DiscoveryProtocol {
    fn default() -> DiscoveryProtocol {
        Self::Lldp
    }
}

