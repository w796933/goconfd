package node

import "encoding/json"

type Node struct {
	Name           string `json:"name"`
	IP             string `json:"ip"`
	CPU            int    `json:"cpu"`
	ServiceAddress string `json:"service_address"`
}

func (n *Node) String() string {
	value, _ := json.Marshal(n)
	return string(value)
}
