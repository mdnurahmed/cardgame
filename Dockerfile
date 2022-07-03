FROM golang:alpine 
WORKDIR /cardgame
COPY . .
RUN apk add build-base
RUN go mod tidy
RUN go install .