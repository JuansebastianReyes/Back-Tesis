package controler_user

import (
	"main/models"
	"main/security"
	"main/util"
	"net/http"
	"strings"
	"time"

	userService "main/services/user.service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/asaskevich/govalidator.v9"
)

func CreateUser(c *fiber.Ctx) error {
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(util.NewJError(err))
	}

	data.Email = util.NormalizeEmail(data.Email)
	if !govalidator.IsEmail(data.Email) {
		return c.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrInvalidEmail))
	}

	exists, err := userService.ReadByEmail(data.Email)
	if err == mongo.ErrNoDocuments {
		if strings.TrimSpace(data.Password) == "" {
			return c.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(util.ErrEmptyPassword))
		}
		data.Password, err = security.EncryptPassword(data.Password)
		if err != nil {
			return c.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		data.CreateAt = time.Now()
		data.UpdateAt = data.CreateAt
		err = userService.Create(data)
		if err != nil {
			return c.
				Status(http.StatusBadRequest).
				JSON(util.NewJError(err))
		}
		return c.
			Status(http.StatusCreated).
			JSON(data)
	}

	if exists != nil {
		err = util.ErrEmailAlreadyExists
	}

	return c.
		Status(http.StatusBadRequest).
		JSON(util.NewJError(err))

}

func GetUsers(c *fiber.Ctx) error {
	users, err := userService.Read()
	if err != nil {
		return err
	}
	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := userService.ReadById(id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := userService.ReadById(id)
	if err != nil {
		return err
	}

	data.Email = util.NormalizeEmail(data.Email)
	if !govalidator.IsEmail(data.Email) {
		return c.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrInvalidEmail))
	}

	data.Password = user.Password
	data.UpdateAt = time.Now()
	err = userService.Update(data, id)
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	err := userService.Delete(id)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}

func UpdateUserPassword(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := userService.ReadById(id)
	if err != nil {
		return err
	}

	if strings.TrimSpace(data.Password) == "" {
		return c.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(util.ErrEmptyPassword))
	}

	data.Password, err = security.EncryptPassword(data.Password)
	if err != nil {
		return c.
			Status(http.StatusBadRequest).
			JSON(util.NewJError(err))
	}

	user.Password = data.Password
	data.UpdateAt = time.Now()
	err = userService.Update(*user, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}

func UpdeteProjectList(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	user, err := userService.ReadById(id)
	if err != nil {
		return err
	}

	user.Projects = data.Projects
	data.UpdateAt = time.Now()
	err = userService.Update(*user, id)
	if err != nil {
		return err
	}

	return c.JSON(user)
}
