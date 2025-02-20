pub enum OS {
    Windows,
    Linux,
    MacOS,
    Unknown,
}

impl OS {
    pub fn current() -> Self {
        match std::env::consts::OS {
            "windows" => OS::Windows,
            "linux" => OS::Linux,
            "macos" => OS::MacOS,
            _ => OS::Unknown,
        }
    }
}
