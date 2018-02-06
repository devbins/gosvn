package gosvn

import "encoding/xml"

type log struct {
	XMLName  xml.Name   `xml:"log"`
	LogEntry []logEntry `xml:"logentry"`
}

type logEntry struct {
	Revision string `xml:"revision,attr"`
	Msg      string `xml:"msg"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
	Paths    paths  `xml:"paths>path"`
}

type paths struct {
	Path     string `xml:",innerxml"`
	Action   string `xml:"action,attr"`
	PropMods string `xml:"prop-mods,attr"`
	TextMods string `xml:"text-mods,attr"`
	Kind     string `xml:"kind,attr"`
}

type Info struct {
	XMLName       xml.Name `xml:"info"`
	Url           string   `xml:"entry>url"`
	RelativeUrl   string   `xml:"entry>relative-url"`
	Root          string   `xml:"entry>repository>root"`
	UUID          string   `xml:"entry>repository>uuid"`
	WcrootAbspath string   `xml:"entry>wc-info>wcroot-abspath"`
	Schedule      string   `xml:"entry>wc-info>schedule"`
	Depth         string   `xml:"entry>wc-info>depth"`
	Commit        commit   `xml:"entry>commit"`
}

type commit struct {
	Revision string `xml:"revision,attr"`
	Author   string `xml:"author"`
	Date     string `xml:"date"`
}

type lists struct {
	XMLName xml.Name `xml:"lists"`
	List    list     `xml:"list"`
}

type list struct {
	Path  string  `xml:"path,attr"`
	Entry []entry `xml:"entry"`
}

type entry struct {
	Kind   string `xml:"kind,attr"`
	Name   string `xml:"name"`
	Size   string `xml:"size"`
	Commit commit `xml:"commit"`
}

type diff struct {
	Paths []diffPath `xml:"paths>path"`
}
type diffPath struct {
	Props string `xml:"props,attr"`
	Kind  string `xml:"kind,attr"`
	Item  string `xml:"item,attr"`
	Path  string `xml:",innerxml"`
}
