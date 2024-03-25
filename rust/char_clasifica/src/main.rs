use std::collections::HashSet;

type CharFn = fn(&char)->bool;
fn main() {

    let functions: [(&str, CharFn, CharFn);4]= [
        ("whitespace", char::is_ascii_whitespace, |&x| char::is_whitespace(x)),
        ("digit", char::is_ascii_digit, |&x| x.is_digit(10)),
        ("alphanum", char::is_ascii_alphabetic, |&x| x.is_alphabetic()),
        ("control", char::is_ascii_control, |&x| x.is_control()),
    ];

    for (name, ascii_f, uni_f) in functions {
        let ascii_set: HashSet<char> = (0u8..255).map(|b| b as char).filter(ascii_f).collect();
        let uni_set: HashSet<char> = (0u8..255).map(|b| b as char).filter(uni_f).collect();
        let ascii_minu_uni = &uni_set -  &ascii_set;
        println!("{name} {diff:?}", name=name, diff = ascii_minu_uni);
    }
}
