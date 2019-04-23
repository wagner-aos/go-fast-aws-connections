package test

import (
	"encoding/json"
	"log"

	"github.com/wagner-aos/go-fast-aws-connections/fac_dynamodb"
)

type POSContract struct {
	PosID      string             `json:"posID"`
	MerchantID string             `json:"merchantID"`
	Products   map[string]Product `json:"products"`
}

//Product - it stores products by POS_ID
type Product struct {
	ProductID       string `json:"productID"`
	Brand           string `json:"brandID"`
	TransactionType string `json:"transactionType"`
	Active          bool   `json:"active"`
}

func main() {
	posContract := POSContract{}
	posContract.PosID = "00000003"
	posContract.MerchantID = "000000000000001"

	//Product MASTER
	masterDebito := createProduct("1011", "MASTER", "002000", true)
	masterCreditoVista := createProduct("1012", "MASTER", "003000", true)
	masterCreditoParceladoLojista := createProduct("1013", "MASTER", "003100", true)
	masterCreditoParceladoEmissor := createProduct("1014", "MASTER", "003800", false)

	//Product VISA
	visaDebito := createProduct("1041", "VISA", "002000", true)
	visaCreditoVista := createProduct("1040", "VISA", "003000", true)
	visaCreditoParceladoLojista := createProduct("1042", "VISA", "003100", false)
	visaCreditoParceladoEmissor := createProduct("1043", "VISA", "003800", true)

	//Map of Products
	products := map[string]Product{}
	products[masterDebito.ProductID] = masterDebito
	products[masterCreditoVista.ProductID] = masterCreditoVista
	products[masterCreditoParceladoEmissor.ProductID] = masterCreditoParceladoEmissor
	products[masterCreditoParceladoLojista.ProductID] = masterCreditoParceladoLojista

	products[visaDebito.ProductID] = visaDebito
	products[visaCreditoVista.ProductID] = visaCreditoVista
	products[visaCreditoParceladoEmissor.ProductID] = visaCreditoParceladoEmissor
	products[visaCreditoParceladoLojista.ProductID] = visaCreditoParceladoLojista

	//Bind to object
	posContract.Products = products
	printObject(posContract)

	profile := "asappay-Dev"
	facdynamodb.Start(profile)
	facdynamodb.PutItem("PaymentProcessorProduct", posContract)
}

func createProduct(id string, brand string, transactionType string, active bool) Product {
	product := Product{}
	product.ProductID = id
	product.Brand = brand
	product.TransactionType = transactionType
	product.Active = active
	return product
}

//JSON PRINT
func printObject(obj interface{}) {
	result, _ := json.MarshalIndent(obj, "", "  ")
	log.Println("\n" + string(result))
}
