# gopot 

Gopot is a honeypot server that logs every login attempt, storing logs and information about the attempt, and never returning a session no matter the password combination. 

### **Why**
I made this project to learn more about credential-stuffing bots, and what it's like being on the receiving end of one. 

Within the first hour of the V1 gopot being live, there were approximately 1500 logs from 19 unique ip addresses. Bots were already hammering it before id even finished testing it myself. Which just goes to show what opening a single port to the internet can do.

### **How**
Gopot uses the gliderlabs/ssh package to start an ssh server, logging every password attempt as one JSON object containing a timestamp, user, password, host ip, and the port:
```json
{"Time":"2026-01-02T10:00:10.447593743Z","Usr":"root","Pwd":"12345","Host":"98.174.59.179","Port":"47124"}
```

To spin up a honeypot operation of your own, go through the following:

### **Usage:**
clone the repo and build the binary:
```bash
git clone https://github.com/leoelsolh/gopot.git
cd gopot
go build -o honeypot
```
Then ensure you generate a host key (ed25519 recommended; RSA and ECDSA also work):
```bash 
ssh-keygen -t ed25519 -f host_key -N ""
```
After that the directory tree should resemble something like this, and you can run the server:
```
├── go.mod
├── main.go
├── README.md
├── host_key
├── host_key.pub
├── src
│   ├── logs.go
│   ├── scoreboard.go
│   └── server.go
└── logfile.log
```
to check it's working as expected, try it from another terminal, and type a couple passwords: 
```bash
ssh -p 2222 test@localhost
```
If you are running it on a separate server make sure the port is open and you run it though systemctl so that it stays alive and survives a reboot.

And I strongly recommend redirecting port `:22` to `:2222` so otherwise you'll be missed by a lot of scanners out there. ⚠ BUT ⚠ Don't do this before moving your real ssh port, otherwise you'll be locked out; edit `Port` in `sshd_config`, and test the new port works before proceeding with the above. 

### **Roadmap:**
My intention is to create a scoreboard of most used passwords. Updated and shipped to a website, maybe something like `gopot.site` So look out for that.

I'm also planning a report function, as in; 'give me a report on this ip'. A flag that will pull a requested ip along with all its attempt combinations. This would be along with other flags like chosen port, log file, etc.
For example:
```bash
./honeypot -p 2222 -f v2.log -r 107.172.87.210
```

There is already a `scoreBoard` function which counts the IP addresses that attempted the most authentications. But there is no architecture to extract and rank attempted passwords. *Yet*.