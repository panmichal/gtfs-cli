package gtfs

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

type Route struct {
	Name string
}

type Feed struct {
	routes []Route
}

func (f *Feed) GetRoutes() []Route {
	return f.routes
}

func CreateFeed(files *feedFiles) *Feed {

	routes := createRoutes(files.RouteFile)

	return &Feed{
		routes: routes,
	}
}

func createRoutes(routesContent []byte) []Route {
	var routes []Route
	r := csv.NewReader(strings.NewReader(string(routesContent)))
	_, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//get route name from index 2 of record
		routeName := record[2]
		routes = append(routes, Route{Name: routeName})
		// fmt.Println(record)
	}

	return routes
}
