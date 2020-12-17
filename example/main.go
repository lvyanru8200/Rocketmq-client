package main

import (
	"context"
	"fmt"
	"rocketmq-client-go/admin"
	"rocketmq-client-go/primitive"
)

func main() {
	nameSrvAddr := []string{"10.21.7.6:32666"}
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	json, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-1"))
	/*	json,_:=testAdmin.GetConsumerRunningInfo(context.Background(),admin.WithConsumerBrokerAddr("10.21.7.6:43213"),admin.WithConsumerGroup("testGroup"))*/
	/*json,_:=testAdmin.GetBrokerRuntimeInfo(context.Background(),"10.21.7.6:32666","10.21.7.6:43213")*/
	/*json,_:=testAdmin.GetConsumeStats(context.Background(),"10.21.7.6:54199")*/
	fmt.Println(json)

}
