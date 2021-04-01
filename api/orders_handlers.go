package api

import (
	"github.com/gofiber/fiber/v2"
)

func (application *application) GetOrders(ctx *fiber.Ctx) error {
	session, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	if userId != getSessionUserId(session) {
		return fiber.ErrForbidden
	}
	offset, err := getQueryOffset(ctx)
	if err != nil {
		return err
	}
	limit, err := getQueryLimit(ctx)
	if err != nil {
		return err
	}
	orders, err := application.orderService.GetOrders(userId, offset, limit)
	if err != nil {
		return err
	}
	return ctx.JSON(orders)
}

func (application *application) GetOrder(ctx *fiber.Ctx) error {
	session, err := application.getSession(ctx)
	if err != nil {
		return err
	}
	userId, err := getUserId(ctx)
	if err != nil {
		return err
	}
	if userId != getSessionUserId(session) {
		return fiber.ErrForbidden
	}
	orderId, err := getOrderId(ctx)
	if err != nil {
		return err
	}
	order, err := application.orderService.GetOrder(orderId)
	if err != nil {
		return err
	}
	if order.UserID != userId {
		return fiber.ErrForbidden
	}
	return ctx.JSON(order)
}
