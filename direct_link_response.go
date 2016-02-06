package ogone

import "encoding/xml"

type DirectLinkResponse struct {
	data map[string]string
}

func NewDirectLinkResponse(xmlContent string) (*DirectLinkResponse, error) {
	r := &DirectLinkResponse{}

	err := xml.Unmarshal([]byte(xmlContent), r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *DirectLinkResponse) IsAuthorised() bool {
	return r.data["STATUS"] == "5"
}

func (e *DirectLinkResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var nodes []*DirectLinkResponse

	var done bool

	for !done {
		t, err := d.Token()

		if err != nil {
			return err
		}
		switch t := t.(type) {
		case xml.StartElement:
			e := &DirectLinkResponse{}
			e.UnmarshalXML(d, t)
			nodes = append(nodes, e)
		case xml.EndElement:
			done = true
		}
	}

	e.data = make(map[string]string, len(start.Attr))

	for _, attr := range start.Attr {
		e.data[attr.Name.Local] = attr.Value
	}

	return nil
}
