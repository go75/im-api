package global

import "im-api/im-ws/entity"

var Channel = make(chan *entity.Request, 1024)