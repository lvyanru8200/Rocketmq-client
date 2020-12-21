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

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/lvyanru8200/rocketmq-client/internal"
	"github.com/lvyanru8200/rocketmq-client/internal/remote"
	"github.com/lvyanru8200/rocketmq-client/primitive"
	"github.com/lvyanru8200/rocketmq-client/rlog"
	"github.com/tidwall/gjson"
	"strconv"
	"sync"
	"time"
)

type Admin interface {
	CreateTopic(ctx context.Context, opts ...OptionCreate) error
	DeleteTopic(ctx context.Context, opts ...OptionDelete) error
	//TODO
	GetTopicsByCluster(ctx context.Context, opts ...OptionTopicList) (*simplejson.Json, error)
	GetBrokerRuntimeInfo(ctx context.Context, nameserver string, broker string) (*simplejson.Json, error)
	GetConsumeStats(ctx context.Context, opts ...OptionGetConsumeStats) (*simplejson.Json, error)
	WipeWritePerm(ctx context.Context, opts ...OptionWipeWritePerm) error
	QueryTopicConsumeByWho(ctx context.Context, opts ...OptionQueryTopicConsume) (*simplejson.Json, error)
	UpdateAndCreateSubscriptionGroup(ctx context.Context, opts ...OptionUpdatesubGroup) error
	UpdateTopic(ctx context.Context, opts ...OptionCreate) error
	GetConsumerRunningInfo(ctx context.Context, opts ...OptionConsumerInfo) (gjson.Result, error)
	DeleteSubscriptionGroup(ctx context.Context, opts ...OptionDeleteSubGroup) error
	GetConsumerOffset(ctx context.Context, opts ...OptionConsumerOffset) (int64, error)
	GetRouteInfo(ctx context.Context, opts ...OptionGetRouteInfo) (*internal.TopicRouteData, error)
	GetConsumeStatsInBroker(ctx context.Context, opts ...OptionConsumeStatsInBroker) (gjson.Result, error)
	GetBrokerClusterInfo(ctx context.Context, nameserver string) (gjson.Result, error)
	UpdateConsumerOffset(ctx context.Context, opts ...OptionUpdateConsumerOffset) error
	Close() error
}

// TODO: move outdated context to ctx
type adminOptions struct {
	internal.ClientOptions
}

type AdminOption func(options *adminOptions)

func defaultAdminOptions() *adminOptions {
	opts := &adminOptions{
		ClientOptions: internal.DefaultClientOptions(),
	}
	opts.GroupName = "TOOLS_ADMIN"
	opts.InstanceName = time.Now().String()
	return opts
}

// WithResolver nameserver resolver to fetch nameserver addr
func WithResolver(resolver primitive.NsResolver) AdminOption {
	return func(options *adminOptions) {
		options.Resolver = resolver
	}
}

type admin struct {
	cli     internal.RMQClient
	namesrv internal.Namesrvs

	opts *adminOptions

	closeOnce sync.Once
}

// NewAdmin initialize admin
func NewAdmin(opts ...AdminOption) (Admin, error) {
	defaultOpts := defaultAdminOptions()
	for _, opt := range opts {
		opt(defaultOpts)
	}

	cli := internal.GetOrNewRocketMQClient(defaultOpts.ClientOptions, nil)
	namesrv, err := internal.NewNamesrv(defaultOpts.Resolver)
	if err != nil {
		return nil, err
	}
	//log.Printf("Client: %#v", namesrv.srvs)
	return &admin{
		cli:     cli,
		namesrv: namesrv,
		opts:    defaultOpts,
	}, nil
}

