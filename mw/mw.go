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

// MerriamWebster contains information for create connection into merriam-webster
type MerriamWebster struct {
	key        string
	dictionary string
}

// Word type contains word defination
type Word struct {
	Word        string
	Version     string
	Definitions []Definition
}

// Definition represents word definition with some description
type Definition struct {
	EWord  string
	Date   string
	Def    []string
	Type   string
	Extera string
}

// New creates merriam-webster object with given key and dictionary
func New(key string, dictionary string) *MerriamWebster {
	return &MerriamWebster{
		key:        key,
		dictionary: dictionary,
	}
}

// Fetch word definition
func (m *MerriamWebster) Fetch(word string) (*Word, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.dictionaryapi.com/api/v1/references/%s/xml/%s?key=%s", m.dictionary, word, m.key))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	var el EntryList
	if err = xml.NewDecoder(resp.Body).Decode(&el); err != nil {
		return nil, err
	}

	var w Word
	w.Version = el.Version
	for _, e := range el.Entries {
		w.Definitions = append(w.Definitions, Definition{
			Date:  e.Date,
			Def:   e.Def,
			Type:  e.Type,
			EWord: e.Word,
		})
	}

	return &w, nil
}
