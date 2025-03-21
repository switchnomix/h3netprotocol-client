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

/// PortConfigAllOfPortSecurity : Port security settings
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct PortConfigAllOfPortSecurity {
    #[serde(rename = "enabled", skip_serializing_if = "Option::is_none")]
    pub enabled: Option<bool>,
    #[serde(rename = "maxMacAddresses", skip_serializing_if = "Option::is_none")]
    pub max_mac_addresses: Option<i32>,
    #[serde(rename = "violationAction", skip_serializing_if = "Option::is_none")]
    pub violation_action: Option<ViolationAction>,
}

impl PortConfigAllOfPortSecurity {
    /// Port security settings
    pub fn new() -> PortConfigAllOfPortSecurity {
        PortConfigAllOfPortSecurity {
            enabled: None,
            max_mac_addresses: None,
            violation_action: None,
        }
    }
}
/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum ViolationAction {
    #[serde(rename = "protect")]
    Protect,
    #[serde(rename = "restrict")]
    Restrict,
    #[serde(rename = "shutdown")]
    Shutdown,
}

impl Default for ViolationAction {
    fn default() -> ViolationAction {
        Self::Protect
    }
}

