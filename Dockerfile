FROM node:18-alpine

ENV NODE_ENV=DEV

COPY ["package.json", "package-lock.json*", "./"]

RUN npm install

RUN NPM RUN DEV

COPY . .

FROM golang:1.19

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /poorhouse

EXPOSE 3000

CMD ["/poorhouse"]
