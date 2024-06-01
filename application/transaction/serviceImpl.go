package transaction

import (
	"fmt"

	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
	"github.com/zakiyalmaya/online-store/infrastructure/repository"
	"github.com/zakiyalmaya/online-store/model"
)

type transactionSvcImpl struct {
	repos *repository.Repositories
}

func NewTransactionService(repos *repository.Repositories) Service {
	return &transactionSvcImpl{repos: repos}
}

func (t *transactionSvcImpl) Checkout(request *model.TransactionRequest) (*model.TransactionResponse, error) {
	// check payment method
	if !request.PaymentMethod.IsValid() {
		return nil, fmt.Errorf("invalid payment method")
	}

	// check shopping cart existence
	cart, err := t.repos.Cart.GetByID(request.CartID)
	if err != nil {
		return nil, fmt.Errorf("error getting cart by id")
	}

	if cart.CustomerID != request.CustomerID {
		return nil, fmt.Errorf("cart does not belong to the customer")
	}

	if cart.Status != cartEnum.CartStatusActive {
		return nil, fmt.Errorf("cart is not active")
	}

	transactionEntity := cart.ToTransactionEntity()
	transactionEntity.PaymentMethod = request.PaymentMethod
	transactionEntity.CustomerID = request.CustomerID

	// create new transaction
	transaction, err := t.repos.Transaction.Create(transactionEntity)
	if err != nil {
		return nil, fmt.Errorf("error creating transaction")
	}

	return transaction.ToResponse(), nil
}

func (t *transactionSvcImpl) GetByID(id int) (*model.TransactionResponse, error) {
	transaction, err := t.repos.Transaction.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting transaction by id")
	}

	return transaction.ToResponse(), nil
}