package main

import (
  "fmt"
  "net/http"
  "net/url"
  "strconv"
  "strings"
)

var likes map[int64]int64

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
- http://golang.org/pkg/net/http/#Error
- http://learntogoogleit.com/post/63098708081/returning-status-codes-in-golang
*/
func (l LikeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  var id, err = strconv.ParseInt(LastPathPart(r.URL), 10, 64)
  if err == nil {
    likes[id] ++
    fmt.Fprintf(w, "%+v", likes)
  } else {
    http.Error(w, fmt.Sprintf("%v", err), 400)
  }
}

func LastPathPart(u *url.URL) string {
  var parts = strings.Split(u.Path, "/")
  return parts[len(parts) - 1]
}

func main() {
  likes = make(map[int64]int64)
  var lh LikeHandler
  http.Handle("/like/", lh)
	http.ListenAndServe("localhost:4000", nil)
}
