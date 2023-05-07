macro_rules! multiply {
    // Edge case
    ($last:expr) => {
        $last
    };

    ($head:expr, $($tail:expr), +) => {
        // Recursive call
        $head *multiply!($($tail),+)
    };
}
fn main() {
    println!("Hello, world!");
    let val = multiply!(2,4,8);
    println!("2*4*8 = {}", val);
}
