package main

import (
	"command_pattern/object"
)

func main() {

	tv := &object.Tv{}
	// 设置开命令，接收者为电视
	onCommand := &object.OnCommand{
		Device: tv,
	}
	// 设置关命令，接收者为电视
	offCommand := &object.OffCommand{
		Device: tv,
	}
	// 设置开按钮 对应 开命令
	onButton := &object.Button{
		Command: onCommand,
	}
	// 执行开按钮
	onButton.Press()
	// 设置关按钮 对应 关命令
	offButton := &object.Button{
		Command: offCommand,
	}
	// 执行关按钮
	offButton.Press()
}
