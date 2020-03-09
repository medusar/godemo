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

	kv := clientv3.NewKV(cli)

	go func() {
		for {
			kv.Put(context.TODO(), "/watch/demo", "haha")
			time.Sleep(1 * time.Second)
			kv.Delete(context.TODO(), "/watch/demo")
		}
	}()

	get, err := kv.Get(context.TODO(), "/watch/demo")
	if err != nil {
		log.Fatal(err)
	}

	if len(get.Kvs) != 0 {
		log.Println("current data:", get.Kvs)
	}
	revision := get.Header.Revision
	log.Println("current revision:", revision)

	watcher := clientv3.NewWatcher(cli)

	ctx, _ := context.WithTimeout(context.TODO(), 20*time.Second)
	watchChan := watcher.Watch(ctx, "/watch/demo")
	//watchChan := watcher.Watch(ctx, "/watch/demo", clientv3.WithRev(revision))
	for wt := range watchChan {
		for _, event := range wt.Events {
			log.Println("event:", event.Type, "value:", string(event.Kv.Value), "create revision:", event.Kv.CreateRevision, "mod revision:", event.Kv.ModRevision)
		}
	}
}
