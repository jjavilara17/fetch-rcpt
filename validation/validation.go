package validation

import (
	"encoding/json"
	"fmt"
	"jjavilara17/rcpt/models"
	"regexp"
	"strconv"
)

const (
	strRegex   = "^\\S+$"
	moneyRegex = "^\\d+\\.\\d{2}$"
	descRegex  = "^[\\w\\s\\-]+$"
	dateFmt    = ""
	timeFmt    = ""
)

func ValidateRegex(rx string, input string) string {
	var msg string
	match, _ := regexp.MatchString(rx, input)
	if !match {
		msg = fmt.Sprintf("Request parameter %s is not valid", input)
	}
	return msg
}

func Validate() {

}

func ValidateRequest(req string) {
	if len(req) == 0 {
		return
	}

	var receipt models.Receipt
	json.Unmarshal([]byte(req), &receipt)
	ValidateRegex(receipt.Retailer, "^\\S+$")
	ValidateRegex(receipt.PurchaseDate, "^\\S+$")
	ValidateRegex(receipt.PurchaseTime, "^\\S+$")

	var total float64
	for _, item := range receipt.Items {
		ValidateRegex(item.Price, "^\\S+$")
		ValidateRegex(item.ShortDescription, "^\\S+$")
		total, _ = strconv.ParseFloat(item.Price, 64)
	}
	rTotal, _ := strconv.ParseFloat(receipt.Total, 64)
	err := total != rTotal
	if err {

	}
}
