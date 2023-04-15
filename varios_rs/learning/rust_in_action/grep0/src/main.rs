use clap::{App, Arg};
use regex::Regex;
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

fn main() -> Result<(), Box<dyn std::error::Error>> {
    let args = App::new("grep0")
        .version("0.1.0")
        .about("ultra simple grep as an exercise")
        .arg(
            Arg::with_name("pattern")
                .help("the pattern to search for")
                .takes_value(true)
                .required(true),
        )
        .arg(
            Arg::with_name("filename")
                .help("input file name")
                .takes_value(true)
                .required(true),
        )
        .get_matches();

    let pattern = args
        .value_of("pattern")
        .expect("pattern argument is required");
    let re = Regex::new(pattern).expect("regular expression must be valid");

    let filename = args.value_of("filename").unwrap_or("-");
    let reader: Box<dyn BufRead> = if filename == "-" {
        let stdin = std::io::stdin();
        Box::new(stdin.lock())
    } else {
        let f = File::open(filename)?;
        Box::new(BufReader::new(f))
    };

    process_line(reader, re)
}

fn process_line<T: BufRead + Sized>(
    reader: T,
    re: Regex,
) -> Result<(), Box<dyn std::error::Error>> {
    for line in reader.lines() {
        let line = line?;
        if re.is_match(&line) {
            println!("{}", line);
        }
    }
    Ok(())
}

// cargo run -- "nar.*[zo]" nariz.txt 