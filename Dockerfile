FROM golang:latest

LABEL maintainer = "ishanak1602@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT 8080

RUN go build 

#Remove Source Files
RUN find . -name "*.go" -type f -delete


EXPOSE ${PORT}

#Run the App
CMD ["./GOREST"]