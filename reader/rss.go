package reader

import (
	"bytes"
	"encoding/xml"
	"html/template"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

// Error message for when the XML is expected to be in Atom format
const errorExpectedAtom = "expected element type <rss> but have <feed>"

// Error message for when the XML is expecting ISO-8859-1 encoding
const errorIso8859 = "xml: encoding \"iso-8859-1\" declared but Decoder.CharsetReader is nil"

// RSS item
type Item struct {
	Title       string        `xml:"title" json:"title"`
	Link        string        `xml:"link" json:"link"`
	Description template.HTML `xml:"description" json:"description"`
	Content     template.HTML `xml:"encoded" json:"content"`
	PubDate     string        `xml:"pubDate" json:"pubDate"`
	Comments    string        `xml:"comments" json:"comments"`
}

// RSS version 2.0 structure
type Rss struct {
	XMLName     xml.Name `xml:"rss"`
	Version     string   `xml:"version,attr"`
	Title       string   `xml:"channel>title"`
	Link        string   `xml:"channel>link"`
	Description string   `xml:"channel>description"`
	PubDate     string   `xml:"channel>pubDate"`
	Items       []Item   `xml:"channel>item"`
}

// Parse Atom feed with the supplied content
func (r *Rss) ParseAtom(content []byte) error {
	a := Atom{}

	d := createDecoder(content)
	err := d.Decode(&a)
	if err != nil {
		return err
	}

	r = a.ToRss()
	return nil
}

// Parse RSS feed from the given URI
func (r *Rss) ParseFeed(content []byte) error {
	d := createDecoder(content)
	err := d.Decode(&r)
	if err != nil {
		// If it appears to be an atom feed go for it
		if err.Error() == errorExpectedAtom {
			return r.ParseAtom(content)
		}
		return err
	}

	if r.Version == "2.0" {
		for i, _ := range r.Items {
			if r.Items[i].Content != "" {
				r.Items[i].Description = r.Items[i].Content
			}
		}
	}

	return nil
}

func createDecoder(content []byte) *xml.Decoder {
	// Create an io.Reader from the array of bytes
	b := bytes.NewReader(content)

	// Create a new XML decoder with the ability to handle multiple character sets
	d := xml.NewDecoder(b)
	d.CharsetReader = charset.NewReader
	return d
}