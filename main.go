package tachyon

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	Vars       map[string]string `short:"s" long:"set" description:"Set a variable"`
	ShowOutput bool              `short:"o" long:"output" description:"Show command output"`
}

func Main(args []string) int {
	args, err := flags.ParseArgs(&opts, args)

	if err != nil {
		fmt.Printf("Error parsing options: %s", err)
		return 1
	}

	if len(args) != 2 {
		fmt.Printf("Usage: tachyon [options] <playbook>\n")
		return 1
	}

	cfg := &Config{ShowCommandOutput: opts.ShowOutput}

	playbook, err := LoadPlaybook(args[1])

	env := &Environment{}
	env.Init(cfg)

	for k, v := range opts.Vars {
		env.Set(k, v)
	}

	err = playbook.Run(env)

	if err != nil {
		fmt.Printf("Error running playbook: %s\n", err)
		return 1
	}

	return 0
}
