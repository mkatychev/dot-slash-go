#!/usr/bin/env bash

ROOT_DIR="$(
  cd "$(dirname "$0")"
  pwd -P
)"

# shellcheck source=./.go/core/bash-cli.inc.sh
source "$ROOT_DIR/.go/core/bash-cli.inc.sh"
# Run the Bash CLI entrypoint
ROOT_DIR="$ROOT_DIR" bcli_entrypoint "$@"
