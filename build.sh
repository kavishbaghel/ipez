#!/bin/bash

echo "Building ipwiz"
rm -f binaries/ip*
go build -o binaries/ipwiz
cp binaries/ipwiz /usr/local/bin
echo "Build completed"
