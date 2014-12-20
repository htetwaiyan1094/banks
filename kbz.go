package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	str "strings"
)

var (
	kbz        = "http://www.kbzbank.com"
	temp       = []string{}
	currencies = []string{}
	buy        = []string{}
	sell       = []string{}
)

func main() {

	f, err := os.Open("kbz.html")
	PanicIf(err)

	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		PanicIf(err)
	}

	doc.Find(".answer tbody tr").Each(func(i int, s *goquery.Selection) {
		s.Find("td").Each(func(u int, t *goquery.Selection) {
			temp = append(temp, t.Text())
		})
	})

	k := Bank{}

	for j, _ := range temp {
		k.Name = "KBZ"
		k.Base = "MMK"
		if j%3 == 0 {
			currencies = append(currencies, str.TrimSpace(temp[j]))
		}
	}

	buy = append(buy, temp[1], temp[4], temp[7])
	sell = append(sell, temp[2], temp[5], temp[8])

	for x, _ := range currencies {
		r := Rate{
			map[string]BuySell{
				currencies[x]: BuySell{buy[x], sell[x]},
			},
		}
		k.Rates = append(k.Rates, r)
	}

	fmt.Println(k)
}

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

type Bank struct {
	Name string
	//Time string
	Base  string
	Rates []Rate
}

type Rate struct {
	BS map[string]BuySell
}

type BuySell struct {
	Buy  string
	Sell string
}
