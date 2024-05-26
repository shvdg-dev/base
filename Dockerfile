# ---- Build Node Front-end ----
FROM node:18.16.0 AS frontend

# Set working directory
WORKDIR /app

# Copy files
COPY . .

# Build and install node modules
RUN npm install
RUN npm run build

# ---- Build Go Back-end ----
FROM golang:1.22 AS backend

# Copy frontend artifacts from previous stage
WORKDIR /app

# Copy frontend artifacts from previous stage
COPY --from=frontend /app/public ./public

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy Go code
COPY . .

# Build Go application
RUN go build -o ./tmp/main.exe ./cmd/app/

# Expose the port on which the application runs
EXPOSE 3000

# Command to run the executable
CMD ["/app/tmp/main.exe"]