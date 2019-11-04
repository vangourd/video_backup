package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {
	args := os.Args
	var config Config
	config = loadConfig()
	switch {
	case len(args) == 1:
		return
	case args[1] == "send_mail":
		sendMail(config)
	default:
		println("You have input an invalid command")
	}
}

type Config struct {
	Smtp struct {
		Server     string
		Port       int
		Sender     string
		Recipients []string
	}
}

func loadConfig() Config {
	var config Config
	bytes := loadFile()
	config = parseYaml(bytes)
	return config
}

func loadFile() []byte {
	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// Get configuration
func parseYaml(data []byte) Config {
	var conf Config
	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return conf
}

func sendMail(conf Config) {
	// Assemble string from config
	svrandport := fmt.Sprintf("%s:%d", conf.Smtp.Server, conf.Smtp.Port)

	conn, err := smtp.Dial(svrandport)
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Mail(conf.Smtp.Sender); err != nil {
		log.Fatal(err)
	}

	// FIX THIS BEFORE COMMITTING
	for _, recipient := range conf.Smtp.Recipients {
		if err := conn.Rcpt(recipient); err != nil {
			log.Fatal(err)
		}
	}

	wc, err := conn.Data()
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprintf(wc, "Did you get this email?")
	if err != nil {
		log.Fatal(err)
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Quit()
	if err != nil {
		log.Fatal(err)
	}
}
