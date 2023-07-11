#!/bin/bash
#
# Build all the CLI tools for all the OS/ARCH combinations.

# Exit on any error
set -e

# Create the binaries directory
mkdir -p binaries

# Start
echo "Building all the CLI tools..."

# Get the last git tag as the version
version=$(git describe --tags $(git rev-list --tags --max-count=1))

# Main program
for os in darwin linux windows; do
  for arch in amd64 arm64; do
    echo "Building for $os/$arch..."
    if [ "$os" = "windows" ]; then
      GOOS=$os GOARCH=$arch go build -o binaries/wabeltools-${version}-$os-$arch.exe
    else
      GOOS=$os GOARCH=$arch go build -o binaries/wabeltools-${version}-$os-$arch
    fi
  done
done

# Done
echo "Done."
