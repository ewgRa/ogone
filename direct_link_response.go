package ogone

import "encoding/xml"

// DirectLinkResponse from DirectLink request
type DirectLinkResponse struct {
	data map[string]string
}

func newDirectLinkResponse(xmlContent string) (*DirectLinkResponse, error) {
	r := &DirectLinkResponse{}

	err := xml.Unmarshal([]byte(xmlContent), r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// IsAuthorised for checking that payment is authorized
func (r *DirectLinkResponse) IsAuthorised() bool {
	return r.data["STATUS"] == "5"
}

// UnmarshalXML for parsing XML answer
func (r *DirectLinkResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	err := d.Skip()

	if err != nil {
		return err
	}

	if start.Name.Local != "ncresponse" {
		return nil
	}

	r.data = make(map[string]string, len(start.Attr))

	for _, attr := range start.Attr {
		r.data[attr.Name.Local] = attr.Value
	}

	return nil
}
