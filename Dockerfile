# Build the Go API
FROM golang:latest AS builder
ADD . /app
WORKDIR /app/server
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w" -a -o /main .

# Build the Vue application
FROM node:alpine AS node_builder
COPY --from=builder /app/client ./
RUN npm install
RUN npm run build

#FInal build for production
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
#When build is run on a vue file, the dist folder is created, copy it to web
COPY --from=node_builder /dist ./web
RUN chmod +x ./main
EXPOSE 8080
CMD ./main