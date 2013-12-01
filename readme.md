likesvcd
========

A proof-of-concept HTTP server in [golang][1], counts "likes".

    go run likesvcd.go &
    curl http://localhost:4000/like/1
    map[1:1]
    curl http://localhost:4000/like/2
    map[1:1 2:1]
    ulimit -n 600 # increases limit on open files
    ab -c 500 -n 500 -r http://127.0.0.1:4000/like/1
    Requests per second:    8220.30 [#/sec] (mean)
    curl http://localhost:4000/like/1
    map[1:502 2:1]

Apache 2.2, by comparison, serving a static file, configured
to with the `prefork` [MPM][2] and `MaxClients 500`:

    echo "herp de derp" > ~/Sites/derp.txt
    ab -c 500 -n 500 -r http://127.0.0.1/~jared/derp.txt
    Requests per second:    4107.08 [#/sec] (mean)

Admittedly, apache has far more burdens (aka. features), and I'm not
qualified to configure it ideally.

[1]: http://golang.org
[2]: http://httpd.apache.org/docs/2.2/mpm.html
