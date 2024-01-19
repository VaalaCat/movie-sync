FROM node:20 AS frontend
WORKDIR /app
COPY . .
RUN npm install && npm run build

FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
COPY --from=frontend /app/out /app/out
RUN go build -o movie-sync-server

FROM alpine
COPY --from=builder /app/movie-sync-server /app/
WORKDIR /app
CMD ["/app/movie-sync-server"]