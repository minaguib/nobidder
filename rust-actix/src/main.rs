use actix_web::{web, App, HttpResponse, HttpServer, Responder};
use openrtb2::BidRequest;

async fn bidder(br: web::Json<BidRequest>) -> impl Responder {
    HttpResponse::NoContent().finish()
}

async fn ping() -> impl Responder {
    HttpResponse::Ok().body("pong")
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
	println!("NOBIDDER RUST-ACTIX Running on http://127.0.0.1:8080/");
    HttpServer::new(|| {
        App::new()
            .route("/ping", web::get().to(ping))
            .service(
                web::resource("/bidder")
                    .app_data(web::JsonConfig::default().limit(16 * 1024))
                    .route(web::post().to(bidder)),
            )
    })
    //.workers(6)
    .bind(("127.0.0.1", 8080))?
    .run()
    .await
}
