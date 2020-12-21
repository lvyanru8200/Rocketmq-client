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

package admin

func updateTopicConfig() TopicConfigCreate {
	opts := TopicConfigCreate{}
	return opts
}
func defaultTopicConfigCreate() TopicConfigCreate {
	opts := TopicConfigCreate{
		DefaultTopic:    "defaultTopic",
		ReadQueueNums:   8,
		WriteQueueNums:  8,
		Perm:            6,
		TopicFilterType: "SINGLE_TAG",
		TopicSysFlag:    0,
		Order:           false,
	}
	return opts
}

type TopicConfigCreate struct {
	Topic           string
	BrokerAddr      string
	DefaultTopic    string
	ReadQueueNums   int
	WriteQueueNums  int
	Perm            int
	TopicFilterType string
	TopicSysFlag    int
	Order           bool
	ClusterName     string
	NamesrvAddr     string
}

type OptionCreate func(*TopicConfigCreate)

func WithTopicClusterName(clustername string) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.ClusterName = clustername
	}
}

func WithTopicCreate(Topic string) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.Topic = Topic
	}
}

func WithBrokerAddrCreate(BrokerAddr string) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.BrokerAddr = BrokerAddr
	}
}

func WithReadQueueNums(ReadQueueNums int) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.ReadQueueNums = ReadQueueNums
	}
}

func WithWriteQueueNums(WriteQueueNums int) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.WriteQueueNums = WriteQueueNums
	}
}

func WithPerm(Perm int) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.Perm = Perm
	}
}

func WithTopicFilterType(TopicFilterType string) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.TopicFilterType = TopicFilterType
	}
}

func WithTopicSysFlag(TopicSysFlag int) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.TopicSysFlag = TopicSysFlag
	}
}

func WithOrder(Order bool) OptionCreate {
	return func(opts *TopicConfigCreate) {
		opts.Order = Order
	}
}

func defaultTopicConfigDelete() TopicConfigDelete {
	opts := TopicConfigDelete{}
	return opts
}

type TopicConfigDelete struct {
	Topic       string
	ClusterName string
	NameSrvAddr []string
	BrokerAddr  string
}

type OptionDelete func(*TopicConfigDelete)

func WithTopicDelete(Topic string) OptionDelete {
	return func(opts *TopicConfigDelete) {
		opts.Topic = Topic
	}
}

func WithBrokerAddrDelete(BrokerAddr string) OptionDelete {
	return func(opts *TopicConfigDelete) {
		opts.BrokerAddr = BrokerAddr
	}
}

func WithClusterName(ClusterName string) OptionDelete {
	return func(opts *TopicConfigDelete) {
		opts.ClusterName = ClusterName
	}
}

func WithNameSrvAddr(NameSrvAddr []string) OptionDelete {
	return func(opts *TopicConfigDelete) {
		opts.NameSrvAddr = NameSrvAddr
	}
}

type OptionTopicList func(*OptionTopicListConfig)

type OptionTopicListConfig struct {
	Cluster    string
	Nameserver string
}

func defaultTopicList() OptionTopicListConfig {
	opts := OptionTopicListConfig{}
	return opts
}

func WithClusterTopicList(cluster string) OptionTopicList {
	return func(opts *OptionTopicListConfig) {
		opts.Cluster = cluster
	}
}

func WithNameserver(nameserver string) OptionTopicList {
	return func(opts *OptionTopicListConfig) {
		opts.Nameserver = nameserver
	}
}

type OptionQueryTopicConsume func(config *QueryTopicConsumeConfig)

type QueryTopicConsumeConfig struct {
	Brokeraddr string
	Topic      string
}

func defaultQueryTopicConsume() QueryTopicConsumeConfig {
	opts := QueryTopicConsumeConfig{}
	return opts
}
func WithBrokeraddr(broker string) OptionQueryTopicConsume {
	return func(opts *QueryTopicConsumeConfig) {
		opts.Brokeraddr = broker
	}
}

func WithTopic(topic string) OptionQueryTopicConsume {
	return func(opts *QueryTopicConsumeConfig) {
		opts.Topic = topic
	}
}

type OptionClusterList func(*ClusterListConfig)

type ClusterListConfig struct {
	NamesvrAddr string
}

func defaultClusterList() ClusterListConfig {
	opts := ClusterListConfig{}
	return opts
}

type OptionUpdatesubGroup func(config *OptionUpdatesubGroupConfig)

