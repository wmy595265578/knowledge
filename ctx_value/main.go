package main

import (
	"context"
	"fmt"
)

func process(ctx  context.Context){
	ret ,ok:= ctx.Value("trace_id").(int)
	if !ok {
		ret = 1111
	}
	fmt.Println(ret)
	s ,_:= ctx.Value("session").(string)
	fmt.Printf(s)
}

func main()  {

	ctx := context.WithValue(context.Background(),"trace_id",595265578)
	ctx = context.WithValue(ctx,"session","asdasdasdsadadsa")
	process(ctx)
}
