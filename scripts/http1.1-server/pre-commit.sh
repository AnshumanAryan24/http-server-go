#!/bin/sh
set -e

echo "=== Running Pre-Commit Hooks (Format & Build) ==="
cd http1.1-server/
echo "From directory:"
pwd

# 1. Check for compilation errors (Fails if the app cannot build)
echo "Checking build..."
go build -o myserver .
rm myserver

# 2. Check for formatting and import issues
# goimports -l lists files that deviate from the standard.
echo "Checking imports and formatting..."
BAD_FILES=$(go run golang.org/x/tools/cmd/goimports@latest -l .)

if [ -n "$BAD_FILES" ]; then
    echo "Format check failed! The following files need formatting or import optimization:"
    echo "$BAD_FILES"
    echo "Run: go run golang.org/x/tools/cmd/goimports@latest -w ."
    exit 1
fi

echo "All checks passed!"
exit 0
