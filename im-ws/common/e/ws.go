package e

import "errors"

// websocket
var AlreadyExistConnErr = errors.New("websocket连接已存在")
var NotExistConnErr = errors.New("连接不存在")