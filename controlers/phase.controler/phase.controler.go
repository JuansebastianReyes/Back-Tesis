package phase_controler

import (
	"main/models"
	phaseService "main/services/phase.service"
	projectService "main/services/project.service"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gofiber/fiber/v2"
)

func CreatePhase(c *fiber.Ctx) error {
	projectId := c.Params("projectId")

	var data models.Phase
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}
	data.ID = primitive.NewObjectID()

	project, err := projectService.GetProjectById(projectId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	project.Phases = append(project.Phases, data.ID.Hex())
	err = projectService.UpdateProject(*project, projectId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	data.UpdateAt = time.Now()
	err = phaseService.CreatePhase(data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(data)
}

func GetPhaseById(c *fiber.Ctx) error {
	id := c.Params("id")

	phase, err := phaseService.GetPhaseById(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}
	return c.JSON(phase)
}

func GetPhaseListByProjectId(c *fiber.Ctx) error {
	projectId := c.Params("projectId")
	phases := []*models.Phase{}

	project, err := projectService.GetProjectById(projectId)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	for _, phaseId := range project.Phases {
		phase, err := phaseService.GetPhaseById(phaseId)
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(err)
		}
		phases = append(phases, phase)
	}

	return c.JSON(phases)
}

func GetPhaseList(c *fiber.Ctx) error {

	phases, err := phaseService.GetPhaseList()
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(phases)
}

func UpdatePhase(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.Phase
	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	data.ID = primitive.NewObjectID()
	data.UpdateAt = time.Now()
	err = phaseService.UpdatePhase(data, id)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(data)
}

func GetPhaseMembers(c *fiber.Ctx) error {
	id := c.Params("id")

	members, err := phaseService.GetPhaseMembers(id)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	return c.JSON(members)
}

func AddMemberPhase(c *fiber.Ctx) error {
	phaseId := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	err = phaseService.AddMemberPhase(phaseId, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(data)
}

func RemoveMemberPhase(c *fiber.Ctx) error {
	phaseId := c.Params("id")
	var data models.User

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	err = phaseService.RemoveMemberPhase(phaseId, data.ID.Hex())
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(err)
	}

	return c.JSON(data)
}
