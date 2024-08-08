
#[test]
fn test_connect_to_peer() {
    let node = crate::tools::synthetic_node::SyntheticNode::new();
    let peer_address = "/ip4/35.237.66.77/tcp/7777/p2p/12D3KooWR8ikUDiinyE5wgdYiqsdLfJRsBDYKGii6L3oyoipVEaV";
    
    let result = node.connect(peer_address);
    
    assert!(result.is_ok(), "Failed to connect: {:?}", result.err().unwrap());
}