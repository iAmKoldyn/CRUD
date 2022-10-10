package models

import (
	"encoding/json"
	"time"
)

type Orders struct {
	Id         string          `json:"id"`
	Create_dtm time.Time       `json:"create_dtm"`
	Order_id   string          `json:"order_id"`
	Phone      string          `json:"phone"`
	Name       string          `json:"name"`
	Address    string          `json:"address"`
	Menu       json.RawMessage `json:"menu"`
	Total_item int             `json:"total_item"`
	Pay        int             `json:"pay"`
}

type Query struct {
	Phone string `json:"phone"`
	Date  string `json:"date"`
}
