package articles

import (
	"github.com/go-martini/martini"
	"github.com/madhums/go-martini-mgo-demo/models"
	"github.com/martini-contrib/render"
	"gopkg.in/mgo.v2"
)

/**
 * List
 */

func List(r render.Render, params martini.Params, db *mgo.Database) {

	var articles []models.Article

	err := db.C("articles").Find(nil).All(&articles)

	if err != nil {
		r.Error(400)
	}

	r.HTML(200, "articles/list", articles)
}
