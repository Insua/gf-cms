package boot

import (
	"gf-cms/app/service/template/function"

	"github.com/gogf/gf/frame/g"
)

func init() {
	g.View().BindFunc("title", function.Title)
	g.View().BindFunc("mix", function.Mix)
}
