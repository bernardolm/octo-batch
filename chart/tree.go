package chart

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/bernardolm/octo-batch/debug"
)

func BuildTreeChart(data []opts.TreeData) error {
	// debug.Print("data", data)

	var TreeNodes = []*opts.TreeData{
		{
			Name: "Node1",
			Children: []*opts.TreeData{
				{
					Name: "Chield1",
				},
			},
		},
		{
			Name: "Node2",
			Children: []*opts.TreeData{
				{
					Name: "Chield1",
				},
				{
					Name: "Chield2",
				},
				{
					Name: "Chield3",
				},
			},
		},
		{
			Name:      "Node3",
			Collapsed: true,
			Children: []*opts.TreeData{
				{
					Name: "Chield1",
				},
				{
					Name: "Chield2",
				},
				{
					Name: "Chield3",
				},
			},
		},
	}

	var Tree = []opts.TreeData{
		{
			Name:     "Root",
			Children: TreeNodes,
		},
	}

	page := components.NewPage()

	graph := charts.NewTree()

	graph.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "basic tree example"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)

	debug.Print("Tree", Tree)

	graph.AddSeries("treeeeeee", Tree).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)

	graph1 := charts.NewTree()
	graph1.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "basic tree example"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)

	graph1.AddSeries("VAI!", data).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)

	graph2 := charts.NewTree()
	graph2.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Width: "100%", Height: "95vh"}),
		charts.WithTitleOpts(opts.Title{Title: "basic tree example"}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
	)

	debug.Print("data", data)

	graph2.AddSeries("Hurb", data).
		SetSeriesOptions(
			charts.WithTreeOpts(
				opts.TreeChart{
					Layout:           "orthogonal",
					Orient:           "LR",
					InitialTreeDepth: -1,
					Leaves: &opts.TreeLeaves{
						Label: &opts.Label{Show: true, Position: "right", Color: "Black"},
					},
				},
			),
			charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}),
		)

	page.AddCharts(graph, graph1, graph2)

	f, err := os.Create("html/tree.html")
	if err != nil {
		return err
	}

	if err := page.Render(io.MultiWriter(f)); err != nil {
		return err
	}

	return nil
}
