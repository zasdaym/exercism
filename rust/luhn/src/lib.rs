pub fn is_valid(code: &str) -> bool {
    if code.chars().any(|c| !c.is_ascii_digit() && !c.is_whitespace()) {
        return false;
    }

    let digits: Vec<u32> = code.chars().filter_map(|c| c.to_digit(10)).collect();

    if digits.len() < 2 {
        return false;
    }

    let sum: u32 = digits
        .into_iter()
        .rev()
        .enumerate()
        .map(|(i, d)| transform(i, d))
        .sum();

    sum % 10 == 0
}

fn transform(i: usize, number: u32) -> u32 {
    if i % 2 == 1 {
        let doubled = number * 2;
        if doubled > 9 { doubled - 9 } else { doubled }
    } else {
        number
    }
}
