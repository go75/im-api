package utils

import (
	"encoding/json"
	"im-api/im-ws/entity"
	"im-api/im-ws/manager"

	"github.com/gorilla/websocket"
)

func Send(connID uint32, id uint16, option uint8, processId uint32, payload any) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	obj := entity.Obj {
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: string(data),
	}
	data, err = json.Marshal(obj)
	if err != nil {
		return err
	}
	err  = manager.Send(connID, data)
	return err
}

func SendStr(connID uint32, id uint16, option uint8, processId uint32, payload string) error {
	conn := manager.Get(connID)
	obj := entity.Obj {
		ID: id,
		Type: option,
		ProcessId: processId,
		Payload: payload,
	}
	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	err = conn.WriteMessage(websocket.TextMessage, data)
	if err != nil {
		return err
	}
	return nil
}