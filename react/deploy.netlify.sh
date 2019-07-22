#!/bin/bash
set -e

yarn run build

echo "/*    /index.html   200" >> ./build/_redirects

timestamp=$(date +%s)

netlify deploy -d ./build -m $timestamp -p