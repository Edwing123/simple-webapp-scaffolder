package scaffold

import (
	"bytes"
	"text/template"

	"github.edwing123.com/ruth/pkg/resources"
)

func newTemplate() *template.Template {
	tmp := template.Must(
		template.New("views").ParseFS(
			resources.ViewsFS,
			"resources/views/*.go.html",
		),
	)

	return tmp
}

var (
	layoutTmp = newTemplate()
)

func executeLayout(ctx map[string]any) (string, error) {
	var buff bytes.Buffer

	err := layoutTmp.ExecuteTemplate(&buff, "layout", ctx)
	if err != nil {
		return "", err
	}

	return buff.String(), nil
}
