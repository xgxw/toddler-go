#!/bin/bash

_term() {
echo "Caught SIGTERM signal!"
    kill -TERM "$child" 2>/dev/null
}

trap _term SIGTERM

if [ -f toddle.upload ]; then
    mv toddle.upload toddle
fi

./toddle server --config=config.yaml &
child=$!
wait "$child"
