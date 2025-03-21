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

/// VlanQos : QoS settings for the VLAN
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct VlanQos {
    #[serde(rename = "priority", skip_serializing_if = "Option::is_none")]
    pub priority: Option<i32>,
    #[serde(rename = "bandwidthLimit", skip_serializing_if = "Option::is_none")]
    pub bandwidth_limit: Option<i32>,
    #[serde(rename = "dscp", skip_serializing_if = "Option::is_none")]
    pub dscp: Option<i32>,
}

impl VlanQos {
    /// QoS settings for the VLAN
    pub fn new() -> VlanQos {
        VlanQos {
            priority: None,
            bandwidth_limit: None,
            dscp: None,
        }
    }
}

