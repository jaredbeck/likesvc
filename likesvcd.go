package main

import (
  "fmt"
  "net/http"
  "net/url"
  "strings"
)

var likes map[string]int

/*
`LikeHandler` implements `http.Handler`
- http://tour.golang.org/#57
- http://golang.org/pkg/net/http/
*/
type LikeHandler struct{}

/*
`ServeHTTP` is a method defined *on* the `LikeHandler` type.
`LikeHandler` is the *method receiver*.  The receiver can be "by
reference" (with a `*`) or "by value".
- http://tour.golang.org/#50
- http://tour.golang.org/#52
*/
func (l LikeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  var id = LastPathPart(r.URL)
  likes[id] ++

  /* when printing structs, the plus flag (%+v) adds field names
  http://golang.org/pkg/fmt */
  fmt.Fprintf(w, "%+v", likes)
}

func LastPathPart(u *url.URL) string {
  var parts = strings.Split(u.Path, "/")
  return parts[len(parts) - 1]
}

func main() {
  likes = make(map[string]int)
  var lh LikeHandler
  http.Handle("/like/", lh)
	http.ListenAndServe("localhost:4000", nil)
}
