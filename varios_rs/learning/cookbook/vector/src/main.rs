fn main() {
    println!("Hello, world!");

    let fruits = vec!["apple", "tomato", "pear"];
    println!("fruits: {fruits:?}");

    let mut fruits = Vec::new();
    fruits.push("apple");
    fruits.push("tomato");
    fruits.push("pear");
    println!("fruits: {fruits:?}");

    let last = fruits.pop();
    if let Some(last) = last {
        println!("removed '{last}' from {fruits:?}");
    }

    fruits.insert(1, "grape");
    println!("fruits after insertion {fruits:?}");

    fruits.swap(0, 1);
    println!("fruits after swap {fruits:?}");

    let first = fruits.first();
    if let Some(first) = first {
        println!("first: {first} at {fruits:?}");
    }

    let last = fruits.last();
    if let Some(last) = last {
        println!("last: {last} at {fruits:?}");
    }

    if let Some(second) = fruits.get(1) {
        println!("second: {second} at {fruits:?}");
    }

    let second = fruits[1];
    println!("second: {second} at {fruits:?}");

    let bunch_of_zeros = [0; 5];
    dbg!(bunch_of_zeros);

    let mut nums = vec![1, 2, 3, 4, 5];
    nums.remove(1);
    dbg!(nums);

    let mut names = vec!["Adelfa", "Australia", "Laurel", "Abelia", "x", "Artrosis"];
    names.retain(|e| e.starts_with('A'));
    dbg!(&names);
    dbg!(names.contains(&"Antonio"));

    let mut nums = vec![
        1, 2, 2, 2, 3, 3, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 6, 7, 7, 8, 8, 8, 8, 9,
    ];
    nums.dedup();
    dbg!(nums);

    let mut nums = vec![1, 1, 2, 1, 2, 2, 2, 2, 1, 1, 1, 3, 3, 3, 1, 2];
    nums.dedup();
    dbg!(nums);

    let mut nums = vec![1, 3, 8, 5, 4, 4, 2, 3, 1, 9, 3, 8, 6];
    println!("unsorted nums:   {nums:?}");
    nums.sort();
    println!("sorted nums:     {nums:?}");
    nums.dedup();
    println!("dedup-ed nums:   {nums:?}");
    nums.reverse();
    println!("reverse-ed nums: {nums:?}");

    let mut alphabet = vec!['a', 'b', 'c', 'd', 'e', 'f'];
    for letter in alphabet.drain(1..=3) {
        print!("{letter} ");
    }
    println!();
    println!("alphabet after drain: {alphabet:?}");

    let mut nevera = vec!["cerveza", "mayonesa", "pepinillos"];
    println!("¿está vacía ({:?})?: {}", nevera, nevera.is_empty());
    nevera.clear();
    println!(
        "despues de clear, ¿está vacía ({:?})?: {}",
        nevera,
        nevera.is_empty()
    );

    let mut colors = vec!["red", "blue", "green", "yellow", "pink", "white"];
    println!("colors before split_off: {colors:?}");
    let mut half = colors.split_off(colors.len() / 2);
    println!("colors: {colors:?}");
    println!("half: {half:?}");
    colors.append(&mut half);
    println!("colors after append: {colors:?}");
    println!("half after append: {half:?}");

    let mut stuff = vec!["1", "2", "3", "4", "5", "6"];
    println!("original stuff {stuff:?}");
    let stuff_to_insert = vec!["a", "b", "c"];
    let removed_stuff: Vec<_> = stuff.splice(1..4, stuff_to_insert).collect();
    println!("spliced stuff {stuff:?}");
    println!("removed stuff {removed_stuff:?}");

    let mut large_vector: Vec<i32> = Vec::with_capacity(1_000_000);
    large_vector.push(1);
    println!("large vector len: {}", large_vector.len());
    println!("large vector capacity: {}", large_vector.capacity());
    large_vector.shrink_to_fit();
    println!("shirinked large vector len: {}", large_vector.len());
    println!(
        "shrinked large vector capacity: {}",
        large_vector.capacity()
    );

    // Remove some item, replacing it with the last.
    // It works in O(1).
    let mut nums = vec![1, 2, 3, 4];
    println!("nums: {nums:?}");
    nums.swap_remove(1);
    println!("nums: {nums:?}");
}
