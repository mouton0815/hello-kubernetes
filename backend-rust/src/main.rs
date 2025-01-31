mod redis_client;

use std::env;
use std::sync::{Arc, Mutex};
use axum::extract::{Path,State};
use axum::http::StatusCode;
use axum::routing::get;
use axum::Router;
use crate::redis_client::RedisClient;

const REDIS_HOST: &str = "localhost";

// TODO: Remove all the unwraps!

type MutexSharedState = Arc<Mutex<RedisClient>>; // TODO: Name, maybe

#[tokio::main]
async fn main() {
    let redis_host = match env::var("redisHost") {
        Ok(host) => host,
        Err(_) => REDIS_HOST.to_string()
    };
    println!("Redis host: {}", redis_host);
    let redis = RedisClient::new(redis_host.as_str()).unwrap();
    spawn_http_server(Arc::new(Mutex::new(redis))).await;
}

async fn spawn_http_server(state: MutexSharedState) {
    println!("Spawn HTTP server");
    let app = Router::new()
        .route("/{*key}", get(rest_handler))
        .with_state(state);

    let listener = tokio::net::TcpListener::bind("0.0.0.0:8080").await.unwrap();
    axum::serve(listener, app).await.unwrap()
}

async fn rest_handler(State(state): State<MutexSharedState>, Path(path): Path<String>) -> Result<String, StatusCode> {
    let greeting_label = "#greetingLabel#"; // TODO: Move to env
    let mut guard = state.lock().unwrap();
    match (*guard).incr("foo") { // TODO: "backend-golang-counter"
        Ok(value) => Ok(format!("[RUSTIC] {} {} (call #{})\n", greeting_label, path, value).to_string()),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR)
    }
}
