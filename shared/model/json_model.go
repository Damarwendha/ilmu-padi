package model

// Struct Statuss
type Status struct {
	Code    int `json:"code"`
	Message any `json:"message"`
}

// Struct Single Response
type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

// Struct Paged Response
type PagedResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data"`
	Paging Paging        `json:"paging"`
}