type OptionUpdatesubGroupConfig struct {
	BrokerAddr             string
	ClusterName            string
	GroupName              string
	BrokerId               string
	NamesrvAddr            string
	RetryQueueNums         int
	RetryMaxTimes          int64
	ConsumeEnable          bool
	ConsumeFromMinEnable   bool
	ConsumeBroadcastEnable bool
}

func defaultOptionUpdatesubGroup() OptionUpdatesubGroupConfig {
	opts := OptionUpdatesubGroupConfig{}
	return opts
}
func WithGroupName(groupname string) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.GroupName = groupname
	}
}

func WithGroupClusterName(clustername string) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.ClusterName = clustername
	}
}

func WithConsumeFromMin(t bool) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.ConsumeFromMinEnable = t
	}
}

func WithConsumeBroadcast(t bool) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.ConsumeBroadcastEnable = t
	}
}

func WithNamesrvAddr(nameserver string) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.NamesrvAddr = nameserver
	}
}

func WithBrokerAddr(broker string) OptionUpdatesubGroup {
	return func(opts *OptionUpdatesubGroupConfig) {
		opts.BrokerAddr = broker
	}
}

type OptionConsumerInfo func(config *OptionConsumerInfoConfig)

type OptionConsumerInfoConfig struct {
	ConsumerGroup string
	BrokerAddr    string
}

func defaultOptionConsumerInfo() OptionConsumerInfoConfig {
	opts := OptionConsumerInfoConfig{}
	return opts
}

func WithConsumerGroup(group string) OptionConsumerInfo {
	return func(opts *OptionConsumerInfoConfig) {
		opts.ConsumerGroup = group
	}
}

func WithConsumerBrokerAddr(broker string) OptionConsumerInfo {
	return func(opts *OptionConsumerInfoConfig) {
		opts.BrokerAddr = broker
	}
}

type OptionDeleteSubGroup func(*OptionDeleteSubGroupConfig)

type OptionDeleteSubGroupConfig struct {
	GroupName   string
	ClusterName string
	BrokerAddr  string
}

func defaultDeleteSubGroup() OptionDeleteSubGroupConfig {
	opts := OptionDeleteSubGroupConfig{}
	return opts
}

func WithDeleteSubGroup(group string) OptionDeleteSubGroup {
	return func(opts *OptionDeleteSubGroupConfig) {
		opts.GroupName = group
	}
}

func WithDeleteCluster(cluster string) OptionDeleteSubGroup {
	return func(opts *OptionDeleteSubGroupConfig) {
		opts.ClusterName = cluster
	}
}

func WithDeletebrokeraddr(broker string) OptionDeleteSubGroup {
	return func(opts *OptionDeleteSubGroupConfig) {
		opts.BrokerAddr = broker
	}
}

type OptionConsumerOffset func(*OptionConsumerOffsetConfig)

type OptionConsumerOffsetConfig struct {
	ConsumerGroup string
	Topic         string
	QueueId       int
	BrokerAddr    string
}

func defaultConsumerOffset() OptionConsumerOffsetConfig {
	opts := OptionConsumerOffsetConfig{}
	return opts
}

func WithConsumeroffsetBrokerAddr(broker string) OptionConsumerOffset {
	return func(opts *OptionConsumerOffsetConfig) {
		opts.BrokerAddr = broker
	}
}

func WithConsumerOffsetGroup(group string) OptionConsumerOffset {
	return func(opts *OptionConsumerOffsetConfig) {
		opts.ConsumerGroup = group
	}
}

func WithConsumerOffsetTopic(topic string) OptionConsumerOffset {
	return func(opts *OptionConsumerOffsetConfig) {
		opts.Topic = topic
	}
}

func WithQueId(queid int) OptionConsumerOffset {
	return func(opts *OptionConsumerOffsetConfig) {
		opts.QueueId = queid
	}
}

type OptionQueryConsumeQueueConfig struct {
	Topic         string
	QueueId       string
	Index         int
	Count         int
	ConsumerGroup string
}

type OptionQueryConsumeQueue func(*OptionQueryConsumeQueueConfig)

func defaultOptionQueryConsumeQueue() OptionQueryConsumeQueueConfig {
	opts := OptionQueryConsumeQueueConfig{
		Index:   8,
		QueueId: "1",
	}
	return opts
}

func WithQueryTopic(topic string) OptionQueryConsumeQueue {
	return func(opts *OptionQueryConsumeQueueConfig) {
		opts.Topic = topic
	}
}

func WithQueryConsumerGroup(group string) OptionQueryConsumeQueue {
	return func(opts *OptionQueryConsumeQueueConfig) {
		opts.ConsumerGroup = group
	}
}

type OptionWipeWritePermConfig struct {
	BrokerName  string
	NamesrvAddr string
}

