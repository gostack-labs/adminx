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
	var mailConfig = configs.Get().Mail
	err = mail.Send(&mail.Options{
		MailHost: mailConfig.Host,
		MailPort: mailConfig.Port,
		MailUser: mailConfig.User,
		MailPass: mailConfig.Pass,
		MailTo:   email,
		Subject:  subject,
		Body:     body,
	})
	return err
}

func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	return vc.Store.Check(key, answer, false)
}

func (vc *VerifyCode) generateVerifyCode(key string) string {
	code := cast.ToString(gofakeit.Number(100000, 999999))
	err := vc.Store.Set(key, code)
	if err != nil {
		return ""
	}
	return code
}
