#[derive(Debug, PartialEq, Eq)]
pub enum Comparison {
    Equal,
    Sublist,
    Superlist,
    Unequal,
}

pub fn sublist<T: PartialEq>(_first_list: &[T], _second_list: &[T]) -> Comparison {
    if _first_list.len() == _second_list.len() {
        return if _first_list == _second_list {
            Comparison::Equal
        } else {
            Comparison::Unequal
        };
    }

    let (short_list, long_list, result) = match _first_list.len() < _second_list.len() {
        true => (_first_list, _second_list, Comparison::Sublist),
        false => (_second_list, _first_list, Comparison::Superlist),
    };

    if short_list.is_empty()
        || long_list
            .windows(short_list.len())
            .any(|window| window == short_list)
    {
        result
    } else {
        Comparison::Unequal
    }
}
