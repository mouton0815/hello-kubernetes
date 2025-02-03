mod redis_client;
mod http_server;

use std::env;
use crate::redis_client::RedisClient;
use crate::http_server::spawn_http_server;

#[tokio::main]
async fn main() {
    env_logger::init();

    let redis_host = get_env("redisHost", "localhost");
    let greeting_label = get_env("greetingLabel", "#greetingLabel#");

    let redis = RedisClient::new(redis_host.as_str(), "backend-rust-counter").unwrap();
    spawn_http_server(redis, greeting_label).await.unwrap();
}

fn get_env(name: &str, default: &str) -> String {
    match env::var(name) {
        Ok(value) => value,
        Err(_) => default.to_string()
    }
}
