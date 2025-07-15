pub fn annotate(garden: &[&str]) -> Vec<String> {
    garden
        .iter()
        .enumerate()
        .map(|(y, row)| {
            row.chars()
                .enumerate()
                .map(|(x, chr)| match chr {
                    '*' => chr,
                    _ => annotation(garden, y, x),
                })
                .collect()
        })
        .collect()
}

fn annotation(garden: &[&str], y: usize, x: usize) -> char {
    let rows = garden.len();
    let cols = garden[0].len();

    let y_start = y.saturating_sub(1);
    let y_end = y.saturating_add(1).min(rows - 1);

    let x_start = x.saturating_sub(1);
    let x_end = x.saturating_add(1).min(cols - 1);

    let mut count = 0;
    for i in y_start..=y_end {
        for j in x_start..=x_end {
            if i == y && j == x {
                continue;
            }
            if garden[i].chars().nth(j) == Some('*') {
                count += 1;
            }
        }
    }

    if count == 0 {
        ' '
    } else {
        (count + b'0') as char
    }
}
