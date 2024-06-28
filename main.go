package main

import (
	"golang.org/x/crypto/ocsp"

	"crypto/x509"
	"fmt"
	"io"
	"os"
)

func main() {
	ocspResp, err := os.OpenFile("ocsp.resp", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	ocspBytes, err := io.ReadAll(ocspResp)
	if err != nil {
		fmt.Println(err)
		return
	}

	rootFile, err := os.OpenFile("root.cer", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	rootBytes, err := io.ReadAll(rootFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	rootCert, err := x509.ParseCertificate(rootBytes)
	if err != nil {
		fmt.Println(err)
		return
	}

	resp, err := ocsp.ParseResponseForCert(ocspBytes, nil, rootCert)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)
	fmt.Println(resp.Certificate.Raw)
}
