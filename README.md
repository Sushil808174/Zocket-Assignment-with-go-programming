# Project Management System.
## Description
The Project Management Application is a backend system built using the Go programming language. This system focuses on simplicity, good unit tests, and a working solution for managing product data. It comprises API endpoints to receive and store product details in a MySQL database. Additionally, there's an image analysis component that handles tasks such as downloading and compressing product images, updating the database with the compressed image paths.

## Technologies Used:
- Go (Golang)
- MySQL
- Gin (Go web framework)
- GORM (Go Object-Relational Mapping library)

# Getting Started:

## Prerequisites:
- Go installed
- MySQL database set up

## Installation:
- Clone the repository: git clone `https://github.com/Sushil808174/Zocket-Assignment-with-go-programming`
- Navigate to the project directory: cd `project_management_application`
- Run the application: go run main.go

## API Endpoints:
1. Create User
- Endpoint: `http://localhost:8888/user`
- Description: Create a new user in the database.
- Request Example:
    ```bash
      {
          "name": "John Doe",
          "mobile": "1234567890",
          "latitude": 37.7749,
          "longitude": -122.4194
      }

2. Create Product
- Endpoint: `http://localhost:8888/product`
- Description: Create a new product in the database.
- Request Example:
     ```bash
       {
          "user_id": "1",
          "product_name": "Sample Product",
          "product_description": "Sample Description",
          "product_images": "image1.jpg",
          "product_price": 19.99
      }




