use std::env;
fn main() {
    println!("Got following parameters: ");
    for arg in env::args() {
        println!("- {}", arg);
    }

    println!("iterator");
    let mut args = env::args();
    if let Some(arg) = args.next() {
        println!("The path to this program is: {}", arg);
    }
    if let Some(arg) = args.next() {
        println!("The first parameter is: {}", arg);
    }
    if let Some(arg) = args.next() {
        println!("The second parameter is: {}", arg);
    }

    println!("vector");
    let args: Vec<_> = env::args().collect();
    println!("The path to this program is: {}", args[0]);
    if args.len() > 1 {
        println!("The first parameter is: {}", args[1]);
    }
    if args.len() > 2 {
        println!("The second parameter is: {}", args[2]);
    }
}

// use clap
