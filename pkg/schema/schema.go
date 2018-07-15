package schema

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	strmangle "github.com/volatiletech/sqlboiler/strmangle"
)

// this map is based on sdformat/sdf/schema/<version>/types.sdf
// it contains also simple types not included in file above
var typeSdfToGo = map[string]string{
	"double":       "float64",
	"unsigned int": "uint32",
	"int":          "int32",
	"vector3":      "string",
	"quaternion":   "string",
	"vector2d":     "string",
	"vector2i":     "string",
	"pose":         "string",
	"time":         "float64",
	"color":        "string",
}

// API used in templates
type API struct {
	Types []*Element `xml:"-"`
}

// Element represents sdf schema element
type Element struct {
	Name        string     `xml:"name,attr"`
	ID          string     `xml:"-"`
	GoName      string     `xml:"-"`
	Type        string     `xml:"type,attr"`
	GoType      string     `xml:"-"`
	XMLType     string     `xml:"-"`
	Ref         string     `xml:"ref,attr"`
	Required    string     `xml:"required,attr"`
	Default     string     `xml:"default,attr"`
	Attributes  []*Element `xml:"attribute"`
	Children    []*Element `xml:"element"`
	Includes    []*Include `xml:"include"`
	Description string     `xml:"description"`
	Filename    string     `xml:"-"`
}

// Include represents sdf schema import
type Include struct {
	Filename string `xml:"filename,attr"`
	Required string `xml:"required,attr"`
}

func trimExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}

func (api *API) getElementByFileName(fileName string) (e *Element, err error) {
	for _, t := range api.Types {
		if t.Filename == fileName {
			return t, nil
		}
	}
	return nil, fmt.Errorf("Element with filename %s not found", fileName)
}

func (api *API) getElementByID(id string) (e *Element, err error) {
	for _, t := range api.Types {
		if t.ID == id {
			return t, nil
		}
	}
	return nil, fmt.Errorf("Element with ID %s not found", id)
}

func (e *Element) resolveIncludes(api *API) error {
	for _, c := range e.Children {
		c.resolveIncludes(api)
	}
	for _, i := range e.Includes {
		definition, err := api.getElementByFileName(i.Filename)
		if err != nil {
			return err
		}
		e.Children = append(e.Children, definition)
	}
	return nil
}

func (e *Element) resolveGoNames() {
	if e.ID != "" {
		e.GoName = strmangle.TitleCase(e.ID)
	} else {
		e.GoName = strmangle.TitleCase(e.Name)
	}

	for _, c := range e.Children {
		c.resolveGoNames()

	}
	for _, a := range e.Attributes {
		a.resolveGoNames()
	}
}

func (e *Element) resolveXMLType() {
	if len(e.Attributes)+len(e.Children)+len(e.Includes) == 0 {
		e.XMLType = "Simple"
	}
	if len(e.Children)+len(e.Includes) != 0 {
		e.XMLType = "Complex"
	} else {
		e.XMLType = "SimpleWithAttribute"
	}
}

func (e *Element) resolveGoTypes(api *API) error {
	e.resolveXMLType()
	t, found := typeSdfToGo[e.Type]
	if found {
		e.GoType = t
	} else {
		e.GoType = e.Type
	}
	if e.Ref != "" {
		t, err := api.getElementByID(e.Ref)
		if err != nil {
			return err
		}
		e.GoType = "*" + t.GoName
		e.Type = t.GoName
	}
	if e.ID != "" {
		e.GoType = "*" + e.GoName
		e.Type = e.GoName
	}
	if e.Required == "*" {
		e.GoType = "[]" + e.GoType
	}
	for _, c := range e.Children {
		err := c.resolveGoTypes(api)
		if err != nil {
			return err
		}
	}
	for _, a := range e.Attributes {
		err := a.resolveGoTypes(api)
		if err != nil {
			return err
		}
	}
	return nil
}

func (api *API) resolveProperties() error {
	for _, t := range api.Types {
		t.ID = trimExtension(t.Filename)
		t.resolveGoNames()
		err := t.resolveGoTypes(api)
		if err != nil {
			return err
		}
		err = t.resolveIncludes(api)
		if err != nil {
			return err
		}

	}
	return nil
}

// MakeAPI creates API based on schema
func MakeAPI(dir string) (api *API, _ error) {
	api = &API{}

	re := regexp.MustCompile(`.*\.sdf$`)
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}

		if re.MatchString(path) == false {
			return nil
		}

		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		fileName := f.Name()

		e := &Element{
			Filename: fileName,
		}
		err = xml.Unmarshal(bytes, e)
		if err != nil {
			return err
		}

		api.Types = append(api.Types, e)

		return nil
	})
	if err != nil {
		return nil, err
	}

	err = api.resolveProperties()
	if err != nil {
		return nil, err
	}

	return api, nil
}
