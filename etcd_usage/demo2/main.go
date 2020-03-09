package main

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	//get a lease
	lease := clientv3.NewLease(cli)
	grant, err := lease.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}
	leaseId := grant.ID

	// keep the lease alive for 5 more seconds
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	// extend lease lifetime
	//alive, err := lease.KeepAlive(context.TODO(), leaseId)
	alive, err := lease.KeepAlive(ctx, leaseId)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case kar := <-alive:
				if kar == nil {
					log.Println("lease expired")
					return
				} else {
					log.Println("lease extended, id:", kar.ID)
				}
			}
		}
	}()

	kv := clientv3.NewKV(cli)
	//put data with lease
	_, err = kv.Put(context.TODO(), "/ttl/lock1", "", clientv3.WithLease(leaseId))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("put with lease successfully")
	for {
		resp, err := kv.Get(context.TODO(), "/ttl/lock1")
		if err != nil {
			log.Fatal(err)
		}
		if len(resp.Kvs) > 0 {
			log.Println(resp.Kvs)
		} else {
			log.Println("key removed")
			break
		}
		time.Sleep(1 * time.Second)
	}
}
