package articles

import (
	"github.com/go-martini/martini"
	"github.com/madhums/go-martini-mgo-demo/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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
