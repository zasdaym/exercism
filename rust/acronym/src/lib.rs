pub fn abbreviate(phrase: &str) -> String {
    let mut result = String::new();
    let mut last_was_upper = false;

    for (i, c) in phrase.chars().enumerate() {
        if c.is_alphabetic() || c.is_whitespace() || c == '-' {
            if i > 0 && c.is_uppercase() && !last_was_upper {
                result.push(' ');
            }
            result.push(c);
            last_was_upper = c.is_uppercase();
        }
    }

    result
        .split(|c: char| c.is_whitespace() || c == '-')
        .filter(|s| !s.is_empty())
        .map(|s| s.chars().next().unwrap().to_uppercase().to_string())
        .collect()
}
