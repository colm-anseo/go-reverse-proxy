#!/bin/bash

# example on how to pass in the various parameters

# note: google example does not work, try instead your favorite SPA web-app URL instead

go build && \
	LOCAL_ADDR=":9999" \
	LOCAL_CERT_FILE="" \
	LOCAL_KEY_FILE="" \
	REMOTE_ADDR="https://www.google.com" \
	REMOTE_NAME="" \
	REMOTE_CERT_FILE="" \
		./go-reverse-proxy

