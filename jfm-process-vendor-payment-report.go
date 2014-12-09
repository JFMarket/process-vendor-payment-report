package main

import (
	"encoding/csv"
	"github.com/codegangsta/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"sort"
	"strconv"
	// "fmt"
)

func main() {
	m := martini.Classic()

	m.Use(martini.Static("app"))

	m.Use(render.Renderer())

	m.Get("/", func(r render.Render) {
		r.HTML(200, "index", nil)
	})

	m.Post("/upload", upload)

	m.Run()
}

type SoldItem struct {
	SalesRecordId    string
	SaleDate         string
	CustomerRecordId string
	ItemDescription  string
	Department       string
	Category         string
	UPCCode          string
	StoreCode        string
	UnitPrice        float64
	Quantity         float64
	TotalPrice       float64
	Discount         float64
	Total            float64
	CostOfGoodsSold  float64
	RegisterNumber   int
	SupplierName     string
}

func SoldItemFromStrings(s []string) *SoldItem {
	return &SoldItem{
		s[0],
		s[1],
		s[2],
		s[3],
		s[4],
		s[5],
		s[6],
		s[7],
		parseFloat(s[8]),
		parseFloat(s[9]),
		parseFloat(s[10]),
		parseFloat(s[11]),
		parseFloat(s[12]),
		parseFloat(s[13]),
		parseInt(s[14]),
		s[15],
	}
}

type SaleDataByVendor struct {
	Name  string
	Total float64
}

type ByName []SaleDataByVendor

func (n ByName) Len() int           { return len(n) }
func (n ByName) Swap(i, j int)      { n[i], n[j] = n[j], n[i] }
func (n ByName) Less(i, j int) bool { return n[i].Name < n[j].Name }

func upload(req *http.Request, r render.Render) {
	file, _, err := req.FormFile("csv")
	if err != nil {
		// Respond with JSON error that can be displayed on page
		r.JSON(500, map[string]interface{}{"error": err})
		return
	}
	defer file.Close()

	c := csv.NewReader(file)

	lines, err := c.ReadAll()
	if err != nil {
		// Respond with JSON error that can be displayed on page
		r.JSON(500, map[string]interface{}{"error": err})
		return
	}

	// One line is the header
	items := make([]*SoldItem, len(lines)-1)

	// Skip header line
	for i := 1; i < len(lines); i++ {
		items[i-1] = SoldItemFromStrings(lines[i])
	}

	// Group items by farmers
	farmers := make(map[string][]*SoldItem)

	for _, item := range items {
		farmers[item.SupplierName] = append(farmers[item.SupplierName], item)
	}

	// Calculate totals and only send names and totals to the client
	var data []SaleDataByVendor

	for n, items := range farmers {
		var total float64

		for _, i := range items {
			total = total + i.Total
		}

		data = append(data, SaleDataByVendor{n, total})
	}

	sort.Sort(ByName(data))

	r.JSON(200, data)
}

func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

func parseInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
