use crate::enums::OS;
use crate::infra::system::{execute_command, get_shutdown_command};
use crate::presentation::messages::ERROR_SHUTDOWN_FAILED;

pub fn shutdown() -> Result<(), String> {
    let os = OS::current();
    let args = get_shutdown_command(&os)?;
    execute_command(&args).map_err(|e| format!("{} {}", ERROR_SHUTDOWN_FAILED, e))
}
