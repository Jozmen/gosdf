package generate

import (
	"path/filepath"

	"github.com/Jozmen/gosdf/pkg/schema"
	log "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var source string

// GenerateCmd code generation command
var GenerateCmd = &cobra.Command{
	Use:  "generate",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		api, err := schema.MakeAPI(source)
		if err != nil {
			panic(err)
		}
		schema.ApplyTemplates(filepath.Join(schema.TemplateRoot, "template_config.yml"), api)
	},
}

func init() {
	GenerateCmd.Flags().StringVarP(&source, "source", "s", "", "Schema directory")
	GenerateCmd.MarkFlagRequired("source")

	log.SetFormatter(&prefixed.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}
