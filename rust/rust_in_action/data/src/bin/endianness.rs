fn main() {
    let bytes: [u8; 2] = [0x01, 0x00];

    let a: i16 = unsafe { std::mem::transmute(bytes) };
    let be = i16::from_be_bytes(bytes);
    let le = i16::from_le_bytes(bytes);
    let ne = i16::from_ne_bytes(bytes);
    println!("transmute      {a:10}");
    println!("big endian     {be:10}");
    println!("little endian  {le:10}");
    println!("native endian  {ne:10}");
    match ne {
        x if x == be => println!("It seems we are on 'big-endianer'"),
        x if x == le => println!("It seems we are on 'little-endianer'"),
        x => panic!("where does {x} come from???"),
    }
    println!();
}
