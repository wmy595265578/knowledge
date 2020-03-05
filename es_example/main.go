package main

import (
	"fmt"
	elastic "gopkg.in/olivere/elastic.v2"
)

type Tweet struct {
	User string
	Message string
}

func main()  {
	client,err :=elastic.NewClient(elastic.SetSniff(false),elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		fmt.Println("connect es error", err)
		return
	}

	fmt.Println("conn es succ")
	tweet := Tweet{User: "olivere", Message: "Take Five"}
	for i:=2;i<=100;i++ {
		_, err = client.Index().
			Index("twitter").
			Type("tweet").
			Id(fmt.Sprintf("%d",i)).
			BodyJson(tweet).
			Do()

		if err != nil {
			// Handle error
			panic(err)
			return
		}
	}
	fmt.Println("insert succ")
}

