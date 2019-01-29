package structures

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

//Envelop structure used for communication
type Envelop struct {
	XMLName    xml.Name `xml:"Data"`
	Encryption bool     `xml:"Head>encryption"`
	SessionID  string   `xml:"Head>session"`
	Body       []byte   `xml:"Body"`
	Key        []byte   `xml:"Key,omitempty"`
	CHeck      string   `xml:"Chech,omitempty"`
}

//FromEnvelop extract structure from request
func (e *Envelop) FromEnvelop(r *http.Request) error {
	//var out Envelop
	data, err := ioutil.ReadAll(r.Body)
	err = xml.Unmarshal(data, &e)
	return err
}

//ToEnvelop pack structure to response
func (e *Envelop) ToEnvelop(env Envelop) ([]byte, error) {
	out, err := xml.MarshalIndent(env, "	", "		")
	return out, err
}
