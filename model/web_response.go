package model

type WebResponse struct {
	Data   []OrderResponse `json:"data"`
	Status []Status        `json:"status"`
}

type Status struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
