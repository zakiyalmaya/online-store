package cart

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application/cart"
	"github.com/zakiyalmaya/online-store/model"
	"github.com/zakiyalmaya/online-store/utils"
)

type Controller struct {
	cartSvc cart.Service
}

func NewCartController(cartSvc cart.Service) *Controller {
	return &Controller{cartSvc: cartSvc}
}

func (c *Controller) Add(ctx *fiber.Ctx) error {
	createRequest := &model.CreateCartRequest{}
	if err := ctx.BodyParser(createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	customerID := ctx.Locals("user_id").(int)
	createRequest.CustomerID = customerID

	if err := utils.Validator(createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	cart, err := c.cartSvc.Create(createRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(cart))
}

func (c *Controller) GetAll(ctx *fiber.Ctx) error {
	customerID := ctx.Locals("user_id").(int)
	status := ctx.Query("status")
	var statusInt *int
	if status != "" {
		statusParsed, err := strconv.Atoi(status)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid status"))
		}
		statusInt = &statusParsed
	}

	getRequest := &model.GetCartRequest{
		CustomerID: customerID,
		Status:     statusInt,
	}

	if err := utils.Validator(getRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	cart, err := c.cartSvc.GetByParams(getRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(cart))
}
	