//! Command line interface.
use cdk_config::Config;
use clap::Parser;
use cli::Cli;
use execute::Execute;
use std::path::PathBuf;
use std::process::Command;
use std::sync::Arc;

mod cli;
mod logging;

const CDK_CLIENT_PATH: &str = "cdk-client";
const CDK_ERIGON_PATH: &str = "cdk-erigon";

fn main() -> anyhow::Result<()> {
    dotenvy::dotenv().ok();

    let cli = Cli::parse();

    match cli.cmd {
        cli::Commands::Run { cfg } => run(cfg)?,
    }

    Ok(())
}

/// This is the main node entrypoint.
///
/// This function starts everything needed to run an Agglayer node.
/// Starting by a Tokio runtime which can be used by the different components.
/// The configuration file is parsed and used to configure the node.
///
/// This function returns on fatal error or after graceful shutdown has
/// completed.
pub fn run(cfg: PathBuf) -> anyhow::Result<()> {
    // Load the configuration file
    let config: Arc<Config> = Arc::new(toml::from_str(&std::fs::read_to_string(cfg)?)?);

    // Run the node passing the parsed config values as flags
    let mut command = Command::new(CDK_CLIENT_PATH);

    let output = command.execute_output().unwrap();

    if let Some(exit_code) = output.status.code() {
        if exit_code == 0 {
            println!("Ok.");
        } else {
            eprintln!("Failed.");
        }
    } else {
        eprintln!("Interrupted!");
    }

    // Initialize the logger
    logging::tracing(&config.log);

    Ok(())
}