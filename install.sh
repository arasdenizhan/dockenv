#!/usr/bin/env bash
set -e

REPO="arasdenizhan/dockenv"
BINARY="dockenv"

command -v curl >/dev/null 2>&1 || {
    echo "Error: curl is required but not installed."
    exit 1
}

OS="$(uname -s)"
ARCH="$(uname -m)"

case "$OS" in
    Linux) OS="linux" ;;
    Darwin) OS="darwin" ;;
    *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
    x86_64) ARCH="amd64" ;;
    aarch64|arm64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

URL="https://github.com/$REPO/releases/latest/download/${BINARY}-${OS}-${ARCH}"

echo "Downloading $BINARY for $OS/$ARCH..."
curl -fL "$URL" -o "$BINARY"

chmod +x "$BINARY"

# install dir logic
if [ -w "/usr/local/bin" ]; then
    INSTALL_DIR="/usr/local/bin"
else
    INSTALL_DIR="$HOME/.local/bin"
    mkdir -p "$INSTALL_DIR"
fi

mv "$BINARY" "$INSTALL_DIR/$BINARY"

echo ""
echo "dockenv installed to $INSTALL_DIR"

# PATH kontrol
if [[ ":$PATH:" != *":$INSTALL_DIR:"* ]]; then
    echo "Add this to your shell config:"
    echo "export PATH=\"$INSTALL_DIR:\$PATH\""
fi

echo ""
"$INSTALL_DIR/$BINARY" --help