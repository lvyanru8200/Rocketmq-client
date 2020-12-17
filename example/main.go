package main

import (
	"context"
	"rocketmq-client-go/admin"
	"rocketmq-client-go/primitive"
)

func main() {
	nameSrvAddr := []string{"10.21.7.6:32666"}
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	/*		json1, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-1"))*/
	/*		fmt.Println(json1)*/
	/*	json2, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-2"))*/
	/*	fmt.Println(json2)*/
	//topiclist:=json.Get("topicList").MustStringArray()
	/*json,_:=testAdmin.GetConsumerRunningInfo(context.Background(),admin.WithConsumerBrokerAddr("10.21.7.6:54199"),admin.WithConsumerGroup("testGroup"))*/
	/*	json,_:=testAdmin.GetBrokerRuntimeInfo(context.Background(),"10.21.7.6:32666","10.21.7.6:43213")*/
	/*	json,_:=testAdmin.GetConsumeStats(context.Background(),"10.21.7.6:54199")*/
	/*		json,_:=testAdmin.QueryTopicConsumeByWho(context.Background(),admin.WithBrokeraddr("10.21.7.5:47505"),admin.WithTopic("%RETRY%testGroup"))*/
	/*	json,_:=testAdmin.GetRouteInfo(context.Background(),admin.WithNamesrvAddrRoute("10.21.7.6:32666"),admin.WithTopicRoute("testcr"))*/
	/*	json,_:=testAdmin.GetConsumeStats(context.Background(),admin.WithConsumerGroupGetConsume("testGroup"),admin.WithTopicGroupGetConsume("testcr"),admin.WithBrokerAddrGetConsume("10.21.7.5:47505"))*/
	testAdmin.GetConsumeStatsInBroker(context.Background(), admin.WithConsumerStatsInIsOrder(true), admin.WithConsumeStatsInBrokerBroker("10.21.7.5:47505"))
	//fmt.Println(json)

}
