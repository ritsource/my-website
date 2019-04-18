# Run Client inside a Docker Container

# ENVs for React App (Console) 
export REACT_APP_API_URL="http://localhost:4001"
export REACT_APP_WEBSITE_URL="http://localhost:4001"

# Build React Bundle
yarn install
yarn run build

# Start Express server
yarn run server