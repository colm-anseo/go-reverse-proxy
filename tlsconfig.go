package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// createTLSConfigFromCertFile generates a proper TLS config from the provide
// certificate file.
//
// Use this method and *NOT* the dreaded/insecure/lazy method of:
//
//	tls.Config{InsecureSkipVerify: true} // <-- DO NOT USE THIS!
//
func createTLSConfigFromCertFile(trustFile string) (*tls.Config, error) {

	var rootCAs *x509.CertPool

	if trustFile != "" {

		rootCAs = x509.NewCertPool()

		certs, err := ioutil.ReadFile(trustFile)
		if err != nil {
			return nil, fmt.Errorf("Failed to read trust cert '%s': %v", trustFile, err)
		}

		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			return nil, fmt.Errorf("Failed to add trust cert: %s", trustFile)
		}
	} else {
		var err error
		rootCAs, err = x509.SystemCertPool()
		if err != nil {
			return nil, fmt.Errorf("Failed to load system trust-cert pool:", err)
		}
	}

	return &tls.Config{RootCAs: rootCAs}, nil
}
