package e

import "errors"

// token
var ExpiredTokenErr = errors.New("token过期")
var InvaildTokenErr = errors.New("token无效")