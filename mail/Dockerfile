# STAGE 1: Build the binary
FROM amazonlinux:2 AS builder

# 1. Install Go and build dependencies
# Amazon Linux 2 provides Go 1.21+ via amazon-linux-extras or yum
RUN yum update -y && \
    yum install -y golang gcc sqlite-devel tar gzip

# 2. Set up working directory
WORKDIR /app

# 3. Copy go.mod and go.sum and download dependencies
# Doing this before copying the whole source saves time with Docker caching
COPY go.mod go.sum ./
RUN go mod download

# 4. Copy the rest of the source code
COPY . .

# 5. Build the binary
# We keep CGO_ENABLED=1 for SQLite support
# This will now link against GLIBC 2.26 because of the base image
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 \
    go build -a -installsuffix cgo -o mail .

# STAGE 2: (Optional) Create a tiny runtime image
# If you just want the binary, you can stop at Stage 1
FROM amazonlinux:2
WORKDIR /app
COPY --from=builder /app/mail .

# Ensure the binary is executable
RUN chmod +x ./mail

ENTRYPOINT ["./mail"]