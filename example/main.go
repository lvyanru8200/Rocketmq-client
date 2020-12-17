package main

import (
	"context"
	"fmt"
	"github.com/lvyanru8200/rocketmq-client-go/admin"
	"github.com/lvyanru8200/rocketmq-client-go/primitive"
)

func main() {
	nameSrvAddr := []string{"10.21.7.6:32666"}
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	/*		json1, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-1"))*/
	/*		fmt.Println(json1)*/
	/*	json2, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-2"))*/
	/*	fmt.Println(json2)*/
	//topiclist:=json.Get("topicList").MustStringArray()
	//json,_:=testAdmin.GetConsumerRunningInfo(context.Background(),admin.WithConsumerBrokerAddr("10.21.7.6:54199"),admin.WithConsumerGroup("testGroup"))
	//json,_:=testAdmin.GetBrokerRuntimeInfo(context.Background(),"10.21.7.6:32666","10.21.7.6:43213")
	//json,_:=testAdmin.QueryTopicConsumeByWho(context.Background(),admin.WithBrokeraddr("10.21.7.5:47505"),admin.WithTopic("testcr"))
	json1, _ := testAdmin.GetRouteInfo(context.Background(), admin.WithNamesrvAddrRoute("10.21.7.6:32666"), admin.WithTopicRoute("testcr"))
	//json,_:=testAdmin.GetConsumeStats(context.Background(),admin.WithConsumerGroupGetConsume("testGroup"),admin.WithTopicGroupGetConsume("testcr"),admin.WithBrokerAddrGetConsume("10.21.7.5:47505"))
	json2, _ := testAdmin.GetConsumeStatsInBroker(context.Background(), admin.WithConsumerStatsInIsOrder(true), admin.WithConsumeStatsInBrokerBroker("10.21.7.6:54199"))
	p := json2.Get("totalDiff").Int()
	fmt.Println(p)
	//p:=json.Get("groupList").MustStringArray()
	fmt.Println(json1)
	fmt.Println(json2)

	//fmt.Println(json)

}
