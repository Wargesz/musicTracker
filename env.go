package main

import (
	"log"
	"os"
	"strings"
)

var env = map[string]string{}

func loadEnv() {
	raw_content, err := os.ReadFile(".env")
	if err != nil {
		log.Fatal("cannot open .env file")
	}
	iter := strings.SplitSeq(string(raw_content), "\n")
	for line := range iter {
		if !strings.Contains(line, "=") {
			continue
		}
		params := strings.Split(line, "=")
		env[params[0]] = params[1]
	}
}
