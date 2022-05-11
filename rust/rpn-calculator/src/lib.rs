#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut operands: Vec<i32> = Vec::new();

    for input in inputs {
        match input {
            CalculatorInput::Value(x) => operands.push(*x),
            _ => {
                if operands.len() < 2 {
                    return None;
                }
                let operand_b = operands.pop().unwrap();
                let operand_a = operands.pop().unwrap();
                let result: i32 = match input {
                    CalculatorInput::Add => operand_a + operand_b,
                    CalculatorInput::Subtract => operand_a - operand_b,
                    CalculatorInput::Multiply => operand_a * operand_b,
                    CalculatorInput::Divide => operand_a / operand_b,
                    CalculatorInput::Value(_) => panic!("operand in operator stack"),
                };
                operands.push(result);
            }
        }
    }

    if operands.len() == 1 {
        return Some(operands.pop().unwrap());
    }

    return None;
}
