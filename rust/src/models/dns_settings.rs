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

/// DnsSettings : DNS configuration
#[derive(Clone, Default, Debug, PartialEq, Serialize, Deserialize)]
pub struct DnsSettings {
    #[serde(rename = "servers", skip_serializing_if = "Option::is_none")]
    pub servers: Option<Vec<String>>,
    #[serde(rename = "domain", skip_serializing_if = "Option::is_none")]
    pub domain: Option<String>,
    #[serde(rename = "searchDomains", skip_serializing_if = "Option::is_none")]
    pub search_domains: Option<Vec<String>>,
}

impl DnsSettings {
    /// DNS configuration
    pub fn new() -> DnsSettings {
        DnsSettings {
            servers: None,
            domain: None,
            search_domains: None,
        }
    }
}

