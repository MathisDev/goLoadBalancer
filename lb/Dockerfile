FROM golang:1.22

WORKDIR /home

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN apt-get update && apt-get install -y iputils-ping

COPY . .

CMD ["go","run","."]
