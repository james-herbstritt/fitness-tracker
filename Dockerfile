# syntax=docker/dockerfile:1

FROM golang:1.20

ARG CLIENT_SECRET
ARG CLIENT_ID
ARG REDIRECT_URL
ARG AUTH_URL
ARG ACCREF_TOKEN_URL

ENV CLIENT_SECRET $CLIENT_SECRET
ENV CLIENT_ID $CLIENT_ID
ENV REDIRECT_URL $REDIRECT_URL
ENV AUTH_URL $AUTH_URL
ENV ACCREF_TOKEN_URL $ACCREF_TOKEN_URL

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080

# Run
CMD ["/docker-gs-ping"]
