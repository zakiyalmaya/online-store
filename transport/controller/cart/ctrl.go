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

func (c *Controller) Create(ctx *fiber.Ctx) error {
	createRequest := &model.CreateCartRequest{}
	if err := ctx.BodyParser(createRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	customerID := ctx.Locals("user_id").(int)
	if customerID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid customer id"))
	}
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
	if customerID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid customer id"))
	}

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

func (c *Controller) Delete(ctx *fiber.Ctx) error {
	customerID := ctx.Locals("user_id").(int)
	if customerID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid customer id"))
	}

	cartItemID := ctx.Params("cart_item_id")
	cartItemIDInt, err := strconv.Atoi(cartItemID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid cart item id"))
	}

	if err := c.cartSvc.Delete(&model.DeleteCartRequest{
		CustomerID: customerID,
		CartItemID: cartItemIDInt,
	}); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(nil))
}
