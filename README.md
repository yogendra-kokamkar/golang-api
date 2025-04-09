# User Balance API

A RESTful API built with Go that handles user authentication and provides balance-related information.

## Description
This API service provides secure endpoints to manage user authentication and retrieve balance information. It implements middleware for authorization to ensure secure access to user data.

## Table of Contents
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [API Endpoints](#api-endpoints)

## Technologies Used
- Go (Golang)
- Simple String from Mock Database for authentication

## Prerequisites
- Go 1.16 or higher
- Mock Database (Created in the same project)
- Git

## Installation
```bash
# Clone the repository
git clone [your-repository-url]

# Install dependencies
go mod download

# Navigate to the project directory
cd cmd/api

# Run the application
go run main.go
```

## API Endpoints
```bash
http://localhost:8000/account/coins?username=Yogendra
```
Note - Add Authentication in Headers.

Details Found (Correct Credentials)
![image](https://github.com/user-attachments/assets/c01037b3-04c5-4ba4-989c-a2aa8f542e45)

Details Not Found (Incorrect Credentials)
![image](https://github.com/user-attachments/assets/09f9dcd0-7768-43d6-823c-5ee7bb5525bd)

