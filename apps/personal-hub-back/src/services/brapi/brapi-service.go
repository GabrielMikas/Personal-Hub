package brapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func BrapiService() {
	err := godotenv.Load()
	url := "https://brapi.dev/api/v2/stocks/quote?symbols=TAEE11"

	req, _ := http.NewRequest("GET", url, nil)
	k := os.Getenv("BRAPI_KEY")
	a := "Bearer " + k
	req.Header.Add("Authorization", a)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		panic(err)
	}

	price := response.Results[0].Data.RegularMarketPrice
	fmt.Printf("TAEE11 price: %.2f\n", price)
	position := price * 800
	fmt.Printf("Your position currently is %.2f\n", position)
}
