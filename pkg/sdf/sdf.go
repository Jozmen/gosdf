package sdf

import (
	"encoding/xml"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Sdf SDF Document
type Sdf struct {
	XMLName struct{} `xml:"sdf"`
	Root    `yaml:"sdf"`
}

// XMLToYaml Converts XML SDF Model to Yaml
func XMLToYaml(input, output string) {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	sdf := Sdf{}
	err = xml.Unmarshal(bytes, &sdf)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err = yaml.Marshal(sdf)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(output, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// YamlToXML Converts Yaml SDF Model to XML
func YamlToXML(input, output string) {
	bytes, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal(err)
	}
	sdf := Sdf{}
	err = yaml.Unmarshal(bytes, &sdf)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err = xml.Marshal(sdf)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(output, bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
