package main

import (
	"github.com/gin-gonic/gin"
	"github.com/silsgah/gowebportal/httpd/handler"
	"github.com/silsgah/gowebportal/platform/newsfeed"
)

func main() {
	feed := newsfeed.New()
	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/ping", handler.NewsfeedPost(feed))
	r.Run()
	// feed := newsfeed.New()
	// fmt.Println(feed)
	// feed.Add(newsfeed.Item{"Hello", "Someone out there"})
	// fmt.Println(feed)
	// r := mux.NewRouter()
	// r.HandleFunc("/hello", handler).Methods("GET")
	// http.ListenAndServe(":8080", r)
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World!")
// }
