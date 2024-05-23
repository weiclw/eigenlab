package options

import (
    "flag"
    "os"
)

type Options struct {
    RedirectInput bool
    ActionFile string
}

// Always extract information from env first.
func optionsFromEnv(opts *Options) {
    _, redirect_input_yes := os.LookupEnv("qemu_redirect_input")
    if redirect_input_yes {
        opts.RedirectInput = true
    }

    opts.ActionFile = os.Getenv("qemu_action_file")
}

// Each of the instance will represent a command line argment in the function below
type cmdArg struct {
    value interface{}
    visited bool
    name string
    comment string
}

// Commandline flags shall override values from env.
func optionsFromFlags(opts *Options) {
    redirectInput := cmdArg{false, false, "redirect_input", "redirect so that it can run script"}
    actionFile := cmdArg{"", false, "action_file", "path of action script"}

    if val, ok := redirectInput.value.(bool); ok {
        flag.BoolVar(&val, redirectInput.name, val, redirectInput.comment)
        redirectInput.value = val
    }

    if val, ok := actionFile.value.(string); ok {
        flag.StringVar(&val, actionFile.name, val, actionFile.comment)
        actionFile.value = val
    }

    // Check which flags have been specified.
    flag.Visit(func(f *flag.Flag) {
        if f.Name == redirectInput.name {
            redirectInput.visited = true
        } else if f.Name == actionFile.name {
            actionFile.visited = true
        }
    })

    // This function automatically handles parsing error and may exit the program as well.
    flag.Parse()

    if redirectInput.visited {
        opts.RedirectInput = redirectInput.value.(bool)
    }

    if actionFile.visited {
        opts.ActionFile = actionFile.value.(string)
    }
}

func GetOptionsOnce(opts *Options) {
    optionsFromEnv(opts)
    optionsFromFlags(opts)
}
