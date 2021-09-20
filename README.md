# product-api

To use the API run ```go run main.go``` command in the terminal, and send requests to ```localhost:8080/products```

## Endpoints

- ```/products```, Method: **GET**, get a list of all products
- ```/products/{id}```, Method: **GET**, get a single product with the given id
- ```/products```, Method: **POST**, add a product
- ```/products/{id}```, Method: **PUT**, update a product with the given id
- ```/products/{id}```, Method: **DELETE**, delete a product with the given id
