package handlers

import (
	"net/http"

	"github.com/razzkumar/gofullstack/volume3/section4/gopherface/common"

	"go.isomorphicgo.org/go/isokit"
)

func FeedHandler(env *common.Env) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := make(map[string]string)
		m["PageTitle"] = "Feed"
		env.TemplateSet.Render("feed_page", &isokit.RenderParams{Writer: w, Data: m})
	})
}
