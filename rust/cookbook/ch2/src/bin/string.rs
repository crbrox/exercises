fn main() {
    let mut s = String::new();
    s.push('H');
    s.push('i');
    println!("s: {s}");

    let s = "Hello".to_string();
    dbg!(s);

    let s = String::from("Hola");
    dbg!(s);

    let s = "🦀 y ¡olé! 🏳️ ☯™";
    for c in s.chars() {
        if c.is_whitespace() {
            continue;
        }
        println!("<{c}> <{c:?}>");
    }
    dbg!(s);

    let mut s = "Hello".to_string();
    s.push_str(" world!");
    dbg!(s);

    let s = "ôÿàã y̆";
    for c in s.chars() {
        println!("<{c}> <{c:?}>");
    }
    dbg!(s);

    let (first, second) = "HelloThere".split_at(5);
    println!("first: {first}; second: {second}");

    let haiku = "
    Day after day, alone on the hill
    The man with the foolish grin is keeping perfectly still
    But nobody wants to know him
    They can see that he's just a fool
    And he never gives an answer
    ";
    for line in haiku.lines() {
        println!("*{}.", line);
    }

    for s in "Never;Give;Up".split(';') {
        println!("{s}");
    }

    let string = "::Hi::There::";
    println!("string to split {string}");
    // Añade cadenas vacias para los separadores al principio y al final
    let s: Vec<_> = string.split("::").collect();
    println!("split {s:?}");

    let s: Vec<_> = string.split_inclusive("::").collect();
    println!("split_inclusive {s:?}");

    let s: Vec<_> = string.split_terminator("::").collect();
    println!("split_terminator {s:?}");

    for s in "I'm2fast4you".split_inclusive(char::is_numeric) {
        print!("{s} ");
    }
    println!();

    for s in "It's not your fault, it's mine".splitn(3, char::is_whitespace) {
        println!("{s} ");
    }

    for c in "The Rust Programming Language 2nd ed.".matches(char::is_uppercase) {
        print!("{c}");
    }
    println!();

    let saying = "The early bird gets the worm";
    let starts_with_the = saying.starts_with("The");
    println!(r#"Does "{saying}" start with "The"? {starts_with_the} "#);
    let ends_with_worm = saying.ends_with("worm");
    println!(r#"Does "{saying}" end with "worm"? {ends_with_worm} "#);
    let contains_bird = saying.contains("bird");
    println!(r#"Does "{saying}" contain "bird"? {contains_bird} "#);

    let a_lot_of_whitespace = "       I love     spaaace      Sure!";
    let s: Vec<_> = a_lot_of_whitespace.split(' ').collect();
    println!("{s:?}");
    let s: Vec<_> = a_lot_of_whitespace.split_whitespace().collect();
    println!("{s:?}");

    let input = "     P3ngu1n\n\n";
    println!("trim: {:?}", input.trim());
    println!("trim start: {:?}", input.trim_start());
    println!("trim end: {:?}", input.trim_end());

    let num = "123".parse::<i32>();
    println!("{num:?}");
    let num = "123a".parse::<i32>();
    println!("{num:?}");
    let num: i32 = "123".parse().unwrap();
    println!("{num:?}");

    let s = "My dad is the best dad";
    println!("{}", s.replace("dad", "mom"));
    println!("{}", s.to_lowercase());
    println!("{}", s.to_uppercase());
}
