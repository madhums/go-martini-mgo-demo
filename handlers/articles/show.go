package articles

import (
	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/github.com/go-martini/martini"
	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/github.com/martini-contrib/render"
	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/gopkg.in/mgo.v2"
	"github.com/madhums/go-martini-mgo-demo/Godeps/_workspace/src/gopkg.in/mgo.v2/bson"
	"github.com/madhums/go-martini-mgo-demo/models"
)

/**
 * Show
 */

func Show(params martini.Params, r render.Render, db *mgo.Database) {

	article := models.Article{}
	oId := bson.ObjectIdHex(params["_id"])

	err := db.C("articles").FindId(oId).One(&article)

	if err != nil {
		r.HTML(400, "400", err)
	}

	r.HTML(200, "articles/show", article)
}
