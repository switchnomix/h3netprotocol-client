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
pub struct UpdateMetricThresholdsRequestCpu {
    #[serde(rename = "warning", skip_serializing_if = "Option::is_none")]
    pub warning: Option<f64>,
    #[serde(rename = "critical", skip_serializing_if = "Option::is_none")]
    pub critical: Option<f64>,
}

impl UpdateMetricThresholdsRequestCpu {
    pub fn new() -> UpdateMetricThresholdsRequestCpu {
        UpdateMetricThresholdsRequestCpu {
            warning: None,
            critical: None,
        }
    }
}

