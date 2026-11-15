package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"github.com/gliderlabs/ssh"
)

type Log struct {
	Time time.Time
	IP   net.Addr
	Usr  string
	Pwd  string
}

func loadLogFile(fileName string) (*os.File, error) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644, // file permissions
	)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func startServer() {
	hdlr := func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	}

	auth := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		logAttempt(ctx.RemoteAddr(), ctx.User(), password)
		return false
	})

	log.Fatal(ssh.ListenAndServe(":2222", hdlr,
		ssh.HostKeyFile("host_key"),
		auth,
	))
}

func logAttempt(ip net.Addr, usr string, pwd string) {
	e := Log{
		Time: time.Now(),
		IP:   ip,
		Usr:  usr,
		Pwd:  pwd,
	}
	b, _ := json.Marshal(e)
	log.Println(string(b))
}

func main() {
	f, err := loadLogFile("logfile.log")
	if err != nil {
		log.Printf("Failed to open log file: %s", err)
		return
	}
	defer f.Close()

	log.SetOutput(f) // Set log output to *os.File
	log.SetFlags(0)

	startServer()
}
