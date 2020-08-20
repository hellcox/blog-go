// 网页三要素
package global

type Tkd struct {
	T, K, D string
}

type AllTkd struct {
	Home          Tkd // 首页
	Tools         Tkd // 工具列表
	ToolMd5       Tkd // MD5
	ToolIp        Tkd // IP
	ToolTime      Tkd // 时间
	ToolBase64    Tkd // base64 Encode/Decode
	ToolUrlEncode Tkd // urlEncode/urlDecode
	ToolColor     Tkd // 网页颜色
}

func initTKD() {
	TKD = AllTkd{
		Home: Tkd{
			T: "小豆豆博客-纪录心动时刻！",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		Tools: Tkd{
			T: "工具集合-小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolMd5: Tkd{
			T: "MD5加解密-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolIp: Tkd{
			T: "IP地址-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolTime: Tkd{
			T: "时间转换-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolBase64: Tkd{
			T: "BASE64加解密-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolUrlEncode: Tkd{
			T: "URL Encode/Decode-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
		ToolColor: Tkd{
			T: "颜色表-工具|小豆豆博客",
			K: "小豆豆博客,博客,个人博客,小豆豆,小小豆,豆豆,原创博客,互联网,自媒体",
			D: "小豆豆个人博客，是一个记录自己生活点滴、互联网技术的原创独立博客(www.iuxiao.com)。",
		},
	}
}
