package verifycode

import (
	"log"
	"sync"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gostack-labs/adminx/configs"
	"github.com/gostack-labs/adminx/internal/repository/redis"
	"github.com/gostack-labs/adminx/pkg/mail"
	"github.com/spf13/cast"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		cache, err := redis.New()
		if err != nil {
			log.Fatal("redis new err:", err)
		}
		internalVerifyCode = &VerifyCode{
			Store: &redisStore{cache: cache},
		}
	})
	return internalVerifyCode
}

func (vc *VerifyCode) SendEmail(email string) error {
	code := vc.generateVerifyCode(email)

	subject, body, err := newHTMLEmail(code)
	if err != nil {
		return err
	}
	err = mail.Send(&mail.Options{
		MailHost: configs.Config.Mail.Host,
		MailPort: configs.Config.Mail.Port,
		MailUser: configs.Config.Mail.User,
		MailPass: configs.Config.Mail.Pass,
		MailTo:   email,
		Subject:  subject,
		Body:     body,
	})
	return err
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := cast.ToString(gofakeit.Number(100000, 999999))
	vc.Store.Set(key, code)
	return code
}
