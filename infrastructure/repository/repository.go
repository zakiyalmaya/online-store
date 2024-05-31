package repository

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Repositories struct {
	db  *sqlx.DB
}

func NewRespository(db *sqlx.DB) *Repositories {
	return &Repositories{
		db:        db,
	}
}

func DBConnection() *sqlx.DB {
	db, err := sqlx.Open("sqlite3", "./online_store.db")
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		return nil
	}

	createTableCustomer(db)
	createTableCategories(db)
	createTableProduct(db)
	createTableShoppingCart(db)
	createTableCartItems(db)
	createTableTransaction(db)
	createTableTransactionDetails(db)
	return db
}

func createTableCustomer(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		username VARCHAR(255) UNIQUE NOT NULL,
		email VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		phone_number VARCHAR(255) NOT NULL,
		address TEXT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating table customers: ", err.Error())
	}
}

func createTableCategories(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Panicln("error creating table categories: ", err.Error())
	}
}

func createTableProduct(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(255) NOT NULL,
		description TEXT NULL,
		price DECIMAL(10, 2) NOT NULL,
		stock_quantity INT NOT NULL,
		category_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (category_id) REFERENCES categories(id)
	)`)
	if err != nil {
		log.Panicln("error creating table product: ", err.Error())
	}
}

func createTableShoppingCart(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS shopping_carts (
		id INT PRIMARY KEY AUTO_INCREMENT,
		customer_id INT NOT NULL,
		status int NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	)`)
	if err != nil {
		log.Panicln("error creating table shopping_cart: ", err.Error())
	}
}


func createTableCartItems(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS cart_items (
		id INT PRIMARY KEY AUTO_INCREMENT,
		shopping_cart_id INT NOT NULL,
		product_id INT NOT NULL,
		quantity INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (shopping_cart_id) REFERENCES shopping_carts(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	)`)
	if err != nil {
		log.Panicln("error creating table cart_items: ", err.Error())
	}
}

func createTableTransaction(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS transactios (
		id INT PRIMARY KEY AUTO_INCREMENT,
		idempotency_key VARCHAR(255) UNIQUE NOT NULL,
		shopping_cart_id INT NOT NULL,
		customer_id INT NOT NULL,
		status INT NOT NULL,
		total_amount DECIMAL(10, 2) NOT NULL,
		payment_method INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (shopping_cart_id) REFERENCES shopping_carts(id),
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	)`)
	if err != nil {
		log.Panicln("error creating table transactions: ", err.Error())
	}
}

func createTableTransactionDetails(db *sqlx.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS transactio_details (
		id INT PRIMARY KEY AUTO_INCREMENT,
		transaction_id INT NOT NULL,
		product_id INT NOT NULL,
		quantity INT NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (transaction_id) REFERENCES transactions(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	)`)
	if err != nil {
		log.Panicln("error creating table transaction_details: ", err.Error())
	}
}