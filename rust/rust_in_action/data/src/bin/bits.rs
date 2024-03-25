fn main() {
    let a: u16 = 50115;
    let b: i16 = -15421;
    println!("a: {a:016b} {a:10}");
    println!("b: {b:016b} {b:10}");
    println!();

    let a: u16 = u16::from_be_bytes([0xFF, 0x80]);
    let b: i16 = i16::from_be_bytes([0xFF, 0x80]);
    println!("a: {a:016b} {a:6}"); // std::fmt::Binary
    println!("b: {b:016b} {b:6}");
    println!();

    let a: u16 = u16::from_le_bytes([0xFF, 0x80]);
    let b: i16 = i16::from_le_bytes([0xFF, 0x80]);
    println!("a: {a:016b} {a:6}"); // std::fmt::Binary
    println!("b: {b:016b} {b:6}");
    println!();

    let a: f32 = 42.42;
    let frankentype: u32 = unsafe { std::mem::transmute(a) };
    println!("frankentype {frankentype}");
    println!("frankentype {frankentype:032b}");
    let b: f32 = unsafe { std::mem::transmute(frankentype) };
    println!("a == b; {a} == {b} ({})", a == b);

    let a: f32 = 42.42;
    let bits: u32 = a.to_bits();
    println!("bits        {bits}");
    println!("bits        {bits:032b}");
    let b: f32 = f32::from_bits(bits);
    println!("a == b; {a} == {b} ({})", a == b);
}
