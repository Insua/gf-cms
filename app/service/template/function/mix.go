package function

import (
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gres"
	"github.com/gogf/gf/util/gconv"
)

func Mix(url string) string {
	var j *gjson.Json
	var err error
	if gfile.IsFile("public/mix-manifest.json") {
		j, err = gjson.Load("public/mix-manifest.json")
	} else {
		mixFile := gres.GetContent("public/mix-manifest.json")
		j, err = gjson.LoadContent(mixFile)
	}
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
