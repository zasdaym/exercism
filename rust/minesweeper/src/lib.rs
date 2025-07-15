pub fn annotate(minefield: &[&str]) -> Vec<String> {
    minefield
        .iter()
        .enumerate()
        .map(|(y, row)| {
            row.chars()
                .enumerate()
                .map(|(x, c)| match c {
                    '*' => '*',
                    _ => match nearby_mines(minefield, y, x) {
                        0 => ' ',
                        n => (n + b'0') as char,
                    },
                })
                .collect()
        })
        .collect()
}

fn nearby_mines(minefield: &[&str], y: usize, x: usize) -> u8 {
    let y_start = y.saturating_sub(1);
    let y_end = (y + 1).min(minefield.len() - 1);
    let x_start = x.saturating_sub(1);
    let x_end = (x + 1).min(minefield[0].len() - 1);

    let mut count = 0;
    for i in y_start..=y_end {
        for j in x_start..=x_end {
            let is_self = i == y && j == x;
            if is_self {
                continue;
            }
            if minefield[i].chars().nth(j) == Some('*') {
                count += 1;
            }
        }
    }
    count
}
