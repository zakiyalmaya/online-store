package transaction

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/zakiyalmaya/online-store/model"
	cartEnum "github.com/zakiyalmaya/online-store/constant/cart"
)

type transactonRepoImpl struct {
	db *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) Repository {
	return &transactonRepoImpl{db: db}
}

func (t *transactonRepoImpl) Create(transaction *model.TransactionEntity) (*model.TransactionEntity, error) {
	tx, err := t.db.Beginx()
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}
	
	res, err := tx.NamedExec(`INSERT INTO transactions (idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method) VALUES (:idempotency_key, :customer_id, :shopping_cart_id, :status, :total_amount, :payment_method)`, transaction)
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	transactionID, err := res.LastInsertId()
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	for _, detail := range transaction.Details {
		detail.TransactionID = int(transactionID)
		_, err = tx.NamedExec(`INSERT INTO transaction_details (transaction_id, product_id, quantity, price) VALUES (:transaction_id, :product_id, :quantity, :price)`, detail)
		if err != nil {
			tx.Rollback()
			log.Println("errorRepository: ", err.Error())
			return nil, err
		}
	}

	_, err = tx.Exec(`UPDATE shopping_carts SET status = ? WHERE id = ?`, cartEnum.CartStatusPending, transaction.CartID)
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	return t.getByID(int(transactionID))
}

func (t *transactonRepoImpl) getByID(id int) (*model.TransactionEntity, error) {
	transaction := &model.TransactionEntity{}
	query := "SELECT id, idempotency_key, customer_id, shopping_cart_id, status, total_amount, payment_method, created_at, updated_at FROM transactions WHERE id = ?"
	err := t.db.Get(transaction, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	details := []*model.TransactionDetailEntity{}
	query = "SELECT td.id, td.transaction_id, td.product_id, p.name AS product_name, td.quantity, td.price, td.created_at, td.updated_at FROM transaction_details AS td JOIN products AS p ON td.product_id = p.id WHERE td.transaction_id = ? ORDER BY td.id"
	err = t.db.Select(&details, query, id)
	if err != nil {
		log.Println("errorRepository: ", err.Error())
		return nil, err
	}

	transaction.Details = details
	return transaction, nil
}

func (t *transactonRepoImpl) GetByID(id int) (*model.TransactionEntity, error) {
	return t.getByID(id)
}
