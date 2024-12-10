package main

import "github.com/bluesky0724/jsonstream"

func main() {
	jsonstream.JSON2CSV("url", "https://open.gsa.gov/data.json", "output.csv", ".dataset", []string{"modified", "publisher.name", "publisher.subOrganizationOf.name", "contactPoint.fn", "keyword"})
}
