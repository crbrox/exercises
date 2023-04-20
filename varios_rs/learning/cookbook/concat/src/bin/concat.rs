fn main() {
    by_moving();
    by_cloning();
    by_mutating();
    
}

fn by_moving() {
    let hello = "hello ".to_string();
    let world = "world!";

    let hello_world = hello + world; // `hello` moved due to usage in operator
    println!("{}", hello_world);
    // println!("{}", hello);
}

fn by_cloning() {
    let hello = "hello ".to_string();
    let world = "world!";

    let hello_world = hello.clone() + world;

    println!("{}", hello_world);
    println!("{}", hello);
}

fn by_mutating() {
    let mut hello = "hello ".to_string();
    let world = "world!";

    hello.push_str(world);

    println!("{}", hello);
}