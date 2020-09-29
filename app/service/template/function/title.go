package function

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func Title(title interface{}) string {
	if title == nil {
		return g.Cfg().GetString("viewer.Title")
	}
	return gconv.String(title)
}
