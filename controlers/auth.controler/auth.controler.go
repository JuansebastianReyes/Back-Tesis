package controler_auth

import (
	"fmt"
	"log"
	"main/models"
	"main/security"
	"main/util"
	"net/http"

	userService "main/services/user.service"

	"github.com/gofiber/fiber/v2"
)

func SignIn(ctx *fiber.Ctx) error {
	var input models.User
	err := ctx.BodyParser(&input)
	if err != nil {
		return ctx.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	input.Email = util.NormalizeEmail(input.Email)
	user, err := userService.ReadByEmail(input.Email)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}

	err = security.VerifyPassword(user.Password, input.Password)
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(util.ErrInvalidCredentials))
	}
	token, err := security.NewToken(user.ID.Hex())
	if err != nil {
		log.Printf("%s signin failed: %v\n", input.Email, err.Error())
		return ctx.
			Status(http.StatusUnauthorized).
			JSON(util.NewJError(err))
	}
	return ctx.
		Status(http.StatusOK).
		JSON(fiber.Map{
			"user":  user,
			"token": fmt.Sprintf("Bearer %s", token),
		})
}
