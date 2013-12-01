likesvcd
========

A proof-of-concept HTTP server in [golang][1].

    go run likesvcd.go &
    curl http://localhost:4000/like/1
    map[1:1]
    curl http://localhost:4000/like/2
    map[1:1 2:1]
    ab -c 500 -n 500 -r http://127.0.0.1:4000/like/1
    Requests per second:    8220.30 [#/sec] (mean)
    curl http://localhost:4000/like/1
    map[1:502 2:1]

Testing
-------

    ulimit -n 600 # increases limit on open files
    ab -c 500 -n 500 -r http://127.0.0.1:4000/like/2

`ab` is an apache utility.  `-r` means "Don't exit on socket receive errors"

[1]: http://golang.org
