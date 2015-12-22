package proxy

import (
    "github.com/parnurzeal/gorequest"

    "errors"
    "encoding/xml"
    "encoding/json"
)

type Proxy struct {
    apiKey  string
    apiArea string
}

type ProxyItems struct {
    Item    []ProxyItem    `xml:"item"`
}

type ProxyItem struct {
    Ip          string  `xml:"ip"`
    RequestTime string  `xml:"requesttime"`
    Area        string  `xml:"area"`
    AreanNme    string  `xml:"areaname"`
    Http        int     `xml:"http"`
    Https       int     `xml:"https"`
    Anonymous   bool    `xml:"anonymous"`
    Type        string  `xml:"type"`
    Twoch       int     `xml:"twoch"`
}

func NewProxy(key string, area string) *Proxy {
    return &Proxy{
        apiKey:  key,
        apiArea: area,
    }
}

func (p *Proxy) Get() (string, error) {
    request := gorequest.New()

    resp, body, errs := request.
        Get("http://www.getproxy.jp/proxyapi").
        Query("ApiKey=" + p.apiKey).
        Query("area=" + p.apiArea).
        Query("sort=requesttime").
        Query("orderby=desc").
        Query("page=1").
        End()

    if errs != nil {
        return "", errs[0]
    }

    if resp.Status != "200 OK" {
        return "", errors.New("Proxy api return status code: " + resp.Status)
    }

    return body, nil
}

func (p *Proxy) ParseXML(proxyXml string) (*ProxyItems, error) {
    proxyItems := &ProxyItems{}

    if err := xml.Unmarshal([]byte(proxyXml), proxyItems); err != nil {
        return nil, err
    }

    return proxyItems, nil
}

func (p *Proxy) ToJSON(proxyXml string) (string, error) {
    proxyItems := &ProxyItems{}

    if err := xml.Unmarshal([]byte(proxyXml), proxyItems); err != nil {
        return "", err
    }

    proxyJson, err := json.Marshal(proxyItems)

    if err != nil {
        return "", err
    }

    return string(proxyJson), nil
}
