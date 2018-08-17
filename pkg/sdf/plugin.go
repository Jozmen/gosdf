package sdf

import (
	"encoding/json"
	"encoding/xml"
	"strings"

	log "github.com/sirupsen/logrus"
)

type attr struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type element struct {
	Name     string     `json:"name,omitempty"`
	Value    string     `json:"value,omitempty"`
	Attr     []*attr    `json:"attributes,omitempty"`
	Children []*element `json:"children,omitempty"`
}

// MarshalXML Marshals Plugin element with user-defined tags, not specified in sdf schema
func (p *Plugin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return nil
}

func newElement(start *xml.StartElement) *element {
	e := &element{
		Children: []*element{},
		Name:     start.Name.Local,
	}
	for _, a := range start.Attr {
		e.Attr = append(e.Attr, &attr{Name: a.Name.Local, Value: a.Value})
	}
	return e
}

func (p *Plugin) unmarshalXML(data *[]*element, e *element, d *xml.Decoder, start *xml.StartElement) error {
	log.Debugf("START: %v", start)
	*data = append(*data, e)

	var t xml.Token
	var err error

	for t, err = d.Token(); start.End() != t; t, err = d.Token() {
		if err != nil {
			return err
		}
		switch v := t.(type) {
		case xml.StartElement:
			log.Debugf("START: %v", v)
			newElem := newElement(&v)
			err = p.unmarshalXML(&e.Children, newElem, d, &v)
			if err != nil {
				return err
			}
		case xml.EndElement:
			log.Debugf("END: %v", v)
		case xml.CharData:
			e.Value = strings.TrimSpace(string(v))
			if e.Value != "" {
				log.Debugf("CHAR: [%s]", e.Value)
			}
		}
	}

	log.Debugf("END: %v", t)
	return nil
}

// UnmarshalXML Unmarshals Plugin element with user-defined tags, not specified in sdf schema
func (p *Plugin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	data := []*element{}
	e := newElement(&start)
	err := p.unmarshalXML(&data, e, d, &start)
	if err != nil {
		return err
	}

	for _, a := range e.Attr {
		switch a.Name {
		case "name":
			p.Name = a.Value
		case "filename":
			p.Filename = a.Value
		}
	}

	for _, e := range data {
		b, _ := json.MarshalIndent(e, "", "  ")
		log.Debugf("%s\n", string(b))
		p.copyData = append(p.copyData, e)
	}

	return nil
}
