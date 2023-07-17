package main

import (
	"strings"

	yml "github.com/buildkite/yaml"
	compiler "github.com/go-vela/server/compiler/native"
	"github.com/go-vela/types/yaml"
)

func modifyPipeline(cfg *Config, modReq *compiler.ModifyRequest) (*yaml.Build, error) {
	// unmarshal the modification request into a yaml build
	build := new(yaml.Build)
	err := yml.Unmarshal([]byte(modReq.Pipeline), &build)
	if err != nil {
		return nil, err
	}

	// analyze/modify the pipeline
	for _, step := range build.Steps {
		if strings.Contains(step.Image, "artifactory") {
			step.Name = "this step deserves free credentials"
		}
	}

	return build, nil
}
