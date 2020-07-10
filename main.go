package main

import (
	"log"
	"net/http"
	"os"
	"fmt"
)

// usage: <program> <ssl-directory> <cert-crt-filename> <cert-key-filename>
func main() {

	currentDir, _ := os.Getwd()
	fmt.Println("currentDir:" + currentDir)

	sslCertCrtName, sslCertKeyName := setCertPath()

	// create file server handler
	fs := http.FileServer(http.Dir(currentDir))

	// start HTTP server with `fs` as the default handler
	fmt.Println("server run with :9000 and :443")
	log.Fatal(http.ListenAndServeTLS(":443", sslCertCrtName, sslCertKeyName, fs))
	log.Fatal(http.ListenAndServe(":9000", fs))

}

func setCertPath() (string ,string) {
	args := os.Args[1:]
	homeDir, err := os.UserHomeDir()
	if err != nil {log.Fatal("error in getting home directory")}
	fmt.Println("home directory: " + homeDir)

	sslCertDir := homeDir + "/ssl-local-cert"
	sslCertCrtName := "test.iw.com.pem"
	sslCertKeyName := "test.iw.com-key.pem"

	if (len(args) >= 1) { sslCertDir = args[0] }
	if (len(args) >= 2) { sslCertCrtName = args[1] }
	if (len(args) >= 3) { sslCertKeyName = args[2] }

	if (string(sslCertCrtName[0]) != "/") { sslCertCrtName = "/" + sslCertCrtName }
	if (string(sslCertKeyName[0]) != "/") { sslCertKeyName = "/" + sslCertKeyName }
	sslCertCrtPath := sslCertDir + sslCertCrtName
	sslCertKeyPath := sslCertDir + sslCertKeyName

	fmt.Println(sslCertCrtPath)
	fmt.Println(sslCertKeyPath)

	return sslCertCrtPath, sslCertKeyPath
}
