#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"

echo "$root" "$workspace"

bubbledir="$workspace/src/github.com/bubblenet"
if [ ! -L "$bubbledir/bubble" ]; then
    mkdir -p "$bubbledir"
    cd "$bubbledir"
    ln -s ../../../../../. bubble
    cd "$root"
fi

echo "ln -s success."

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$bubbledir/bubble"
PWD="$bubbledir/bubble"

# Launch the arguments with the configured environment.
exec "$@"
