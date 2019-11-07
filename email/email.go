// Package email provides simplified smpt wrapper for sending notifications
package email

import (
	"net/smtp"
	"log"
	"fmt"

	"video_backup/config"
)

// Notifier stores data and methods necessary to send mail
type Notifier struct {
	server 		string
	port		int
	recipients	[]string
	sender		string
}

// NewNotifier constructs an Email Notifier convenience wrapper
func NewNotifier(conf *config.Config) (*Notifier, error) {
	var ntfy Notifier
	ntfy.server = conf.SMTP.Server
	ntfy.port = conf.SMTP.Port
	ntfy.recipients = conf.SMTP.Recipients
	ntfy.sender = conf.SMTP.Sender
	return &ntfy, nil
}

// Send a quick email from a notifier
func (n *Notifier) Send(subject string, message string) (error) {

	svrandport := fmt.Sprintf("%s:%d", n.server, n.port)
	msg := fmt.Sprintf("From: %s\r\n" + 
		   "To: %s\r\n" +
		   "Subject: %s\r\n" + 
		   "Mime-Version: 1.0\r\n" +
		   "Content-Type: Text/HTML; charset=ISO-8859-1\r\n" + 
		   "Content-Transferr-Encoding: QUOTED-PRINTABLE\r\n\r\n" + 
			"<h1>%s</h1>\r\n" +
			"<strike>Strikethrought test</strike>\r\n", n.sender, n.recipients[0],subject,message)
		   

	conn, err := smtp.Dial(svrandport)
	if err != nil {
		log.Fatal(err)
		return err
	}

	if err := conn.Mail(n.sender); err != nil {
		log.Fatal(err)
		return err
	}

	for _, recipient := range n.recipients {
		if err := conn.Rcpt(recipient); err != nil {
			log.Fatal(err)
			return err
		}
	}

	wc, err := conn.Data()
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = fmt.Fprintf(wc, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = wc.Close()
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = conn.Quit()
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}