#!/bin/sh
set -e

# Just
INSTALL_DIR="${HOME}/bin"
mkdir -p "${INSTALL_DIR}"
tmp="$(mktemp)"
trap 'rm -f "$tmp"' EXIT
curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh -o "$tmp"
if [ -n "${JUST_TAG:-}" ]; then
  bash "$tmp" --to "${INSTALL_DIR}" --tag "${JUST_TAG}"
else
  bash "$tmp" --to "${INSTALL_DIR}"
fi
if [ -x "${INSTALL_DIR}/just" ] && ! command -v just >/dev/null 2>&1; then
  echo "Note: ${INSTALL_DIR} is not on your PATH. Add it to use 'just': export PATH=\"${INSTALL_DIR}:\$PATH\""
fi

# Golangci-lint
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(go env GOPATH)/bin" v2.11.4
