package services

import (
	"log"
	"time"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type LocalMailService struct{}

func NewLocalMailService() *LocalMailService {
	return &LocalMailService{}
}

func (s *LocalMailService) SendEmail(to string, subject string, body string) error {
	log.Printf("Local Mail Service: Sending email to: %s, subject: %s, body: %s\n", to, subject, body)

	// Simulate email sending delay
	time.Sleep(1 * time.Second)

	return nil
}
