package svn

import "encoding/xml"

type log struct {
	XMLName  xml.Name   `xml:"log`
	logEntry []logEntry `xml:logentry`
}

type logEntry struct {
	XMLAttr xml.Attr `xml:"revision"`
	Msg     string   `xml:"msg"`
	Author  string   `xml:"author"`
	Date    string   `xml:"date`
	Paths   path     `xml:"paths`
}

type path struct {
	Path     []string `xml:path`
	Action   string   `xml:"action,attr"`
	Propmods string   `xml:"prop-mods,attr"`
	Textmods string   `xml:"text-mods,attr"`
	Kind     string   `xml:"kind,attr"`
}

type info struct {
	XMLName       xml.Name `xml:"info"`
	Url           string   `xml:"entry>url"`
	RelativeUrl   string   `xml:"entry>relative-url"`
	Root          string   `xml:"entry>repository>root"`
	UUID          string   `xml:"entry>repository>uuid"`
	WcrootAbspath string   `xml:"entry>wc-info>wcroot-abspath"`
	Schedule      string   `xml:"entry>wc-info>schedule"`
	Depth         string   `xml:"entry>wc-info>depth"`
	Revision      string   `xml:"entry>commit,attr"`
	Author        string   `xml:"entry>commit>author"`
	Date          string   `xml:"entry>commit>date"`
}
