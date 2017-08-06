package cluster

import "errors"

// Cluster 是集群对象
type Cluster struct {
	Name      string  // 集群名称
	NodeCount int     // 节点数量
	Nodes     []*Node // 节点数组
	Master    Node    // 管理节点
}

// Regist 注册节点
func (c *Cluster) Regist(n *Node) error {
	for _, node := range c.Nodes {
		if node.ID == n.ID {
			return errors.New("exists node")
		}
	}
	c.Nodes = append(c.Nodes, n)
	return nil
}

// Remove 移除节点
func (c *Cluster) Remove(n *Node) error {
	return c.RemoveByID(n.ID)
}

// RemoveByID 通过Id删除节点
func (c *Cluster) RemoveByID(nodeID string) error {
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
