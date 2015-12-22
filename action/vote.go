package action

import (
    "github.com/parnurzeal/gorequest"

    "errors"
    "net"
    "net/url"
    "fmt"
    "time"
)

type Vote struct {
    agent   *gorequest.SuperAgent
    proxy   string
}

func NewVote() *Vote {
    return &Vote{
        agent: gorequest.New(),
    }
}

func (v *Vote) Proxy(proxy string) *Vote {
    v.proxy = proxy
    return v
}

func (v *Vote) Run() (string, error) {
    fmt.Println("-- Fetch page")

    _, err := v.fetchPage()

    if err != nil {
        return "", err
    }else{
        fmt.Println("-- Vote target")

        _, err := v.voteTarget()

        if err != nil {
            return "", err
        }
    }

    return "", nil
}

func (v *Vote) fetchPage() (string, error) {
    request := v.agent.Timeout(time.Duration(9000)*time.Millisecond).Proxy("http://" + v.proxy)

    resp, body, errs := request.
        Get("http://www.mache.tv/m/chat_room.php").
        Query("e=queen2012").
        Set("Host", "www.mache.tv").
        Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8").
        Set("Referer", "https://www.facebook.com/").
        Set("Accept-Language", "en-US,en;q=0.8").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.3").
        End()

    if errs != nil {
        error := errs[0]

        if _, ok := error.(*url.Error); ok {
            return "", errors.New("Can not connect to proxy server")
        }

        if _, ok := error.(*net.OpError); ok {
            return "", errors.New("Timeout in the proxy server")
        }

        return "", error
    }

    if resp.Status != "200 OK" {
        return "", errors.New("Action [vote::fetchPage] return status code: " + resp.Status)
    }

    return body, nil
}

func (v *Vote) voteTarget() (string, error) {
    resp, body, errs := v.agent.Timeout(time.Duration(9000)*time.Millisecond).Proxy("http://" + v.proxy).
        Post("http://www.mache.tv/m/chat_room.php").
        Query("e=queen2012").
        Set("Host", "www.mache.tv").
        Set("Accept", "*/*").
        Set("Origin", "http://www.mache.tv").
        Set("X-Requested-With", "XMLHttpRequest").
        Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_11_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/47.0.2526.80 Safari/537.3").
        Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8").
        Set("DNT", "1").
        Set("Referer", "http://www.mache.tv/m/chat_room.php?e=queen2012").
        Set("Accept-Language", "en-US,en;q=0.8").
        Send("mode=vote_complete").
        Send("ajax=1").
        Send("p_id=136").
        End()

    if errs != nil {
        error := errs[0]

        if _, ok := error.(*url.Error); ok {
            return "", errors.New("Can not connect to proxy server")
        }

        if _, ok := error.(*net.OpError); ok {
            return "", errors.New("Timeout in the proxy server")
        }

        return "", error
    }

    if resp.Status != "200 OK" {
        return "", errors.New("Action [vote::voteTarget] return status code: " + resp.Status)
    }

    return body, nil
}
