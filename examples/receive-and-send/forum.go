package main

import (
	"fmt"

	"github.com/miaoming3/botgo/dto"
	"github.com/miaoming3/botgo/event"
)

// ThreadEventHandler 论坛主贴事件
func ThreadEventHandler() event.ThreadEventHandler {
	return func(event *dto.WSPayload, data *dto.WSThreadData) error {
		fmt.Println(event, data)
		return nil
	}
}
