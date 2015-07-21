package middlewares

import (
	"fmt"
	"os"

	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/github.com/go-martini/martini"
	. "github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/github.com/tj/go-debug"
	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/gopkg.in/mgo.v2"
)

/**
 * Connect to mongo and make `db *mgo.Database` available for all handlers
 */

func Connect() martini.Handler {

	debug := Debug("middlewares:connect")
	uri := os.Getenv("MONGODB_URL")

	if uri == "" {
		uri = "mongodb://localhost:27017/magazine_app"
	}

	mInfo, err := mgo.ParseURL(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	session.SetSafe(&mgo.Safe{})
	debug("Connecting to %s", uri)

	return func(c martini.Context) {
		s := session.Clone()
		c.Map(s.DB(mInfo.Database))
	}
}
