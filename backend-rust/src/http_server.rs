use std::sync::{Arc, Mutex};
use axum::extract::{Path, State};
use axum::http::StatusCode;
use axum::Router;
use axum::routing::get;
use tokio::net::TcpListener;
use crate::redis_client::RedisClient;

struct SharedState {
    redis: RedisClient,
    label: String
}

impl SharedState {
    fn new(redis: RedisClient, greeting_label: String) -> Self {
        Self{ redis, label: greeting_label }
    }
}

type MutexSharedState = Arc<Mutex<SharedState>>;

pub async fn spawn_http_server(redis: RedisClient, greeting_label: String) -> Result<(), std::io::Error> {
    println!("Spawn HTTP server");
    let state = Arc::new(Mutex::new(SharedState::new(redis, greeting_label)));
    let app = Router::new()
        .route("/{*key}", get(rest_handler))
        .with_state(state);

    let listener = TcpListener::bind("0.0.0.0:8080").await?;
    axum::serve(listener, app).await?;
    Ok(())
}

async fn rest_handler(State(state): State<MutexSharedState>, Path(path): Path<String>) -> Result<String, StatusCode> {
    let mut guard = state.lock().unwrap();
    match (*guard).redis.incr() {
        Ok(value) => Ok(format!("[RUSTIC] {} {} (call #{})\n", (*guard).label, path, value)),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR)
    }
}
