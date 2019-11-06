package handlers

import (
	"fmt"
	"github.com/fatih/color"
	"time"

	"github.com/abiosoft/ishell"
	"github.com/go-resty/resty/v2"
)

type CacheClient struct {
	host     string
	username string
	password string
	client   *resty.Client
}

var (
	cyan    = color.New(color.FgCyan).SprintFunc()
	boldRed = color.New(color.FgRed, color.Bold).SprintFunc()
)

func NewCacheClient(host, username, password string) *CacheClient {
	client := resty.New().
		SetDisableWarn(true).
		SetRetryCount(3).
		SetRetryWaitTime(time.Second)

	return &CacheClient{
		host:     "http://" + host,
		username: username,
		password: password,
		client:   client,
	}
}

func (c *CacheClient) getClient() *resty.Client {
	return c.client.
		SetHostURL(c.host).
		SetBasicAuth(c.username, c.password)
}

func (c *CacheClient) SetLogin() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		ctx.ShowPrompt(false)
		defer ctx.ShowPrompt(true)

		ctx.Println("Let's login")
		ctx.Print("Username: ")
		c.username = ctx.ReadLine()
		ctx.Print("Password: ")
		c.password = ctx.ReadPassword()
	}
}

func (c *CacheClient) SetHost() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		ctx.Println("Enter hostname and port")
		ctx.Print("Host: ")
		host := ctx.ReadLine()
		ctx.Print("Port: ")
		port := ctx.ReadLine()
		c.host = "http://" + host + ":" + port
		ctx.Println(cyan("Hostname set to", c.host))
	}
}

func (c *CacheClient) Keys() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		req := c.getClient().NewRequest()
		resp, err := req.Get("cache/keys")
		if err != nil {
			ctx.Println(boldRed("error: ", err))
		}
		ctx.Println(cyan(string(resp.Body())))
	}
}

func (c *CacheClient) Key() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		req := c.getClient().NewRequest()
		var (
			resp             *resty.Response
			err              error
			key, internalKey string
		)
		switch len(ctx.Args) {
		case 1:
			key = ctx.Args[0]
			resp, err = req.Get(fmt.Sprintf("cache/key/%s", key))
		case 2:
			internalKey = ctx.Args[1]
			resp, err = req.Get(fmt.Sprintf("cache/key/%s/%s", key, internalKey))
		default:
			ctx.Println(boldRed("too many arguments"))
			return
		}
		if err != nil {
			ctx.Println(boldRed("error: ", err))
		}
		ctx.Println(cyan(string(resp.Body())))
	}
}

func (c *CacheClient) Set() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		req := c.getClient().NewRequest()
		var (
			resp                *resty.Response
			err                 error
			vartype, key, value string
		)
		if len(ctx.Args) != 3 {
			ctx.Println(boldRed("arguments error"))
			return
		}

		vartype = ctx.Args[0]
		key = ctx.Args[1]
		value = ctx.Args[2]

		switch vartype {
		case "string", "list", "dictionary":
			req.SetBody(value).
				SetHeader("Content-Type", "application/json")
			resp, err = req.Put(fmt.Sprintf("cache/set/%s/%s", vartype, key))
		default:
			ctx.Println(boldRed("type error"))
			return
		}
		if err != nil {
			ctx.Println(boldRed("error: ", err))
		}
		ctx.Println(cyan(string(resp.Body())))
	}
}

func (c *CacheClient) Remove() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		req := c.getClient().NewRequest()
		var (
			resp *resty.Response
			err  error
			key  string
		)
		key = ctx.Args[0]
		resp, err = req.Delete(fmt.Sprintf("cache/remove/%s", key))
		if err != nil {
			ctx.Println(boldRed("error: ", err))
		}
		ctx.Println(cyan(string(resp.Body())))
	}
}
