#!/usr/bin/env bash

docker run -it --rm -w /node -v "$(pwd):/node" ghcr.io/zeta-chain/gosec:2.21.4-zeta2 -exclude-generated -exclude-dir testutil ./...
