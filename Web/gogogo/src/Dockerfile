FROM golang:alpine AS builder

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.* .
RUN go mod tidy

COPY . .
RUN go build -o .

FROM alpine

WORKDIR /app

COPY --from=builder /app/gogogo .
COPY ./templates ./templates

COPY flag.txt .
RUN chmod 444 flag.txt
RUN mv flag.txt /flag_$(head -c 10 /dev/urandom | sha1sum | cut -d ' ' -f 1).txt

# make new user ctf
RUN adduser -D ctf
RUN chown -R ctf:ctf /app

USER ctf

EXPOSE 1337

CMD ["./gogogo"]