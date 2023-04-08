package sift

import (
	"github.com/go-echarts/go-echarts/v2/opts"
)

func PeopleIntoSankeyChart(people []person) ([]opts.SankeyNode, []opts.SankeyLink) {
	var sankeyNodes = []opts.SankeyNode{}
	for _, v := range people {
		sankeyNodes = append(sankeyNodes, opts.SankeyNode{
			Name: v.Id,
		})
	}

	var sankeyLinks = []opts.SankeyLink{}
	for _, v := range people {
		if v.TotalReportCount == 0 {
			v.TotalReportCount = 1
		}

		sankeyLinks = append(sankeyLinks, opts.SankeyLink{
			Source: v.TeamLeaderId,
			Target: v.Id,
			Value:  10 * float32(v.TotalReportCount),
		})
	}

	return sankeyNodes, sankeyLinks
}
