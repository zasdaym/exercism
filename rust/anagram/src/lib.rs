use std::collections::{HashMap, HashSet};

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let word_lower = word.to_lowercase();
    let target = encode_frequency(&word_lower);

    possible_anagrams
        .iter()
        .filter(|candidate| {
            let candidate_lower = candidate.to_lowercase();
            candidate_lower.len() == word_lower.len()
                && candidate_lower != word_lower
                && encode_frequency(&candidate_lower) == target
        })
        .copied()
        .collect()
}

fn encode_frequency(word: &str) -> HashMap<char, usize> {
    let mut frequency = HashMap::new();
    for c in word.chars() {
        *frequency.entry(c).or_insert(0) += 1;
    }
    frequency
}
