use std::time::Instant;
use rand::Rng;

struct Point(f64, f64);

fn distance(a: &Point, b: &Point) -> f64 {
    let (dx, dy) = (a.0 - b.0, a.1 - b.1);
    let sq = dx * dx + dy * dy;
    sq.sqrt()
    //return hypot (dx, dy);
}

fn main() {
    let mut rng = rand::thread_rng();

    //double d = atof (argv[2]);
    let d = 0.5;
    let mut cnt: usize = 0;
    const N: usize = 100_000; //= atoi (argv[1]);
    let mut v = Vec::with_capacity(N);
    {
        let start = Instant::now();
        for _ in 0..N {
            v.push(Point(rng.gen::<f64>(), rng.gen::<f64>()));
        }
        let time_taken = start.elapsed();
        println!("filling: {} millis", time_taken.as_millis());
    }

    {
        let start = Instant::now();
        {
            for (i, p1) in v.iter().enumerate() {
                for p2 in v[i+1..N].iter()  {
                    if distance(p1, p2) < d {
                        cnt += 1;
                    }
                }
            }
        }
        let time_taken = start.elapsed();
        println!("searching:backtrace {} millis \n", time_taken.as_millis());
    }

    println!("{} edges shorter than {}", cnt, d);
}
