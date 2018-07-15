package schema

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/flosch/pongo2"
	yaml "gopkg.in/yaml.v2"
)

const (
	templateRoot = "tools/templates"
)

type templateConfig struct {
	TemplateFilePath string `yaml:"template_file"`
	OutputFilePath   string `yaml:"output_file"`
}

func ensureOutputDirectory(path string) error {
	return os.MkdirAll(filepath.Dir(path), 0644)
}

func loadConfig(path string) (config []*templateConfig, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *templateConfig) applyTemplate(api *API) error {
	t, err := pongo2.FromFile(filepath.Join(templateRoot, c.TemplateFilePath))
	if err != nil {
		return err
	}

	if err = ensureOutputDirectory(c.OutputFilePath); err != nil {
		return err
	}

	output, err := t.Execute(pongo2.Context{"types": api.Types})
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.OutputFilePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

// ApplyTemplates loads template config and generates code using api
func ApplyTemplates(configPath string, api *API) error {
	config, err := loadConfig(configPath)
	if err != nil {
		return err
	}

	for _, c := range config {
		err = c.applyTemplate(api)
		if err != nil {
			return err
		}
	}

	return nil
}
