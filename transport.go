package main

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"
)

func createHttpTransport(tc *tls.Config) http.RoundTripper {

	// see https://golang.org/src/net/http/transport.go?s=3607:10104#L42
	//
	// or if lines don't match up (different go version etc.) look for:
	//
	//	"var DefaultTransport RoundTripper"
	//
	// for the default values with http.DefaultTransport

	rt := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	// note: this is safer than rt := http.DefaultTransport
	// as the default transport may have been altered by any other
	// package you import (since it is a public global)

	// add in the caller's TLS config
	rt.TLSClientConfig = tc

	return rt
}
