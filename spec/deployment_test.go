package spec

import (
    "testing"
)

var yamlToTest1 string = `
apiVersion: 1
kind: deployment
`

var extraFields string = `
kind: deployment
xyz: foo
`

var totallyDifferentFields string = `
abc: hello
bar: foo
`


func TestParseSimpleYaml(t *testing.T) {
    var config Deployment
    err := ParseYaml([]byte(yamlToTest1), &config)
    if err != nil {
        t.Error(err)
    }

    if config.APIVersion != "1" {
        t.Error("ApiVersion mismatch")
    }
}

func TestExtraFieldInYaml(t *testing.T) {
   var config Deployment
    err := ParseYaml([]byte(extraFields), &config)
    if err != nil {
        t.Error(err)
    }
}

func TestTotallyDifferentFields(t *testing.T) {
   var config Deployment
    err := ParseYaml([]byte(totallyDifferentFields), &config)
    if err != nil {
        t.Error(err)
    }
}
