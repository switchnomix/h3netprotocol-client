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
pub struct ResponseMessage {
    #[serde(rename = "type")]
    pub r#type: Type,
    /// ISO 8601 formatted timestamp with timezone
    #[serde(rename = "timestamp")]
    pub timestamp: String,
    /// Schema version
    #[serde(rename = "schemaVersion")]
    pub schema_version: String,
    #[serde(rename = "status")]
    pub status: Status,
    /// Optional details about the response.
    #[serde(rename = "message", skip_serializing_if = "Option::is_none")]
    pub message: Option<String>,
}

impl ResponseMessage {
    pub fn new(r#type: Type, timestamp: String, schema_version: String, status: Status) -> ResponseMessage {
        ResponseMessage {
            r#type,
            timestamp,
            schema_version,
            status,
            message: None,
        }
    }
}
/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Type {
    #[serde(rename = "RESPONSE")]
    Response,
}

impl Default for Type {
    fn default() -> Type {
        Self::Response
    }
}
/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Status {
    #[serde(rename = "success")]
    Success,
    #[serde(rename = "failure")]
    Failure,
}

impl Default for Status {
    fn default() -> Status {
        Self::Success
    }
}

