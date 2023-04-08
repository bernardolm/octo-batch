package sift

import (
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/bernardolm/octo-batch/debug"
)

var (
	nodeTable = map[string]*opts.TreeData{}
	root      []opts.TreeData
)

func PeopleIntoTreeChart(items []person) []opts.TreeData {
	add := func(id, name, parentId string) {
		node := &opts.TreeData{
			Name:     name,
			Children: []*opts.TreeData{},
		}

		if parentId == "" {
			debug.Print("node without parent", node)

			root[0].Children = append(root[0].Children, node)
		} else {
			parent, ok := nodeTable[parentId]
			if !ok {
				return
			}

			parent.Children = append(parent.Children, node)
		}

		nodeTable[id] = node
	}

	scan := func() {
		for _, v := range items {
			add(v.Id, v.FirstName, v.TeamLeaderId)
		}
	}

	root = append(root, opts.TreeData{
		Name: "root",
	})

	scan()

	debug.Print("root", root)

	return nil
}
