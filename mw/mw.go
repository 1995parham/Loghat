/*
 * +===============================================
 * | Author:        Parham Alvani <parham.alvani@gmail.com>
 * |
 * | Creation Date: 01-12-2017
 * |
 * | File Name:     mw.go
 * +===============================================
 */

// Package mw is merriam webster api client package
package mw

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// MerriamWebster type contains word defination
type MerriamWebster struct {
	Word    string
	Version string
}

// Word represents word definition with some description
type Word struct {
	Date   string
	Def    string
	Extera string
}

// New creates merriam-webster object contains given word definition
func New(word string) *MerriamWebster {
	return &MerriamWebster{
		Word: word,
	}
}

// Fetch word definition
func (m *MerriamWebster) Fetch() error {
	resp, err := http.Get(fmt.Sprintf("https://www.dictionaryapi.com/api/v1/references/collegiate/xml/%s?key=8159d5d5-7a05-4e4d-a1d5-001b81ceb618", m.Word))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	var el EntryList
	if err = xml.NewDecoder(resp.Body).Decode(&el); err != nil {
		return err
	}

	m.Version = el.Version
	for _, e := range el.Entries {
		fmt.Println(e.Def)
	}

	return nil
}
