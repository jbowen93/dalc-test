#!/bin/bash

set -e 

if [ "$1" = 'dalc-test' ]; then
    exec ./"$@" "--"
fi

exec "$@"
