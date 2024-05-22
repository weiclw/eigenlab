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

// Commandline flags shall override values from env.
func optionsFromFlags(opts *Options) {
    var(
        redirectInput bool
        actionFile string
    )

    redirectInputVisited := false
    actionFileVisited := false

    redirectInputName := "redirect_input"
    actionFileName := "action_file"


    flag.BoolVar(&redirectInput, redirectInputName, false, "redirect so that it can run script")
    flag.StringVar(&actionFile, actionFileName, "", "path of action script file")

    // Check which flags have been specified.
    flag.Visit(func(f *flag.Flag) {
        if f.Name == redirectInputName {
            redirectInputVisited = true
        } else if f.Name == actionFileName {
            actionFileVisited = true
        }
    })

    // This function automatically handles parsing error and may exit the program as well.
    flag.Parse()

    if redirectInputVisited {
        opts.RedirectInput = redirectInput
    }

    if actionFileVisited {
        opts.ActionFile = actionFile
    }
}

func GetOptionsOnce(opts *Options) {
    optionsFromEnv(opts)
    optionsFromFlags(opts)
}
