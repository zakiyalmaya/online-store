# ONLINE STORE APP
An application that allow consumers to purchase products.

This application is build using clean architecture and the following tech stack:
- Language: **Golang**
- Framework: **Fiber**
- Database: **SQLite**
- NoSQL: **Redis**

ERD diagram can be found [here](./ERDOnlineStoreApp.jpg)
1. Table
    - **customers**: Contains customer details.
    - **categories**: Contains product category details.
    - **products**: Contains product details.
    - **shopping_carts**: Contains shopping cart details.
    - **cart_items**: Contains items in the shopping cart.
    - **transactions**: Contains transaction details when a customer checks out.
    - **transaction_details**: Contains details of each product in a transaction.

2. Relationships
    - **customers** to **shopping_carts**: One customer can have multiple shopping carts.
    - **shopping_carts** to **cart_items**: One shopping cart can contain multiple cart items.
    - **cart_items** to **products**: Each cart item is linked to one product.
    - **products** to **categories**: Each product belongs to one category.
    - **shopping_carts** to **transactions**: One shopping cart can be associated with one transaction upon checkout.
    - **transactions** to **transaction_details**: One transaction can have multiple transaction details.
    - **transaction_details** to **products**: Each transaction detail is linked to one product.

`config.go` is a configuration file that contains credential database values used by the application.

## GETTING STARTED

1. Via GitHub Repository
    - Clone the git repository
    ```sh
    https://github.com/zakiyalmaya/online-store.git
    ```
    
    - Make sure you have installed Redis and are running Redis locally. If not, you can install it using the link here: https://redis.io/docs/latest/operate/oss_and_stack/install/

    - Run the `.\main.go` file using this command
    ```sh
    go run .\main.go
    ```

2. Via docker image
    - Pull the Docker image from Docker Hub
    ```sh
    docker pull zakiyalmaya/online-store_app:latest
    docker pull zakiyalmaya/redis:6.2
    ```
    
    - Create a Custom Network. Ensure that both the Redis container and the app container are on the same Docker network.
    ```sh
    docker network create online-store-network
    ```

    - Run the app container
    ```sh
    docker run -d --name redis --network online-store-network zakiyalmaya/redis:6.2
    docker run -d --name app --network online-store-network zakiyalmaya/online-store_app:latest
    ```

    - Run the `.\online_store_app` file using this command
    ```sh
    .\online_store_app
    ```
    
    - Visit http://localhost:3000 in your browser

## API CONTRACT

### Customer Service

1. **Register** 

    `POST /customer`

    ```sh
    curl --location 'http://localhost:3000/customer' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "name": "John Doe",
        "username": "johndoe",
        "password": "John-123",
        "email": "john@example.com",
        "phone_number": "+625793710781",
        "address": "Jakarta"
    }'
    ```

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | name | string | Y | name of the customer | 
        | username | string | Y | username of the customer |
        | password | string | Y | password of the customer |
        | email | string | Y | email of the customer |
        | phone_number | string | Y | phone number of the customer |
        | address | string | Y | address of the customer |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |

        example:

        ```sh
        HTTP/1.1 201 Created
        {
            "message": "success"
        }
        ```

        ```sh
        HTTP/1.1 400 Bad Request
        {
            "message": "Key: 'CustomerRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag"
        }
        ```

        ```sh
        HTTP/1.1 500 Internal Server Error
        {
            "message": "error creating customer"
        }
        ```

