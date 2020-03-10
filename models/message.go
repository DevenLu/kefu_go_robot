package models

// Message struct
// BizType  消息类型  先以字符串形式声明吧
// 1 video 				视频
// 2 text  				文本
// 3 photo				图片
// 4 transfer   		转接
// 5 knowledge   		知识内容
// 6 handshake   		用户请求与握手
// 7 pong   		    对方正在输入中
// 8 end        		结束消息
// 9 timeout        	超时会话关闭
// 10 cancel        	撤回消息（暂时还没做预留先）
// 11 contacts        	用户列表（客服端）
// 12 welcome			欢迎语
// 13 system			系统消息
// 14 search_knowledge	检索关键词知识库消息
// 15 key				消息key
type Message struct {
	ID              int64  `json:"id"`               // 消息ID
	FromAccount     int64  `json:"from_account"`     // 发送人账号
	ToAccount       int64  `json:"to_account"`       // 接收人账号
	BizType         string `json:"biz_type"`         // 消息类型
	Version         string `json:"version"`          // 版本号，预留暂时无用
	Timestamp       int64  `json:"timestamp"`        // 服务端消息消息时间
	Key             int64  `json:"key"`              // key
	TransferAccount int64  `json:"transfer_account"` // 转接到客户的账号
	Platform        int64  `json:"platform"`         // 此消息来自哪个平台（即渠道）
	Payload         string `json:"payload"`          // 消息内容
	Read            int    `json:"read"`             // 是否已读消息0已读1未读
	Delete          int    `json:"delete"`           // 是否已删除消息0 ro 1 1已删除
}
