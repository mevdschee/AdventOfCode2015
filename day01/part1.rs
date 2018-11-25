use std::io::{self, Read};

fn main() {
    let mut input = String::new();
    let stdin = io::stdin();
    stdin.lock().read_to_string(&mut input).expect("Could not read stdin");
    let output: i32 = input.chars()
        .map(|item| match item {
            '(' => 1,
            ')' => -1,
            _ => 0
        })
        .sum();
    println!("{}",output);
}
