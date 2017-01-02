package yum

import (
	"fmt"
	"time"
)

// PackageEntry is a RPM package as defined in a yum repository database.
type PackageEntry struct {
	db *PrimaryDatabase

	Key         int
	Arch        string               `xml:"arch"`
	Size        PackageEntrySize     `xml:"size"`
	Checksums   PackageEntryChecksum `xml:"checksum"`
	Location    PackageEntryLocation `xml:"location"`
	PackageName string               `xml:"name"`
	Versions    PackageEntryVersion  `xml:"version"`
	Time        PackageEntryTime     `xml:"time"`
	Summary     string               `xml:"summary"`
	Url         string               `xml:"url"`
	Packager    string               `xml:"packager"`
}

type PackageEntrySize struct {
	Package   int64 `xml:"type,attr"`
	Installed int64 `xml:"installed,attr"`
	Archive   int64 `xml:"archive,attr"`
}

type PackageEntryVersion struct {
	Epoch   int    `xml:"epoch,attr"`
	Version string `xml:"ver,attr"`
	Release string `xml:"rel,attr"`
}

// PackageEntryChecksum is the XML element of a package metadata file which
// describes the checksum required to validate a package.
type PackageEntryChecksum struct {
	Type  string `xml:"type,attr"`
	Pkgid string `xml:"pkgid,attr"`
	Hash  string `xml:",chardata"`
}

// PackageEntryChecksum is the XML element of a package metadata file which
// describes the checksum required to validate a package.
type PackageEntryTime struct {
	File  int64 `xml:"file,attr"`
	Build int64 `xml:"build,attr"`
}

// RepoDatabaseLocation represents the URI, relative to a package repository,
// of a repository database.
type PackageEntryLocation struct {
	Href string `xml:"href,attr"`
}

// PackageEntries is a slice of PackageEntry structs.
type PackageEntries []PackageEntry

// String reassembles package metadata to form a standard rpm package name;
// including the package name, version, release and architecture.
func (c PackageEntry) String() string {
	return fmt.Sprintf("%s-%s-%s.%s", c.Name(), c.Version(), c.Release(), c.Architecture())
}

// LocationHref is the location of the package, relative to the parent
// repository.
func (c *PackageEntry) LocationHref() string {
	return c.Location.Href
}

func (c *PackageEntry) Checksum() (string, error) {
	return c.Checksums.Hash, nil
}

func (c *PackageEntry) ChecksumType() string {
	return c.Checksums.Type
}

func (c *PackageEntry) PackageSize() int64 {
	return c.Size.Package
}

func (c *PackageEntry) InstallSize() int64 {
	return c.Size.Installed
}

func (c *PackageEntry) ArchiveSize() int64 {
	return c.Size.Archive
}

func (c *PackageEntry) Name() string {
	return c.PackageName
}

func (c *PackageEntry) Version() string {
	return c.Versions.Version
}

func (c *PackageEntry) Release() string {
	return c.Versions.Release
}

func (c *PackageEntry) Architecture() string {
	return c.Arch
}
func (c *PackageEntry) Epoch() int {
	return c.Versions.Epoch
}

func (c *PackageEntry) BuildTime() time.Time {
	return time.Unix(c.Time.Build, 0)
}
