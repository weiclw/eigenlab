package spec

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

func ParseYaml(data []byte, config *Deployment) error {
    unmarshalError := yaml.Unmarshal(data, config)
    if unmarshalError != nil {
        return unmarshalError
    } else {
        return nil
    }
}

func ReadYaml(path string, config *Deployment) error {
    yamlFile, readFileError := ioutil.ReadFile(path)
    if readFileError != nil {
        return readFileError
    } else {
        return ParseYaml(yamlFile, config)
    }
}
