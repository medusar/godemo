package main

import (
	"context"
	"fmt"
	"os/exec"
	"time"
)

type result struct {
	output []byte
	err    error
}

func main() {

	resultChan := make(chan *result, 1024)

	ctx, cancelFun := context.WithCancel(context.TODO())
	go func() {
		cmd := exec.CommandContext(ctx, "/bin/bash", "-c", "sleep 5; ls -l")
		output, err := cmd.CombinedOutput()
		resultChan <- &result{output: output, err: err}
	}()

	time.Sleep(time.Second * 1)
	//cancelFun()
	_ = cancelFun

	result := <-resultChan
	fmt.Println(result.err, string(result.output))
}
