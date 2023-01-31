use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut word_count: HashMap<&str, i32> = HashMap::new();
    for word in magazine {
        *word_count.entry(word).or_default() += 1
    }
    for word in note {
        *word_count.entry(word).or_default() -= 1
    }
    word_count.values().all(|count| *count >= 0)
}
