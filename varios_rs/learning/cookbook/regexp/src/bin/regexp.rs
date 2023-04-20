use regex::{Regex, RegexBuilder};

fn main() {
    println!("Hello, world!");

    let date_regex = Regex::new(r"^\d{2}.\d{2}.\d{4}$").expect("date regex");
    let date = "15.10.2017";
    let is_date = date_regex.is_match(date);
    println!("Is '{date}' a date? {is_date}");

    let date_regex = Regex::new(r"(\d{2}).(\d{2}).(\d{4})").expect("date regex");
    let text_with_dates = "Alan Turing was born on 23.06.1912 and \
    died on 07.06.1954. \
    A movie about his life called 'The Imitation Game' 
    came out on \
    14.11.2017";
    for cap in date_regex.captures_iter(text_with_dates) {
        println!("Found date {}", &cap[0]);
        // Notice that the year is in the capture group indexed at 1.
        // This is because the entire match is stored in the capture group at index 0.
        println!("Year: {}, Month: {}, Day: {}", &cap[3], &cap[2], &cap[1]);
    }

    println!("Original text:\n{}\n", text_with_dates);
    let text_indian_dates = date_regex.replace_all(text_with_dates, "$1-$2-$3");
    println!("In indian format:\n{}\n", text_indian_dates);

    // ?P<somename> gives a capture group a name.
    let date_regex = Regex::new(r"(?P<day>\d{2}).(?P<month>\d{2}).(?P<year>\d{4})").unwrap();
    let text_with_american_dates = date_regex.replace_all(text_with_dates, "$month/$day/$year");
    println!("In american format:\n{}\n", text_with_american_dates);

    let rust_regex = Regex::new(r"(?i)rust").unwrap();
    println!("Do we match RuSt? {}", rust_regex.is_match("RuSt?"));

    // It is considered good style to use raw strings in *every* regex
    // even when it doesn't have any backslashes.  
    let rust_regex = RegexBuilder::new(r"rust").case_insensitive(true).build().unwrap();
    println!("Do we match RuSt? {}", rust_regex.is_match("RuSt?"));

}
