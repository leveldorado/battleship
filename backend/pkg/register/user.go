package register

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"golang.org/x/crypto/bcrypt"
)

type CreateProfileForm struct {
	FirstName string
	LastName  string
	Password  string
	Email     string
}

type Service interface {
	Create(ctx context.Context, f CreateProfileForm) error
}

type Account struct {
	ID              string
	FirstName       string
	LastName        string
	Email           string
	EmailVerifiedAt time.Time
	CreatedAt       time.Time
}

type accountRepository interface {
	Create(ctx context.Context, doc Account) error
	GetByEmail(ctx context.Context, email string) (Account, error)
}

type Password struct {
	AccountID      string
	HashedPassword string
}

type passwordRepository interface {
	Save(ctx context.Context, doc Password) error
}

type mailSender interface {
	Send(toEmail, toName, subject, body string) error
}

type VerifyEmailCode struct {
	AccountID string
	Email     string
	Code      string
	SendAt    time.Time
}

type verifyEmailCodeRepository interface {
	Save(ctx context.Context, doc VerifyEmailCode) error
}

type service struct {
	accountRepo         accountRepository
	passwordRepo        passwordRepository
	mailSender          mailSender
	verifyEmailCodeRepo verifyEmailCodeRepository
	endpoint            string
}

type ErrDuplicateEmail struct{}

func (ErrDuplicateEmail) Error() string {
	return "duplicate email"
}

type ErrNotFound struct{}

func (ErrNotFound) Error() string {
	return "not found"
}

type ErrEmailIsNotVerified struct{}

func (s *service) Create(ctx context.Context, f CreateProfileForm) error {
	_, err := s.accountRepo.GetByEmail(ctx, f.Email)
	if err == nil {
		return ErrDuplicateEmail{}
	}
	if !errors.As(err, &ErrNotFound{}) {
		return fmt.Errorf(`failed to get account by email: [email: %s, err: %w]`, f.Email, err)
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(f.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf(`failed to generate password hash: [err: %w]`, err)
	}
	id := uuid.New().String()
	if err = s.passwordRepo.Save(ctx, Password{AccountID: id, HashedPassword: string(hash)}); err != nil {
		return fmt.Errorf(`failed to save password: [err: %w]`, err)
	}
	acc := Account{
		ID:        id,
		FirstName: f.FirstName,
		LastName:  f.LastName,
		Email:     f.Email,
		CreatedAt: time.Now().UTC(),
	}
	if err = s.accountRepo.Create(ctx, acc); err != nil {
		return fmt.Errorf(`failed to create account: [acc: %+v, err: %w]`, acc, err)
	}
	doc := VerifyEmailCode{
		AccountID: id,
		Email:     f.Email,
		Code:      uuid.New().String(),
		SendAt:    time.Now().UTC(),
	}
	if err = s.verifyEmailCodeRepo.Save(ctx, doc); err != nil {
		return fmt.Errorf(`failed to save veriy email code: [err: %w]`, err)
	}
	msg := fmt.Sprintf(`%s/api/v1/verify-email?email=%s&code=%s`, f.Email, doc.Code)
	err = s.mailSender.Send(f.Email, f.FirstName+" "+f.LastName, "Please verify email by following link in message", msg)
	if err != nil {
		return fmt.Errorf(`failed to send email: [err: %w]`, err)
	}
	return nil
}
