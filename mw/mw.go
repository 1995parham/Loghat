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

// MerriamWebster type contains word defination
type MerriamWebster struct {
	Word string
}

// New creates merriam-webster object contains given word definition
func New(word string) *MerriamWebster {
	return &MerriamWebster{
		Word: word,
	}
}

// Fetch word definition
func (m *MerriamWebster) Fetch() error {
	return nil
}
