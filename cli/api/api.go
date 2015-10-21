package api

import (
    "fmt"
    "net/http"
    neturl "net/url"
    "io/ioutil"
    "math/rand"
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


func rndLG() (float64, float64) {
    min := rand.Float64() * 2 - 1 // make it between [-1.0, 1.0)
    max := rand.Float64() * 2 - 1

    if max < min {
        min, max = max, min
    }

    return min, max
}

var sampleUsers []string = []string {
    "Frodo",
    "Gandalf",
    "Aragorn",
    "Gimli",
    "Sam",
}

func getPweet(fname string) string {
    pweet, err := ioutil.ReadFile(fname)
    if err != nil {
        fmt.Println(err)
        return ""
    }
    return string(pweet)
}

var sampleBodies []string = []string {
    getPweet("pweets/hr"),
    getPweet("pweets/limits"),
    getPweet("pweets/mushrooms"),
    getPweet("pweets/orchestras"),
    getPweet("pweets/sheik"),
}

func rndUserBody() (string, string) {
    ui := rand.Int31n(int32(len(sampleUsers)))
    bi := rand.Int31n(int32(len(sampleBodies)))
    u := sampleUsers[ui]
    b := sampleBodies[bi]

    return u, b
}

// Function to stress endpoints.
// Api calls are performed concurrently.
// `ng` the number of gets to perform
// `np` the number of posts to perform
// It returns a channel of booleans containing if the api call
// has been succesfull or not
func (a *Api) Stress(ng int, np int) <-chan bool {
    out := make(chan bool)
    ctrl := make(chan int)

    // gets
    go func() {
        for i := 0; i < ng; i++ {
            go func() {
                _, err := a.Get(rndLG())
                out <- err == nil
                ctrl <- 0 // say to control channel I finished
            }()
        }
    }()

    // posts
    go func(){
        for i := 0; i < np; i++ {
            go func() {
                _, err := a.Post(rndUserBody())
                out <- err == nil
                ctrl <- 0 // say to control channel I finished
            }()
        }
    }()

    // control routine
    // closes channels in the end
    go func() {
        for i := 0; i < np + ng; i++ {
            <-ctrl
        }
        close(ctrl)
        close(out)
    }()

    return out
}