2. **Login**

    `POST /customer/login`

    ```sh
    curl --location 'http://localhost:3000/customer/login' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "username": "johndoe",
        "password": "John-123"
    }'
    ```

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | username | string | Y | username of the customer |
        | password | string | Y | password of the customer |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | object | N | response data |
        | username | string | Y | username of the customer |
        | nama | string | Y | name of the customer |
        | token | string | Y | token authentication |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": {
                "username": "minniemouse",
                "name": "Minnie Mouse",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMyNDI3fQ.MlOzYXadNkWyZWKNiFYux3xSs5kHNr1mKnv3lZ5I6dQ"
            }
        }
        ```

        ```sh
        HTTP/1.1 400 Bad Request
        {
            "message": "wrong password"
        }
        ```

3. **Logout**

    `POST /customer/logout`

    ```sh
    curl --location --request POST 'http://localhost:3000/customer/logout' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2tleW1vdXNlIiwiZXhwIjoxNzE3MTQ4ODQwfQ.QUvJ8C9Bj38GAg5TrR-IqPAXAoz3cHy3oNTlwwSgPM0'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | success message |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success"
        }
        ```

        ```sh
        HTTP/1.1 401 Unauthorized
        {
            "message": "Invalid or expired token"
        }
        ```

## Category Service

1. **Create**

    `POST /category`

    ```sh
    curl --location 'http://localhost:3000/category' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2tleW1vdXNlIiwiZXhwIjoxNzE3MTQ4ODQwfQ.QUvJ8C9Bj38GAg5TrR-IqPAXAoz3cHy3oNTlwwSgPM0' \
    --data '{
        "name": "Automotive"
    }'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | name | string | Y | name of the category |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |

        example:

        ```sh   
        HTTP/1.1 201 Created
        {
            "message": "success"
        }
        ```

        ```sh
        HTTP/1.1 400 Bad Request
        {
            "message": "Key: 'CategoryRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
        }
        ```

2. **Get All**

    `GET /categories`

    ```sh
    curl --location 'http://localhost:3000/categories' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2tleW1vdXNlIiwiZXhwIjoxNzE3MTQ5MDUwfQ.EaD9nL70z232j0O28YlZo4VoKYaNYlJIg5tA98BG36c'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | array | N | list of categories |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": [
                {
                    "id": 1,
                    "name": "Automotive"
                }
            ]
        }
        ```

        ```sh
        HTTP/1.1 401 Unauthorized
        {
            "message": "Invalid or expired token"
        }
        ```

## Product Service

