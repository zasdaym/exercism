pub fn annotate(minefield: &[&str]) -> Vec<String> {
    let height = minefield.len() as i32;
    (0..height)
        .map(|i| {
            let width = minefield[i as usize].len() as i32;
            (0..width)
                .map(|j| {
                    if minefield[i as usize].as_bytes()[j as usize] == b'*' {
                        '*'
                    } else {
                        match neighbors(i, j)
                            .iter()
                            .copied()
                            .filter(|&(y, x)| y >= 0 && x >= 0 && y < height && x < width)
                            .filter(|&(y, x)| minefield[y as usize].as_bytes()[x as usize] == b'*')
                            .count()
                        {
                            0 => ' ',
                            n => (n as u8 + '0' as u8) as char,
                        }
                    }
                })
                .collect()
        })
        .collect()
}

fn neighbors(y: i32, x: i32) -> Vec<(i32, i32)> {
    vec![
        (y - 1, x - 1),
        (y - 1, x),
        (y - 1, x + 1),
        (y, x - 1),
        (y, x + 1),
        (y + 1, x - 1),
        (y + 1, x),
        (y + 1, x + 1),
    ]
}
