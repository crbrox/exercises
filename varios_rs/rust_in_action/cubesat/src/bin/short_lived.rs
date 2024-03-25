#[derive(Debug)]
struct CubeSat {
    id: u64,
}

#[derive(Debug)]
struct Mailbox {
    messages: Vec<Message>,
}

#[derive(Debug, Default)]
struct Message {
    to: u64,
    content: String,
}

struct GroundStation;

impl Mailbox {
    fn post(&mut self, msg: Message) {
        self.messages.push(msg);
    }
    fn deliver(&mut self, recipient: &CubeSat) -> Option<Message> {
        for i in 0..self.messages.len() {
            if self.messages[i].to == recipient.id {
                let msg = self.messages.remove(i);
                return Some(msg);
            }
        }
        None
    }
}

impl GroundStation {
    fn connect(sat_id: u64) -> CubeSat {
        CubeSat { id: sat_id }
    }
    fn send(mailbox: &mut Mailbox, msg: Message) {
        mailbox.post(msg);
    }
    fn fetch_sat_ids() -> Vec<u64> {
        vec![1, 2, 3]
    }
}

impl CubeSat {
    fn recv(&self, mailbox: &mut Mailbox) -> Option<Message> {
        mailbox.deliver(&self)
    }
}

fn main() {
    let mut mailbox = Mailbox { messages: vec![] };
    let sat_ids = GroundStation::fetch_sat_ids();
    for sat_id in sat_ids {
        let sat = GroundStation::connect(sat_id);
        let msg = Message {
            to: sat.id,
            content: format!("hello to sat-{sat_id}!"),
        };
        GroundStation::send(&mut mailbox, msg);
    }

    let sat_ids = GroundStation::fetch_sat_ids();

    for &sat_id in &sat_ids {
        let sat = GroundStation::connect(sat_id);
        let msg = sat.recv(&mut mailbox).unwrap_or_default();
        println!("{sat:?} got  {}->{:#?}", msg.to, msg.content);
    }

    // no more messages
    for &sat_id in &sat_ids {
        let sat = GroundStation::connect(sat_id);
        let msg = sat.recv(&mut mailbox).unwrap_or_default();
        println!("{sat:?} got  {}->{:#?}", msg.to, msg.content);
    }
}
