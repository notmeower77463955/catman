#!/bin/bash

set -e

GOOS=linux GOARCH=amd64 go build -o catman
sudo mv catman /usr/local/bin/
sudo chmod +x /usr/local/bin/catman

echo "Installation complete"
