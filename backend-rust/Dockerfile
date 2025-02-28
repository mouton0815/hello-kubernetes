FROM rust:1.85

ENV redisHost=redis
ENV RUST_LOG=info

RUN mkdir /app
COPY Cargo.toml /app
COPY src /app/src

WORKDIR /app
RUN cargo build

CMD ["cargo", "run"]