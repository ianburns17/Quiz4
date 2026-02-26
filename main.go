package main

import (
	"fmt"
	"net/smtp"
)

type EmailService struct{}

func (e *EmailService) Send(to string, message string) error {
	auth := smtp.PlainAuth("", "myemail@gmail.com", "password", "smtp.gmail.com")
	return smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"myemail@gmail.com",
		[]string{to},
		[]byte(message),
	)
}

type UserService struct {
	emailService EmailService // ❌ Concrete dependency
}

func (u *UserService) Register(email string) {
	fmt.Println("User registered:", email)
	u.emailService.Send(email, "Welcome!")
}

func main() {
	userService := UserService{}
	userService.Register("test@example.com")
}

/*

package main

import "fmt"

type EmailSender interface {
	Send(to string, message string) error
}

// Real implementation
type SMTPEmailService struct{}

func (s *SMTPEmailService) Send(to string, message string) error {
	fmt.Println("Sending real email to:", to)
	return nil
}

// Mock implementation (for testing)
type MockEmailService struct{}

func (m *MockEmailService) Send(to string, message string) error {
	fmt.Println("Mock email sent to:", to)
	return nil
}

type UserService struct {
	emailSender EmailSender // ✅ Depends on abstraction
}

// Constructor Injection
func NewUserService(sender EmailSender) *UserService {
	return &UserService{
		emailSender: sender,
	}
}

func (u *UserService) Register(email string) {
	fmt.Println("User registered:", email)
	u.emailSender.Send(email, "Welcome!")
}

func main() {
	// Swap implementations easily
	smtp := &SMTPEmailService{}
	userService := NewUserService(smtp)

	userService.Register("test@example.com")
}

*/
