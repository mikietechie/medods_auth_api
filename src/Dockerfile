# Date Created		    28 Nov 2024
# Author				Mike Z
# Email				    mzinyoni7@outlook.com
# Status				Looking for a job!
# Description			A Golang Auth API

FROM golang:latest
WORKDIR /app
COPY . /app
RUN go mod download
EXPOSE 8000

# RUN go run internal/cmd/su/main.go
ENTRYPOINT [ "go", "run", "main.go" ]