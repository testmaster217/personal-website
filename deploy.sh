#!/bin/bash

# This script builds and packages the NextJS app before deploying
# it to a remote server.

# Setup. Read in parameters from "deploy-params.yml"
# which is a sibling of the repo folder.
# (Relative paths inside the YAML file are relative to this script.)
eval "$(cat ../deploy-params.yml | sed 's/: /=/g' | sed 's/^/export /')"

# Build the project.
npm run build

# Make a .tar.gz file with the needed files and folders.
tar -czvf ../personal-website-build.tar.gz .next/ node_modules/ public/ package.json package-lock.json next.config.ts

# SFTP the .tar.gz file onto the server.
# cd ..
sftp -i "${ssh_key_path}" $server_user@$server_host <<< 'put ../personal-website-build.tar.gz'

# SSH into the server to run the second half of the script.
cat ./deploy-remote.sh | ssh -i "${ssh_key_path}" $server_user@$server_host