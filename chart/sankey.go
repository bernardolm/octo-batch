package chart

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func BuildSankeyChart(nodes []opts.SankeyNode, links []opts.SankeyLink) error {
	graph := charts.NewSankey()
	graph.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Sankey-basic-example",
		}),
	)

	graph.AddSeries("sankey", nodes, links,
		charts.WithLabelOpts(opts.Label{Show: true})).
		SetSeriesOptions(
			charts.WithLineStyleOpts(opts.LineStyle{
				Color:     "source",
				Curveness: 0.5,
			}),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)

	page := components.NewPage()
	page.AddCharts(graph)

	f, err := os.Create("html/sankey.html")
	if err != nil {
		return err
	}

	if err := page.Render(io.MultiWriter(f)); err != nil {
		return err
	}

	return nil
}
