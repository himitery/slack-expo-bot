FROM golang

WORKDIR /app
COPY . ./app

RUN go mod download
RUN go build main.go

EXPOSE 3000

CMD ["go", "./main"]