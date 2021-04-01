package api

import (
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"

	"fiber-service/domain/model"
)

func (application *application) GetUsers(ctx *fiber.Ctx) error {
	_, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	offset, err := getQueryOffset(ctx)
	if err != nil {
		return err
	}
	limit, err := getQueryLimit(ctx)
	if err != nil {
		return err
	}
	users, err := application.userService.GetUsers(offset, limit)
	if err != nil {
		return err
	}
	return ctx.JSON(users)
}

func (application *application) GetUserById(ctx *fiber.Ctx) error {
	_, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	user, err := application.userService.GetUser(userId)
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (application *application) DeleteUserById(ctx *fiber.Ctx) error {
	_, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	return application.userService.DeleteUser(userId)
}

func (application *application) GetUser(ctx *fiber.Ctx) error {
	session, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	user, err := application.userService.GetUser(getSessionUserId(session))
	if err != nil {
		return err
	}
	return ctx.JSON(user)
}

func (application *application) CreateUser(ctx *fiber.Ctx) error {
	user := model.User{}
	err := json.Unmarshal(ctx.Body(), &user)
	if err != nil {
		return err
	}
	err = application.userService.CreateUser(&user)
	if err != nil {
		return err
	}
	// return ctx.JSON(user)
	return ctx.SendString(strconv.FormatInt(user.ID, 10))
}

func (application *application) GetUserOrders(ctx *fiber.Ctx) error {
	_, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	offset, err := getQueryOffset(ctx)
	if err != nil {
		return err
	}
	limit, err := getQueryLimit(ctx)
	if err != nil {
		return err
	}
	userOrders, err := application.userService.GetUserOrders(offset, limit)
	if err != nil {
		return err
	}
	return ctx.JSON(userOrders)
}
