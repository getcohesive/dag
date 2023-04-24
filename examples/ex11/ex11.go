package main

import (
	"context"
	"fmt"
	"github.com/getcohesive/dag"
	"time"
)

func main() {
	d := dag.New()
	d.Spawns(f1, f2, f3)

	_ = d.RunAsync(context.Background(), func(err error) {
		if err != nil {
			fmt.Println("Error: ", err.Error())
		}
		fmt.Println("Done")
	})
	time.Sleep(300 * time.Millisecond)
	err := d.Stop(context.Background())
	if err != nil {
		fmt.Println("Error while stopping: ", err.Error())
		return
	}
}

func f1(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		println("f1", i)
	}
	return nil
}

func f2(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		println("f2", i)
	}
	return nil
}

func f3(ctx context.Context) error {
	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Millisecond)
		println("f3", i)
	}
	return nil
}
