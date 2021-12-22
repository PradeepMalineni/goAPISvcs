package main

import (
	"fmt"
	"time"

	"github.com/PradeepMalineni/goAPISvcs/v1/datalayer"
	"github.com/PradeepMalineni/goAPISvcs/v1/events"
	"github.com/PradeepMalineni/goAPISvcs/v1/timedata"
)

func main() {
	a()
	data := events.Aexport("Sent1")

	fmt.Println(data)
	fmt.Println(timedata.GetTime())
	//getdata.CallHttpData()

	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))

	client, ctx, cacel, err := datalayer.Connect("mongodb://0.0.0.0:27017")
	if err != nil {
		panic(err)
	}
	defer datalayer.Close(client, ctx, cacel)

	datalayer.Ping(client, ctx)

}
func expensiveCall() {
	time.Sleep(10 * time.Second)
}
func a() {
	fmt.Println("In func a")
	defer b()
	fmt.Println("Ends a")
}

func b() {
	fmt.Println("in b")
}
