package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"time"
)
const filename="machineinfo.txt"

func main() {
	for {
		v, _ := mem.VirtualMemory()
		c, _ := cpu.Info()
		cc,_:=cpu.Percent(1*time.Second,true)
		// almost every return value is a struct
		fmt.Println("        Machine Performance Infomaiton")
		//io.WriteString(f,"        Machine Performance Infomaiton");
		fmt.Printf("MEM   : %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)
		//f.Write("MEM   : "+v.Total+", Free:"+v.Free+", UsedPercent:"+v.UsedPercent+"%\n")
		if len(c) > 1&&len(cc)>1 {

			for index, sub_cpu := range c {
				modelname := sub_cpu.ModelName
				cores := sub_cpu.Cores
				fmt.Printf("CPU   : %v   %v cores, UsedPercent:%f%% \n", modelname, cores,cc[index])
			}
		} else {
			sub_cpu := c[0]
			modelname := sub_cpu.ModelName
			cores := sub_cpu.Cores
			//fmt.Printf("        CPU       : %v   %v cores \n", modelname, cores)
			fmt.Printf("CPU   : %v   %v cores, UsedPercent:%f%% \n", modelname, cores,cc[0])


		}
		// convert to JSON. String() is also implemented
		//fmt.Println(v)
		time.Sleep(2*time.Second)
	}
}