#!/bin/bash

# This script runs on the remote server to unpack and run the built application.
# It may be merged into the "deploy.sh" script to reduce complecxity.

# Setup
export PATH=$PATH:/home/admin/.local/share/fnm
eval "`fnm env`"

# Extract the files
mkdir -p ./personal-website-build
tar -xzvf ./personal-website-build.tar.gz -C ./personal-website-build

# Move into the resulting folder
cd ./personal-website-build

# Install all packages
npm install

# Start the program with PM2, or restart if an older version is already running.
[ -z "$(pm2 id "npm run start" | grep -e "[[:digit:]]")" ] && pm2 start "npm run start" --watch
