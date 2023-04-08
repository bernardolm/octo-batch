package main

import (
	"log"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/opts"

	"github.com/bernardolm/octo-batch/chart"
	"github.com/bernardolm/octo-batch/config"
)

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func main() {
	_ = config.Init()

	// ctx := context.Background()

	// Github
	// repos, err := github.RepositoriesListAll(ctx)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// for _, r := range repos {
	// 	pp.Println(r.GetFullName())
	// }

	// if err := github.ActivitySetRepositoriesSubscription(ctx, repos); err != nil {
	// 	log.Panic(err)
	// }
	// Github - end

	// Just Sift
	// if err := sift.Start(); err != nil {
	// 	log.Panic(err)
	// }

	// people, err := sift.PeopleListAll(ctx)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// Just Sift - end

	// Charts
	// nodes, link := sift.PeopleIntoSankeyChart(people)
	// if err != nil {
	// 	log.Panic(err)
	// }

	// if err := chart.BuildSankeyChart(nodes, link); err != nil {
	// 	log.Panic(err)
	// }

	// data := sift.PeopleIntoTreeChart(people)
	// if err != nil {
	// 	log.Panic(err)
	// }

	data := []opts.TreeData{
		{
			Name: "JR",
			Children: []*opts.TreeData{
				{
					Name: "Zanaca",
					Children: []*opts.TreeData{
						{
							Name: "Be",
							Children: []*opts.TreeData{
								{
									Name: "Fernanda",
								},
								{
									Name: "Vini",
								},
							},
						},
						{
							Name: "Bruno",
							Children: []*opts.TreeData{
								{
									Name: "André",
									Children: []*opts.TreeData{
										{
											Name: "Spaniol",
										},
										{
											Name: "Vini",
										},
									},
								},
								{
									Name: "Zafas",
								},
							},
						},
					},
				},
			},
		},
	}
	if err := chart.BuildTreeChart(data); err != nil {
		log.Panic(err)
	}
	// Charts - end

	fs := http.FileServer(http.Dir("html"))
	log.Println("running server at http://localhost:8089")
	go log.Fatal(http.ListenAndServe("localhost:8089", logRequest(fs)))

	// c := make(chan error)

	// go func() {
	// 	port := ":8089"
	// 	log.Println("running server at http://localhost" + port)
	// 	if err := http.ListenAndServe(port, logRequest(fs)); err != nil {
	// 		c <- err
	// 	}
	// }()

	// for {
	// 	select {
	// 	case err := <-c:
	// 		panic(err)
	// 	default:
	// 		// nothing
	// 	}
	// }
}
