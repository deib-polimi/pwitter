package api

import (
    "fmt"
    "net/http"
    neturl "net/url"
    "io/ioutil"
)

func getBody(resp *http.Response, err error) (string, error) {
    if err != nil {
        return "", err
    }

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil

}

type Api struct {
    baseUrl string
}

func New(host string, port int) *Api {
    return &Api{baseUrl: fmt.Sprintf("http://%s:%d", host, port)}
}


func (a *Api) Get(min float64, max float64) (string, error) {
    url := fmt.Sprintf("%s/pweets?lte=%f&gte=%f", a.baseUrl, max, min)

    resp, err := http.Get(url)
    return getBody(resp, err)
}

func (a *Api) Post(user string, pbody string) (string, error) {
    url := fmt.Sprintf("%s/pweets", a.baseUrl)
    data := neturl.Values{"user": {user}, "body": {pbody}}

    resp, err := http.PostForm(url, data)
    return getBody(resp, err)
}
