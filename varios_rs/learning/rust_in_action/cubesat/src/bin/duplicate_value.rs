// Often frowned upon, however, but it can be useful in a pinch
// std::clone::Clone
// std::marker::Copy

#[derive(Debug, Clone)]
struct CubeSat {
    #[allow(unused)]
    id: u64,
}

#[derive(Debug, Clone)]
enum StatusMessage {
    Ok,
}

fn check_status(sat: CubeSat) -> StatusMessage {
    println!("checking status for {sat:?}");
    StatusMessage::Ok
}

fn main() {
    let sat_a = CubeSat { id: 0 };
    let a_status = check_status(sat_a.clone());
    println!("a: {:?}", a_status);

    let a_status = check_status(sat_a.clone());
    println!("a: {:?}", a_status);

    // still own sat_a
    println!("sat_a {sat_a:?}");
}
