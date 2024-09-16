package utils

import (
	jwtservice "gofiber-boilerplate/modules/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetFiberJwtClaims(c *fiber.Ctx) jwtservice.JwtClaim {
	return c.Locals("user").(*jwt.Token).Claims.(jwtservice.JwtClaim)
}

func GetFiberJwtUserId(c *fiber.Ctx) (id uuid.UUID, err error) {
	idString, err := GetFiberJwtClaims(c).GetSubject()
	if err != nil {
		return
	}

	id, err = uuid.Parse(idString)
	return
}
