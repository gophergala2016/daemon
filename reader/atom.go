package reader

import (
	"encoding/xml"
	"html/template"
)

type Link struct {
	Href string `xml:"href,attr"`
}

// Atom author
type Author struct {
	Name  string `xml:"name"`
	Email string `xml:"email"`
}

// Atom entry
type Entry struct {
	Title   string `xml:"title"`
	Summary string `xml:"summary"`
	Content string `xml:"content"`
	Id      string `xml:"id"`
	Updated string `xml:"updated"`
	Link    Link   `xml:"link"`
	Author  Author `xml:"author"`
}

// Atom version 1.0 structure
type Atom struct {
	XMLName  xml.Name `xml:"http://www.w3.org/2005/Atom feed"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Id       string   `xml:"id"`
	Updated  string   `xml:"updated"`
	Rights   string   `xml:"rights"`
	Link     Link     `xml:"link"`
	Author   Author   `xml:"author"`
	Entries  []Entry  `xml:"entry"`
}

// Convert from Atom to RSS
func (a *Atom) ToRss() *Rss {
	r := Rss{
		Title:       a.Title,
		Link:        a.Link.Href,
		Description: a.Subtitle,
		PubDate:     a.Updated,
	}
	r.Items = make([]Item, len(a.Entries))

	for i, entry := range a.Entries {
		r.Items[i].Title = entry.Title
		r.Items[i].Link = entry.Link.Href
		if entry.Content == "" {
			r.Items[i].Description = template.HTML(entry.Summary)
		} else {
			r.Items[i].Description = template.HTML(entry.Content)
		}
	}

	return &r
}
