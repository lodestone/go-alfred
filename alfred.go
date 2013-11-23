package Alfred

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
)

var noResultString string

type GoAlfred struct {
	bundleID  string
	results   items
	dataDir   string
	bundleDir string
	id        string
}

type AlfredIcon struct {
	Filename string `xml:",chardata"`
	Type     string `xml:"type,attr,omitempty"`
}

type item struct {
	XMLName      xml.Name   `xml:"item"`
	Uid          string     `xml:"uid,attr"`
	Arg          string     `xml:"arg,attr"`
	Type         string     `xml:"type,attr,omitempty"`
	Valid        string     `xml:"valid,attr,omitempty"`
	AutoComplete string     `xml:"autocomplete,attr,omitempty"`
	Title        string     `xml:"tittle"`
	SubTitle     string     `xml:"subtitle"`
	Icon         AlfredIcon `xml:"icon"`
}

type items struct {
	XMLName xml.Name `xml:"items"`
	Results []item
}

func NewAlfred(id string) *GoAlfred {
	ga := new(GoAlfred)
	ga.init(id)
	return ga
}

func (ga *GoAlfred) SetNoResultTxt(title string) {
	noResultString = title
}

func (ga *GoAlfred) XML() (output []byte, err error) {
	output, err = ga.results.toXML()
	if err != nil {
		output = nil
	}
	return output, nil
}

func (ga *GoAlfred) WriteToAlfred() {
	output := ga.results.toXML()
	os.Stdout.Write(output)
}

func (ga *GoAlfred) init(id string) {
	ga.id = id
	// Get bundleid
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("go-alfred: Can't initiate: %v", err)
	}
	plistfn := path.Join(pwd, "info.plist")
	_, err = os.Stat(plistfn)
	if err != nil {
		fmt.Println("It's working", plistfn)
	}

func (ga *GoAlfred) AddItem(title, subtitle, valid, auto, rtype, arg string,
	icon AlfredIcon) {
	if title == "" {
		title = noResultString
	}
	r := item{Arg: arg, Type: rtype, Valid: valid, AutoComplete: auto,
		Title: title, SubTitle: subtitle}
	r.Icon = icon
	ga.results.Results = append(ga.results.Results, r)
}

func (results *items) toXML() (output []byte, err error) {
	output, err = xml.MarshalIndent(results, "", "  ")
	if err != nil {
		output = nil
	}
	return output, err
}

func NewIcon(fn, itype string) (ico AlfredIcon) {
	return AlfredIcon{Filename: fn, Type: "icontype"}
}
