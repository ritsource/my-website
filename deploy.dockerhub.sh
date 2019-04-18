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

# Docker builds
docker build -t ritwik310/my-website-api -f ./api/Dockerfile ./api
docker build -t ritwik310/my-website-nginx -f ./nginx/Dockerfile ./nginx
docker build -t ritwik310/my-website-console -f ./console/Dockerfile ./console
docker build -t ritwik310/my-website-renderer -f ./renderer/Dockerfile ./renderer

# Push to docker hub
docker push ritwik310/my-website-api
docker push ritwik310/my-website-nginx
docker push ritwik310/my-website-console
docker push ritwik310/my-website-renderer