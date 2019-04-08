#!/bin/bash

HOST=""
PORT=""

echo | openssl s_client -showcerts -servername "${HOST}" -connect "${HOST}:${PORT}"  2>/dev/null
