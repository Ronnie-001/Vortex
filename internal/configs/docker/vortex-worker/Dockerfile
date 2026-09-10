#------------------ VORTEX WORKER Dockerfile --------------------
FROM golang:1.26.5 AS build

# Each contianer has their own internal filesystem; set the 
# working directory to /app for INTERNAL in the container use.
WORKDIR /src

# Download all of the dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy over the source code.
COPY . ./

CMD ["go", "run", "/cmd/vortex-worker/main.go"]
