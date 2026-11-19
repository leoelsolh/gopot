package src

import (
	"fmt"
	"io"
	"log"

	"github.com/gliderlabs/ssh"
)

func StartServer() {
	hdlr := func(s ssh.Session) {
		io.WriteString(s, fmt.Sprintf("Hello %s\n", s.User()))
	}

	auth := ssh.PasswordAuth(func(ctx ssh.Context, password string) bool {
		LogAttempt(ctx.RemoteAddr().String(), ctx.User(), password)
		return false
	})

	log.Fatal(ssh.ListenAndServe(":2222", hdlr,
		ssh.HostKeyFile("host_key"),
		auth,
	))
}
