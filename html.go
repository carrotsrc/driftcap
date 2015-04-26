package main
import "encoding/xml"
import "io"
import "strings"

type HtmlAnchor struct {
	Href string `xml:"href,attr"`
}

type HtmlLink struct {
	Href string `xml:"href,attr"`
}

type HtmlScript struct {
	Src string `xml:"src,attr"`
}

func isLocal(target string, url string) (bool) {

	if strings.HasPrefix(url, target) || strings.HasPrefix(url, "/") {
		return true
	} else {
		return false
	}
}

func parseElement(token interface{}, decoder *xml.Decoder, target string) (res *SiteResource) {

	res = nil

	switch se := token.(type) {
	case xml.StartElement:
		if se.Name.Local == "a" {
			var a HtmlAnchor
			decoder.DecodeElement(&a, &se)

			res = new(SiteResource)

			res.ResourceType = Page
			if isLocal(target, a.Href) {
				res.ResourceLocation = Local
			} else {
				res.ResourceLocation = Remote
			}

			res.Ref = a.Href
		} else
		if se.Name.Local == "link" {
			var link HtmlLink
			decoder.DecodeElement(&link, &se)

			res = new(SiteResource)

			res.ResourceType = Asset
			if isLocal(target, link.Href) {
				res.ResourceLocation = Local
			} else {
				res.ResourceLocation = Remote
			}

			res.Ref = link.Href

		} else
		if se.Name.Local == "script" {
			var script HtmlScript
			decoder.DecodeElement(&script, &se)

			res = new(SiteResource)

			res.ResourceType = Asset
			if isLocal(target, script.Src) {
				res.ResourceLocation = Local
			} else {
				res.ResourceLocation = Remote
			}

			res.Ref = script.Src
		}
	}

	return res

}
func parseLinks(body io.ReadCloser, target string) ([]*SiteResource) {
	dec := xml.NewDecoder(body)
	resources := []*SiteResource{}

	for {
		t, _ := dec.Token()
		if t == nil {
			break
		}

		res := parseElement(t, dec, target)
		if res == nil {
			continue
		}

		resources = append(resources, res)
	}

	return resources
}
