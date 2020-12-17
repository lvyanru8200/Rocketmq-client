package main

import (
	"context"
	"rocketmq-client-go/admin"
	"rocketmq-client-go/primitive"
)

func main() {
	nameSrvAddr := []string{"10.21.7.6:32666"}
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	/*	json, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-2"))
		topiclist:=json.Get("topicList").MustStringArray()*/
	/*	json,_:=testAdmin.GetConsumerRunningInfo(context.Background(),admin.WithConsumerBrokerAddr("10.21.7.6:43213"),admin.WithConsumerGroup("testGroup"))*/
	/*json,_:=testAdmin.GetBrokerRuntimeInfo(context.Background(),"10.21.7.6:32666","10.21.7.6:43213")*/
	/*json,_:=testAdmin.GetConsumeStats(context.Background(),"10.21.7.6:54199")*/
	/*	json,_:=testAdmin.QueryTopicConsumeByWho(context.Background(),admin.WithBrokeraddr("10.21.7.7:54262"),admin.WithTopic("testcr"))*/
	/*testAdmin.GetRouteInfo(context.Background(),admin.WithNamesrvAddrRoute("10.21.7.6:32666"),admin.WithTopicRoute("testcr"))*/
	/*	fmt.Println(json)*/

}
