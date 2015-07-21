package articles

import (
	"fmt"

	"gopkg.in/mgo.v2"

	"github.com/madhums/go-martini-mgo-demo/models"
	"github.com/martini-contrib/render"
)

/**
 * Get add/edit form
 */

func AddEdit(r render.Render) {
	r.HTML(200, "articles/form", nil)
}

/**
 * Add
 */

func Add(article models.Article, r render.Render, db *mgo.Database) {
	fmt.Println(article)
	err := db.C("articles").Insert(article)

	if err != nil {
		r.HTML(400, "400", err)
	} else {
		r.Redirect("/")
	}

}
