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
pub struct TelemetryMessageMetrics {
    #[serde(rename = "packetCount")]
    pub packet_count: i32,
    #[serde(rename = "errorCount")]
    pub error_count: i32,
    #[serde(rename = "deviceSpecificMetrics", skip_serializing_if = "Option::is_none")]
    pub device_specific_metrics: Option<Box<models::DeviceMetrics>>,
    #[serde(rename = "timestamp", skip_serializing_if = "Option::is_none")]
    pub timestamp: Option<String>,
    #[serde(rename = "interval", skip_serializing_if = "Option::is_none")]
    pub interval: Option<i32>,
    #[serde(rename = "sequence", skip_serializing_if = "Option::is_none")]
    pub sequence: Option<i32>,
}

impl TelemetryMessageMetrics {
    pub fn new(packet_count: i32, error_count: i32) -> TelemetryMessageMetrics {
        TelemetryMessageMetrics {
            packet_count,
            error_count,
            device_specific_metrics: None,
            timestamp: None,
            interval: None,
            sequence: None,
        }
    }
}

