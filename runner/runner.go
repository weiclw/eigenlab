package runner

import (
    "bufio"
    "eigenlab/options"
    "fmt"
    "io"
    "os"
    "os/exec"
    "time"
)

func issueCommand(wr *io.PipeWriter, delay time.Duration, cmd []byte) {
    time.Sleep(delay)
    _, err := wr.Write(cmd)
    if err != nil {
        fmt.Println("Failed to write to pipe: ", err)
    }
}

func readActionFile(path string) ([]string, error) {
     file, err := os.Open(path)
     if err != nil {
         fmt.Println("failed to read from file: ", path)
         return []string{}, err
     }

     defer file.Close()
     actions := []string{}
     scanner := bufio.NewScanner(file)
     for scanner.Scan() {
         cmd := scanner.Text() + "\n"
         actions = append(actions, cmd)
     }

     return actions, nil
}

func asyncInputs(redirect_input_yes bool, action_file string, wr *io.PipeWriter) {
    defer wr.Close()

    // No need to continue if auto-run is not needed.
    if !redirect_input_yes {
        return
    }

    cmd_list, read_err := readActionFile(action_file)
    if read_err != nil {
        return
    }

    // Wait for initial prompt.
    time.Sleep(10*1000*time.Millisecond)

    for i := 0; i < len(cmd_list); i++ {
        issueCommand(wr, 5*1000*time.Millisecond, []byte(cmd_list[i]))
    }
}

func Run(x []string) {
     opts := options.NewOptions()
     options.GetOptionsOnce(opts)

     redirectInput := false
     if redirect, got := options.GetOptionsValue[bool](opts, options.RedirectInputFlag); !got || !redirect {
         fmt.Println("Do not redirect input")
     } else {
         redirectInput = true
     }

     actionFile := ""
     if file, got := options.GetOptionsValue[string](opts, options.ActionFileFlag); got {
         actionFile = file
     }

     remainings := x[1:]
     cmd := exec.Command(x[0], remainings...)

     rd, wr := io.Pipe()
     defer rd.Close()

     cmd.Stdout = os.Stdout
     cmd.Stderr = os.Stderr

     if redirectInput {
         cmd.Stdin = rd
     } else {
         cmd.Stdin = os.Stdin
     }

     fmt.Println("Before launching go routine")
     go asyncInputs(redirectInput, actionFile, wr)

     err := cmd.Run()

     if err != nil {
         fmt.Println("Something goes wrong: ", err)
     }
}
