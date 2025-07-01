pub fn annotate(minefield: &[&str]) -> Vec<String> {
    if minefield.is_empty() {
        return vec![];
    }
    let rows = minefield.len();
    let cols = minefield[0].len();

    let mut result = Vec::with_capacity(rows);
    for _ in 0..rows {
        result.push(vec![' '; cols]);
    }

    for (i, row) in minefield.iter().enumerate() {
        for (j, cell) in row.chars().enumerate() {
            if cell == '*' {
                result[i][j] = '*';
            } else {
                let count = count_mines(i, j, minefield);
                if count > 0 {
                    result[i][j] = (count as u8 + b'0') as char;
                }
            }
        }
    }

    result
        .into_iter()
        .map(|row| row.into_iter().collect())
        .collect()
}

fn count_mines(row: usize, col: usize, minefield: &[&str]) -> i32 {
    let mut count = 0;

    for i in row.saturating_sub(1)..=row.saturating_add(1) {
        for j in col.saturating_sub(1)..=col.saturating_add(1) {
            if i < minefield.len() && j < minefield[0].len() && (i != row || j != col) {
                if minefield[i].chars().nth(j).unwrap() == '*' {
                    count += 1;
                }
            }
        }
    }

    count
}
