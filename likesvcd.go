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
*/
func (l LikeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  var id, err = strconv.ParseInt(LastPathPart(r.URL), 10, 64)
  if err == nil {
    likes[id] ++
  } else {
    fmt.Fprintf(w, "%v\n", err)
  }

  /* when printing structs, the plus flag (%+v) adds field names
  http://golang.org/pkg/fmt */
  fmt.Fprintf(w, "%+v", likes)
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
