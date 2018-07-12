package main

import (
	"runtime"
	"LearningGo/CPUInfo/Utility"
	"fmt"
)

var iterations = 10000

func main() {
	runtime.GOMAXPROCS(1)
	processCompletion := make(chan bool)
	//go CPUMonitor()
	for i := 0; i < iterations; i++ {
		go loop(processCompletion, i)
	}
	<- processCompletion
}

func loop (processCompletion chan bool, i int) {
	count := 0
	for j := 0; ; j++ {
		count += j
		handleCpuUsage()
	}
	if i == iterations - 1 {
		processCompletion <- true
	}
}

/*func CPUMonitor() {
	fmt.Println("Monitoring CPU Usage")
		for {
			select {
				case <-time.After(1 * time.Second):
					handleCpuUsage()
			}
		}
}*/

func handleCpuUsage() {
	currentCpuCount := float64(runtime.GOMAXPROCS(0))
	cpuUsage := MacCpuUtility.GoCpuUsage()
	fmt.Println("In")
	if (100/currentCpuCount - cpuUsage) <= 5 {
		fmt.Println("CPU usage reached limit: ", cpuUsage, " Increasing processors to ", runtime.GOMAXPROCS(0) + 1)
		runtime.GOMAXPROCS(runtime.GOMAXPROCS(0) + 1)
	} else {
		fmt.Println("CPU usage: ", cpuUsage)
	}
	fmt.Println("Out")
}