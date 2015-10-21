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

var sampleBodies []string = []string {
    "There is only one Lord of the Ring, only one who can bend it to his will. And he does not share power.",
    "That there’s some good in this world, Mr. Frodo, and it’s worth fighting for.",
    "Oh, it’s quite simple. If you are a friend, you speak the password, and the doors will open.",
    "A day may come when the courage of men fails, but it is not this day.",
    "The Ring has awoken, it’s heard its masters call.",
    "We swears, to serve the master of the Precious. We will swear on the Precious!",
    "You are the luckiest, the canniest, and the most reckless man I ever knew. Bless you, laddie.",
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
