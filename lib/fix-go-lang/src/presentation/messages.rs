pub const ERROR_UNKNOWN_OS: &str = "Can't fix Go Lang";

pub const ERROR_COMMAND_FAILED: &str =
    "Failed to execute {:?}: {}. Maybe I should rewrite this in Go...";

pub const ERROR_SHUTDOWN_FAILED: &str = "Couldn't shut down the system: {}. \
     Guess I'll have to manually turn it off...";

pub fn print_error(message: &str) {
    eprintln!("{}", message);
}
