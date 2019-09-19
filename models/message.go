package models

const (
	//TextElem 文本消息
	TextElem = "TIMTextElem"
	//LocationElem 地理位置消息
	LocationElem = "TIMLocationElem"
	//FaceElem 表情消息
	FaceElem = "TIMFaceElem"
	//CustomElem 自定义消息，当接收方为 iOS 系统且应用处在后台时，此消息类型可携带除文本以外的字段到 APNs。一条组合消息中只能包含一个 TIMCustomElem 自定义消息元素。
	CustomElem = "TIMCustomElem"
	//SoundElem 语音消息。（服务端集成 Rest API 不支持发送该类消息）
	SoundElem = "TIMSoundElem"
	//ImageElem 图像消息。（服务端集成 Rest API 不支持发送该类消息）
	ImageElem = "TIMImageElem"
	//FileElem 文件消息。（服务端集成 Rest API 不支持发送该类消息）
	FileElem = "TIMFileElem"
	//VideoFileElem is 视频消息。（服务端集成 Rest API 不支持发送该类消息）
	VideoFileElem = "TIMVideoFileElem"
)

type Message struct {
	CallbackCommand string
	FromAccount     string `From_Account`
	ToAccount       string `To_Account`
	MsgSeq          int
	MsgTime         int64
	MsgBody         []MsgBody
}

type MsgBody struct {
	MsgType    string
	MsgContent interface{}
}
