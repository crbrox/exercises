use std::sync;
use std::sync::mpsc;
use std::thread;

pub struct ThreadPool {
    workers: Vec<Worker>,
    sender: mpsc::Sender<Message>,
}

impl ThreadPool {
    /// Create a new thread pool.
    ///
    /// The size is the number of threads in the pool.
    ///
    /// # Panics
    ///
    /// The new function will panic if size is zero.
    pub fn new(size: usize) -> ThreadPool {
        assert!(size > 0);
        let (sender, receiver) = mpsc::channel();
        let receiver = sync::Arc::new(sync::Mutex::new(receiver));
        let mut workers = Vec::with_capacity(size);
        for i in 0..size {
            workers.push(Worker::new(i, sync::Arc::clone(&receiver)))
        }

        ThreadPool { workers, sender }
    }
    pub fn execute<F>(&self, f: F)
    where
        F: FnOnce() + Send + 'static,
    {
        let job = Box::new(f);
        self.sender.send(Message::NewJob(job)).unwrap();
    }

    // mala idea desacoplarlo del join.
    // mejor como en TRPL book, todo en drop, y llamar drop si
    // se quiere más control
    pub fn stop(&self) {
        panic!("esto es una mala idea");
        /*
        for _ in 0..self.workers.len() {
            self.sender.send(Message::Terminate).unwrap()
        }
        */
    }
}

struct Worker {
    id: usize,
    thread: Option<thread::JoinHandle<()>>,
}

impl Worker {
    fn new(id: usize, receiver: Receiver) -> Worker {
        let thread = Some(thread::spawn(move || loop {
            let msg = receiver.lock().unwrap().recv().unwrap();
            match msg {
                Message::NewJob(job) => {
                    println!("Worker {id} received a job");
                    job()
                }
                Message::Terminate => {
                    println!("Worker {id} received terminate message");
                    break;
                }
            }
        }));
        Worker { id, thread }
    }
}

type Job = Box<dyn FnOnce() + Send + 'static>;

type Receiver = sync::Arc<sync::Mutex<mpsc::Receiver<Message>>>;

impl Drop for ThreadPool {
    fn drop(&mut self) {
        for _ in 0..self.workers.len() {
            self.sender.send(Message::Terminate).unwrap()
        }
        for worker in &mut self.workers {
            println!("Shutting down worker {}", worker.id);
            if let Some(worker) = worker.thread.take() {
                worker.join().unwrap();
            }
        }
    }
}

enum Message {
    NewJob(Job),
    Terminate,
}
