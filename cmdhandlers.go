package main

import (
	"fmt"
	"github.com/abiosoft/ishell"
	"github.com/go-resty/resty/v2"
)

type CacheClient struct {
	host     string
	username string
	password string
	client   *resty.Client
}

func NewCacheClient(host, username, password string) *CacheClient {
	return &CacheClient{
		host:     host,
		username: username,
		password: password,
		client:   resty.New(),
	}
}

func (c *CacheClient) getClient() *resty.Client {
	return c.client.SetHostURL(c.host).SetBasicAuth(c.username, c.password)
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
	}
}

func (c *CacheClient) Keys() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {
		req := c.getClient().NewRequest()
		resp, err := req.Get("cache/keys")
		if err != nil {
			fmt.Println("error: ", err)
		}
		fmt.Println(string(resp.Body()))
	}
}

func (c *CacheClient) _() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {

	}
}

func (c *CacheClient) _() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {

	}
}

func (c *CacheClient) _() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {

	}
}

func (c *CacheClient) _() func(c *ishell.Context) {
	return func(ctx *ishell.Context) {

	}
}

