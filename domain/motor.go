package domain

type MotorMode int

const (
	Success MotorMode = iota
	NoSuchDev
	ArgErr
	Fail
)

type MotorDomain struct { //
	Protocol int `json:"protocol"`
	Id       int `json:"protidocol"`
	PowerOn  int `json:"powerOn"`
	//位置角度量(一般舵机)
	NumAngel int `json:"numAngel"`
	DenAngel int `json:"denAngel"`
	MaxAngel int `json:"maxAngel"`
	// 编码器量;
	Encode int `json:"encode"`
	//编码器速度
	SpEncode int `json:"spEncode"`
	//占空比
	PwmNum int `json:"pwmNum"`
	PwmDen int `json:"pwmDen"`

	//角速度
	SpNumAngel int       `json:"spNumAngel"`
	SpDenAngel int       `json:"spDenAngel"`
	Mode       MotorMode `json:"mode"`
}
type MotorDomainReply struct {
	Message string `json:"Message"`
}
