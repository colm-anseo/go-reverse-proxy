package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

// default values
// optional env vars will override these values
var (
	localAddr = ":8124"
	localCert = ""
	localKey  = ""

	remoteAddr = "https://www.google.com:443"
	remoteCert = ""
	remoteName = ""
)

func setFromEnv(s *string, k string) {
	if v, ok := os.LookupEnv(k); ok {
		*s = v
	}
}

func main() {

	setFromEnv(&localAddr, "LOCAL_ADDR")
	setFromEnv(&localCert, "LOCAL_CERT_FILE")
	setFromEnv(&localKey, "LOCAL_KEY_FILE")

	setFromEnv(&remoteAddr, "REMOTE_ADDR")
	setFromEnv(&remoteName, "REMOTE_NAME")
	setFromEnv(&remoteCert, "REMOTE_CERT_FILE")

	rpUrl, err := url.Parse(remoteAddr)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Remote Address   :", remoteAddr)

	rp := httputil.NewSingleHostReverseProxy(rpUrl)

	var remoteTLSconfig *tls.Config

	if remoteCert != "" {
		remoteTLSconfig, err = createTLSConfigFromCertFile(remoteCert)
		if err != nil {
			log.Fatal("error creating remote TLS config: ", err)
		}

		if remoteName != "" {
			remoteTLSconfig.ServerName = remoteName
			log.Println("Remote Name      :", remoteName)
		}

		log.Println("Remote Trust     :", remoteCert)
	}

	if remoteTLSconfig != nil {
		rp.Transport = createHttpTransport(remoteTLSconfig)
		log.Println("Remote Transport :", "Custom (per remote name/trust settings)")
	} else {
		log.Println("Remote Transport :", "Default")
	}

	s := &http.Server{
		Addr:    localAddr,
		Handler: rp,
	}

	log.Println("Local Address    :", localAddr)

	if localCert != "" && localKey != "" {
		log.Println("Local TLS        : enabled")
		log.Println("Local Cert       :", localCert)
		log.Println("Local Key        :", localKey)
		log.Fatal(s.ListenAndServeTLS(localCert, localKey))
	} else {
		log.Println("Local TLS        : disabled")
		log.Fatal(s.ListenAndServe())
	}
}
