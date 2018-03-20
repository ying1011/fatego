// Copyright 2014 mqantserver Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"github.com/liangdas/armyant/work"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"fmt"
)

func main()  {
	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("ws://127.0.0.1:3653")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("连接断开", err.Error())
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("连接成功")
	})
	err := this.Connect(opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	msg, err := this.Request("Game/HD_Game", []byte(`{"say":"hello world"}`))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(msg.Payload()))
}