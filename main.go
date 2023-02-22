package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	razorpay "github.com/razorpay/razorpay-go"
)

type PageVariables struct {
	OrderId string
	Email   string
	Name    string
	Amount  string
	Contact string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("*.html")
	router.GET("/", App)
	router.GET("/payment-success", PaymentSuccess)
	router.Run(":8089")
}

func App(c *gin.Context) {

	page := &PageVariables{}
	page.Amount = "11000"
	page.Email = "fazilkp2000@gmail.com"
	page.Name = "FAZIL MUHAMMED kP"
	page.Contact = "9946573296"
	//Create order_id from the server
	client := razorpay.NewClient("rzp_test_vOsKKSWnOE803Q", "JINdUUpdybhJ707mAu37fH84")

	data := map[string]interface{}{
		"amount":   page.Amount,
		"currency": "INR",
		"receipt":  "some_receipt_id",
	}
	body, err := client.Order.Create(data, nil)
	fmt.Println("////////////////reciept", body)
	if err != nil {
		fmt.Println("Problem getting the repository information", err)
		os.Exit(1)
	}

	value := body["id"]

	str := value.(string)
	fmt.Println("str////////////////", str)
	HomePageVars := PageVariables{ //store the order_id in a struct
		OrderId: str,
		Amount:  page.Amount,
		Email:   page.Email,
		Name:    page.Name,
		Contact: page.Contact,
	}

	c.HTML(http.StatusOK, "app.html", HomePageVars)
}

func PaymentSuccess(c *gin.Context) {

	paymentid := c.Query("paymentid")
	orderid := c.Query("orderid")
	signature := c.Query("signature")

	fmt.Println(paymentid, "paymentid")
	fmt.Println(orderid, "orderid")
	fmt.Println(signature, "signature")
}

func PaymentFaliure(c *gin.Context) {

}
