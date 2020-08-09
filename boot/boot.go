package boot

import (
	_ "gf-cms/packed"

	"github.com/gogf/gf/frame/g"
)

func init() {
	g.DB().PingMaster()
	g.DB().PingSlave()
}
