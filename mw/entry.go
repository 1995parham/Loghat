package mw

import "encoding/xml"

// EntryList represents list of entries.
type EntryList struct {
	XMLName     xml.Name `xml:"entry_list"`
	Version     string   `xml:"version,attr"`
	Entries     []Entry  `xml:"entry"`
	Suggestions []string `xml:"suggestion"`
}

// Entry represents merriam webster word definition section
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	ID      string   `xml:"id,attr"`
	Def     []DT     `xml:"def>dt"` // Defining Text
	FL      string   `xml:"fl"`     // Functional Label
	Date    string   `xml:"def>date"`
	Word    string   `xml:"ew"`
}

// DT represents holder for word defintion xml
type DT struct {
	Sx []string `xml:"sx"`
	V  string   `xml:",innerxml"`
}
