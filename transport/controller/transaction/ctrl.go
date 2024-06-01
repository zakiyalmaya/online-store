package transaction

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/zakiyalmaya/online-store/application/transaction"
	"github.com/zakiyalmaya/online-store/model"
	"github.com/zakiyalmaya/online-store/utils"
)

type Controller struct {
	transactionSvc transaction.Service
}

func NewTransactionController(transactionSvc transaction.Service) *Controller {
	return &Controller{transactionSvc: transactionSvc}
}

func (t *Controller) Checkout(ctx *fiber.Ctx) error {
	transactionRequest := &model.TransactionRequest{}
	if err := ctx.BodyParser(transactionRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	customerID := ctx.Locals("user_id").(int)
	if customerID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid customer id"))
	}

	transactionRequest.CustomerID = customerID
	if err := utils.Validator(transactionRequest); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse(err.Error()))
	}

	transaction, err := t.transactionSvc.Checkout(transactionRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusCreated).JSON(model.HTTPSuccessResponse(transaction))
}

func (t *Controller) GetByID(ctx *fiber.Ctx) error {
	transactionID := ctx.Query("id")
	transactionIDInt, err := strconv.Atoi(transactionID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(model.HTTPErrorResponse("invalid transaction id"))
	}

	transaction, err := t.transactionSvc.GetByID(transactionIDInt)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(model.HTTPErrorResponse(err.Error()))
	}

	return ctx.Status(fiber.StatusOK).JSON(model.HTTPSuccessResponse(transaction))
}