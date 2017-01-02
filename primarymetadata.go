package yum

import (
	"encoding/xml"
	"fmt"
	"io"
)

// PrimaryMetadata represents the metadata XML file for a RPM/Yum repository. It
// contains packages available in the repository.
type PrimaryMetadata struct {
	XMLName       xml.Name `xml:"metadata"`
	XMLNS         string   `xml:"xmlns,attr"`
	PackagesCount int      `xml:"packages,attr"`

	Packages PackageEntries `xml:"package"`
}

// ReadPrimaryMetadata loads a primary.xml file from the given io.Reader and returns
// a pointer to the resulting PrimaryMetadata struct.
func ReadPrimaryMetadata(r io.Reader) (*PrimaryMetadata, error) {
	md := PrimaryMetadata{
		Packages: make([]PackageEntry, 0),
	}

	decoder := xml.NewDecoder(r)
	err := decoder.Decode(&md)

	if err != nil {
		return nil, fmt.Errorf("Error decoding primary metadata: %v", err)
	}

	return &md, nil
}
