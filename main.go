package main

import (
	"bytes"
	"log"

	packerSSH "github.com/hashicorp/packer-plugin-sdk/sdk-internals/communicator/ssh"
	"golang.org/x/crypto/ssh"
)

func main() {
	addr := "192.168.121.65:22"
	comm, err := packerSSH.New(addr, &packerSSH.Config{
		Connection: packerSSH.ConnectFunc("tcp", addr),
		SSHConfig: &ssh.ClientConfig{
			User:            "vagrant",
			Auth:            []ssh.AuthMethod{ssh.Password("vagrant")},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	err = comm.Upload("C:/packer-test.txt", bytes.NewReader([]byte("hello")), nil)
	if err != nil {
		log.Fatalf("failed to upload file: %v", err)
	}
}
