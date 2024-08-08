use crate::protocol::starknet;
use tokio::runtime::Runtime;
use std::error::Error;

pub struct SyntheticNode;

impl SyntheticNode {
    pub fn new() -> Self {
        SyntheticNode
    }

    pub fn connect(&self, address: &str) -> Result<(), Box<dyn Error>> {
        let runtime = Runtime::new().expect("Failed to create Tokio runtime");
        runtime.block_on(async {
            starknet::initialize_p2p(address).await
        })
    }
}