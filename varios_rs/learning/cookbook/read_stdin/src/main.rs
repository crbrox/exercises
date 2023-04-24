use std::io;
use std::io::prelude::*;

fn main() {
    println!("Hello, world!");
    print_single_line("Please enter your forename: ");
    let forename = read_line_iter();
    print_single_line("Please enter your surname: ");
    let surname = read_line_buffer();
    print_single_line("Please enter your age: ");
    let age = read_number();
    println!("Hello, {age} year old human named {forename} {surname}")
}

fn print_single_line(text: &str) {
    print!("{text}");
    io::stdout().flush().expect("should flush to stdout");
}

fn read_line_iter() -> String {
    let stdin = io::stdin();
    let input = stdin.lock().lines().next();
    input
        .expect("should be at least one line in input")
        .expect("should be read the line")
        .trim()
        .to_string()
}

fn read_line_buffer() -> String {
    let mut input = String::new();
    io::stdin()
        .read_line(&mut input)
        .expect("should read one line");
    input.trim().to_string()
}

fn read_number() -> i32 {
    let stdin = io::stdin();
    loop {
        for line in stdin.lock().lines() {
            let input = line.expect("should read a line");
            match input.trim().parse::<i32>() {
                Ok(num) => return num,
                Err(e) => println!("Failed to read number: {e}"),
            }
        }
    }
}
