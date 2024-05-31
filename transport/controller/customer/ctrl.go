package customer

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application/customer"
	"github.com/zakiyalmaya/online-store/model"
	"github.com/zakiyalmaya/online-store/utils"
)

type Controller struct {
	customerSvc customer.Service
}

func NewCategoryController(customerSvc customer.Service) *Controller {
	return &Controller{customerSvc: customerSvc}
}

func (c *Controller) Register(ctx *fiber.Ctx) error {
	customerRequest := &model.CustomerRequest{}
	if err := ctx.BodyParser(customerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	if err := utils.Validator(customerRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	err := c.customerSvc.Register(customerRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}	

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(nil))
}

func (c *Controller) Login(ctx *fiber.Ctx) error {
	authRequest := &model.AuthRequest{}
	if err := ctx.BodyParser(authRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	if err := utils.Validator(authRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	response, err := c.customerSvc.Login(authRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(response))
}

func (c *Controller) Logout(ctx *fiber.Ctx) error {
	username := ctx.Locals("username").(string)
	if err := c.customerSvc.Logout(username); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(nil))
}