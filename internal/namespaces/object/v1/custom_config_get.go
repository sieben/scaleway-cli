package object

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func configGetCommand() *core.Command {
	type getArgs struct {
		Region scw.Region
		Type   s3tool
		Name   string
	}

	return &core.Command{
		Namespace: "object",
		Resource:  "config",
		Verb:      "get",
		Short:     "Generate a S3 tool configuration file",
		Long:      "Generate a S3 tool configuration file.",
		ArgsType:  reflect.TypeOf(getArgs{}),
		ArgSpecs: []*core.ArgSpec{
			{
				Name:       "type",
				Short:      "Type of S3 tool you want to generate a config for",
				Required:   true,
				EnumValues: supportedTools.ToStringArray(),
			},
			{
				Name:     "name",
				Short:    "Name of the s3 remote you want to generate",
				Required: false,
				Default: func() (value string, doc string) {
					return "scaleway", "scaleway"
				},
			},
			core.RegionArgSpec(scw.RegionFrPar, scw.RegionNlAms),
		},
		Examples: []*core.Example{
			{
				Short:   "Generate a s3cmd config file for Paris region",
				Request: `{"type": "s3cmd", "region": "fr-par"}`,
			},
			{
				Short:   "Generate a rclone config file for default region",
				Request: `{"type": "rclone"}`,
			},
			{
				Short:   "Generate a mc (minio) config file for default region",
				Request: `{"type": "mc"}`,
			},
		},
		SeeAlsos: []*core.SeeAlso{
			{
				Short:   "Install a S3 tool configuration file",
				Command: "scw object config install",
			},
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*getArgs)
			region := args.Region
			name := args.Name

			config, err := newS3Config(ctx, region, name)
			if err != nil {
				return "", err
			}
			return config.getConfigFile(args.Type)
		},
	}
}