mod redis_client;
mod http_server;

use std::env;
use crate::redis_client::RedisClient;
use crate::http_server::spawn_http_server;

const REDIS_HOST: &str = "localhost";

#[tokio::main]
async fn main() {
    let redis_host = match env::var("redisHost") {
        Ok(host) => host,
        Err(_) => REDIS_HOST.to_string()
    };
    println!("Redis host: {}", redis_host);
    let redis = RedisClient::new(redis_host.as_str()).unwrap();
    spawn_http_server(redis).await.unwrap();
}

