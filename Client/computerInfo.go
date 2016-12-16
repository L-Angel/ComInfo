package main

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"time"
	"net/http"
	"encoding/json"
)
const filename="machineinfo.txt"
func responseInfo(w http.ResponseWriter,r *http.Request){
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	cc,_:=cpu.Percent(1*time.Second,true)
	result:=make(map[string]interface{})
	result["meminfo"]=v
	result["cpuinfo"]=c
	result["cpures"]=cc
	rr,_:=json.Marshal(result)
	fmt.Fprint(w,string(rr))

}
func main() {
	http.HandleFunc("/",responseInfo)
	http.ListenAndServe(":2552",nil)
}