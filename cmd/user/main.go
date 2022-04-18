package main

import (
	"app/workflow/user"
	"context"
	"fmt"
	"go.temporal.io/sdk/worker"

	"go.temporal.io/sdk/client"
)

const (
	wfid = "mleow-0"
	q    = "dev.user.svc"
)

func main() {
	// The all-in-one dev for User service ..
	fmt.Println("START run ===========================> *****")
	c, err := client.NewClient(client.Options{
		//HostPort: "temporal:7233",
	})
	if err != nil {
		panic(err)
	}
	defer c.Close()

	m := user.MembershipSvc{}
	// Start the workflow ..
	go func() {
		wf, exerr := c.ExecuteWorkflow(context.Background(),
			client.StartWorkflowOptions{
				ID:        wfid,
				TaskQueue: q,
			}, m.MembershipWorkflow,
			user.MembershipReq{})
		if exerr != nil {
			panic(exerr)
		}
		rerr := wf.Get(context.Background(), nil)
		if rerr != nil {
			panic(rerr)
		}
	}()

	// Start the worker ..
	w := worker.New(c, q, worker.Options{})
	w.RegisterWorkflow(m.MembershipWorkflow)
	w.Run(worker.InterruptCh())
	// Below is an incomplete simplified version of Run
	//w.Start()
	//// Block till done ..
	//termChan := worker.InterruptCh()
	//<-termChan
	//w.Stop()

}
