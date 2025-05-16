#!/bin/bash

# This script runs on the remote server to unpack and run the built application.
# It may be merged into the "deploy.sh" script to reduce complecxity.

# Extract the files
mkdir -p ./personal-website-build
tar -xzvf ./personal-website-build.tar.gz -C ./personal-website-build

# Restart the website's daemon.
sudo systemctl restart admin-collinvesel_me.service