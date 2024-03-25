type Message = String;

#[derive(Debug)]
struct MailBox {
    messages: Vec<Message>,
}

#[derive(Debug)]
struct CubeSat {
    #[allow(unused)]
    id: u64,
    mailbox: MailBox,
}

impl CubeSat {
    fn recv(&mut self) -> Option<Message> {
        self.mailbox.messages.pop()
    }
}

struct GroundStation;

impl GroundStation {
    fn send(to: &mut CubeSat, msg: Message) {
        to.mailbox.messages.push(msg);
    }
}

fn main() {
    println!("Hello, world!");

    // let base = GroundStation {};

    let mut sat_a = CubeSat {
        id: 11,
        mailbox: MailBox {
            messages: Vec::new(),
        },
    };
    println!("t0: {:#?}", sat_a);

    GroundStation::send(&mut sat_a, Message::from("Hello there!"));
    println!("t1: {:#?}", sat_a);

    let msg = sat_a.recv().unwrap_or("<empty msg>".to_string());
    println!("\n>>>>>>>>  msg -> {msg}\n");
    println!("t2: {:#?}", sat_a);
}
