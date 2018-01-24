package svn

import "encoding/xml"

type logEntry struct {
	XMLName xml.Name `xml:"logentry"`
	Msg     string   `xml:"msg"`
	Author  string   `xml:"author"`
}
