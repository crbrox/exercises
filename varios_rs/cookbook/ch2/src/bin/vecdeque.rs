use std::collections::VecDeque;

fn main() {
    // A VecDequeue is best thought as a FIFO queue

    let mut orders = VecDeque::new();

    println!("un cliente ha pedido ostras");
    orders.push_back("ostras");

    println!("un cliente ha pedido pescaito");
    orders.push_back("pescaito");

    let prepared = orders.pop_front();
    if let Some(prepared) = prepared {
        println!("{} ¡listo!", prepared);
    }

    println!("un cliente ha pedido tabla de quesos");
    orders.push_back("tabla de quesos");

    let prepared = orders.pop_front();
    if let Some(prepared) = prepared {
        println!("{} ¡listo!", prepared);
    }

    println!("un cliente ha pedido tortilla de patata");
    orders.push_back("tortilla de patata");

    let prepared = orders.pop_front();
    if let Some(prepared) = prepared {
        println!("{} ¡listo!", prepared);
    }

    let prepared = orders.pop_front();
    if let Some(prepared) = prepared {
        println!("{} ¡listo!", prepared);
    }

    println!();

    // You can freely swtich your pushing
    // from front to back and viceversa
    let mut sentence = VecDeque::new();
    sentence.push_back("meneallo");
    sentence.push_front("es");
    sentence.push_back("amigo");
    sentence.push_front("Peor");
    sentence.push_back("Sancho");
    println!("sentence: {:?}", sentence);

    sentence.pop_front();
    sentence.push_front("Mejor");
    sentence.pop_back();
    sentence.push_back("Sánchez");
    println!("sentence: {:?}", sentence);

    println!();

    {
        let mut some_queue = VecDeque::with_capacity(5);
        some_queue.push_back("A");
        some_queue.push_back("B");
        some_queue.push_back("C");
        some_queue.push_back("D");
        some_queue.push_back("E");
        println!("some_queue {:?}", some_queue);

        some_queue.swap_remove_back(2);
        println!("some_queue.swap_remove_back(2) {:?}", some_queue);
    }
    {
        let mut other_queue = VecDeque::from(["A", "B", "C", "D", "E"]);
        println!("other_queue {:?}", other_queue);
        other_queue.swap_remove_front(2);
        println!("other_queue.swap_remove_front(2) {:?}", other_queue);
    }
}
