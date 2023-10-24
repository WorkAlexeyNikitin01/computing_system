package main

import (
	"lab/graphApiForServices/startServerGraph"
	"lab/orderService/startOrder"
	"lab/productService/startProduct"
	"lab/storeroomService/startStoreroom"
)

func main() {
	go startServerGraph.Start()
	go startOrder.Start()
	go startProduct.Start()
	startStoreroom.Start()
}