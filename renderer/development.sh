echo "Starting Rendering-Server in Development-mode..."

export DEV_MODE="true"
export API_URL="http://localhost:8080"

# Running Go app
go run *.go
