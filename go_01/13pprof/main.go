package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// 使用命令行工具查看pprof
// go tool pprof cpu.pprof

func fb(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return fb(n-1) + fb(n-2)
}

var isCPUPprof bool
var isMemPprof bool

func main() {
	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()

	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			log.Fatal("create file failed,err:", err)
		}
		pprof.StartCPUProfile(file)
		defer func() {
			pprof.StopCPUProfile()
			file.Close()
		}()
	}

	n := fb(20)
	fmt.Println(n)

	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			log.Fatal("create file failed,err:", err)
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}
