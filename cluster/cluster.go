package cluster

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var C Cluster
var nodeId int

func init() {
	// 初始化配置
	InitConfig()
	//初始化clcluster
	C = Cluster{
		Name: ClusterName,
	}
	if Master == "" {
		C.Master = &Node{
			Addr: NodeAddr,
		}
		C.Regist(C.Master)
	} else {
		RegistCurrent(&Node{
			Addr: NodeAddr,
		})
	}
	if C.Master != nil && C.Master.Addr == NodeAddr {
		go C.heartCheck()
	}
}

// Cluster 是集群对象
type Cluster struct {
	Name   string  // 集群名称
	Nodes  []*Node // 节点数组
	Master *Node   // 管理节点
}

// Regist 注册节点
func (c *Cluster) Regist(n *Node) error {
	for _, node := range c.Nodes {
		if node.Addr == n.Addr {
			return errors.New("exists node")
		}
	}
	nodeId++
	n.ID = nodeId
	c.Nodes = append(c.Nodes, n)
	return nil
}

// Remove 移除节点
func (c *Cluster) Remove(n *Node) error {
	return c.RemoveByID(n.ID)
}

// RemoveByID 通过Id删除节点
func (c *Cluster) RemoveByID(nodeID int) error {
	nodes := make([]*Node, 0)
	for _, node := range c.Nodes {
		if node.ID == nodeID {
			continue
		}
		nodes = append(nodes, node)
	}
	if len(c.Nodes) == len(nodes)+1 {
		c.Nodes = nodes
		return nil
	}
	return errors.New("Node is not exists")
}

func (c *Cluster) heartCheck() {
	for {
		for _, node := range c.Nodes {
			if node.Addr == c.Master.Addr {
				continue
			}
			result := node.heartCheck()
			fmt.Printf("Node: %v, addr: %s, check result: %v\n", node.ID, node.Addr, result)
			if !result {
				c.Remove(node)
				fmt.Printf("Node: %v, addr: %s, removed\n", node.ID, node.Addr)
			}
		}
		time.Sleep(5 * time.Second)
	}
}

type Result struct {
	Message string `message`
}

func RegistCurrent(n *Node) {
	for {
		requstBytes, err := json.Marshal(n)
		if err != nil {
			fmt.Printf("node config %s is error\n", string(requstBytes))
			break
		}
		resp, err := http.Post("http://"+Master+"/cluster/regist", "application/json", strings.NewReader(string(requstBytes)))
		if err != nil {
			time.Sleep(time.Second * 2)
			fmt.Printf("master is %s, node add failed, will be try again 2 second later \n", Master)
			continue
		}
		var result *Result
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			time.Sleep(time.Second * 2)
			fmt.Printf("master is %s, node add failed, error message is %s, will be try again 2 second later \n", Master, fmt.Sprintf("%s", err))
			continue
		}
		err = json.Unmarshal(body, &result)
		if err != nil {
			time.Sleep(time.Second * 2)
			fmt.Printf("master is %s, node add failed, error message is %s, will be try again 2 second later \n", Master, fmt.Sprintf("%s", err))
			continue
		}
		if "success" == result.Message {
			fmt.Printf("master is %s, node add %s success\n", Master, n.Addr)
		} else {
			time.Sleep(time.Second * 2)
			fmt.Printf("master is %s, node add failed, error message is %s, will be try again 2 second later \n", Master, result.Message)
		}
		break
	}
}
