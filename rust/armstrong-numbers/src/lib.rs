pub fn is_armstrong_number(num: u32) -> bool {
    let mut n = num;
    let mut exp = 0;
    while n > 0 {
        n /= 10;
        exp += 1;
    }

    let mut n = num;
    let mut sum = 0;
    while n > 0 {
        let digit = n % 10;
        sum += digit.pow(exp);
        n /= 10;
    }

    sum == num
}
