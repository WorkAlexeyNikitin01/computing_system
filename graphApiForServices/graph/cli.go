package graph

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Cli struct {
	Cli *http.Client
}

func NewCli() Cli {
	return Cli{&http.Client{}}
}

type Product struct {
	Id int `json:"id"`
	Code string `json:"code"`
	Price float64 `json:"price"`
	Name string `json:"name"`
	Description  string `json:"description"`
}

func(c *Cli) GetProductRest(code string) Product {
	body := map[string]any{
		"code": code,
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:18080/api/v1/get-product/", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := c.Cli.Do(req)
	respBody, _ := io.ReadAll(resp.Body)
	var p Product
	_ = json.Unmarshal(respBody, &p)
	return p
}

func(c *Cli) CreateProduct(name string, price float64, code string) Product {
	body := map[string]any{
		"code": code,
		"name": name,
		"price": price,
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:18080/api/v1/create-product/", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := c.Cli.Do(req)
	respBody, _ := io.ReadAll(resp.Body)
	var p Product
	_ = json.Unmarshal(respBody, &p)
	log.Println(p.Code)
	return p
}

type Order struct {
	Id       string  `json:"Id"`
	Name     string  `json:"name"`
	Quantity int     `json:"quantity"`
	Code     string  `json:"code"`
	Price    float64 `json:"price"`
	Total    float64 `json:"total"`
}

type DataOrder struct {
	Data Order `json:"data"`
}

func(c *Cli) CreateOrder(code string, price float64, name string, quantity int) Order {
	body := map[string]any{
		"quantity": quantity,
		"code": code,
		"name": name,
		"price": price,
	}
	data, _ := json.Marshal(body)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:18081/api/v1/create-order/", bytes.NewReader(data))
	req.Header.Add("Content-Type", "application/json")
	resp, _ := c.Cli.Do(req)
	respBody, _ := io.ReadAll(resp.Body)
	var or DataOrder
	_ = json.Unmarshal(respBody, &or)
	return or.Data
}
