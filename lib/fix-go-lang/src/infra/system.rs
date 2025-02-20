use crate::enums::OS;
use crate::presentation::messages::{ERROR_SHUTDOWN_FAILED, ERROR_UNKNOWN_OS};
use std::process::Command;

pub fn get_shutdown_command(os: &OS) -> Result<Vec<&'static str>, &'static str> {
    match os {
        OS::Windows => Ok(vec!["shutdown", "/s", "/t", "0"]),
        OS::Linux | OS::MacOS => Ok(vec!["shutdown", "-h", "now"]),
        OS::Unknown => Err(ERROR_UNKNOWN_OS),
    }
}

pub fn execute_command(args: &[&str]) -> Result<(), String> {
    Command::new(args[0])
        .args(&args[1..])
        .spawn()
        .map_err(|e| format!("{} {}", ERROR_SHUTDOWN_FAILED, e))
        .map(|_| ())
}
