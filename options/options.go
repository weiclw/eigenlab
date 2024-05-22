package options

import (
    "flag"
    "os"
)

type Options struct {
    RedirectInput bool
    ActionFile string
}


func OptionsFromEnv(opts *Options) {
    _, redirect_input_yes := os.LookupEnv("qemu_redirect_input")
    if redirect_input_yes {
        opts.RedirectInput = true
    }

    opts.ActionFile = os.Getenv("qemu_action_file")
}

func OptionsFromFlags(opts *Options) {
    flag.BoolVar(&opts.RedirectInput, "redirect_input", false, "redirect so that it can run script")
    flag.StringVar(&opts.ActionFile, "action_file", "", "path of action script file")

    // This function automatically handles parsing error and may exit the program as well.
    flag.Parse()
}
