use std::sync::{Arc, Mutex};
use axum::extract::{Path, State};
use axum::http::StatusCode;
use axum::Router;
use axum::routing::get;
use log::info;
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
    info!("Spawn HTTP server ...");
    let state = Arc::new(Mutex::new(SharedState::new(redis, greeting_label)));
    let router = Router::new()
        .route("/{*_}", get(rest_handler))
        .with_state(state);

    let listener = TcpListener::bind("0.0.0.0:8080").await?;
    let handle = tokio::spawn(async move {
        axum::serve(listener, router).await.unwrap() // Panic accepted
    });
    tokio::task::yield_now().await; // Yield execution to allow the server task to start
    info!("HTTP Server is up and running");
    let _ = handle.await?;
    Ok(())
}

async fn rest_handler(State(state): State<MutexSharedState>, Path(path): Path<String>) -> Result<String, StatusCode> {
    let mut guard = state.lock().unwrap();
    match (*guard).redis.incr() {
        Ok(value) => {
            let msg = format!("[RUSTIC] {} {} (call #{})", (*guard).label, path, value);
            info!("{}", msg);
            Ok(msg)
        },
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR)
    }
}
