package gtfs

import "fmt"

type Route struct {
	Name string
}

type Feed struct {
	routes []Route
}

func (f *Feed) Routes() []Route {
	return f.routes
}

func CreateFeed(files feedFiles) *Feed {
	for files.routeFile.Scan() {
		fmt.Println(files.routeFile.Text())
	}

	return &Feed{
		routes: []Route{
			{Name: "Route 1"},
			{Name: "Route 2"},
			{Name: "Route 3"},
		},
	}
}
