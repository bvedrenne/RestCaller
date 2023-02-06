package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/mkideal/cli"
)

func main() {
	if err := cli.Root(root,
		cli.Tree(help),
		cli.Tree(get),
		cli.Tree(post),
		cli.Tree(header),
		cli.Tree(put),
		cli.Tree(delete),
		cli.Tree(option),
		cli.Tree(patch),
	).Run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

var help = cli.HelpCommand("display help information")

type argH struct {
	Help bool `cli:"h,help" usage:"show help"`
}

func (argv *argH) AutoHelp() bool {
	return argv.Help
}

var root = &cli.Command{
	Desc: "Launch HTTP command to a server",
	Argv: func() interface{} { return new(argH) },
}

type argGet struct {
	Url    string            `cli:"*u,url" usage:"Url to get"`
	Header map[string]string `cli:"H" usage:"header option"`
}

var get = &cli.Command{
	Name: "GET",
	Desc: "Launch HTTP GET command on url",
	Argv: func() interface{} { return new(argGet) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argGet)
		res := call("GET", argv.Url, nil, argv.Header)
		printStatus(res)
		printHeader(res)
		fmt.Println()
		printBody(res)
		return nil
	},
}

type argPost struct {
	Url     string            `cli:"*u,url" usage:"Url to get"`
	Header  map[string]string `cli:"H" usage:"header option"`
	Content string            `cli:"*c,content" usage:"Post content"`
}

var post = &cli.Command{
	Name: "POST",
	Desc: "Launch HTTP POST command on url",
	Argv: func() interface{} { return new(argPost) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argPost)
		res := call("POST", argv.Url, strings.NewReader(argv.Content), argv.Header)
		printStatus(res)
		printHeader(res)
		fmt.Println()
		printBody(res)
		return nil
	},
}

var header = &cli.Command{
	Name: "HEAD",
	Desc: "Launch HTTP HEADER command on url",
	Argv: func() interface{} { return new(argGet) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argGet)

		res := call("HEAD", argv.Url, nil, argv.Header)
		printStatus(res)
		printHeader(res)
		return nil
	},
}

var put = &cli.Command{
	Name: "PUT",
	Desc: "Launch HTTP PUT command on url",
	Argv: func() interface{} { return new(argPost) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argPost)
		res := call("PUT", argv.Url, strings.NewReader(argv.Content), argv.Header)
		printStatus(res)
		printHeader(res)
		fmt.Println()
		printBody(res)
		return nil
	},
}

type argDelete struct {
	Url     string            `cli:"*u,url" usage:"Url to get"`
	Header  map[string]string `cli:"H" usage:"header option"`
	Content string            `cli:"c,content" usage:"Post content"`
}

var delete = &cli.Command{
	Name: "DELETE",
	Desc: "Launch HTTP DELETE command on url",
	Argv: func() interface{} { return new(argDelete) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argDelete)
		var content io.Reader
		if len(argv.Content) != 0 {
			content = strings.NewReader(argv.Content)
		}
		res := call("DELETE", argv.Url, content, argv.Header)
		printStatus(res)
		printHeader(res)
		fmt.Println()
		printBody(res)
		return nil
	},
}

var option = &cli.Command{
	Name: "OPTIONS",
	Desc: "Launch HTTP OPTIONS command on url",
	Argv: func() interface{} { return new(argGet) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argGet)

		res := call("OPTIONS", argv.Url, nil, argv.Header)
		printStatus(res)
		printHeader(res)
		return nil
	},
}

var patch = &cli.Command{
	Name: "PATCH",
	Desc: "Launch HTTP PUT command on url",
	Argv: func() interface{} { return new(argPost) },
	Fn: func(ctx *cli.Context) error {
		argv := ctx.Argv().(*argPost)
		res := call("PATCH", argv.Url, strings.NewReader(argv.Content), argv.Header)
		printStatus(res)
		printHeader(res)
		fmt.Println()
		printBody(res)
		return nil
	},
}

func call(method string, url string, content io.Reader, header map[string]string) *http.Response {
	req, err := http.NewRequest(method, url, content)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	for key, val := range header {
		req.Header.Set(key, val)
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}

	return res
}

func printStatus(res *http.Response) {
	fmt.Printf("%s\n", res.Status)
}

func printHeader(res *http.Response) {
	for key, val := range res.Header {
		fmt.Printf("%s: %s\n", key, val)
	}
}

func printBody(res *http.Response) {
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	fmt.Println(string(body))
}
