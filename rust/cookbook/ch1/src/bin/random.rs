
fn main() {
    let random_num1 = rand::random::<i32>();
    println!("random_num1 == {random_num1}");

    let random_num2: i32 = rand::random();
    println!("random_num2 == {random_num2}");

    let random_char: char = rand::random();
    println!("random_char == {random_char}");

    let random_bool: bool = rand::random();
    println!("random_bool == {random_bool}");

    use rand::Rng;

    let mut rng = rand::thread_rng();
    if rng.gen() {
        println!("Este mensaje tiene un 50% de posibilidades de escribirse");
    }

    let random_num3 = rng.gen_range(0..=100);
    println!("random_num3 == {random_num3}");

    let random_float = rng.gen_range(0.0..=1.0);
    println!("random_float == {random_float}");

    use rand::distributions::{Alphanumeric, DistString};
    let string = Alphanumeric.sample_string(&mut rand::thread_rng(), 32);
    println!("string == {string}");
}
