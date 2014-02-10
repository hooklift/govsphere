package soap

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type Envelope struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	EncodingStyle string   `xml:"http://schemas.xmlsoap.org/soap/encoding/ encodingStyle,attr"`
	Header        Header   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Body          Body     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type Header struct {
	Header interface{}
}

type Body struct {
	Body  string
	Fault Fault
}

type Fault struct {
	faultcode   string `xml:"http://schemas.xmlsoap.org/soap/envelope/ faultcode"`
	faultstring string `xml:"faultstring"`
	faultactora string `xml:"faultactor"`
	detail      string `xml:"detail"`
}

type Client struct {
	url, soapAction string
	tls             bool
}

func NewClient(url, soapAction string, tls bool) *Client {
	return &Client{
		url:        url,
		tls:        tls,
		soapAction: soapAction,
	}
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func (s *Client) Call(request interface{}, response interface{}) error {
	envelope := Envelope{
		Header:        Header{},
		EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
	}

	reqXml, err := xml.Marshal(request)
	if err != nil {
		return err
	}

	envelope.Body = Body{
		Body: string(reqXml),
	}

	buffer := &bytes.Buffer{}

	encoder := xml.NewEncoder(buffer)
	//encoder.Indent("  ", "    ")

	err = encoder.Encode(envelope)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, buffer)
	req.Header.Add("Content-Type", "text/xml; charset=\"utf-8\"")
	req.Header.Add("SOAPAction", s.soapAction)
	req.Header.Set("User-Agent", "govsphere/1.0")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.tls,
		},
		Dial: dialTimeout,
	}

	client := &http.Client{Transport: tr}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	respEnvelope := &Envelope{}

	err = xml.Unmarshal(body, respEnvelope)
	if err != nil {
		return err
	}

	if respEnvelope.Body.Body == "" {
		log.Printf("%#v\n", respEnvelope.Body)
		return nil
	}

	err = xml.Unmarshal([]byte(respEnvelope.Body.Body), response)
	if err != nil {
		return err
	}

	return nil
}
