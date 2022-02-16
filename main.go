package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	program := os.Args[0]

	if os.Args[1] == "help" {
		usage(program)
		os.Exit(0)
	}

	if len(os.Args) < 3 {
		usage(program)
		os.Exit(1)

	}

	datas := os.Args[1:]

	URL := getUrl(datas[0], datas[1])

	resp, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		panic(err)
	}

	result := doc.Find(".result__BigRate-sc-1bsijpp-1").Text()

	fmt.Println("result:", result)
}

// usage ...
func usage(program string) {
	fmt.Printf("Usage: %s <money1> <money2>\n", program)
}

func getUrl(arg1, arg2 string) string {
	return "https://www.xe.com/currencyconverter/convert/?Amount=1&From=" + arg1 + "&To=" + arg2
}
