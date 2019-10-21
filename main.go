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
	load_config()
	switch {
	case len(args) == 1:
		return
	case args[1] == "send_mail":
		send_mail()
	default:
		println("You have inputted an invalid command")
	}
}

type Config struct {
	smtp_server string
	smtp_port   int
}

func load_config() {
	b := load_file()
	yaml := parse_yaml(b)
	fmt.Printf("SMTP Server is: ", yaml)
}

func load_file() []byte {
	f, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return f
}

// Get configuration
func parse_yaml(data []byte) Config {
	conf := Config{}

	err := yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	return conf
}

func send_mail(conf Config) {
	conn, err := smtp.Dial("server:25")
	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Mail("from@email.net"); err != nil {
		log.Fatal(err)
	}

	if err := conn.Rcpt("to@email.net"); err != nil {
		log.Fatal(err)
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
