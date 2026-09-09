#------------------ VORTEXD Dockerfile --------------------
FROM golang:1.26.5 AS builder

# Each contianer has their own internal filesystem; set the 
# working directory to /app for INTERNAL in the container use.
WORKDIR /app

# Download all of the dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy over the source code.
COPY . ./

CMD ["/cmd/vortexd/main.go"]
