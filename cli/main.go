package main

import (
    "os"
    "fmt"
    "github.com/codegangsta/cli"
)

func main() {
    api := Api{}
    app := cli.NewApp()
    app.Name = "pwitter"
    app.Usage = "Pwitter command line client"

    app.Commands = []cli.Command {
        {
            Name: "get",
            Usage: "Get Pweets",
            Flags: []cli.Flag {
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
                api.Get(min, max)
            },
        },
        {
            Name: "post",
            Usage: "Create a new Pweet",
            Flags: []cli.Flag {
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
                api.Post(user, body)
            },
        },
    }

    app.Run(os.Args)
}
