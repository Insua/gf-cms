package function

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/util/gconv"
)

func Mix(url string) string {
	if !gfile.IsFile("public/mix-manifest.json") {
		panic("dos not have mix manifest file")
	}

	j, err := gjson.Load("public/mix-manifest.json")
	if err != nil {
		panic(err)
	}

	j.SetViolenceCheck(true)
	if p := j.Get(url); p == nil {
		panic("wrong file path")
	} else {
		return gconv.String(p)
	}
}
