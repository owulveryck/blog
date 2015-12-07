package main

import (
	"encoding/json"
	"fmt"
)

type Input struct {
	Name    string `json:name,omitempty`
	Digraph []int  `json:digraph`
	Nodes   []Node `json:nodes`
}

type Node struct {
	ID       int               `json:id`
	Name     string            `json:name,omitempty`
	Engine   string            `json:engine:omitempty` // The execution engine (ie ansible, shell); aim to be like a shebang in a shell file
	Artifact string            `json:artifact`
	Args     []string          `json:args,omitempty`   // the arguments of the artifact, if needed
	Inputs   map[string]string `json:input,omitempty`  // the key is the node that gives the information, and  the value, its output (always a string)
	Outputs  map[string]string `json:output,omitempty` // the key is the name of the parameter, the value its value (always a string)
}

func main() {
	test := Input{"Test",
		[]int{0, 1, 0, 0, 1, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 1, 0, 0, 0, 1, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			1, 1, 1, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 1, 0,
		},
		[]Node{
			{0, "a", "ansible", "myplaybook.yml", nil, nil, nil},
			{1, "b", "ansible", "myplaybook1.yml", nil, nil,
				map[string]string{
					"output1": "",
				},
			},
			{2, "c", "ansible", "myplaybook2.yml", nil,
				map[string]string{
					"b": "output1",
				}, nil},
			{3, "d", "ansible", "myplaybook3.yml", nil, nil, nil},
			{4, "e", "ansible", "myplaybook4.yml", nil, nil, nil},
			{5, "f", "ansible", "myplaybook5.yml", nil, nil, nil},
			{6, "g", "ansible", "myplaybook6.yml", nil, nil, nil},
			{7, "h", "ansible", "myplaybook7.yml", nil, nil, nil},
		},
	}
	o, err := json.MarshalIndent(test, "  ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", o)
}
