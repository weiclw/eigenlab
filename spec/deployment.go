package deployment

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
)

type Deployment struct {
    APIVersion string `yaml:"apiVersion"`
    Kind string `yaml:"kind"`
    Metadata struct {
        Name string `yaml:"name"`
        Labels struct {
            App string `yaml:"app"`
        } `yaml:"app"`
    } `yaml:"metadata"`
    Spec struct {
        Name string `yaml:"name"`
        Image string `yaml:"image"`
        ActionScript string `yaml:"action_script"`
    } `yaml:"spec"`
}

func ReadYaml(path string, config *Deployment) error {
    yamlFile, readFileError := ioutil.ReadFile(path)
    if readFileError != nil {
        return readFileError
    }

    unmarshalError := yaml.Unmarshal(yamlFile, config)
    if unmarshalError != nil {
        return unmarshalError
    }

    return nil
}
