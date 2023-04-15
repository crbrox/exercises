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

    let filename = args
        .value_of("filename")
        .expect("filename argument is required");
    let f = File::open(filename)?;
    let bf = BufReader::new(f);

    for line in bf.lines() {
        let line = line?;
        if re.is_match(&line) {
            println!("{}", line);
        }
    }
    Ok(())
}
