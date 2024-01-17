FROM alpine

COPY movie-sync-server-linux-amd64 /app/bin/movie-sync-server-linux-amd64

COPY dist /app/asset

WORKDIR /app/bin

CMD [ "/app/bin/movie-sync-server-linux-amd64" ]