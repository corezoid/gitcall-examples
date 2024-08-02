use jsonrpc_core::IoHandler as JsonRpcCoreIoHandler;
use jsonrpc_core::Params as JsonRpcCoreParams;
use jsonrpc_http_server::{hyper, ServerBuilder};
use libc::c_int;
use serde_json::json;
use signal_hook::consts::signal::*;
use std::env;

use serde_json;
use std::io::Error;
use std::thread;
use std::time::{SystemTime, UNIX_EPOCH};

#[cfg(not(feature = "extended-siginfo"))]
use signal_hook::iterator::Signals;

fn main() {
    wait_sign().unwrap();

    let mut addr: String = "0.0.0.0:".to_owned();
    let port = env::var("GITCALL_PORT").expect("$GITCALL_PORT is not set");
    addr.push_str(&port);

    let mut io = JsonRpcCoreIoHandler::default();
    io.add_method("Usercode.Run", |_params: JsonRpcCoreParams| async {
        let contacts = json!({
            "rust": "Hello, world!"
        });
        Ok(contacts)
    });
    let server = ServerBuilder::new(io)
        .request_middleware(|request: hyper::Request<hyper::Body>| {
            let id = request
                .headers()
                .get("x-request-id")
                .map(|h| h.to_str().unwrap_or("").to_owned())
                .unwrap_or("".to_owned());

            eprintln!(
                "[req] time={:?} id={:?}",
                SystemTime::now()
                    .duration_since(UNIX_EPOCH)
                    .unwrap()
                    .as_millis(),
                id
            );

            let result = request.into();

            eprintln!(
                "[res] time={:?} id={:?}",
                SystemTime::now()
                    .duration_since(UNIX_EPOCH)
                    .unwrap()
                    .as_millis(),
                id
            );

            result
        })
        .start_http(&addr.parse().unwrap())
        .expect("Server must start with no issues.");

    eprintln!("Listening on {:?}", addr);
    server.wait();
}

fn wait_sign() -> Result<(), Error> {
    const SIGNALS: &[c_int] = &[SIGTERM, SIGINT, SIGQUIT];
    let mut sigs = Signals::new(SIGNALS)?;
    thread::spawn(move || {
        for signal in &mut sigs {
            eprintln!("Received signal {:?}", signal);
            std::process::exit(0);
        }
    });
    Ok(())
}
