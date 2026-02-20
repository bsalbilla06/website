#!/bin/sh
set -e

# Detect OS and Architecture
OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture names
case $ARCH in
    x86_64) ARCH="x64" ;;
    aarch64) ARCH="arm64" ;;
    arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# Map OS names
case $OS in
    linux) 
        OS="linux"
        # Check for musl (common on Alpine)
        if ldd /bin/ls 2>&1 | grep -q musl; then
            ARCH="${ARCH}-musl"
        fi
        ;;
    darwin) OS="macos" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

# Construct binary name
BINARY_NAME="tailwindcss-$OS-$ARCH"
DOWNLOAD_URL="https://github.com/tailwindlabs/tailwindcss/releases/latest/download/$BINARY_NAME"

echo "Downloading Tailwind CSS standalone binary from $DOWNLOAD_URL..."
curl -sLO "$DOWNLOAD_URL"
chmod +x "$BINARY_NAME"
mv "$BINARY_NAME" /usr/local/bin/tailwindcss
echo "Tailwind CSS installed successfully."
