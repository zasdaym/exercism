const MAX_VALUE: usize = 104_744;

pub fn nth(n: u32) -> u32 {
    let primes = calculate_primes();
    let mut count = 0;
    let mut result = 2;
    while count < n {
        result += 1;
        if primes[result as usize] {
            count += 1;
        }
    }
    result
}

fn calculate_primes() -> Vec<bool> {
    let mut result = vec![true; MAX_VALUE];
    for i in 2..MAX_VALUE {
        if !result[i] {
            continue;
        }
        for j in (i + i..MAX_VALUE).step_by(i) {
            result[j] = false;
        }
    }
    result
}
