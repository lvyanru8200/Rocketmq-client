package main

import (
	"context"
	"github.com/lvyanru8200/rocketmq-client/admin"
	"github.com/lvyanru8200/rocketmq-client/primitive"
)

func main() {
	nameSrvAddr := []string{"10.21.7.6:32284"}
	testAdmin, _ := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	/*		json1, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-1"))*/
	/*		fmt.Println(json1)*/
	/*	json2, _ := testAdmin.GetTopicsByCluster(context.Background(), admin.WithNameserver("10.21.7.6:32666"), admin.WithClusterTopicList("raft-2"))*/
	/*	fmt.Println(json2)*/
	//topiclist:=json.Get("topicList").MustStringArray()
	//json,_:=testAdmin.GetConsumerRunningInfo(context.Background(),admin.WithConsumerBrokerAddr("10.21.7.6:54199"),admin.WithConsumerGroup("testGroup"))
	/*	json,_:=testAdmin.GetBrokerRuntimeInfo(context.Background(),"10.21.7.6:32011","10.21.7.7:56042")*/
	//json,_:=testAdmin.QueryTopicConsumeByWho(context.Background(),admin.WithBrokeraddr("10.21.7.5:47505"),admin.WithTopic("testcr"))
	//	json1, _ := testAdmin.GetRouteInfo(context.Background(), admin.WithNamesrvAddrRoute("10.21.7.6:32011"), admin.WithTopicRoute("testcr"))
	//json,_:=testAdmin.GetConsumeStats(context.Background(),admin.WithConsumerGroupGetConsume("testGroup"),admin.WithTopicGroupGetConsume("testcr"),admin.WithBrokerAddrGetConsume("10.21.7.5:47505"))
	/*	json2, _ := testAdmin.GetConsumeStatsInBroker(context.Background(), admin.WithConsumerStatsInIsOrder(true), admin.WithConsumeStatsInBrokerBroker("10.21.7.6:53080"))
		p := json2.Get("totalDiff").Int()*/
	//fmt.Println(p)
	//p:=json.Get("groupList").MustStringArray()
	/*	fmt.Println(strconv.FormatInt(p,2))*/
	//fmt.Println(json2)
	/*	err:=testAdmin.UpdateAndCreateSubscriptionGroup(context.Background(),admin.WithBrokerAddr("10.21.7.7:53334"),admin.WithGroupName("testGroup"))
		fmt.Println(err)*/
	/*	j,err:=testAdmin.QueryTopicConsumeByWho(context.Background(),admin.WithBrokeraddr("10.21.7.7:56042"),admin.WithTopic("testcr"))
		if err!=nil{
			fmt.Println(err)
		}
		fmt.Println(j.Get("groupList").MustStringArray()[0])*/
	/*j:=testAdmin.CreateTopic(context.Background(),admin.WithTopicCreate("testcr"),admin.WithTopicFilterType("SINGLE_TAG"),admin.WithBrokerAddrCreate("10.21.7.7:53334"))*/

	testAdmin.GetBrokerClusterInfo(context.Background(), "10.21.7.6:32011")
	//fmt.Println(j)

}
