use axum::{
    http::StatusCode,
    response::IntoResponse,
    routing::{get, post},
    Json, Router,
};
use openrtb2::BidRequest;
use std::net::SocketAddr;

async fn bidder(Json(br): Json<BidRequest>) -> impl IntoResponse {
    (StatusCode::NO_CONTENT, "")
}

async fn ping() -> impl IntoResponse {
    (StatusCode::OK, "pong")
}
#[tokio::main(worker_threads = 6)]
async fn main() {
    // initialize tracing
    tracing_subscriber::fmt::init();

    println!("NOBIDDER RUST-AXUM Running on http://127.0.0.1:8080/");
    let app = Router::new()
        .route("/bidder", post(bidder))
        .route("/ping", get(ping));

    let addr = SocketAddr::from(([127, 0, 0, 1], 8080));
    tracing::debug!("listening on {}", addr);
    axum::Server::bind(&addr)
        .serve(app.into_make_service())
        .await
        .unwrap();
}
