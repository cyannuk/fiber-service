package api

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"fiber-service/domain/model"
)

func (application *application) Login(ctx *fiber.Ctx) error {
	login := model.Login{}
	err := json.Unmarshal(ctx.Body(), &login)
	if err != nil {
		return err
	}
	user, err := application.userService.FindUser(login.Email)
	if err != nil {
		return err
	}
	if user.Password != login.Password {
		return fiber.ErrUnauthorized
	}
	s, err := application.sessionStore.Get(ctx)
	if err != nil {
		return err
	}
	s.Set("user_id", user.ID)
	return s.Save()
}

func (application *application) getSession(ctx *fiber.Ctx) (*session.Session, error) {
	s, err := application.sessionStore.Get(ctx)
	if err != nil {
		return nil, err
	}
	if s.Fresh() {
		_ = s.Destroy()
		return nil, fiber.ErrUnauthorized
	}
	return s, nil
}

func getSessionUserId(session *session.Session) int64 {
	return session.Get("user_id").(int64)
}
