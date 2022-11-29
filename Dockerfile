FROM golang:1.18-alpine AS build
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /src/bin/cake-api-service .


FROM alpine:3.9 AS cake-api-service
RUN apk add ca-certificates
COPY --from=build /src/bin/cake-api-service /bin/cake-api-service

EXPOSE 80
WORKDIR ~
CMD ["cake-api-service"]