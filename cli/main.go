package main

import (
    "os"
    "fmt"
    "github.com/affo/pwitter/api"

    "github.com/codegangsta/cli"
)

func main() {
    app := cli.NewApp()
    app.Name = "pwitter"
    app.Usage = "Pwitter command line client"

    ipFlag := cli.StringFlag {
        Name: "ip",
        Usage: "The IP address of the web server",
        Value: "localhost",
    }

    portFlag := cli.IntFlag {
        Name: "port, p",
        Usage: "The port of the web server",
        Value: 5000,
    }

    newApi := func(c *cli.Context) *api.Api {
        ip := c.String("ip")
        port := c.Int("port")
        return api.New(ip, port)
    }

    app.Commands = []cli.Command {
        {
            Name: "get",
            Usage: "Get Pweets",
            Flags: []cli.Flag {
                ipFlag,
                portFlag,
                cli.Float64Flag {
                    Name: "max, M",
                    Usage: "Max polarity for Pweets returned",
                    Value: 1.0,
                },
                cli.Float64Flag {
                    Name: "min, m",
                    Usage: "Min polarity for Pweets returned",
                    Value: -1.0,
                },
            },
            Action: func(c *cli.Context) {
                min := c.Float64("min")
                max := c.Float64("max")
                r, err := newApi(c).Get(min, max)

                if err != nil {
                    fmt.Println(err)
                    return
                }
                fmt.Println(r)
            },
        },
        {
            Name: "post",
            Usage: "Create a new Pweet",
            Flags: []cli.Flag {
                ipFlag,
                portFlag,
                cli.StringFlag {
                    Name: "user, u",
                    Usage: "User name for this Pweet",
                    Value: "Anonymous",
                },
                cli.StringFlag {
                    Name: "body, b",
                    Usage: "The body of this Pweet",
                },
            },
            Action: func(c *cli.Context) {
                user := c.String("user")
                body := c.String("body")
                if len(body) == 0 {
                    fmt.Println("Body is required")
                    cli.ShowCommandHelp(c, "post")
                    return
                }
                r, err := newApi(c).Post(user, body)

                if err != nil {
                    fmt.Println(err)
                    return
                }
                fmt.Println(r)
            },
        },
    }

    app.Run(os.Args)
}
