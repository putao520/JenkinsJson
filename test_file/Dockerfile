FROM ubuntu
RUN mkdir /app
COPY ./target/release/hello-rust /app/
EXPOSE 8080
WORKDIR /app
CMD ["./hello-rust"]