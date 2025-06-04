package main

import (
	"backend-interview/exception"
	"backend-interview/helper"
	"backend-interview/model"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	router := httprouter.New()
	router.PanicHandler = exception.ErrorHandler

	router.GET("/orders", GetOrders)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	err := godotenv.Load()
	if err != nil {
		panic("error loading .env file")
	}

	conn, err := initDB()
	if err != nil {
		panic("failed to connect to db")
	}

	defer conn.Close(context.Background())

	runningSampleData(conn)

	slog.Info("start a server")

	server := http.Server{
		Addr:    os.Getenv("GO_PORT"),
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		slog.Error("failed to start http server", "error", err)
		panic("failed to start http server")
	}
}

func initDB() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		return nil, err
	}

	return conn, nil
}

func runningSampleData(conn *pgx.Conn) {
	ctx := context.Background()

	buyer := model.User{
		ID:           "11111111-1111-1111-1111-111111111111",
		Username:     "andhika",
		Email:        "andhika@gmail.com",
		ContactPhone: "012345678",
		Password:     "andhika",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	buyer2 := model.User{
		ID:           "22222222-2222-2222-2222-222222222222",
		Username:     "andhikaferdiansyah",
		Email:        "andhikaferdiansyah@gmail.com",
		ContactPhone: "012345678",
		Password:     "andhikaferdiansyah",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	seller := model.User{
		ID:           "33333333-3333-3333-3333-333333333333",
		Username:     "ferdiansyah",
		Email:        "ferdiansyah@gmail.com",
		ContactPhone: "012345678",
		Password:     "ferdiansyah",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	_, err := conn.Exec(ctx, `
		INSERT INTO users (id, username, email, contact_phone, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO NOTHING;
	`, buyer.ID, buyer.Username, buyer.Email, buyer.ContactPhone, buyer.Password, buyer.CreatedAt, buyer.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert users", "error", err)
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO users (id, username, email, contact_phone, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO NOTHING;
	`, buyer2.ID, buyer2.Username, buyer2.Email, buyer2.ContactPhone, buyer2.Password, buyer2.CreatedAt, buyer2.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert users", "error", err)
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO users (id, username, email, contact_phone, password, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)
		ON CONFLICT (id) DO NOTHING;
	`, seller.ID, seller.Username, seller.Email, seller.ContactPhone, seller.Password, seller.CreatedAt, seller.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert users", "error", err)
	}

	sellerProduct := model.Product{
		ID:          "44444444-4444-4444-4444-444444444444",
		UserID:      seller.ID,
		Name:        "Gaming Mouse",
		Description: "This is a gaming mouse",
		Price:       50.50,
		Quantity:    10,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO products (id, user_id, name, description, price, quantity, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (id) DO NOTHING;
	`, sellerProduct.ID, sellerProduct.UserID, sellerProduct.Name, sellerProduct.Description, sellerProduct.Price, sellerProduct.Quantity, sellerProduct.CreatedAt, sellerProduct.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert product", "error", err)
	}

	order := model.Order{
		ID:         "55555555-5555-5555-5555-555555555555",
		ProductID:  sellerProduct.ID,
		CustomerID: buyer.ID,
		SellerID:   seller.ID,
		CreatedBy:  buyer.Username,
		Amount:     30.30,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	order2 := model.Order{
		ID:         "66666666-6666-6666-6666-666666666666",
		ProductID:  sellerProduct.ID,
		CustomerID: buyer2.ID,
		SellerID:   seller.ID,
		CreatedBy:  buyer2.Username,
		Amount:     99.10,
		Status:     0,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO orders (id, product_id, customer_id, seller_id, created_by,amount, status, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
		ON CONFLICT (id) DO NOTHING;
	`, order.ID, order.ProductID, order.CustomerID, order.SellerID, order.CreatedBy, order.Amount, order.Status, order.CreatedAt, order.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert order", "error", err)
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO orders (id, product_id, customer_id, seller_id, created_by,amount, status, created_at, updated_at) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)
		ON CONFLICT (id) DO NOTHING;
	`, order2.ID, order2.ProductID, order2.CustomerID, order2.SellerID, order2.CreatedBy, order2.Amount, order2.Status, order2.CreatedAt, order2.UpdatedAt)
	if err != nil {
		slog.Error("failed to insert order", "error", err)
	}

	transaction := model.Transaction{
		ID:                "77777777-7777-7777-7777-777777777777",
		OrderID:           order.ID,
		TransactionMethod: "Debit Card",
		Amount:            60.60,
		Status:            1,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	transaction2 := model.Transaction{
		ID:                "88888888-8888-8888-8888-888888888888",
		OrderID:           order2.ID,
		TransactionMethod: "Credit Card",
		Amount:            70.90,
		Status:            0,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}
	_, err = conn.Exec(ctx, `
		INSERT INTO transactions (id, order_id, transaction_method, amount, status, created_at) VALUES($1,$2,$3,$4,$5,$6)
		ON CONFLICT (id) DO NOTHING;
	`, transaction.ID, transaction.OrderID, transaction.TransactionMethod, transaction.Amount, transaction.Status, transaction.CreatedAt)
	if err != nil {
		slog.Error("failed to insert transaction", "error", err)
	}

	_, err = conn.Exec(ctx, `
		INSERT INTO transactions (id, order_id, transaction_method, amount, status, created_at) VALUES($1,$2,$3,$4,$5,$6)
		ON CONFLICT (id) DO NOTHING;
	`, transaction2.ID, transaction2.OrderID, transaction2.TransactionMethod, transaction2.Amount, transaction2.Status, transaction2.CreatedAt)
	if err != nil {
		slog.Error("failed to insert transaction", "error", err)
	}
}

func GetOrders(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	conn, err := initDB()
	if err != nil {
		slog.Error("failed to init db")
	}
	defer conn.Close(context.Background())

	orders, err := GetOrderData(conn)
	if err != nil {
		panic("order not found")
	}
	orderResponse := make([]model.OrderResponse, len(orders))
	statusMap := make(map[int]string)

	for i := range orders {
		orderResponse[i].ID = orders[i].ID
		orderResponse[i].ProductID = orders[i].ProductID
		orderResponse[i].ProductName, _ = GetProductName(conn, orderResponse[i].ProductID)
		orderResponse[i].Amount, orderResponse[i].TransactionDate, _ = GetAmountAndTransactionDate(conn, orderResponse[i].ID)
		orderResponse[i].CustomerName, _ = GetCustomerName(conn, orders[i].CustomerID)
		orderResponse[i].CreateBy = orderResponse[i].CustomerName
		orderResponse[i].CreateOn = orders[i].CreatedAt

		orderResponse[i].Status = orders[i].Status

		if _, exists := statusMap[orderResponse[i].Status]; !exists {
			if orderResponse[i].Status == 0 {
				statusMap[0] = "SUCCESS"
			} else if orderResponse[i].Status == 1 {
				statusMap[1] = "FAILED"
			}
		}
	}

	var statusList []model.Status
	for id, name := range statusMap {
		statusList = append(statusList, model.Status{
			Id:   id,
			Name: name,
		})
	}

	webResponse := model.WebResponse{
		Data:   orderResponse,
		Status: statusList,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func GetOrderData(conn *pgx.Conn) ([]model.Order, error) {
	rows, err := conn.Query(context.Background(),
		`SELECT id,product_id,customer_id,status,created_by,created_at FROM orders`,
	)
	if err != nil {
		slog.Error("failed to query into orders table", "error", err)
		panic("failed to query into orders table")
	}

	defer rows.Close()

	var orders []model.Order
	hasData := false

	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.ID, &order.ProductID, &order.CustomerID, &order.Status, &order.CreatedBy, &order.CreatedAt)
		if err != nil {
			slog.Error("failed to scan query result", "error", err)
			panic("failed to scan query result")
		}
		hasData = true
		orders = append(orders, order)
	}

	if hasData == false {
		slog.Error("order not found")
		return nil, errors.New("order not found")
	}

	return orders, nil
}

func GetProductName(conn *pgx.Conn, productID string) (string, error) {
	row, err := conn.Query(context.Background(),
		`SELECT name FROM products WHERE id=$1`, productID,
	)
	if err != nil {
		slog.Error("failed to query into orders table", "error", err)
		panic("failed to query into orders table")
	}

	defer row.Close()

	var productName string
	if row.Next() {
		err = row.Scan(&productName)
		if err != nil {
			slog.Error("failed to scan query result", "error", err)
			panic("failed to query into product table")
		}

		return productName, nil
	} else {
		return "", errors.New("product not found")
	}
}

func GetCustomerName(conn *pgx.Conn, customerID string) (string, error) {
	row, err := conn.Query(context.Background(),
		`SELECT username FROM users WHERE id=$1`, customerID,
	)

	if err != nil {
		slog.Error("failed to query into orders table", "error", err)
	}

	defer row.Close()

	var customerName string
	if row.Next() {
		err = row.Scan(&customerName)
		if err != nil {
			slog.Error("failed to scan query result", "error", err)
			panic("failed to query into users table")
		}

		return customerName, nil
	} else {
		return "", errors.New("user not found")
	}
}

func GetAmountAndTransactionDate(conn *pgx.Conn, orderID string) (string, time.Time, error) {
	row, err := conn.Query(context.Background(),
		`SELECT amount,created_at FROM transactions WHERE order_id=$1`, orderID,
	)

	if err != nil {
		slog.Error("failed to query into orders table", "error", err)
	}

	defer row.Close()

	var amount string
	var transactionDate time.Time

	if row.Next() {
		err = row.Scan(&amount, &transactionDate)
		if err != nil {
			slog.Error("failed to scan query result", "error", err)
			panic("failed to query into transactions table")
		}

		return amount, transactionDate, nil
	} else {
		return "", time.Time{}, errors.New("transaction not found")
	}
}
