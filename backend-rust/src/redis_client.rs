use std::{thread, time};
use redis::{Client, Commands, Connection, RedisResult};

pub struct RedisClient {
    conn: Connection,
    key: &'static str
}

impl RedisClient {
    pub fn new(host: &str, key: &'static str) -> RedisResult<Self> {
        let address = format!("redis://{}", host);
        let client = Client::open(address.as_str())?;
        // Redis may not be up and running yet, so try at most three times to connect
        let mut attempts = 1;
        loop {
            match client.get_connection() {
                Ok(conn) => {
                    println!("Connection to Redis established - attempt {}", attempts);
                    return Ok(Self { conn, key })
                }
                Err(err) => {
                    println!("{} - attempt {}", err.to_string(), attempts); // TODO: Error log
                    if attempts >= 3 {
                        return Err(err)
                    }
                }
            }
            let duration = time::Duration::from_secs(3);
            thread::sleep(duration);
            attempts += 1;
        }
    }

    pub fn incr(&mut self) -> RedisResult<u32> {
        self.conn.incr(self.key, 1)
    }
}
