package common

import (
	"bytes"
	"errors"
	"io"
	"os/exec"
)

// Pipeline is a collection of commands to execute chaning the output
// of each command to the input of the next successive command.
type Pipeline []*exec.Cmd

// Run executes the pipeline.
func (pipeline Pipeline) Run(in io.Reader) ([]byte, error) {
	// Require at least one command to execute
	n := len(pipeline)
	if n == 0 {
		return nil, errors.New("No commands found in the pipeline")
	}

	// Declare some buffers to store shit
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Assign the first input
	if in != nil {
		pipeline[0].Stdin = in
	}

	// Pass the input of each command to the output of the last
	last := n - 1
	for i, cmd := range pipeline[:last] {
		var err error
		if pipeline[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return nil, err
		}
		cmd.Stderr = &stderr
	}

	// Assign the proper buffers to the last item
	pipeline[last].Stdout, pipeline[last].Stderr = &stdout, &stderr

	// Start each command
	for _, cmd := range pipeline {
		if err := cmd.Start(); err != nil {
			return stderr.Bytes(), err
		}
	}

	// Wait for each command to complete
	for _, cmd := range pipeline {
		if err := cmd.Wait(); err != nil {
			return stderr.Bytes(), err
		}
	}

	if stderr.Len() > 0 {
		return stderr.Bytes(), errors.New("Errors occurred during processing")
	}

	// Return the output collected
	return stdout.Bytes(), nil
}
