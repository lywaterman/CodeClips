package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/client"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"net"
	"time"
)

func IsOpened(host string, port int) bool {

	timeout := 5 * time.Second
	target := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", target, timeout)
	if err != nil {
		return false
	}

	if conn != nil {
		conn.Close()
		return true
	}

	return false
}

func main() {
	fmt.Println("sfsfs")
	v, _ := mem.VirtualMemory()
	fmt.Println(v)

	//vP, _ := disk.Partitions(true)

	//fmt.Println(vP)

	vU, _ := disk.Usage("/home")
	fmt.Println(vU)

	dockerRunning := false
	//processes, _ := process.Processes()
	//for _, p := range processes {
	//	name, _ := p.Name()
	//	if strings.Contains(name, "docker") {
	//		dockerRunning = true
	//		break
	//	}
	//}

	//fmt.Println(dockerRunning)

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

	ping, err := cli.Ping(ctx)

	if err != nil {
		dockerRunning = false
	} else {
		dockerRunning = true
	}
	fmt.Println(ping)
	fmt.Println(err)
	fmt.Println(dockerRunning)

	isOpened := IsOpened("127.0.0.1", 6445)

	fmt.Println(isOpened)
}
