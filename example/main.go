package main

import "numa"

func main() {
    cpus, _:= numa.GetNodesCoresInfo()
    fmt.Println(cpus)
    ram, _ := numa.GetNodesMemoryInfo()
    fmt.Println(ram)
    nics,_ := numa.GetNodesNicsInfo()
    fmt.Println(nics)
}

