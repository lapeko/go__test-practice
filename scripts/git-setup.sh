#!/usr/bin/env bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(realpath "$SCRIPT_DIR/..")"
cd "$REPO_ROOT"

git config core.hooksPath .github/hooks
chmod +x .github/hooks/commit-msg
