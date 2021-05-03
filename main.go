package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	file_name := "cryptocoinmarketcap.csv"

	file, err := os.Create(file_name)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", file_name, err)
		return
	}

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"Name", "Symbol", "Price (USD)", "Volume (USD)", "Market capacity (USD)", "Change (1h)", "Change (24h)", "Change (7d)"})

	c := colly.NewCollector()

	c.OnHTML(".cmc-table-row", func(e *colly.HTMLElement) {
		crypto_name := e.ChildText(".cmc-table__column-name")
		// if crypto_name != "" {
		// 	fmt.Println(crypto_name)
		// }

		crypto_sym := e.ChildText(".cmc-table__cell--sort-by__symbol")
		// if crypto_sym != "" {
		// 	fmt.Println(crypto_sym)
		// }

		crypto_price := e.ChildText(".cmc-table__cell--sort-by__price")
		// if crypto_price != "" {
		// 	fmt.Println(crypto_price)
		// }

		crypto_volume := e.ChildText(".cmc-table__cell--sort-by__volume-24-h")
		// if crypto_volume != "" {
		// 	fmt.Println(crypto_volume)
		// }

		crypto_market_cap := e.ChildText(".cmc-table__cell--sort-by__market-cap")
		// if crypto_market_cap != "" {
		// 	fmt.Println(crypto_market_cap)
		// }

		crypto_percent_1h := e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h")
		// if crypto_percent_1h != "" {
		// 	fmt.Println(crypto_percent_1h)
		// }

		crypto_percent_24h := e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h")
		// if crypto_percent_24h != "" {
		// 	fmt.Println(crypto_percent_24h)
		// }

		crypto_percent_7d := e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d")
		// if crypto_percent_7d != "" {
		// 	fmt.Println(crypto_percent_7d)
		// }

		writer.Write([]string{
			crypto_name,
			crypto_sym,
			crypto_price,
			crypto_volume,
			crypto_market_cap,
			crypto_percent_1h,
			crypto_percent_24h,
			crypto_percent_7d,
		})
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")

	log.Printf("Scraping finished, check file %q for result\n", file_name)
}
