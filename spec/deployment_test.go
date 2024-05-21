package spec

import (
    "testing"
)

var yamlToTest1 string = `
apiVersion: 1
kind: deployment
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
