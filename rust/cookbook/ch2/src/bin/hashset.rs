use std::collections::HashSet;

fn main() {
    const DON_QUIJOTE: &str = "El Ingenioso Hidalgo Don Quijote de La Mancha";
    const MOMO: &str = "Momo";
    const MUNDO: &str = "El Mundo Perdido";
    const MARTE: &str = "Una Princesa de Marte";

    let mut books = HashSet::new();
    dbg!(books.insert(DON_QUIJOTE));
    dbg!(books.insert(MOMO));
    dbg!(books.insert(MUNDO));
    dbg!(books.insert(MARTE));
    println!();
    dbg!(books.insert(MUNDO));
    dbg!(books.insert(MARTE));
    println!();
    dbg!(books.contains(MUNDO));
    println!();
    dbg!(books.remove(MARTE));
    dbg!(books.contains(MARTE));
    dbg!(books.remove(MARTE));
    println!("\n-----");
    let uno_cinco: HashSet<_> = (1..=5).collect();
    let cinco_diez: HashSet<_> = (5..11).collect();
    let seis_ocho: HashSet<_> = (6..=8).collect();
    dbg!(&uno_cinco);
    dbg!(&cinco_diez);
    dbg!(&seis_ocho);

    println!();
    dbg!(uno_cinco.is_disjoint(&seis_ocho));
    dbg!(cinco_diez.is_disjoint(&seis_ocho));
    dbg!(cinco_diez.is_subset(&seis_ocho));
    dbg!(cinco_diez.is_superset(&seis_ocho));
    dbg!(seis_ocho.is_subset(&cinco_diez));
    dbg!(seis_ocho.is_superset(&cinco_diez));
    println!();
    dbg!(&uno_cinco.union(&cinco_diez));
    println!();
    let nueve_doce: HashSet<_> = (9..13).collect();
    dbg!(&cinco_diez.intersection(&nueve_doce));
    dbg!(&cinco_diez.symmetric_difference(&nueve_doce));
    dbg!(&nueve_doce.difference(&cinco_diez));
    dbg!(&cinco_diez.difference(&nueve_doce));
    println!();
}
