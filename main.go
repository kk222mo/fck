package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs"`
	Date    string   `xml:"Date,attr"`
	Name    string   `xml:"name,attr"`
	Valutes []struct {
		ID       string `xml:"ID,attr"`
		NumCode  string `xml:"NumCode"`
		CharCode string `xml:"CharCode"`
		Nominal  string `xml:"Nominal"`
		Name     string `xml:"Name"`
		Value    string `xml:"Value"`
	} `xml:"Valute"`
}

var draw string = `
   ____________
  |            |
  |      8     |
  |    88888   |
  |      8     |
  |      8     |
  |      8     |
  |            |` + "\033[32;1m" + `
 \` + "\033[37;1m" + `|     RIP    |` + "\033[32;1m" + `//
\\` + "\033[37;1m" + `|____РУБЛЬ___|` + "\033[32;1m" + `///
        `

func getRand(a int, b int) int {
	if b-a == 0 {
		return a
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(b-a) + a
}

func getRandError() string {
	errors := [4]string{"Все, центробанк сдох, расходимся.", "Пиздец наступил", "Блять.", "Блять сука"}
	return errors[getRand(0, len(errors))]
}

func main() {
	resp, err := http.Get("http://www.cbr.ru/scripts/XML_daily_eng.asp")
	if err != nil {
		fmt.Println(getRandError())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = body[strings.Index(string(body), ">")+1:]
	if err != nil {
		fmt.Println(getRandError())
		return
	}
	v := ValCurs{}
	xml.Unmarshal(body, &v)
	var res_values map[string]string = make(map[string]string, 0)
	for _, vv := range v.Valutes {
		res_values[vv.CharCode] = vv.Value
	}
	usd := res_values["USD"]
	fmt.Printf("\n\n\033[1;31m  USD:    %v\n", usd)
	fmt.Printf("  EUR:    %v\033[0m\n", res_values["EUR"])
	fmt.Printf("\033[37;1m")
	fmt.Println(draw)
	fmt.Printf("\033[0m")
}
