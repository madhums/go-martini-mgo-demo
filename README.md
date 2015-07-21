## go-martini-mgo-demo

A simple CRUD app in [golang](http://golang.org) and [mongodb](http://mongodb.org) using [martini](https://github.com/go-martini/martini), [mgo](https://github.com/go-mgo/mgo) and godep.

[Demo](http://go-martini-mgo-demo.herokuapp.com)

For the reloading of source files, I am using [gin](https://github.com/codegangsta/gin)

## Installation

For development, you can just clone the repo and run

```
$ go get && go install && PORT=7000 DEBUG=* gin -p 9000 run
```

Then visit `localhost:9000`

If you just want to run it and see then

```sh
$ go get github.com/madhums/go-martini-mgo-demo
$ PORT=7000 go-martini-mgo-demo # should start listening on port 7000
```

you can also specify `MONGODB_URL` env variable and it will connect to it.

## Todo

- Add features like content negotiation
- Image uploads
- REST API
- vendor experiment

Thanks to [go-martini](https://github.com/go-martini/martini), [go-mgo](https://github.com/go-mgo/mgo) and [godep](https://github.com/tools/godep)
