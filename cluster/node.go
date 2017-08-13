package cluster

import (
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// Node 节点对象
type Node struct {
	ID        int    // 节点id
	Addr      string // 地址
	TaskCount int    // 任务数
	tryCount  int
}

func (n *Node) heartCheck() bool {
	result := false
	client := getClinet()
	resp, err := client.Get("http://" + n.Addr + "/cluster/heart")
	if err == nil {
		_, err = ioutil.ReadAll(resp.Body)
		result = (err == nil && resp.StatusCode == 200)
		if result {
			n.tryCount = 0
		}
	} else {
		result = false
	}
	if !result {
		n.tryCount++
	}
	return n.tryCount < NodeCheckTryTimes || result
}

func getClinet() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*2)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * 2))
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * 2,
		},
	}
	return client
}
