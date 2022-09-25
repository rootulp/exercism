/// various log levels
#[derive(Clone, PartialEq, Eq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
    Debug,
}

/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
    let label = label(level);
    format!("[{label}]: {message}")
}
pub fn info(message: &str) -> String {
    log(LogLevel::Info, message)
}
pub fn warn(message: &str) -> String {
    log(LogLevel::Warning, message)
}
pub fn error(message: &str) -> String {
    log(LogLevel::Error, message)
}

fn label(level: LogLevel) -> String {
    match level {
        LogLevel::Info => "INFO".to_string(),
        LogLevel::Warning => "WARNING".to_string(),
        LogLevel::Error => "ERROR".to_string(),
        LogLevel::Debug => "DEBUG".to_string(),
    }
}
