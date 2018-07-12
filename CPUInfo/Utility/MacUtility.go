package MacCpuUtility

import (
	"bytes"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type Process struct {
	pid int
	cpu float64
}

func GoCpuUsage() float64 {
	cmd := exec.Command("ps", "aux")
	cmd1 := exec.Command("grep", "goland")
	var out bytes.Buffer
	//cmd.Stdout = &out
	var err error
	out, _, err = Pipeline(cmd, cmd1)
	if err != nil {
		log.Fatal(err)
	}
	processes := make([]*Process, 0)
	for {
		line, err := out.ReadString('\n')
		if err!=nil {
			break;
		}
		tokens := strings.Split(line, " ")
		ft := make([]string, 0)
		for _, t := range(tokens) {
			if t!="" && t!="\t" {
				ft = append(ft, t)
			}
		}
		pid, err := strconv.Atoi(ft[1])
		if err!=nil {
			continue
		}
		cpu, err := strconv.ParseFloat(ft[2], 64)
		if err!=nil {
			log.Fatal(err)
		}
		processes = append(processes, &Process{pid, cpu})
	}
	return processes[0].cpu
}

func Pipeline(cmds ...*exec.Cmd) (pipeLineOutput bytes.Buffer, collectedStandardError []byte, pipeLineError error) {

	// Collect the output from the command(s)
	var output bytes.Buffer
	var stderr bytes.Buffer

	last := len(cmds) - 1
	for i, cmd := range cmds[:last] {
		var err error
		// Connect each command's stdin to the previous command's stdout
		if cmds[i+1].Stdin, err = cmd.StdoutPipe(); err != nil {
			return bytes.Buffer{}, nil, err
		}
		// Connect each command's stderr to a buffer
		cmd.Stderr = &stderr
	}

	// Connect the output and error for the last command
	cmds[last].Stdout, cmds[last].Stderr = &output, &stderr

	// Start each command
	for _, cmd := range cmds {
		if err := cmd.Start(); err != nil {
			return output, stderr.Bytes(), err
		}
	}

	// Wait for each command to complete
	for _, cmd := range cmds {
		if err := cmd.Wait(); err != nil {
			return output, stderr.Bytes(), err
		}
	}

	// Return the pipeline output and the collected standard error
	return output, stderr.Bytes(), nil
}