#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack = vec![0; 0];
    for input in inputs {
        match input {
            CalculatorInput::Value(number) => stack.push(*number),
            CalculatorInput::Add => {
                let b = stack.pop()?;
                let a = stack.pop()?;
                stack.push(a + b);
            }
            CalculatorInput::Subtract => {
                let b = stack.pop()?;
                let a = stack.pop()?;
                stack.push(a - b);
            }
            CalculatorInput::Multiply => {
                let b = stack.pop()?;
                let a = stack.pop()?;
                stack.push(a * b);
            }
            CalculatorInput::Divide => {
                let b = stack.pop()?;
                let a = stack.pop()?;
                stack.push(a / b);
            }
        }
    }
    match stack.len() {
        1 => stack.pop(),
        _ => None,
    }
}
