package main

import (
	"encoding/json"
	"gobyexample/service"
	"net/http"
	"strings"
	"sync"
)

var responseGoByExample map[string]string

func mainRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	path := strings.Split(r.URL.Path, "/")[1]
	if path == "" {
		once := sync.Once(sync.Once{})
		once.Do(func() {
			if responseGoByExample == nil {
				dataMain := service.ScrapeMain()
				jsons, err := json.Marshal(map[string]any{
					"message": "Scrape Go By Example",
					"total":   len(dataMain),
					"data":    dataMain,
				})
				if err != nil {
					w.WriteHeader(500)
				}

				responseGoByExample = dataMain
				w.WriteHeader(200)
				w.Write(jsons)
			} else {
				json, err := json.Marshal(map[string]any{
					"message": "Scrape Go By Example",
					"total":   len(responseGoByExample),
					"data":    responseGoByExample,
					"cache":   true,
				})
				if err != nil {
					w.WriteHeader(500)
				}

				w.WriteHeader(200)
				w.Write(json)
			}
		})
	} else {
		if responseGoByExample != nil {
			w.Header().Del("Content-Type")
			lq := responseGoByExample[path]

			q := service.ScrapeSubMain(lq)
			w.WriteHeader(200)
			w.Write([]byte(q))
		} else {
			w.Header().Del("Content-Type")
			dataMain := service.ScrapeMain()
			responseGoByExample = dataMain
			lq := dataMain[path]

			q := service.ScrapeSubMain(lq)
			w.WriteHeader(200)
			w.Write([]byte(q))
		}
	}

}
func main() {
	http.HandleFunc("/", mainRoute)
	http.ListenAndServe(":8080", nil)
}
