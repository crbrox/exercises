const N: usize = 10_000_000;

fn main() {
    let mut a = vec![true; N];
    for i in 2..N {
        if a[i] {
            let mut j = i;
            while i * j < N {
                a[i * j] = false;
                j += 1;
            }
        }
    }

    let mut count = 0;
    for i in 2..N {
        if a[i] {
           count += 1;
        }
    }
    println!("{} primes", count);
}
