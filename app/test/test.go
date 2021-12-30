package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

//
func main()  {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	seedUrl := "https://" + "zt-express.com"
	resp, err := client.Get(seedUrl)
	fmt.Println(1233)
	defer resp.Body.Close()
	fmt.Println(1234)

	if err != nil {
		fmt.Println(2, seedUrl)
		return
	}

	if resp.TLS != nil && len(resp.TLS.PeerCertificates) > 0 {
		certInfo := resp.TLS.PeerCertificates[0]
		fmt.Println(certInfo.NotAfter, 111)
		return
	} else {
		fmt.Println(resp.TLS, 444)
	}
	fmt.Println(3333)
	return
}