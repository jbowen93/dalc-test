#!/bin/bash

# Start dalc-test
docker run --platform linux/amd64 --network=docker_localnet --env DEST_IP=$DEST_IP --env DEST_PORT=$DEST_PORT --env BLOCK_HEIGHT=$BLOCK_HEIGHT jbowen93/dalc-test:test
