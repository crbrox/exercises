use std::env;

fn main() {
    println!("Listing all env vars:");
    for (key, val) in env::vars() {
        println!("{key}: {val}");
    }

    println!();
    let key = "PORT";
    println!("Setting env var {key}");
    env::set_var(key, "8080");
    print_env_var(key);

    println!("Removing env var {key}");
    env::remove_var(key);
    print_env_var(key);

    println!();
    let redis_addr = env::var("REDIS_ADDR").unwrap_or("localhost:6379".into());
    println!("redis_addr: {redis_addr}");
}

fn print_env_var(key: &str) {
    match env::var(key) {
        Ok(val) => println!("{key}: {val}"),
        Err(e) => println!("Couldn't print env var {key}: {e}"),
    }
}

// See https://lib.rs/crates/dotenvy
// See https://lib.rs/crates/config
