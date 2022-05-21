package action

import (
	"github.com/fyyang/goctl-api2proto/generate"
	"github.com/urfave/cli/v2"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Generator(ctx *cli.Context) error {
	fileName := ctx.String("filename")

	if len(fileName) == 0 {
		fileName = "filename.proto"
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	return generate.Do(fileName, p)
}
