package user

import (
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/utils"

	"github.com/gofiber/fiber/v2"
)

type UserJwtMiddleware interface {
	jwt.JwtMiddleware
	IsAdmin(c *fiber.Ctx) error
}

type userMiddlewareImpl struct {
	jwtService jwt.JwtService
}

func NewUserJwtMiddleware(jwtService jwt.JwtService) UserJwtMiddleware {
	return &userMiddlewareImpl{
		jwtService: jwtService,
	}
}

// impl `UserJwtMiddleware` start

func (service *userMiddlewareImpl) IsAdmin(c *fiber.Ctx) error {
	if utils.GetFiberJwtClaims(c)["is_admin"] != true {
		return fiber.NewError(401, "Unauthenticated")
	}

	return c.Next()
}

// impl `UserJwtMiddleware` end

// impl `jwt.JwtMiddleware` start

func (service *userMiddlewareImpl) CanAccess(c *fiber.Ctx) error {
	return service.jwtService.CanAccess(c, jwtIssuer)
}
func (service *userMiddlewareImpl) CanRefresh(c *fiber.Ctx) error {
	return service.jwtService.CanRefresh(c, jwtIssuer)
}

// impl `jwt.JwtMiddleware` end
