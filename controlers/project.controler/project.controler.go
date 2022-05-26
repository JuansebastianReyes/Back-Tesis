package project_controler

import (
	"main/models"
	projectServices "main/services/project.service"
	user_service "main/services/user.service"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateProject(c *fiber.Ctx) error {
	var data models.Project

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	data.CreateAt = time.Now()
	data.UpdateAt = time.Now()
	err = projectServices.CreateProject(data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	user, err := user_service.ReadById(data.LeaderId)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	user.Projects = append(user.Projects, data.ID.Hex())
	err = user_service.Update(*user, user.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(data)
}

func GetProjectList(c *fiber.Ctx) error {
	projects, err := projectServices.GetProjectList()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(projects)
}

func GetProyectListForLeader(c *fiber.Ctx) error {

	leaderId := c.Params("leaderId")
	projects, err := projectServices.GetProyectListForLeader(leaderId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(projects)
}

func GetProjectListForUser(c *fiber.Ctx) error {
	userId := c.Params("userId")
	projects, err := projectServices.GetProjectListForUser(userId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(projects)
}

func GetProjectById(c *fiber.Ctx) error {
	id := c.Params("id")

	project, err := projectServices.GetProjectById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.JSON(project)
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.Project

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	data.UpdateAt = time.Now()
	err = projectServices.UpdateProject(data, id)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	return c.JSON(data)
}

func GetProjectMembers(c *fiber.Ctx) error {
	id := c.Params("id")
	members, err := projectServices.GetProjectMembers(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(members)
}

func GetProjectPhases(c *fiber.Ctx) error {
	id := c.Params("id")
	phases, err := projectServices.GetProjectPhases(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.JSON(phases)
}

func AddMemberProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	err = projectServices.AddMemberProject(id, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	data.Projects = append(data.Projects, id)
	err = user_service.Update(data, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	return c.JSON(data)
}

func RemovedMemberProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	err = projectServices.RemovedMemberProject(id, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	//delete project from user
	for i := range data.Projects {
		if data.Projects[i] == id {
			data.Projects = append(data.Projects[:i], data.Projects[i+1:]...)
		}
	}
	err = user_service.Update(data, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	return c.JSON(data)
}

func GetProjectMembersNotInProject(c *fiber.Ctx) error {
	id := c.Params("id")
	members, err := projectServices.GetProjectMembersNotInProject(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(members)
}
