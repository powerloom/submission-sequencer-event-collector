#!/bin/bash

# Ensure Go environment is set up 🌱
echo "🔧 Setting up Go environment..."
export PATH=$PATH:$(go env GOPATH)/bin

echo ""

# Step 1: Build the project ⚙️
echo "🔨 Building the project..."
go build ./...
if [ $? -eq 0 ]; then
    echo "✅ Build successful!"
else
    echo "❌ Build failed!"
    exit 1
fi

echo ""

# Step 2: Run tests 🧪
echo "Running tests..."
go test ./...
if [ $? -eq 0 ]; then
    echo "✅ All tests passed!"
else
    echo "❌ Some tests failed!"
    exit 1
fi

echo "" 

# Step 3: Notify the committer that the process is complete 
echo "🎉 Project built and tests completed!"
