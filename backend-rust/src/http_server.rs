use std::sync::{Arc, Mutex};
use axum::extract::{Path, State};
use axum::http::StatusCode;
use axum::Router;
use axum::routing::get;
use tokio::net::TcpListener;
use crate::redis_client::RedisClient;

type MutexSharedState = Arc<Mutex<RedisClient>>;

pub async fn spawn_http_server(redis: RedisClient) -> Result<(), std::io::Error> {
    println!("Spawn HTTP server");
    let state = Arc::new(Mutex::new(redis));
    let app = Router::new()
        .route("/{*key}", get(rest_handler))
        .with_state(state);

    let listener = TcpListener::bind("0.0.0.0:8080").await?;
    axum::serve(listener, app).await?;
    Ok(())
}

async fn rest_handler(State(state): State<MutexSharedState>, Path(path): Path<String>) -> Result<String, StatusCode> {
    let greeting_label = "#greetingLabel#"; // TODO: Move to env
    let mut guard = state.lock().unwrap();
    match (*guard).incr("foo") { // TODO: "backend-golang-counter"
        Ok(value) => Ok(format!("[RUSTIC] {} {} (call #{})\n", greeting_label, path, value).to_string()),
        Err(_) => Err(StatusCode::INTERNAL_SERVER_ERROR)
    }
}
