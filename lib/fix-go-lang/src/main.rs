mod app;
mod enums;
mod infra;
mod presentation;

use app::shutdown::shutdown;
use presentation::messages::print_error;

fn main() {
    if let Err(e) = shutdown() {
        print_error(&e);
    }
}
