package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

type Log struct {
	Time time.Time
	Hst  string `json:"Host"`
	Prt  string `json:"Port"`
	Usr  string
	Pwd  string
}

func loadLogFile(fn string) (*os.File, error) {
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644, // file permissions
	)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func logAttempt(ip, usr, pwd string) {
	h, p, err := net.SplitHostPort(ip)
	if err != nil {
		log.Println("Count not split host:port string: %s", err)
	}
	e := Log{
		Time: time.Now(),
		Hst:  h,
		Prt:  p,
		Usr:  usr,
		Pwd:  pwd,
	}
	b, _ := json.Marshal(e)
	log.Println(string(b))
}

func parseLog(fn string) ([]Log, error) {
	f, err := os.Open(fn)
	if err != nil {
		log.Printf("Could not load file: %s", err)
		return nil, err
	}
	defer f.Close()

	s := bufio.NewScanner(f)

	var data []Log
	for s.Scan() {
		var line Log
		l := s.Bytes()
		err = json.Unmarshal(l, &line)
		if err != nil {
			log.Printf("Could not unpack line content: %s", err)
			return nil, err
		}
		data = append(data, line)
	}
	return data, nil
}

func scoreBoard(line []Log) {
	count := map[string]int{}
	for _, l := range line {
		count[l.Hst]++
	}
	fmt.Println(count)
}
