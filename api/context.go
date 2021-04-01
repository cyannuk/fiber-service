package api

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"fiber-service/api/errors"
)

func getPathParameter(ctx *fiber.Ctx, name string) (int64, error) {
	parameter := ctx.Params(name)
	if len(parameter) > 0 {
		return strconv.ParseInt(parameter, 10, 64)
	}
	return 0, errors.ErrNoParameter
}

func getQueryParameter(ctx *fiber.Ctx, name string) (int64, error) {
	parameter := ctx.Query(name)
	if len(parameter) > 0 {
		return strconv.ParseInt(parameter, 10, 64)
	}
	return 0, errors.ErrNoParameter
}

func getQueryOffset(ctx *fiber.Ctx) (int64, error) {
	offset, err := getQueryParameter(ctx, "offset")
	if err != nil {
		if err != errors.ErrNoParameter {
			return 0, err
		}
		offset = 0
	} else {
		if offset < 0 {
			return 0, errors.ErrInvalidParameter
		}
	}
	return offset, nil
}

func getQueryLimit(ctx *fiber.Ctx) (int64, error) {
	limit, err := getQueryParameter(ctx, "limit")
	if err != nil {
		if err != errors.ErrNoParameter {
			return 0, err
		}
		limit = 50
	} else {
		if limit <= 0 {
			return 0, errors.ErrInvalidParameter
		}
	}
	return limit, nil
}

func getUserId(ctx *fiber.Ctx) (int64, error) {
	userId, err := getPathParameter(ctx, "user_id")
	if err != nil {
		return 0, err
	} else {
		if userId <= 0 {
			return 0, errors.ErrInvalidParameter
		}
	}
	return userId, nil
}

func getOrderId(ctx *fiber.Ctx) (int64, error) {
	orderId, err := getPathParameter(ctx, "order_id")
	if err != nil {
		return 0, err
	} else {
		if orderId <= 0 {
			return 0, errors.ErrInvalidParameter
		}
	}
	return orderId, nil
}
