use anyhow::Result;
use log::info;
use r2d2::Pool;
use redis::{Client, Commands};

pub struct RedisClient {
    pool: Pool<Client>,
    key: &'static str
}

impl RedisClient {
    pub fn new(host: &str, key: &'static str) -> Result<Self> {
        let address = format!("redis://{}", host);
        let client = Client::open(address.as_str())?;
        let pool = Pool::builder().build(client)?;
        info!("Redis client created");
        return Ok(Self { pool, key })
    }

    pub fn incr(&mut self) -> Result<u32> {
        let mut conn = self.pool.get()?;
        Ok(conn.incr(self.key, 1)?)
    }
}
