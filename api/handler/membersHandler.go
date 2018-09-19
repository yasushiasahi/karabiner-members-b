package handler

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var url = "https://www.karabiner.tech"

// Member is screbed data
type Member struct {
	Img   string `json:"img"`
	Title string `json:"title"`
	Name  string `json:"name"`
}

// MembersHandler handles api for users
type MembersHandler struct{}

// ServeHTTP ...
func (h MembersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	w.Header().Set("Access-Control-Allow-Origin", "*")

	switch r.Method {
	case "GET":
		h.get(w, r)
		return
	default:
		http.NotFound(w, r)
	}
}

func (h MembersHandler) get(w http.ResponseWriter, r *http.Request) {
	ms, err := scrape()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(ms)
	if err != nil {
		http.Error(w, "failed to Encode json"+err.Error(), http.StatusInternalServerError)
	}
}

func scrape() ([]Member, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("failed to get response from www.karabiner.tecj")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var ms []Member
	doc.Find(".memberListItem").Each(func(i int, s *goquery.Selection) {
		m := Member{Img: url}
		src, _ := s.Find(".memberImg").Find("img").Attr("src")
		m.Img += src
		m.Title = s.Find(".nameTitle").Text()
		m.Name = s.Find(".memberName").Text()
		ms = append(ms, m)
	})
	return ms, nil
}