type OptionWipeWritePerm func(*OptionWipeWritePermConfig)

func defaultOptionWipeWritePerm() OptionWipeWritePermConfig {
	opts := OptionWipeWritePermConfig{}
	return opts
}

func WithBrokerName(brokername string) OptionWipeWritePerm {
	return func(opts *OptionWipeWritePermConfig) {
		opts.BrokerName = brokername
	}
}

func WithWipeWritePermNamesvr(nameserver string) OptionWipeWritePerm {
	return func(opts *OptionWipeWritePermConfig) {
		opts.NamesrvAddr = nameserver
	}
}

type OptionGetRouteInfoConfig struct {
	Topic       string
	NamesrvAddr string
}

type OptionGetRouteInfo func(*OptionGetRouteInfoConfig)

func defaultOptionGetRouteInfo() OptionGetRouteInfoConfig {
	opts := OptionGetRouteInfoConfig{}
	return opts
}

func WithTopicRoute(topic string) OptionGetRouteInfo {
	return func(opts *OptionGetRouteInfoConfig) {
		opts.Topic = topic
	}
}

func WithNamesrvAddrRoute(nameserver string) OptionGetRouteInfo {
	return func(opts *OptionGetRouteInfoConfig) {
		opts.NamesrvAddr = nameserver
	}
}

type OptionGetConsumeStatsConfig struct {
	ConsumerGroup string
	Topic         string
	BrokerAddr    string
}

type OptionGetConsumeStats func(*OptionGetConsumeStatsConfig)

func defaultOptionGetConsumeStats() OptionGetConsumeStatsConfig {
	opts := OptionGetConsumeStatsConfig{}
	return opts
}

func WithBrokerAddrGetConsume(broker string) OptionGetConsumeStats {
	return func(opts *OptionGetConsumeStatsConfig) {
		opts.BrokerAddr = broker
	}
}

func WithConsumerGroupGetConsume(group string) OptionGetConsumeStats {
	return func(opts *OptionGetConsumeStatsConfig) {
		opts.ConsumerGroup = group
	}
}

func WithTopicGroupGetConsume(topic string) OptionGetConsumeStats {
	return func(opts *OptionGetConsumeStatsConfig) {
		opts.Topic = topic
	}
}

type OptionConsumeStatsInBrokerConfig struct {
	IsOrder    bool
	BrokerAddr string
}

type OptionConsumeStatsInBroker func(config *OptionConsumeStatsInBrokerConfig)

func defaultOptionConsumeStatsInBroker() OptionConsumeStatsInBrokerConfig {
	opts := OptionConsumeStatsInBrokerConfig{}
	return opts
}

func WithConsumeStatsInBrokerBroker(broker string) OptionConsumeStatsInBroker {
	return func(opts *OptionConsumeStatsInBrokerConfig) {
		opts.BrokerAddr = broker
	}
}

func WithConsumerStatsInIsOrder(b bool) OptionConsumeStatsInBroker {
	return func(opts *OptionConsumeStatsInBrokerConfig) {
		opts.IsOrder = b
	}
}

type OptionUpdateConsumerOffsetConfig struct {
	ConsumerGroup string
	Topic         string
	QueueId       int
	CommitOffset  int64
	BrokerAddr    string
}

type OptionUpdateConsumerOffset func(config *OptionUpdateConsumerOffsetConfig)

func defaultOptionUpdateConsumerOffset() OptionUpdateConsumerOffsetConfig {
	opts := OptionUpdateConsumerOffsetConfig{}
	return opts
}

func WithUpdateoffsetConsumerGroup(group string) OptionUpdateConsumerOffset {
	return func(opts *OptionUpdateConsumerOffsetConfig) {
		opts.ConsumerGroup = group
	}
}

func WithUpdateoffsetTopic(topic string) OptionUpdateConsumerOffset {
	return func(opts *OptionUpdateConsumerOffsetConfig) {
		opts.Topic = topic
	}
}

func WithUpdateoffsetQueueId(queueid int) OptionUpdateConsumerOffset {
	return func(opts *OptionUpdateConsumerOffsetConfig) {
		opts.QueueId = queueid
	}
}

func WithUpdateoffsetCommitoffset(commitoffset int64) OptionUpdateConsumerOffset {
	return func(opts *OptionUpdateConsumerOffsetConfig) {
		opts.CommitOffset = commitoffset
	}
}
func WithUpdateoffsetBrokerAddr(broker string) OptionUpdateConsumerOffset {
	return func(opts *OptionUpdateConsumerOffsetConfig) {
		opts.BrokerAddr = broker
	}
}
