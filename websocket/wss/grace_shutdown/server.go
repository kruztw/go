package main

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket" //这里使用的是 gorilla 的 websocket 库
)

func generateCert(host string) error {
	priv, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		return fmt.Errorf("Failed to generate private key: %v", err)
	}

	keyUsage := x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment

	notBefore := time.Now()
	notAfter := notBefore.Add(365 * 24 * time.Hour)

	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return fmt.Errorf("Failed to generate serial number: %v", err)
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"kruztw.inc"},
		},
		NotBefore: notBefore,
		NotAfter:  notAfter,

		KeyUsage:              keyUsage,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	if ip := net.ParseIP(host); ip != nil {
		template.IPAddresses = append(template.IPAddresses, ip)
	} else {
		template.DNSNames = append(template.DNSNames, host)
	}

	derBytes, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return fmt.Errorf("Failed to create certificate: %v", err)
	}

	certpemfd, err := os.OpenFile("cert.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Failed to open key.pem for writing: %v", err)
	}

	if err := pem.Encode(certpemfd, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes}); err != nil {
		return fmt.Errorf("Failed to write data to cert.pem: %v", err)
	}

	if err := certpemfd.Close(); err != nil {
		return fmt.Errorf("Error closing cert.pem: %v", err)
	}

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return fmt.Errorf("Unable to marshal private key: %v", err)
	}

	keypemfd, err := os.OpenFile("key.pem", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("Failed to open key.pem for writing: %v", err)
	}

	if err := pem.Encode(keypemfd, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes}); err != nil {
		return fmt.Errorf("Failed to write data to key.pem: %v", err)
	}

	if err := keypemfd.Close(); err != nil {
		return fmt.Errorf("Error closing key.pem: %v", err)
	}

	return nil
}

func main() {
	if err := generateCert("localhost"); err != nil {
		return
	}

	upgrader := websocket.Upgrader{}
	srv := &http.Server{Addr: "127.0.0.1:8888"}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		connect, err := upgrader.Upgrade(writer, request, nil)
		if nil != err {
			log.Println(err)
			return
		}

		defer connect.Close()

		go tickWriter(connect)

		for {
			messageType, messageData, err := connect.ReadMessage()
			if nil != err {
				log.Println(err)
				break
			}
			switch messageType {
			case websocket.TextMessage:
				fmt.Println(string(messageData))
			case websocket.BinaryMessage:
				fmt.Println(messageData)
			case websocket.CloseMessage:
			case websocket.PingMessage:
			case websocket.PongMessage:
			default:

			}
		}
	})

	go func() {
		err := srv.ListenAndServeTLS("cert.pem", "key.pem")
		if nil != err {
			log.Println(err)
			return
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	<-quit
	fmt.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server Shutdown: %v", err)
		return
	}
}

func tickWriter(connect *websocket.Conn) {
	for {
		err := connect.WriteMessage(websocket.TextMessage, []byte("from server to client"))
		if nil != err {
			log.Println(err)
			break
		}

		time.Sleep(time.Second)
	}
}
