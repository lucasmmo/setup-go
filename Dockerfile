FROM golang:1.17 as builder

LABEL mainteiner="Lucas Monteiro"

WORKDIR /build

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o app cmd/app/main.go


FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /build/app .

EXPOSE 3000

CMD [ "./app" ]