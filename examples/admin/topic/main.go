/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/primitive"
)

func main() {
	//topic := "9999"
	//clusterName := "DefaultCluster"
	nameSrvAddr := []string{"10.21.7.5:30021"}
	//brokerAddr := "10.21.7.6:54874"

	testAdmin, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(nameSrvAddr)))
	if err != nil {
		fmt.Println(err.Error())
	}

	//create topic
	/*err = testAdmin.CreateTopic(
		context.Background(),
		admin.WithTopicCreate(topic),
		admin.WithBrokerAddrCreate(brokerAddr),
	)
	if err != nil {
		fmt.Println("Create topic error:", err.Error())
	}*/

	//deletetopic
	/*	err = testAdmin.DeleteTopic(
			context.Background(),
			admin.WithTopicDelete(topic),
			admin.WithBrokerAddrDelete(brokerAddr),
			//admin.WithNameSrvAddr(nameSrvAddr),
		)
		if err != nil {
			fmt.Println("Delete topic error:", err.Error())
		}*/

	/*opt,err:=testAdmin.GetBrokerClusterInfo(
	context.Background(),
	"10.21.7.5:30021",
	"10.21.7.6:54874",
	)*/
	/*	testAdmin.GetConsumeStats(
		context.Background(),
		"10.21.7.6:54874",
	)*/
	/*json,err:=testAdmin.TopicList(context.Background(),"10.21.7.5:30021")
	fmt.Println(json)*/
	/*fmt.Println(json.Get("topicList").GetIndex(0))*/
	err = testAdmin.WipeWritePerm(context.Background(), "n11", "10.21.7.5:30021")
	if err != nil {
		fmt.Println(err)
	}
	err = testAdmin.Close()
	if err != nil {
		fmt.Printf("Shutdown admin error: %s", err.Error())
	}
}
