package models

type Receipt struct {
	Retailer     string `json:"shortDescription"`
	PurchaseDate string `json:"PurchaseDate"`
	PurchaseTime string `json:"PurchaseTime"`
	Items        []Item `json:"Items"`
	Total        string `json:"Total"`
}

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type Response struct {
	Description string `json:"Description"`
	Content     string `json:"Content"`
}
