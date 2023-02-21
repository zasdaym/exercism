pub fn annotate(minefield: &[&str]) -> Vec<String> {
    vec![]
}

fn neighbor_coordinates(i: usize, j: usize) -> Vec<(usize, usize)> {
    vec![
        (i - 1, j - 1),
        (i - 1, j),
        (i - 1, j + 1),
        (i, j - 1),
        (i, j + 1),
        (i + 1, j - 1),
        (i + 1, j),
        (i + 1, j + 1),
    ]
}
