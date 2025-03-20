use webtransport_quinn as webtransport;
use quinn::Connection;
use serde::{Serialize, Deserialize};
use crate::error::Error;

pub struct WebTransportClient {
    conn: Connection,
}

impl WebTransportClient {
    pub async fn connect(url: &str) -> Result<Self, Error> {
        // Implementation will be added later
        unimplemented!()
    }
}
