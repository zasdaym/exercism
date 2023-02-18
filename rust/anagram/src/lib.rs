use std::collections::HashSet;

pub fn anagrams_for<'a>(word: &str, possible_anagrams: &[&'a str]) -> HashSet<&'a str> {
    let lowercased_word = lowercase_chars(word);
    let normalized_word = sorted_chars(&lowercased_word);
    possible_anagrams
        .iter()
        .copied()
        .filter(|candidate| {
            let lowercased_candidate = lowercase_chars(candidate);
            let normalized_candidate = sorted_chars(&lowercased_candidate);
            lowercased_word != lowercased_candidate && normalized_word == normalized_candidate
        })
        .collect()
}

fn lowercase_chars(word: &str) -> Vec<char> {
    word.chars().flat_map(|c| c.to_lowercase()).collect()
}

fn sorted_chars(chars: &Vec<char>) -> Vec<char> {
    let mut result = chars.clone();
    result.sort_unstable();
    return result;
}
