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

/// InterfaceSpeed : Interface speed in standard units
/// Interface speed in standard units
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum InterfaceSpeed {
    #[serde(rename = "10Mbps")]
    Variant10Mbps,
    #[serde(rename = "100Mbps")]
    Variant100Mbps,
    #[serde(rename = "1Gbps")]
    Variant1Gbps,
    #[serde(rename = "2.5Gbps")]
    Variant2Period5Gbps,
    #[serde(rename = "5Gbps")]
    Variant5Gbps,
    #[serde(rename = "10Gbps")]
    Variant10Gbps,
    #[serde(rename = "25Gbps")]
    Variant25Gbps,
    #[serde(rename = "40Gbps")]
    Variant40Gbps,
    #[serde(rename = "100Gbps")]
    Variant100Gbps,

}

impl std::fmt::Display for InterfaceSpeed {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
        match self {
            Self::Variant10Mbps => write!(f, "10Mbps"),
            Self::Variant100Mbps => write!(f, "100Mbps"),
            Self::Variant1Gbps => write!(f, "1Gbps"),
            Self::Variant2Period5Gbps => write!(f, "2.5Gbps"),
            Self::Variant5Gbps => write!(f, "5Gbps"),
            Self::Variant10Gbps => write!(f, "10Gbps"),
            Self::Variant25Gbps => write!(f, "25Gbps"),
            Self::Variant40Gbps => write!(f, "40Gbps"),
            Self::Variant100Gbps => write!(f, "100Gbps"),
        }
    }
}

impl Default for InterfaceSpeed {
    fn default() -> InterfaceSpeed {
        Self::Variant10Mbps
    }
}

