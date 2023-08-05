pub fn reply(message: &str) -> &str {
    match message {
        _ if is_question(message) && is_yelling(message) => "Calm down, I know what I'm doing!",
        _ if is_question(message) => "Sure.",
        _ if is_yelling(message) => "Whoa, chill out!",
        _ if is_empty(message) => "Fine. Be that way!",
        _ => "Whatever.",
    }
}

fn is_question(message: &str) -> bool {
    message.trim().ends_with('?')
}

fn is_yelling(message: &str) -> bool {
    message.to_uppercase() == message && message.to_lowercase() != message
}

fn is_empty(message: &str) -> bool {
    message.trim().is_empty()
}
