# TO build API
cd api
echo "Building API (Go)..."
go build
cd ..

# TO build Renderer
cd renderer
echo "Building App-Renderer (Go)..."
go build
cd ..

export CONSOLE_CLIENT_URL="http://localhost:3000"

# Run with Docker-Compose
docker-compose up --build