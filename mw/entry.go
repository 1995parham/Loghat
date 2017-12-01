package mw

import "encoding/xml"

// EntryList represents list of entries.
type EntryList struct {
	XMLName xml.Name `xml:"entry_list"`
	Version string   `xml:"version,attr"`
	Entries []Entry  `xml:"entry"`
}

// Entry represents merriam webster word definition section
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	ID      string   `xml:"id,attr"`
	Def     string   `xml:"def>dt"`
}