1. **Create**

    `POST /product`

    ```sh
    curl --location 'http://localhost:3000/product' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6Im1pY2tleW1vdXNlIiwiZXhwIjoxNzE3MjA0NDI4fQ.lBeFcbO6YWg7ndowWo3a47Qx15gTnEWFLxp4_0HOxK8' \
    --data '{
        "name": "Mothercare Multi Cat Long-Sleeved T-Shirts",
        "price": 459900,
        "stock_quantity": 7,
        "category_id": 2,
        "description": "Baby cloting 3 pcs per pack for 12-18 months old baby girl."
    }'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | name | string | Y | name of the product |
        | price | number | Y | price of the product |
        | stock_quantity | number | Y | stock quantity of the product |
        | category_id | number | Y | category id of the product |
        | description | string | N | description of the product |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |

        example:

        ```sh
        HTTP/1.1 201 Created
        {
            "message": "success"
        }
        ```

        ```sh
        HTTP/1.1 400 Bad Request
        {
            "message": "Key: 'ProductRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
        }
        ```

2. **Get All**

    `GET /products`

    ```sh
    curl --location 'http://localhost:3000/products?category_id=2&page=2&limit=2' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3Mjk1MDI3fQ.oLqQ6Pnvz4PrQ5FJOtcjAYa3VzM_U9HhdEvDYF0PArs'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Query Parameters

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | category_id | number | N | category id of the products |
        | limit | number | N | limit of the products |
        | page | number | N | page of the products |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | array | N | list of products |
        | id | number | Y | id of the product |
        | name | string | Y | name of the product |
        | price | number | Y | price of the product |
        | stock_quantity | number | Y | stock quantity of the product |
        | category | string | Y | category of the product |
        | description | string | N | description of the product |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": [
                {
                    "id": 6,
                    "name": "Mothercare Multi Cat Long-Sleeved T-Shirts",
                    "price": 459900,
                    "stock_quantity": 7,
                    "category": "Fashion",
                    "description": "Baby cloting 3 pcs per pack for 12-18 months old baby girl."
                },
                {
                    "id": 7,
                    "name": "Nylon Solid Color Backpack",
                    "price": 287000,
                    "stock_quantity": 7,
                    "category": "Fashion",
                    "description": "Women Girl Fasihon Nylon Solid Color Backpack. Capacity 20-35L using zipper"
                }
            ]
        }
        ```
        
## Cart Service

1. **Create**

    `POST /cart`

    ```sh
    curl --location 'http://localhost:3000/cart' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMxNjM1fQ.fX3RjSTbH03kAtsGY_QsVwpEKGs2mdOku3PkwAGEYuE' \
    --data '{
        "items": [
        {
            "product_id": 3,
            "quantity": 1
        },
        {
            "product_id": 1,
            "quantity": 1
        }]
    }'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | items | array | Y | list of items |
        | product_id | number | Y | id of the product |
        | quantity | number | Y | quantity of the product |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | object | N | response data |
        | id | number | Y | id of the cart |
        | customer_id | number | Y | id of the customer |
        | status | string | Y | status of the cart |
        | items | array | Y | list of items |
        | id | number | Y | id of the item |
        | product_id | number | Y | id of the product |
        | product_name | string | Y | name of the product |
        | quantity | number | Y | quantity of the product |
        | price | number | Y | price of the product |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": {
                "id": 1,
                "customer_id": 1,
                "status": "active",
                "items": [
                    {
                        "id": 1,
                        "product_id": 1,
                        "product_name": "Mothercare Multi Cat Long-Sleeved T-Shirts",
                        "quantity": 1,
                        "price": 459900
                    },
                    {
                        "id": 2,
                        "product_id": 3,
                        "product_name": "Mothercare Baby Sleeping Pad",
                        "quantity": 1,
                        "price": 350000
                    }
                ]
            }
        }
        ```

2. **Get All By Customer**

    `GET /carts`

    ```sh
    curl --location 'http://localhost:3000/carts?status=1' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMyNDI3fQ.MlOzYXadNkWyZWKNiFYux3xSs5kHNr1mKnv3lZ5I6dQ'
    ```

    - Header

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Query Param

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | status | string | N | status of the cart |

    - Response Body

        | field |type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | array | N | list of carts |
        | id | number | Y | id of the cart |
        | customer_id | number | Y | id of the customer |
        | status | string | Y | status of the cart |
        | items | array | Y | list of items |
        | id | number | Y | id of the item |
        | product_id | number | Y | id of the product |
        | product_name | string | Y | name of the product |
        | quantity | number | Y | quantity of the product |
        | price | number | Y | price of the product |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": [
                {
                    "id": 1,
                    "customer_id": 1,
                    "status": "ACTIVE",
                    "items": [
                        {
                            "id": 1,
                            "product_id": 1,
                            "product_name": "Mothercare Multi Cat Long-Sleeved T-Shirts",
                            "quantity": 1,
                            "price": 459900
                        },
                        {
                            "id": 2,
                            "product_id": 3,
                            "product_name": "Mothercare Baby Sleeping Pad",
                            "quantity": 1,
                            "price": 350000
                        }
                    ]
                }
            ]
        }
        ```

