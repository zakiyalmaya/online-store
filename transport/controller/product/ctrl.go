package product

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application/product"
	"github.com/zakiyalmaya/online-store/model"
	"github.com/zakiyalmaya/online-store/utils"
)

type Controller struct {
	productSvc product.Service
}

func NewProductController(productSvc product.Service) *Controller {
	return &Controller{productSvc: productSvc}
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	categoryID := ctx.Query("category_id")
	var categoryIDInt *int
	if categoryID != "" {
		categoryIDParsed, err := strconv.Atoi(categoryID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid category id"))
		}

		categoryIDInt = &categoryIDParsed
	}

	products, err := c.productSvc.GetAll(categoryIDInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(products))
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	productRequest := &model.ProductRequest{}
	if err := ctx.BodyParser(productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	if err := utils.Validator(productRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	err := c.productSvc.Create(productRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(nil))
}