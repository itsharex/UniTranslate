package translate

// ChatGptTranslateMode ChatGPT 支持
const ChatGptTranslateMode = "ChatGPT"

type ChatGptConfigType struct {
	Key string `json:"key"`
}

// XunFeiTranslateMode 讯飞常用版本
const XunFeiTranslateMode = "XunFei"

// XunFeiNiuTranslateMode 讯飞新版
const XunFeiNiuTranslateMode = "XunFeiNiu"

type XunFeiConfigType struct {
	AppId  string `json:"appId"`
	Secret string `json:"secret"`
	ApiKey string `json:"apiKey"`
}

const TencentTranslateMode = "Tencent"

type TencentConfigType struct {
	Url       string `json:"url"`
	SecretId  string `json:"secretId"`
	SecretKey string `json:"secretKey"`
	Region    string `json:"region"`
}
