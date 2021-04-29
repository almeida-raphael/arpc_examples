package utils

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
)

// LoadCA receives a path to a Certificate Authority Chain and returns a x509 certificate
func LoadCA(caChainFiles ...string) (*x509.CertPool, error) {
	certs := x509.NewCertPool()

	for _, chainFilePath := range caChainFiles {
		pemData, err := ioutil.ReadFile(chainFilePath)
		if err != nil {
			return nil, err
		}
		if !certs.AppendCertsFromPEM(pemData) {
			return nil, fmt.Errorf("cannot load chain certificate")
		}
	}

	return certs, nil
}

func LoadCertificates(certFilePath, pemFilePath string)(*tls.Certificate, error){
	crtFile, err := ioutil.ReadFile(certFilePath)
	if err != nil {
		return nil, err
	}

	pemFile, err := ioutil.ReadFile(pemFilePath)
	if err != nil {
		return nil, err
	}

	certificates, err := tls.X509KeyPair(crtFile, pemFile)
	if err != nil {
		return nil, err
	}

	return &certificates, nil
}