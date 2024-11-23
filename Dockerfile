FROM golang:1.23.3-alpine

# Dockerize'ı yükle
RUN apk add --no-cache curl && \
    curl -sSL https://github.com/jwilder/dockerize/releases/download/v0.6.1/dockerize-linux-amd64-v0.6.1.tar.gz | tar -xzv -C /usr/local/bin

WORKDIR /app

# Go mod dosyasını kopyalıyoruz
COPY go.mod go.sum ./

# Bağımlılıkları indiriyoruz
RUN go mod tidy

# Uygulama dosyalarını kopyalıyoruz
COPY . .

# Uygulama dosyasını derliyoruz
RUN go build -o main .

# PostgreSQL ve MQTT'nin başlatılmasını bekliyoruz
CMD ["dockerize", "-wait", "tcp://postgres:5432", "-wait", "tcp://mqtt:1883", "-timeout", "30s", "/app/main"]
