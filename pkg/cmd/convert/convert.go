package convert

import (
	log "github.com/sirupsen/logrus"
	cobra "github.com/spf13/cobra"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"

	"github.com/Jozmen/gosdf/pkg/sdf"
)

var input string
var output string

// RootCommand root convert command
var RootCommand = &cobra.Command{
	Use:   "convert",
	Long:  ``,
	Short: `SDF Convert Utility`,
}

var xmlToYaml = &cobra.Command{
	Use:   "xty",
	Short: `Converts XML SDF model into Yaml`,
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sdf.XMLToYaml(input, output)
	},
}

var yamlToXML = &cobra.Command{
	Use:   "ytx",
	Short: "Converts Yaml SDF model into XML",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sdf.YamlToXML(input, output)
	},
}

func init() {
	yamlToXML.Flags().StringVarP(&input, "input", "i", "", "Input File")
	xmlToYaml.Flags().StringVarP(&input, "input", "i", "", "Input File")
	yamlToXML.Flags().StringVarP(&output, "output", "o", "", "Output file")
	xmlToYaml.Flags().StringVarP(&output, "output", "o", "", "Output file")

	yamlToXML.MarkFlagRequired("input")
	xmlToYaml.MarkFlagRequired("input")
	yamlToXML.MarkFlagRequired("output")
	xmlToYaml.MarkFlagRequired("output")

	RootCommand.AddCommand(yamlToXML)
	RootCommand.AddCommand(xmlToYaml)

	log.SetFormatter(&prefixed.TextFormatter{})
	log.SetLevel(log.DebugLevel)
}