// CreateTopic create topic.
// TODO: another implementation like sarama, without brokerAddr as input
func (a *admin) CreateTopic(ctx context.Context, opts ...OptionCreate) error {
	cfg := defaultTopicConfigCreate()
	for _, apply := range opts {
		apply(&cfg)
	}

	request := &internal.CreateTopicRequestHeader{
		Topic:           cfg.Topic,
		DefaultTopic:    cfg.DefaultTopic,
		ReadQueueNums:   cfg.ReadQueueNums,
		WriteQueueNums:  cfg.WriteQueueNums,
		Perm:            cfg.Perm,
		TopicFilterType: cfg.TopicFilterType,
		TopicSysFlag:    cfg.TopicSysFlag,
		Order:           cfg.Order,
	}

	cmd := remote.NewRemotingCommand(internal.ReqCreateTopic, request, nil)
	_, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("create topic error", map[string]interface{}{
			rlog.LogKeyTopic:         cfg.Topic,
			rlog.LogKeyBroker:        cfg.BrokerAddr,
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("create topic success", map[string]interface{}{
			rlog.LogKeyTopic:  cfg.Topic,
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	return err
}

// DeleteTopicInBroker delete topic in broker.
func (a *admin) deleteTopicInBroker(ctx context.Context, topic string, brokerAddr string) (*remote.RemotingCommand, error) {
	request := &internal.DeleteTopicRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqDeleteTopicInBroker, request, nil)
	return a.cli.InvokeSync(ctx, brokerAddr, cmd, 5*time.Second)
}

// DeleteTopicInNameServer delete topic in nameserver.
func (a *admin) deleteTopicInNameServer(ctx context.Context, topic string, nameSrvAddr string) (*remote.RemotingCommand, error) {
	request := &internal.DeleteTopicRequestHeader{
		Topic: topic,
	}

	cmd := remote.NewRemotingCommand(internal.ReqDeleteTopicInNameSrv, request, nil)
	return a.cli.InvokeSync(ctx, nameSrvAddr, cmd, 5*time.Second)
}

// DeleteTopic delete topic in both broker and nameserver.
func (a *admin) DeleteTopic(ctx context.Context, opts ...OptionDelete) error {
	cfg := defaultTopicConfigDelete()
	for _, apply := range opts {
		apply(&cfg)
	}
	//delete topic in broker
	if cfg.BrokerAddr == "" {
		a.namesrv.UpdateTopicRouteInfo(cfg.Topic)
		cfg.BrokerAddr = a.namesrv.FindBrokerAddrByTopic(cfg.Topic)
	}

	if _, err := a.deleteTopicInBroker(ctx, cfg.Topic, cfg.BrokerAddr); err != nil {
		if err != nil {
			rlog.Error("delete topic in broker error", map[string]interface{}{
				rlog.LogKeyTopic:         cfg.Topic,
				rlog.LogKeyBroker:        cfg.BrokerAddr,
				rlog.LogKeyUnderlayError: err,
			})
		}
		return err
	}

	//delete topic in nameserver
	if len(cfg.NameSrvAddr) == 0 {
		a.namesrv.UpdateTopicRouteInfo(cfg.Topic)
		cfg.NameSrvAddr = a.namesrv.AddrList()
	}

	for _, nameSrvAddr := range cfg.NameSrvAddr {
		if _, err := a.deleteTopicInNameServer(ctx, cfg.Topic, nameSrvAddr); err != nil {
			if err != nil {
				rlog.Error("delete topic in name server error", map[string]interface{}{
					rlog.LogKeyTopic:         cfg.Topic,
					"nameServer":             nameSrvAddr,
					rlog.LogKeyUnderlayError: err,
				})
			}
			return err
		}
	}
	rlog.Info("delete topic success", map[string]interface{}{
		rlog.LogKeyTopic:  cfg.Topic,
		rlog.LogKeyBroker: cfg.BrokerAddr,
		"nameServer":      cfg.NameSrvAddr,
	})
	return nil
}

func (a *admin) GetTopicsByCluster(ctx context.Context, opts ...OptionTopicList) (*simplejson.Json, error) {
	cfg := defaultTopicList()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.GetTopicsByClusterRequestHeader{
		Cluster: cfg.Cluster,
	}
	cmd := remote.NewRemotingCommand(internal.GetTopicsByCluster, request, nil)
	opout, err := a.cli.InvokeSync(ctx, cfg.Nameserver, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("获取topic列表失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获取topic列表成功", map[string]interface{}{})
	}
	json, err := simplejson.NewJson(opout.Body)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (a *admin) GetBrokerRuntimeInfo(ctx context.Context, nameserver string, broker string) (*simplejson.Json, error) {
	request := &internal.ClusterListRequestHeader{
		NamesvrAddr: nameserver,
		BrokerAddr:  broker,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetBrokerRuntimeInfo, request, nil)
	opout, err := a.cli.InvokeSync(ctx, broker, cmd, 5*time.Second)
	json, err := simplejson.NewFromReader(bytes.NewBuffer(opout.Body))
	if err != nil {
		rlog.Error("获取BrokerInfo失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获取BrokerInfo成功", map[string]interface{}{})
	}
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (a *admin) GetConsumeStats(ctx context.Context, opts ...OptionGetConsumeStats) (*simplejson.Json, error) {
	cfg := defaultOptionGetConsumeStats()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.BrokerConsumeStatRequestHeader{
		ConsumerGroup: cfg.ConsumerGroup,
		Topic:         cfg.Topic,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetConsumeStats, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("获取ConsumeStats失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获取ConsumeStats成功", map[string]interface{}{})
	}
	fmt.Println(string(output.Body))
	json, err := simplejson.NewJson(output.Body)
	if err != nil {
		fmt.Println(err)
	}
	return json, nil
}

func (a *admin) WipeWritePerm(ctx context.Context, opts ...OptionWipeWritePerm) error {
	cfg := defaultOptionWipeWritePerm()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.WipeWritePermOfBrokerRequestHeader{
		Brokername: cfg.BrokerName,
		Nameserver: cfg.NamesrvAddr,
	}
	cmd := remote.NewRemotingCommand(internal.ReqWipeWritePermOfBroker, request, nil)
	_, err := a.cli.InvokeSync(ctx, cfg.NamesrvAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("禁止写失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("禁止写成功", map[string]interface{}{
			rlog.LogKeyBroker: request.Brokername,
		})
	}
	if err != nil {
		return err
	}
	return nil
}

func (a *admin) QueryTopicConsumeByWho(ctx context.Context, opts ...OptionQueryTopicConsume) (*simplejson.Json, error) {
	cfg := defaultQueryTopicConsume()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.QueryTopicConsumeByWhoRequestHeader{
		BrokerAddr: cfg.Brokeraddr,
		Topic:      cfg.Topic,
	}
	cmd := remote.NewRemotingCommand(internal.ReqQueryTopicConsumeByWho, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.Brokeraddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("获取topic对应消费组失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获取topic对应消费组成功", map[string]interface{}{
			rlog.LogKeyBroker: request.Topic,
		})
	}
	json, err := simplejson.NewJson(output.Body)
	if err != nil {
		return nil, err
	}
	return json, nil
}

func (a *admin) UpdateAndCreateSubscriptionGroup(ctx context.Context, opts ...OptionUpdatesubGroup) error {
	cfg := defaultOptionUpdatesubGroup()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.UpdateAndCreateSubscriptionGroupRequestHeader{
		GroupName: cfg.GroupName,
	}
	cmd := remote.NewRemotingCommand(internal.ReqUpdateAndCreateSubscriptionGroup, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("迁移消费组失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("迁移消费组成功", map[string]interface{}{
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	fmt.Println(output.String())
	return err
}

func (a *admin) UpdateTopic(ctx context.Context, opts ...OptionCreate) error {
	cfg := updateTopicConfig()
	for _, apply := range opts {
		apply(&cfg)
	}

	request := &internal.UpdateTopicRequestHeader{
		Topic:           cfg.Topic,
		ClusterName:     cfg.ClusterName,
		TopicFilterType: cfg.TopicFilterType,
	}

	cmd := remote.NewRemotingCommand(internal.ReqCreateTopic, request, nil)
	_, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("更新topic失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("更新topic成功", map[string]interface{}{
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	return err
}

func (a *admin) GetConsumerRunningInfo(ctx context.Context, opts ...OptionConsumerInfo) (gjson.Result, error) {
	cfg := defaultOptionConsumerInfo()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.GetConsumerRunningInfoRequestHeader{
		ConsumerGroup: cfg.ConsumerGroup,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetConsumerRunningInfo, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("获取consumer信息失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获取consumer信息成功", map[string]interface{}{
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	return gjson.Parse(string(output.Body)), err
}

func (a *admin) DeleteSubscriptionGroup(ctx context.Context, opts ...OptionDeleteSubGroup) error {
	cfg := defaultDeleteSubGroup()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.DeleteSubscriptionGroupRequestHeader{
		GroupName:  cfg.GroupName,
		BrokerAddr: cfg.BrokerAddr,
	}
	cmd := remote.NewRemotingCommand(internal.ReqDeleteSubscriptionGroup, request, nil)
	_, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 5*time.Second)
	if err != nil {
		rlog.Error("删除subGroup失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("删除subGroup成功", map[string]interface{}{
			rlog.LogKeyBroker: cfg.BrokerAddr,
		})
	}
	return err
}

func (a *admin) GetConsumerOffset(ctx context.Context, opts ...OptionConsumerOffset) (int64, error) {
	cfg := defaultConsumerOffset()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.QueryConsumerOffsetRequestHeader{
		ConsumerGroup: cfg.ConsumerGroup,
		Topic:         cfg.Topic,
		QueueId:       cfg.QueueId,
	}
	cmd := remote.NewRemotingCommand(internal.ReqQueryConsumerOffset, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 3*time.Second)
	if err != nil {
		rlog.Error("获得consumeroffset失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获得consumeroffset成功", map[string]interface{}{})
	}
	off, err := strconv.ParseInt(output.ExtFields["offset"], 10, 64)
	if err != nil {
		rlog.Error("获取offset失败", map[string]interface{}{
			rlog.LogKeyTopic:         cfg.Topic,
			rlog.LogKeyConsumerGroup: cfg.ConsumerGroup,
		})
	}
	return off, nil
}

func (a *admin) GetRouteInfo(ctx context.Context, opts ...OptionGetRouteInfo) (*internal.TopicRouteData, error) {
	cfg := defaultOptionGetRouteInfo()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.GetRouteInfoRequestHeader{
		Topic: cfg.Topic,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetRouteInfoByTopic, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.NamesrvAddr, cmd, 3*time.Second)
	if err != nil {
		rlog.Error("获得topicRoute失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获得topicRoute成功", map[string]interface{}{})
	}
	routeData := &internal.TopicRouteData{}
	err = routeData.Decode(string(output.Body))
	if err != nil {
		rlog.Warning("decode TopicRouteData error: %s", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
			"topic":                  cfg.Topic,
		})
	}
	return routeData, nil
}

func (a *admin) GetConsumeStatsInBroker(ctx context.Context, opts ...OptionConsumeStatsInBroker) (gjson.Result, error) {
	cfg := defaultOptionConsumeStatsInBroker()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.GetConsumeStatsInBrokerHeader{
		IsOrder: cfg.IsOrder,
	}
	cmd := remote.NewRemotingCommand(internal.ReqGetBrokerConsumeStats, request, nil)
	output, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 3*time.Second)
	if err != nil {
		rlog.Error("获得ConsumeStats失败", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("获得ConsumeStats成功", map[string]interface{}{})
	}
	return gjson.Parse(string(output.Body)), err
}

func (a *admin) GetBrokerClusterInfo(ctx context.Context, nameserver string) (gjson.Result, error) {
	request := &internal.GetBrokerClsuterInfoRequetsHeader{}
	cmd := remote.NewRemotingCommand(internal.ReqGetBrokerClusterInfo, request, nil)
	output, err := a.cli.InvokeSync(ctx, nameserver, cmd, 3*time.Second)
	if err != nil {
		rlog.Error("Fail Get BrokerClusterInfo", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("Successful Get BrokerClusterInfo", map[string]interface{}{})
	}
	return gjson.Parse(string(output.Body)), err
}

func (a *admin) UpdateConsumerOffset(ctx context.Context, opts ...OptionUpdateConsumerOffset) error {
	cfg := defaultOptionUpdateConsumerOffset()
	for _, apply := range opts {
		apply(&cfg)
	}
	request := &internal.UpdateConsumerOffsetRequestHeader{
		ConsumerGroup: cfg.ConsumerGroup,
		Topic:         cfg.Topic,
		CommitOffset:  cfg.CommitOffset,
		QueueId:       cfg.QueueId,
	}
	cmd := remote.NewRemotingCommand(internal.ReqUpdateConsumerOffset, request, nil)
	_, err := a.cli.InvokeSync(ctx, cfg.BrokerAddr, cmd, 3*time.Second)
	if err != nil {
		rlog.Error("Fail  Update ConsumerOffset", map[string]interface{}{
			rlog.LogKeyUnderlayError: err,
		})
	} else {
		rlog.Info("Successful Update ConsumerOffset", map[string]interface{}{})
	}
	return err
}

func (a *admin) Close() error {
	a.closeOnce.Do(func() {
		a.cli.Shutdown()
	})
	return nil
}
