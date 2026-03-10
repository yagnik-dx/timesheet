#!/usr/bin/env bash
set -e

REPO="yagnik-dx/timesheet"
INSTALL_DIR="/usr/local/bin"
BINARY="timesheet"

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Pre-built binary only for Windows; Linux/macOS must build from source
if [[ "$OS" == *"mingw"* || "$OS" == *"cygwin"* || "$OS" == *"msys"* ]]; then
  FILE="timesheet.exe"
  INSTALL_DIR="/usr/bin"
  RAW_URL="https://raw.githubusercontent.com/$REPO/main/bin/$FILE"

  echo "⬇️  Downloading $BINARY from $RAW_URL ..."
  TMP_FILE=$(mktemp)
  curl -fsSL "$RAW_URL" -o "$TMP_FILE"

  echo "🚀 Installing to $INSTALL_DIR/$BINARY ..."
  mv "$TMP_FILE" "$INSTALL_DIR/$BINARY"

  echo "✅ $BINARY installed successfully!"
  echo "Run '$BINARY --help' to get started."
else
  echo "❌ Pre-built binaries are not provided for $OS/$ARCH."
  echo ""
  echo "Build from source:"
  echo "  git clone https://github.com/$REPO.git"
  echo "  cd timesheet"
  echo "  go build -o timesheet ."
  echo "  sudo mv timesheet $INSTALL_DIR/"
  echo ""
  echo "Or install via Go:"
  echo "  git clone https://github.com/$REPO.git && cd timesheet && go install ."
  exit 1
fi
