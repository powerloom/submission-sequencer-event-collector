#!/bin/bash

# Ensure Go environment is set up 
echo "🔧 Setting up Go environment..."
export PATH=$PATH:$(go env GOPATH)/bin

echo ""

# Step 1: Install swag if it's not already installed 
echo "🚨 Checking if swag is installed..."
if ! command -v swag &> /dev/null
then
    echo "❌ swag is not installed. Installing swag..."
    go install github.com/swaggo/swag/cmd/swag@latest
else
    echo "✅ swag is already installed!"
fi

echo "" 

# Step 2: Regenerate Swagger documentation 
echo "🔄 Regenerating Swagger documentation..."
swag init -g pkgs/service/api.go -o pkgs/service/docs
echo "📜 Swagger documentation has been generated!"

echo ""

# Step 3: Build the project 
echo "🔨 Building the project..."
go build ./...
echo "🏗️ Build complete!"

echo "" 

# Step 4: Notify the committer that the process is complete 
echo "🎉 Swagger documentation generated successfully and the project is built!"
