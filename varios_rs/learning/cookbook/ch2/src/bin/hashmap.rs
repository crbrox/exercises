use std::collections::HashMap;

fn main() {
    let mut tv_ratings = HashMap::new();

    tv_ratings.insert("The IT Crowd", 8);
    tv_ratings.insert("13 Reasons Why", 6);
    tv_ratings.insert("House of Cards", 9);
    tv_ratings.insert("Mujeres Desesperadas", 12);

    // Does a key exists?
    let contains_tv_show = tv_ratings.contains_key("House of Cards");
    println!("Did we rate House of Cards? {contains_tv_show}");
    let contains_tv_show = tv_ratings.contains_key("Los Picapiedra");
    println!("Did we rate Los Picapiedra? {contains_tv_show}");

    // Access a value
    if let Some(rating) = tv_ratings.get("Mujeres Desesperadas") {
        println!("I rate Mujeres Desesperadas {rating} out of 10");
    }

    // If we insert a value twice, we overwrite it
    let old_rating = tv_ratings.insert("13 Reasons Why", 4);
    if let Some(old_rating) = old_rating {
        println!("old rating for '13 Reasons Why' was {old_rating}");
    }

    // Remove a key and its value
    let removed_value = tv_ratings.remove("The IT Crowd");
    if let Some(removed_value) = removed_value {
        println!("The removed series had a rating of {removed_value}");
    }

    // Iterating accesses all key and values
    println!("All ratings:");
    for (key, value) in &tv_ratings {
        println!("{key:<25}\t: {value:>3}");
    }
    println!();
    // We can iterate mutably
    println!("All ratings with 100 as a maximun:");
    for (key, value) in &mut tv_ratings {
        *value *= 10;
        println!("{key:<25}\t: {value:>4}");
    }
    println!();

    // Like with the other collections,
    // you can preallocate a size
    // to gain some performance
    let mut age = HashMap::with_capacity(10);
    age.insert("Dory", 9);
    age.insert("Nemo", 5);
    age.insert("Merlin", 10);
    age.insert("Bruce", 9);

    // Iterate over all *keys*
    println!("All names:");
    for name in age.keys() {
        println!("{}", name);
    }
    println!();

    // Iterate over all values
    println!("All values:");
    for age in age.values() {
        println!("{age}");
    }
    println!();

    // Iterate over all values and mutate them
    println!("All values in ten years:");
    for age in age.values_mut() {
        *age += 10;
    }
    for (name, value) in &age {
        println!("{name} will be {value} years old.")
    }

    // entry
    {
        let age_of_coral = age.entry("coral").or_insert(11);
        println!("age_of_coral: {}", age_of_coral);
    }
    let age_of_coral = age.entry("coral").or_insert(15);
    println!("age_of_coral: {}", age_of_coral);
}
