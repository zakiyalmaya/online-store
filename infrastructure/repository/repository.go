package repository

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/zakiyalmaya/online-store/infrastructure/repository/cart"
	"github.com/zakiyalmaya/online-store/infrastructure/repository/category"
	"github.com/zakiyalmaya/online-store/infrastructure/repository/customer"
	"github.com/zakiyalmaya/online-store/infrastructure/repository/product"
	"github.com/zakiyalmaya/online-store/infrastructure/repository/transaction"
)

type Repositories struct {
	db          *sqlx.DB
	RedCl       *redis.Client
	Category    category.Repository
	Customer    customer.Repository
	Product     product.Repository
	Cart        cart.Repository
	Transaction transaction.Repository
}

func NewRepository(db *sqlx.DB, redcl *redis.Client) *Repositories {
	return &Repositories{
		db:          db,
		RedCl:       redcl,
		Category:    category.NewCategoryRepository(db),
		Customer:    customer.NewCustomerRepository(db),
		Product:     product.NewProductRepository(db),
		Cart:        cart.NewCartRepository(db),
		Transaction: transaction.NewTransactionRepository(db),
	}
}

func DBConnection(sqlLite string) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", sqlLite)
	if err != nil {
		log.Panicln("error connecting to database: ", err.Error())
		panic(err)
	}

	createTableCustomer(db)
	createTableCategories(db)
	createTableProduct(db)
	createTableShoppingCart(db)
	createTableCartItems(db)
	createTableTransaction(db)
	createTableTransactionDetails(db)
	createIndexTabelCartItems(db)
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS products (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name VARCHAR(255) NOT NULL,
		description TEXT NULL,
		price DECIMAL(10, 2) NOT NULL,
		stock_quantity INT NOT NULL,
		category_id INTEGER NOT NULL,
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
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
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		shopping_cart_id INT NOT NULL,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		idempotency_key VARCHAR(255) UNIQUE NOT NULL,
		shopping_cart_id INTEGER NOT NULL,
		customer_id INTEGER NOT NULL,
		status INTEGER NOT NULL,
		total_amount DECIMAL(10, 2) NOT NULL,
		payment_method INTEGER NOT NULL,
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
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS transaction_details (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		transaction_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		quantity INTEGER NOT NULL,
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

func createIndexTabelCartItems(db *sqlx.DB) {
	_, err := db.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_product_cart ON cart_items (product_id, shopping_cart_id)`)
	if err != nil {
		log.Panicln("error creating index cart_items: ", err.Error())
	}
}

func RedisClient(redisHost, redisPort string) *redis.Client {
	option := &redis.Options{
		Addr:     redisHost + ":" + redisPort,
		Password: "",
		DB:       0,
	}

	redcl := redis.NewClient(option)
	ctx := context.Background()

	for i := 0; i < 10; i++ {
		_, err := redcl.Ping(ctx).Result()
		if err == nil {
			log.Println("Connected to Redis")
			break
		}
		
		log.Println("Failed to connect to Redis. Retrying...")
		time.Sleep(2 * time.Second)
		if i == 9 {
			log.Panicln("Could not connect to Redis:", err)
		}
	}

	return redcl
}
