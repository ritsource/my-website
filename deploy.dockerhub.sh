# Git Sha for Versioning...
GIT_SHA=$(git rev-parse HEAD)

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
docker build -t ritwik310/my-website-api:latest -t ritwik310/my-website-api:$GIT_SHA  -f ./api/Dockerfile ./api
docker build -t ritwik310/my-website-nginx:latest -t ritwik310/my-website-nginx:$GIT_SHA  -f ./nginx/Dockerfile ./nginx
docker build -t ritwik310/my-website-console:latest -t ritwik310/my-website-console:$GIT_SHA  -f ./console/Dockerfile ./console
docker build -t ritwik310/my-website-renderer:latest -t ritwik310/my-website-renderer:$GIT_SHA  -f ./renderer/Dockerfile ./renderer

# Push to docker hub
docker push ritwik310/my-website-api:latest
docker push ritwik310/my-website-nginx:latest
docker push ritwik310/my-website-console:latest
docker push ritwik310/my-website-renderer:latest

docker push ritwik310/my-website-api:$GIT_SHA
docker push ritwik310/my-website-nginx:$GIT_SHA
docker push ritwik310/my-website-console:$GIT_SHA
docker push ritwik310/my-website-renderer:$GIT_SHA