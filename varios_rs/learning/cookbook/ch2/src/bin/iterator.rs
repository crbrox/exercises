fn main() {
    {
        let nums = 0..10;
        for num in nums {
            print!("{num} ");
        }
        println!();
    }
    {
        let names = vec!["Joe", "Miranda", "Alice"];
        let mut iter = names.iter();

        if let Some(name) = iter.next() {
            println!("first name: {name}");
        }
        if let Some(name) = iter.next() {
            println!("second name: {name}");
        }
        if let Some(name) = iter.next() {
            println!("third name: {name}");
        }
        if let Some(name) = iter.next() {
            println!("fourth name: {name}");
        } else {
            println!("no names left");
        }
    }
    {
        for (index, letter) in "abc".chars().enumerate() {
            println!("#{}. letter in the alphabet: {letter}", index + 1);
        }
    }
    {
        let mut alphabet = "ABCDEFGHIJKLMNOPQRSTUVWYZ".chars();

        let letter = alphabet.nth(2);
        println!("the third letter: {letter:?}");
        let current_first = alphabet.nth(0);
        println!("current 'first' letter: {current_first:?}");
        let current_first = alphabet.nth(0);
        println!("current 'first' letter: {current_first:?}");
        let last_letter = alphabet.last();
        println!("last letter: {last_letter:?}");
        // last() takes ownership of the receiver.
        // this will not work now.
        // let current_first = alphabet.nth(0);
        // println!("current 'first' letter: {current_first:?}");
    }
    {
        let nums: Vec<_> = (1..10).collect();
        println!("nums: {nums:?}");
        let nums: Vec<_> = (1..=10).collect();
        println!("nums: {nums:?}");
        let nums = (-3..=3).rev().collect::<Vec<_>>();
        println!("nums: {nums:?}");
    }
    {
        let all_nums = 1..;
        let nums = all_nums.take(5).collect::<Vec<_>>();
        println!("1.. take(5) nums: {nums:?}");

        let nums = (0..8).skip(2).take(3).collect::<Vec<_>>();
        println!("(0..8).skip(2).take(3) nums: {nums:?}");

        dbg!((0..).take_while(|x| x * x < 9).collect::<Vec<_>>());
    }
    {
        let names = vec!["Alfred", "Andy", "Jose", "luke"];
        let names: Vec<_> = names.iter().skip_while(|n| n.starts_with('A')).collect();
        println!("names: {names:?}");
    }
    {
        let countries = vec![
            "U.S.A", "Germany", "France", "Italy", "India", "Pakistan", "Burma",
        ];
        let countries_with_i: Vec<_> = countries.iter().filter(|c| c.contains('i')).collect();
        dbg!(countries_with_i);

        // Find the first element that satisfies a condition
        let country = countries.iter().find(|c| c.starts_with('I'));
        dbg!(country);

        // Find position instead
        let pos = countries.iter().position(|c| c.starts_with('I'));
        dbg!(pos);

        dbg!(countries.iter().any(|c| c.len() == 5));
        dbg!(countries.iter().all(|c| c.len() == 5));
    }
    {
        dbg!((1..11).sum::<i32>());
        dbg!((1..=3).product::<i32>());
        dbg!((1..11).max());
        dbg!((1..11).min());
        dbg!((1..1).min());
    }
    {
        dbg!((1..=2).cycle().take(5).collect::<Vec<_>>());
        dbg!((1..3).chain(4..6).chain(98..100).collect::<Vec<_>>());
        dbg!((1..5).zip((96..=100).rev()).collect::<Vec<_>>());
        dbg!(('A'..)
            .zip((1..=2).cycle())
            .take(3)
            .map(|(c, n)| format!("{c}-{n}"))
            .collect::<Vec<_>>());
    }
    {
        dbg!((1..=3).map(|n| n.to_string()).collect::<Vec<_>>());
        (1..=3).for_each(|n| println!("n+1 -> {}", n + 1));
        dbg!((1..15)
            .filter_map(|n| if n % 3 == 0 { Some(n * n) } else { None })
            .collect::<Vec<_>>());
    }
    {
        // not necessary currently.
        dbg!((b'A'..b'z'+1).map(|c| c as char).filter(|c| c.is_alphanumeric()).collect::<String>());
        // more concise.
        dbg!(('A'..='z').filter(|c| c.is_alphanumeric()).collect::<String>());
        dbg!(('A'..='z').filter(char::is_ascii_alphanumeric).collect::<String>());
    }

}
