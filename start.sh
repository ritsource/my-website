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

# Run with Docker-Compose
docker-compose up --build