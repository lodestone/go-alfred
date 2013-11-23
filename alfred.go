package Alfred

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path"
)

var noResultString string = "No Result Were Found."

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
	Uid          string     `xml:"uid,attr,omitempty"`
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
		log.Printf("Can't locato info.plist: %v\n", plistfn)
	}
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
	output, err := ga.results.toXML()
	if err != nil {
		output = []byte(fmt.Sprintf("%v", err))
	}
	os.Stdout.Write(output)
}

func (ga *GoAlfred) AddItem(uid, title, subtitle, valid, auto, rtype,
	arg string, icon AlfredIcon, check_valid bool) {

	if title == "" {
		title = noResultString
	}
	r := item{Uid: uid, Arg: arg, Type: rtype, Valid: valid,
		AutoComplete: auto, Title: title, SubTitle: subtitle}
	if check_valid {
		// Make sure the item will work in Alfred as autocomplete if
		// 'auto' parameter is said
		r.make_valid()
	}
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

func (i *item) make_valid() {
	if (i.Valid == "" || i.Valid == "yes") && i.AutoComplete != "" {
		i.Valid = "no"
		i.Arg = ""
	}
}

func NewIcon(fn, itype string) (ico AlfredIcon) {
	return AlfredIcon{Filename: fn, Type: "icontype"}
}
