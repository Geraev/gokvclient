package main

import (
	"fmt"
	"os"

	"github.com/geraev/gokvclient/handlers"

	"github.com/abiosoft/ishell"
	"github.com/jessevdk/go-flags"
)

type options struct {
	Help     bool   `short:"h" long:"help" description:"show help message"`
	Host     string `short:"s" long:"host" description:"host name for GoCache server" default:"http://localhost:8081"`
	Username string `short:"u" long:"username" description:"username"`
	Password string `short:"p" long:"password" description:"password"`
}

var (
	opts        options
	cacheClient *handlers.CacheClient
)

func main() {
	p := flags.NewParser(&opts, flags.Default&^flags.HelpFlag)
	_, err := p.Parse()
	if err != nil {
		fmt.Printf("fail to parse args: %v", err)
		os.Exit(1)
	}
	if opts.Help {
		p.WriteHelp(os.Stdout)
		os.Exit(0)
	}

	cacheClient = handlers.NewCacheClient(opts.Host, opts.Username, opts.Password)

	shell := ishell.New()
	shell.Println("GoCache interactive client")

	// Set login
	shell.AddCmd(&ishell.Cmd{
		Name: "login",
		Help: "set username and password for http basic authentication on endpoint",
		Func: cacheClient.SetLogin(),
	})

	// Set host
	shell.AddCmd(&ishell.Cmd{
		Name: "host",
		Help: "set hostname and port",
		Func: cacheClient.SetHost(),
	})

	// Get all keys
	// Example:
	// >>> keys
	shell.AddCmd(&ishell.Cmd{
		Name: "keys",
		Help: "get all keys in cache",
		LongHelp: `Get all keys
Example:
  keys
`,
		Func: cacheClient.Keys(),
	})

	// Get value for key (and internal key)
	// Example:
	// >>> key <key>
	shell.AddCmd(&ishell.Cmd{
		Name: "key",
		Help: "get value for key (and internal key)",
		LongHelp: `Get value for key (and internal key)
Example:
  key <key>
`,
		Func: cacheClient.Key(),
	})

	// Set or update value
	// Examples:
	// >>> set string new_key '{"value": "string_value", "ttl": 10000}'
	// >>> set list planets '{"value": ["earth","jupiter","saturn"], "ttl": 10000}'
	// >>> set dictionary planets_map '{"value": ["earth":2220,"jupiter":3899,"saturn":23000], "ttl": 10000}'
	shell.AddCmd(&ishell.Cmd{
		Name: "set",
		Help: "set or update value",
		LongHelp: `Set or update value
Examples:
  set string new_key '{"value": "string_value", "ttl": 10000}'
  set list planets '{"value": ["earth","jupiter","saturn"], "ttl": 10000}'
  set dictionary planets_map '{"value": ["earth":2220,"jupiter":3899,"saturn":23000], "ttl": 10000}'
`,
		Func: cacheClient.Set(),
	})

	// Remove key
	// Examples:
	// >>> remove <key>
	shell.AddCmd(&ishell.Cmd{
		Name: "remove",
		Help: "remove key",
		LongHelp: `Remove key
Examples:
  remove <key>
`,
		Func: cacheClient.Remove(),
	})

	// when started with "exit" as first argument, assume non-interactive execution
	if len(os.Args) > 1 && os.Args[1] == "exit" {
		shell.Process(os.Args[2:]...)
	} else {
		// start shell
		shell.Run()
		// teardown
		shell.Close()
	}
}
