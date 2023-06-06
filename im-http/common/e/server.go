package e

import "errors"

var DialServerErr = errors.New("服务调用失败")
var BadRequestErr = errors.New("错误请求")