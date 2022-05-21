package generate

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"log"
	"os"
	"text/template"
)

//go:embed proto.tpl
var apiTemplate string

func Do(filename string, in *plugin.Plugin) error {
	protoMessage, protoService, err := applyGenerate(in)
	if err != nil {
		fmt.Println(err)
	}

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
	}

	t := template.Must(template.New("service").Funcs(funcMap).Parse(apiTemplate))

	output := in.Dir + "/" + filename

	f, err := os.Create(output)
	defer f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	err = t.Execute(f, map[string]interface{}{
		"Message": protoMessage,
		"Service": protoService,
	})
	return err
}
