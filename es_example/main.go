package main


import (
	"fmt"

	elastic "github.com/elastic/go-elasticsearch"
)

type Tweet struct {
	User string
	Message string
}

func main()  {
	client,err :=elastic.
}