3. **Delete cart item By ID**

    `DELETE /cart/{id}`

    ```sh
    curl --location 'http://localhost:3000/cart/1' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMyNDI3fQ.MlOzYXadNkWyZWKNiFYux3xSs5kHNr1mKnv3lZ5I6dQ'
    ```

    - Header

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Path Param

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | id | number | Y | id of the cart item |

    - Response Body

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success"
        }
        ```

## Transaction Service

1. **Checkout**

    `POST /transaction`

    ```sh
    curl --location 'http://localhost:3000/transaction' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMxNjM1fQ.fX3RjSTbH03kAtsGY_QsVwpEKGs2mdOku3PkwAGEYuE' \
    --data '{
        "shopping_cart_id": 6,
        "payment_method" : 4
    }'
    ```

    - Header

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Body

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | shopping_cart_id | number | Y | id of the shopping cart |
        | payment_method | number | Y | id of the payment method |

    - Response Body

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | object | N | response data |
        | id | number | Y | id of the transaction |
        | idempotency_key | string | Y | idempotency key of the transaction |
        | customer_id | number | Y | id of the customer |
        | shopping_cart_id | number | Y | id of the shopping cart |
        | status | string | Y | status of the transaction |
        | total_amount | number | Y | total amount of the transaction |
        | payment_method | number | Y | id of the payment method |
        | transaction_details | array | Y | details of the transaction |
        | id | number | Y | id of the transaction detail |
        | product_id | number | Y | id of the product |
        | product_name | string | Y | name of the product |
        | quantity | number | Y | quantity of the product |
        | price | number | Y | price of the product |

        example:

        ```sh
        HTTP/1.1 201 Created
        {
            "message": "success",
            "data": {
                "id": 5,
                "idempotency_key": "7cf90026-e834-4f2c-bd83-3799ef7b8f6e",
                "customer_id": 2,
                "shopping_cart_id": 6,
                "status": "IN PROGRESS",
                "total_amount": 12000,
                "payment_method": "CASH",
                "transaction_details": [
                    {
                        "id": 5,
                        "product_id": 3,
                        "product_name": "Beng-Beng Share It",
                        "quantity": 1,
                        "price": 12000
                    }
                ]
            }
        }
        ```

        ```sh
        HTTP/1.1 500 Internal Server Error
        {
            "message": "cart does not belong to the customer"
        }
        ```

2. **Get By ID**

    `GET /transaction`

    ```sh
    curl --location 'http://localhost:3000/transaction?id=3' \
    --header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJ1c2VybmFtZSI6Im1pbm5pZW1vdXNlIiwiZXhwIjoxNzE3MjMxNjM1fQ.fX3RjSTbH03kAtsGY_QsVwpEKGs2mdOku3PkwAGEYuE'
    ```

    - Header

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | Authorization | string | Y | token authentication |

    - Request Query Param

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | id | number | Y | id of the transaction |

    - Response Body

        | field | type | required? (Y/N) | description |
        | :---: | :---: | :---: | :---: |
        | message | string | Y | response message |
        | data | object | N | response data |
        | id | number | Y | id of the transaction |
        | idempotency_key | string | Y | idempotency key of the transaction |
        | customer_id | number | Y | id of the customer |
        | shopping_cart_id | number | Y | id of the shopping cart |
        | status | string | Y | status of the transaction |
        | total_amount | number | Y | total amount of the transaction |
        | payment_method | number | Y | id of the payment method |
        | transaction_details | array | Y | details of the transaction |
        | id | number | Y | id of the transaction detail |
        | product_id | number | Y | id of the product |
        | product_name | string | Y | name of the product |
        | quantity | number | Y | quantity of the product |
        | price | number | Y | price of the product |

        example:

        ```sh
        HTTP/1.1 200 OK
        {
            "message": "success",
            "data": {
                "id": 3,
                "idempotency_key": "7cf90026-e834-4f2c-bd83-3799ef7b8f6e",
                "customer_id": 2,
                "shopping_cart_id": 6,
                "status": "IN PROGRESS",
                "total_amount": 12000,
                "payment_method": "CASH",
                "transaction_details": [
                    {
                        "id": 5,
                        "product_id": 3,
                        "product_name": "Beng-Beng Share It",
                        "quantity": 1,
                        "price": 12000
                    }
                ]
            }
        }
        ```

        ```sh
        HTTP/1.1 404 Not Found
        {
            "message": "transaction not found"
        }
        ```
    
            


