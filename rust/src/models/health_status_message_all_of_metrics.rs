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
pub struct HealthStatusMessageAllOfMetrics {
    #[serde(rename = "activeConnections", skip_serializing_if = "Option::is_none")]
    pub active_connections: Option<i32>,
    #[serde(rename = "cpuUsage", skip_serializing_if = "Option::is_none")]
    pub cpu_usage: Option<f64>,
    #[serde(rename = "memoryUsage", skip_serializing_if = "Option::is_none")]
    pub memory_usage: Option<f64>,
}

impl HealthStatusMessageAllOfMetrics {
    pub fn new() -> HealthStatusMessageAllOfMetrics {
        HealthStatusMessageAllOfMetrics {
            active_connections: None,
            cpu_usage: None,
            memory_usage: None,
        }
    }
}

