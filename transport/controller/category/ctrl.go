package category

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application/category"
	"github.com/zakiyalmaya/online-store/model"
	"github.com/zakiyalmaya/online-store/utils"
)

type Controller struct {
	categorySvc category.Service
}

func NewCategoryController(categorySvc category.Service) *Controller {
	return &Controller{categorySvc: categorySvc}
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	categoryRequest := &model.CategoryRequest{}
	if err := ctx.BodyParser(categoryRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	if err := utils.Validator(categoryRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	err := c.categorySvc.Create(categoryRequest.Name)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(nil))
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	categories, err := c.categorySvc.GetAll()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(categories))
}