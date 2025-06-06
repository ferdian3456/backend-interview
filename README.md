# Backend Interview

This project is a http server api for backend interview purposes.  

---

## 🛠 Tech Stack

<p align="left">
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" alt="Go" width="40" height="40"/>
  &nbsp;
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" alt="PostgreSQL" width="40" height="40"/>
  &nbsp;
  <img src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original.svg" alt="Docker" width="40" height="40"/>
</p>

---

## 📄 Entity Relation Diagram (ERD)

<p align="center">
  <img src="ERD.jpg" style="width:900px;height:500px;">
</p>

## 📚 Database Table

Here is a list of database table:
1. User = to save user information such as username, email, contact_phone, etc
2. Product = to save product information such as product name, description, etc
3. Order = to save data related to order such as order amount, status, etc
4. Transaction = to save data after user do transaction, such as 

## 💥 Flow:
```
Seller and buyer create account -> Seller create a product -> Buyer order the product -> Buyer do transaction
```

## 📁 Project Structure
```
├── backend-interview/
│   ├── db/                     -> contains sql script to create table
│       ├── migrations/
│   ├── helper/                 -> contains helper function 
│   ├── model/                  -> contains mapped database entity in go
│   ├── ERD.jpg                 -> entity relationship diagram 
│   ├── go.mod                  -> go module
│   ├── main.go                 -> main file
│   ├── README.md               -> read me
│   ├── Makefile                -> to simplify migration with go migrate
│   ├── docker-compose.yml      -> docker compose to run postgresql contains
│   ├── .env.example            -> environment variable to setup project
│   ├── .gitignore              -> git ignore to ignore .env and .idea
```

## 🚀 Getting Started

To set up and run this project locally, ensure the following are installed:

1. ✅ **Go Programming Language** — [https://golang.org/dl/](https://golang.org/dl/)
2. ✅ **Golang Migrate** — [https://github.com/golang-migrate/migrate](https://github.com/golang-migrate/migrate)
3. ✅ **Docker** — [https://www.docker.com/products/docker-desktop](https://www.docker.com/products/docker-desktop)

## 📦 Setup Project
Follow these steps to start the backend locally:
1. Run postgres database in container
 ```
docker compose up -d
```
2. Run table migration
```
make migrate-up
```
3. Rename .env.example to .env and change enviroment variable to your desire
```
GO_PORT=localhost:8080
POSTGRES_USER=yourusername
POSTGRES_PASSWORD=yourpassword
POSTGRES_DB=myshop
POSTGRES_URL=postgres://yourusername:yourpassword@localhost:5432/myshop
```
4. Run http server
```
go run main.go
```

## 🔌 API Endpoint
**GET** <code>/orders</code>
**Response**
```
{
  "data": [
    {
      "id": "55555555-5555-5555-5555-555555555555",
      "productID": "44444444-4444-4444-4444-444444444444",
      "productName": "Gaming Mouse",
      "amount": "60.60",
      "customerName": "andhika",
      "status": 1,
      "transactionDate": "2025-06-04T22:09:04.957576Z",
      "createBy": "andhika",
      "createOn": "2025-06-04T22:09:04.949972Z"
    },
    {
      "id": "66666666-6666-6666-6666-666666666666",
      "productID": "44444444-4444-4444-4444-444444444444",
      "productName": "Gaming Mouse",
      "amount": "70.90",
      "customerName": "andhikaferdiansyah",
      "status": 0,
      "transactionDate": "2025-06-04T22:09:04.957577Z",
      "createBy": "andhikaferdiansyah",
      "createOn": "2025-06-04T22:09:04.949972Z"
    }
  ],
  "status": [
    {
      "id": 1,
      "name": "FAILED"
    },
    {
      "id": 0,
      "name": "SUCCESS"
    }
  ]
}
```

If you delete 1 order which status is 0 Or Success in db, then the response would looks like this
```
{
  "data": [
    {
      "id": "55555555-5555-5555-5555-555555555555",
      "productID": "44444444-4444-4444-4444-444444444444",
      "productName": "Gaming Mouse",
      "amount": "60.60",
      "customerName": "andhika",
      "status": 1,
      "transactionDate": "2025-06-04T22:09:04.957576Z",
      "createBy": "andhika",
      "createOn": "2025-06-04T22:09:04.949972Z"
    }
  ],
  "status": [
    {
      "id": 1,
      "name": "FAILED"
    }
  ]
}
```
