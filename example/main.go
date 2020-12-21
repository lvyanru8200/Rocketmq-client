package main

import (
	"context"
	"fmt"
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
	/*		json1, _ := testAdmin.GetRouteInfo(context.Background(), admin.WithNamesrvAddrRoute("10.21.7.5:31010"), admin.WithTopicRoute("testcr"))*/
	//json,_:=testAdmin.GetConsumeStats(context.Background(),admin.WithConsumerGroupGetConsume("testGroup"),admin.WithTopicGroupGetConsume("testcr"),admin.WithBrokerAddrGetConsume("10.21.7.5:47505"))
	/*		json2, _ := testAdmin.GetConsumeStatsInBroker(context.Background(), admin.WithConsumerStatsInIsOrder(true), admin.WithConsumeStatsInBrokerBroker("10.21.7.6:53080"))
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

	/*j,_:=testAdmin.GetBrokerClusterInfo(context.Background(), "10.21.7.5:31010")
	fmt.Println(j)
	a:=j.Get("brokerAddrTable").Get("n22").Get("brokerAddrs").String()

	b:=a[1:len(a)-1]
	fmt.Println(b)
	p:=strings.Split(a[1:len(a)-1],":")
	fmt.Println(p[0])
	fmt.Println(p[1][1:])
	fmt.Println(p[2][:len(p[2])-1])
	ee:=fmt.Sprintf(p[1][1:]+":"+p[2][:len(p[2])-1])
	fmt.Println(ee)*/
	/*	err:=testAdmin.DeleteSubscriptionGroup(context.Background(),admin.WithDeleteSubGroup("testGroup"),admin.WithDeletebrokeraddr("10.21.7.7:56042"))*/
	/*	if err!=nil{
			fmt.Println(err)
		}
	*/
	/*	fmt.Println(json1)
		fmt.Println(json1.QueueDataList[0].BrokerName)*/

	p, _ := testAdmin.GetConsumerOffset(context.Background(), admin.WithConsumeroffsetBrokerAddr("10.21.7.6:45612"), admin.WithConsumerOffsetGroup("testupdate"), admin.WithConsumerOffsetTopic("testcr"))
	fmt.Println(p)
	testAdmin.UpdateConsumerOffset(context.Background(), admin.WithUpdateoffsetBrokerAddr("10.21.7.6:45612"), admin.WithUpdateoffsetCommitoffset(p), admin.WithUpdateoffsetConsumerGroup("testupdate"))
}
