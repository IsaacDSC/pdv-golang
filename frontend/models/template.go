package models

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Template struct {
	AppName string `json:"appName"`
	Navbar  struct {
		Logo    string `json:"logo"`
		Actions []struct {
			Title string `json:"title"`
			Path  string `json:"path"`
		} `json:"actions"`
	} `json:"navbar"`
	Footer struct {
		Contact struct {
			LinkLocation  string `json:"linkLocation"`
			LocationTitle string `json:"locationTitle"`
			Telephone     string `json:"telephone"`
			Mailer        string `json:"mailer"`
		} `json:"contact"`
		Operation struct {
			Title string `json:"title"`
			Hours []struct {
				Subtile string `json:"subtile"`
				Hour    string `json:"hour"`
			} `json:"hours"`
		} `json:"operation"`
	} `json:"footer"`
	Pages []struct {
		Home struct {
			Banner struct {
				Title           string `json:"title"`
				Description     string `json:"description"`
				CountActionName string `json:"countActionName"`
				CountActionLink string `json:"countActionLink"`
				Image           string `json:"image"`
			} `json:"banner"`
			Highlights []struct {
				Title    string `json:"title"`
				Discount int    `json:"discount"`
				Image    string `json:"image"`
			} `json:"highlights"`
			Menu struct {
				Categories []string `json:"categories"`
				Products   []struct {
					ID          string  `json:"id"`
					Category    string  `json:"category"`
					Title       string  `json:"title"`
					Description string  `json:"description"`
					Image       string  `json:"image"`
					Price       float64 `json:"price"`
					TypeMoney   string  `json:"typeMoney"`
				} `json:"products"`
			} `json:"menu"`
		} `json:"home"`
	} `json:"pages"`
	Delivery []struct {
		ID           string  `json:"id"`
		State        string  `json:"state"`
		City         string  `json:"city"`
		Neighborhood string  `json:"neighborhood"`
		Price        float32 `json:"price"`
	} `json:"delivery"`
}

func (thisTemplate *Template) GetTemplate() (err error) {
	templateJson, err := os.Open(`/Users/isaacdsc/Solutions/studies/Golang/Github/pdv-golang/frontend/views/template/store.json`)
	if err != nil {
		log.Fatalf(err.Error())
	}
	defer templateJson.Close()
	byteValueJSON, err := io.ReadAll(templateJson)
	json.Unmarshal(byteValueJSON, &thisTemplate)
	return
}
