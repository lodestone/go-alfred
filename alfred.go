package Alfred

import (
	"encoding/xml"
	"fmt"
	"os"
	// "path/filepath"
)

var noResultString string

type GoAlfred struct {
	bundleID  string
	results   items
	dataDir   string
	bundleDir string
}

type alfredIcon struct {
	Type     string `xml:"type,attr,omitempty"`
	Filename string `xml:"ico"`
}

type item struct {
	XMLName      xml.Name `xml:"item"`
	Arg          string   `xml:"arg,attr"`
	Type         string   `xml:"type,attr"`
	Valid        string   `xml:"valid,attr"`
	AutoComplete string   `xml:"autocomplete,attr"`
	Title        string   `xml:"tittle"`
	SubTitle     string   `xml:"subtitle"`
	alfredIcon
}

type items struct {
	XMLName xml.Name `xml:"items"`
	Results []item
}

func NewAlfred() *GoAlfred {
	return new(GoAlfred)
}

func (ga *GoAlfred) SetNoResultTxt(title string) {
	noResultString = title
}

func (ga *GoAlfred) AddItem(title, subtitle, valid, auto, rtype, arg string,
	icon alfredIcon) {
	if title == "" {
		title = noResultString
	}
	r := &item{Arg: arg, Type: rtype, Valid: valid, AutoComplete: auto,
		Title: title, SubTitle: subtitle}
	r.alfredIcon = icon
	ga.results.Results = append(ga.results.Results, r)
}

func (results *items) XML() []byte {
	output, err := xml.MarshalIndent(results, "", "  ")
	if err != nil {
		output = []byte(fmt.Sprintf("alfred.go error: %v\n", err))
	}
	return output
}

func NewIcon(fn, itype string) (ico *alfredIcon) {
	return &alfredIcon{Filename: fn, Type: itype}
}

func WriteToAlfred(results *items) {
	output := results.XML()
	os.Stdout.Write(output)
}
