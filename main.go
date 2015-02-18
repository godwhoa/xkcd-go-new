package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
)

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

func HomeHandle(w http.ResponseWriter, r *http.Request) {
	//Get json data
	res, err := http.Get("http://xkcd.com/info.0.json")
	perror(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	perror(err)
	//Unmarshal json
	comic := Comic{}
	json.Unmarshal(body, &comic)
	//Render template
	t := template.Must(template.ParseFiles("index.html"))
	t.Execute(w, comic)

}

func main() {
	http.HandleFunc("/", HomeHandle)
	http.ListenAndServe(":8080", nil)
}

//Error handling
func perror(err error) {
	if err != nil {
		panic(err)
	}
}
