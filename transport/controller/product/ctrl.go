package product

import (
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
	getRequest, err := getProductParam(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	products, err := c.productSvc.GetAll(getRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(products))
}

func (c *Controller) Create(ctx *fiber.Ctx) error {
	createRequest := &model.CreateProductRequest{}
	if err := ctx.BodyParser(createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	if err := utils.Validator(createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	err := c.productSvc.Create(createRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(nil))
}
