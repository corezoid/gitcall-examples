use libc::c_int;
use std::env;
use jsonrpc_core::IoHandler as JsonRpcCoreIoHandler;
use jsonrpc_core::Params as JsonRpcCoreParams;
use jsonrpc_http_server::ServerBuilder as ServerBuilderHttp;
use serde_json::json;
use signal_hook::consts::signal::*;

use std::thread;
use std::io::Error;
use serde_json;


#[cfg(not(feature = "extended-siginfo"))]
use signal_hook::iterator::Signals;

fn main() -> Result<(), Error>  {

    const SIGNALS: &[c_int] = &[SIGTERM, SIGINT, SIGQUIT];
    let mut sigs = Signals::new(SIGNALS)?;

    thread::spawn(move || {
        for signal in &mut sigs {
            eprintln!("Received signal {:?}", signal);
            std::process::exit(0);
        }
    });

    let mut addr: String = "0.0.0.0:".to_owned();
    let port = env::var("GITCALL_PORT").expect("$GITCALL_PORT is not set");
    addr.push_str(&port);

    let mut io = JsonRpcCoreIoHandler::default();
    io.add_method("Usercode.Run", |_params: JsonRpcCoreParams| {
        let contacts = json!({
            "rust": "Hello, world!"
        });
        Ok(contacts)
    });
    let server = ServerBuilderHttp::new(io)
        .start_http(&addr.parse().unwrap())
        .expect("Server must start with no issues.");

    eprintln!("Listening on {:?}", addr);
    server.wait();

    Ok(())
}
