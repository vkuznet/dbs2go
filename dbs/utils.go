package dbs

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"

	"github.com/vkuznet/x509proxy"
)

// Ckey represents DBS X509 key used by HttpClient
var Ckey string

// Ckey represents DBS X509 cert used by HttpClient
var Cert string

// Ckey represents DBS timeout used by HttpClient
var Timeout int

// client X509 certificates
func tlsCerts(key, cert string) ([]tls.Certificate, error) {
	uproxy := os.Getenv("X509_USER_PROXY")
	uckey := os.Getenv("X509_USER_KEY")
	ucert := os.Getenv("X509_USER_CERT")
	if key != "" {
		uckey = key
	}
	if cert != "" {
		ucert = cert
	}

	// check if /tmp/x509up_u$UID exists, if so setup X509_USER_PROXY env
	u, err := user.Current()
	if err == nil {
		fname := fmt.Sprintf("/tmp/x509up_u%s", u.Uid)
		if _, err := os.Stat(fname); err == nil {
			uproxy = fname
		}
	}

	if uproxy == "" && uckey == "" { // user doesn't have neither proxy or user certs
		return nil, nil
	}
	if uproxy != "" {
		// use local implementation of LoadX409KeyPair instead of tls one
		x509cert, err := x509proxy.LoadX509Proxy(uproxy)
		if err != nil {
			return nil, fmt.Errorf("failed to parse X509 proxy: %v", err)
		}
		certs := []tls.Certificate{x509cert}
		return certs, nil
	}
	x509cert, err := tls.LoadX509KeyPair(ucert, uckey)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user X509 certificate: %v", err)
	}
	certs := []tls.Certificate{x509cert}
	return certs, nil
}

// HttpClient is HTTP client for urlfetch server
func HttpClient(key, cert string, tout int) *http.Client {
	var certs []tls.Certificate
	var err error
	// get X509 certs
	certs, err = tlsCerts(key, cert)
	if err != nil {
		log.Fatal("ERROR ", err.Error())
	}
	timeout := time.Duration(tout) * time.Second
	if len(certs) == 0 {
		if tout > 0 {
			return &http.Client{Timeout: time.Duration(timeout)}
		}
		return &http.Client{}
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{Certificates: certs,
			InsecureSkipVerify: true},
	}
	if tout > 0 {
		return &http.Client{Transport: tr, Timeout: timeout}
	}
	return &http.Client{Transport: tr}
}

// helper function to perform HTTP GET request and return its data
func getData(rurl string) ([]byte, error) {
	var out []byte
	client := HttpClient(Ckey, Cert, Timeout)
	req, err := http.NewRequest("GET", rurl, nil)
	if err != nil {
		return out, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return out, err
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	return data, err
}
