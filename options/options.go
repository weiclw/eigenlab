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
type optionValue struct {
    ptr interface{}
    visited bool
    name string
    comment string
}

type optionType interface {
    bool | string
}

func newOption[V optionType](name string, default_value V, comment string) *optionValue {
    return &optionValue{
        ptr: &default_value,
        visited: false,
        name: name,
        comment: comment,
    }
}

func getOptionPtrOrDie[V optionType](o *optionValue) *V {
    return o.ptr.(*V)
}

// Commandline flags shall override values from env.
func optionsFromFlags(opts *Options) {
    redirectInput := newOption("redirect_input", false, "redirect so that it can run script")
    actionFile := newOption("action_file", "", "path of action script")

    if ptr := getOptionPtrOrDie[bool](redirectInput); true {
        val := *ptr
        flag.BoolVar(ptr, redirectInput.name, val, redirectInput.comment)
    }

    if ptr := getOptionPtrOrDie[string](actionFile); true {
        val := *ptr
        flag.StringVar(ptr, actionFile.name, val, actionFile.comment)
    }

    // This function automatically handles parsing error and may exit the program as well.
    flag.Parse()

    // Check which flags have been specified.
    flag.Visit(func(f *flag.Flag) {
        if f.Name == redirectInput.name {
            redirectInput.visited = true
        } else if f.Name == actionFile.name {
            actionFile.visited = true
        }
    })

    if redirectInput.visited {
        opts.RedirectInput = *getOptionPtrOrDie[bool](redirectInput)
    }

    if actionFile.visited {
        opts.ActionFile = *getOptionPtrOrDie[string](actionFile)
    }
}

func GetOptionsOnce(opts *Options) {
    optionsFromEnv(opts)
    optionsFromFlags(opts)
}
