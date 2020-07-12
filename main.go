package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// usage: <program> <ssl-directory> <cert-crt-filename> <cert-key-filename>
func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("error in getting home directory")
	}
	sslCertDirArg := flag.String("d", homeDir+"/ssl-local-cert", "the directory of SSL cert")
	sslCrtNameArg := flag.String("c", "test.iw.com.pem", "the filename of SSL cert")
	sslKeyNameArg := flag.String("k", "test.iw.com-key.pem", "the filename of SSL key")
	flag.Parse()

	if string((*sslCrtNameArg)[0]) != "/" {
		*sslCrtNameArg = "/" + *sslCrtNameArg
	}
	if string((*sslKeyNameArg)[0]) != "/" {
		*sslKeyNameArg = "/" + *sslKeyNameArg
	}
	sslCertCrtPath := *sslCertDirArg + *sslCrtNameArg
	sslCertKeyPath := *sslCertDirArg + *sslKeyNameArg

	fmt.Println(sslCertCrtPath)
	fmt.Println(sslCertKeyPath)

	// create file server handler serving current directory
	fs := http.FileServer(http.Dir("."))

	// start HTTP server with `fs` as the default handler
	fmt.Println("server run with :9000 and :443")
	log.Fatal(http.ListenAndServeTLS(":443", sslCertCrtPath, sslCertKeyPath, fs))
	log.Fatal(http.ListenAndServe(":9000", fs))

}
