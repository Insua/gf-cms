package migrate

import (
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
	"time"
)

var templateContent = `
-- +migrate Up

-- +migrate Down
`
var tpl *template.Template

func init() {
	tpl = template.Must(template.New("new_migration").Parse(templateContent))
}

func New(fileName string) {
	fileNameWithTimeStamp := fmt.Sprintf("%s-%s.sql", time.Now().Format("20060102150405"), strings.TrimSpace(fileName))
	pathName := path.Join("database/migrations", fileNameWithTimeStamp)
	f, err := os.Create(pathName)

	if err != nil {
		fmt.Println("Create migration files error")
		os.Exit(1)
	}
	defer func() { _ = f.Close() }()

	if err := tpl.Execute(f, nil); err != nil {
		fmt.Println("Create migration files error")
	}
}
