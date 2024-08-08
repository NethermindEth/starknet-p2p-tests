use libp2p::{
    core::{
        muxing::StreamMuxerBox,
        transport::OrTransport,
        upgrade,
    }, dns, futures, identity, noise, relay, swarm, tcp, yamux, Multiaddr, PeerId, Swarm, Transport
};
use libp2p::swarm::SwarmEvent;
use futures::StreamExt;

use std::{error::Error, time::Duration};

pub async fn initialize_p2p(addr: &str) -> Result<(), Box<dyn Error>> {
    // Generate a keypair for the local peer
    let local_key = identity::Keypair::generate_ed25519();
    let local_peer_id = local_key.public().to_peer_id();
    
    // Create a multiaddr for the local peer
    let local_multiaddr: Multiaddr = format!("/ip4/127.0.0.1/tcp/7777/p2p/{}", local_peer_id).parse()?;
    println!("Local peer address: {}", local_multiaddr);

    // Initialize the relay transport
    let (relay_transport, relay_behaviour) = relay::client::new(local_peer_id.clone());

    // Create the transport
    let transport = create_transport(&local_key, relay_transport);

    // Create the Swarm
    let mut swarm = Swarm::new(
        transport,
        relay_behaviour,
        local_peer_id,
        swarm::Config::with_tokio_executor()
            .with_idle_connection_timeout(Duration::from_secs(3600 * 365))
    );

    let multiaddr: Multiaddr = addr.parse()?;

    swarm.dial(multiaddr.clone())?;
    println!("Dialing peer at address: {}", multiaddr);

    // Handle events
    loop {
        match swarm.next().await {
            Some(SwarmEvent::ConnectionEstablished { peer_id, .. }) => {
                println!("Connected to peer: {}", peer_id);
                return Ok(());
            }
            Some(SwarmEvent::OutgoingConnectionError { error, .. }) => {
                return Err(format!("Failed to connect to peer: {:?}", error).into());
            }
            _ => {}
        }
    }
}

pub fn create_transport(
    keypair: &libp2p::identity::Keypair,
    relay_transport: libp2p::relay::client::Transport,
) -> libp2p::core::transport::Boxed<(PeerId, StreamMuxerBox)> {
    let transport = tcp::tokio::Transport::new(tcp::Config::new());
    let transport = OrTransport::new(transport, relay_transport);
    let transport = dns::tokio::Transport::system(transport).unwrap();

    let noise_config =
        noise::Config::new(keypair).expect("Signing libp2p-noise static DH keypair failed.");

    transport
        .upgrade(upgrade::Version::V1)
        .authenticate(noise_config)
        .multiplex(yamux::Config::default())
        .boxed()
}
