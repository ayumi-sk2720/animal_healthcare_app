# Dockerfile
FROM golang:1.21

# ログに出力する時間をJSTにするため、タイムゾーンを設定
ENV TZ /usr/share/zoneinfo/Asia/Tokyo

ENV ROOT=/go/src/app
WORKDIR ${ROOT}

# ModuleモードをON
ENV GO111MODULE=on

COPY . .
EXPOSE 8080

# Airをインストールし、コンテナ起動時に実行する(Hot loader)
RUN go install github.com/cosmtrek/air@latest
# gorm | https://gorm.io/
RUN go get -u gorm.io/gorm
# sqlmigrate
RUN go install github.com/rubenv/sql-migrate/...@latest
CMD ["air"]