echo "Starting Server in Development-mode..."

# Exporting Environment Variables
# Is in development Mode
export DEV_MODE="true"

# Mongo URI for development
export MONGO_URI="mongodb://localhost:27017"

# MongoDB Database name
export DATABASE_NAME="dev_db"

# Console App base URL and Renderer too
export CONSOLE_CLIENT_URL="http://localhost:3000"
export APP_RENDERER_URL="http://localhost:8081"

# Other more secret env variables

# SESSION_KEY = Session key for authentication
# GOOGLE_CLIENT_ID = Client ID for Google Oauth login
# GOOGLE_CLIENT_SECRET = Google client secret

# ADMIN_EMAIL_A = Admin Email No.1
# ADMIN_EMAIL_B = Admin Email No.2

# Running the server
go run *.go