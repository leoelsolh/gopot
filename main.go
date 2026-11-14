package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/gliderlabs/ssh"
)

func startServer() {
	hdlr := func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	}

	auth := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		ip := ctx.RemoteAddr()
		usr := ctx.User()
		pwd := password
		logAttempt(ip, usr, pwd)
		return false
	})

	log.Fatal(ssh.ListenAndServe(":2222", hdlr,
		ssh.HostKeyFile("host_key"),
		auth,
	))
}

func logAttempt(ip net.Addr, usr string, pwd string) {
	log.Printf(
		"%s  %s  '%s'",
		ip, usr, pwd,
	)
}

func loadLogFile() (*os.File, error) {
	f, err := os.OpenFile(
		"logfile.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644, // file permissions
	)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func main() {
	f, err := loadLogFile()
	if err != nil {
		log.Printf("Failed to open log file: %s", err)
		return
	}
	defer f.Close()

	log.SetOutput(f)
	log.SetFlags(log.Ltime)

	startServer()
}
