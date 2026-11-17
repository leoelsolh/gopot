package main

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func main() {
	f, err := loadLogFile("logfile.log")
	if err != nil {
		log.Printf("Failed to open log file: %s", err)
		return
	}
	defer f.Close()

	logs, err := parseLog("logfile.log")
	if err != nil {
		log.Printf("Failed to parse log file: %s", err)
		return
	}
	scoreBoard(logs)
	for _, l := range logs {
		fmt.Printf(
			"Host: %s  Port: %s  User: %s  Password: %s\n",
			l.Hst, l.Prt, l.Usr, l.Pwd,
		)
	}

	log.SetOutput(f) // Set log output to logfile
	log.SetFlags(0)  // No output flags

	fmt.Println("Server started on :2222 ...")
	startServer()
}

func startServer() {
	hdlr := func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	}

	auth := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		logAttempt(ctx.RemoteAddr().String(), ctx.User(), password)
		return false
	})

	log.Fatal(ssh.ListenAndServe(":2222", hdlr,
		ssh.HostKeyFile("host_key"),
		auth,
	))
}
