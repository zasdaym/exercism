pub fn recite(start_bottles: u32, take_down: u32) -> String {
    let mut lines = vec![];
    for i in 0..take_down {
        let n = start_bottles - i;
        lines.push(format!(
            "{} green {} hanging on the wall,",
            number(n),
            bottle(n)
        ));
        lines.push(format!(
            "{} green {} hanging on the wall,",
            number(n),
            bottle(n)
        ));
        lines.push("And if one green bottle should accidentally fall,".to_string());
        lines.push(format!(
            "There'll be {} green {} hanging on the wall.\n",
            number(n - 1).to_lowercase(),
            bottle(n - 1)
        ));
    }
    lines.join("\n")
}

fn number(n: u32) -> String {
    match n {
        1 => "One".to_string(),
        2 => "Two".to_string(),
        3 => "Three".to_string(),
        4 => "Four".to_string(),
        5 => "Five".to_string(),
        6 => "Six".to_string(),
        7 => "Seven".to_string(),
        8 => "Eight".to_string(),
        9 => "Nine".to_string(),
        10 => "Ten".to_string(),
        _ => "No".to_string(),
    }
}

fn bottle(n: u32) -> String {
    if n == 1 {
        "bottle".to_string()
    } else {
        "bottles".to_string()
    }
}
