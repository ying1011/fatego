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

	msg, err := this.Request("Login/HD_Login", []byte(fmt.Sprintf(`{"say":"hello world ","userId":"14"}`)))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(msg.Payload()))

	isRun := true
	num := 0
	var input string
	for isRun  {
		num++


		fmt.Scan(&input)
		fmt.Println("c = ",input, num)

		msg, err := this.Request("Game/HD_Game", []byte(fmt.Sprintf(`{"say":"hello world %s"}`,input)))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(msg.Payload()))

		if(input == "quit"){
			isRun = false
		}

	}


}