package push

type Notification struct {
	AiOpportunity bool                 `json:"ai_opportunity,omitempty"` // 如需采用“智能时机”策略下发通知，必须指定该字段。
	Alert         string               `json:"alert,omitempty"`          // 通知的内容在各个平台上，都可能只有这一个最基本的属性 "alert"。
	Android       *AndroidNotification `json:"android,omitempty"`        // Android通知
}

type AndroidNotification struct {
	Alert             string      `json:"alert"`                        // 通知内容
	Title             string      `json:"title,omitempty"`              // 通知标题
	BuilderID         int         `json:"builder_id,omitempty"`         // 通知栏样式 ID
	ChannelId         string      `json:"channel_id,omitempty"`         // android通知channel_id
	Priority          int         `json:"priority,omitempty"`           // 通知栏展示优先级, 默认为 0，范围为 -2～2。
	Category          string      `json:"category,omitempty"`           // 通知栏条目过滤或排序
	Style             int         `json:"style,omitempty"`              // 通知栏样式类型, 默认为 0，还有 1，2，3 可选
	AlertType         int         `json:"alert_type,omitempty"`         // 通知提醒方式, 可选范围为 -1～7
	BigText           string      `json:"big_text,omitempty"`           // 大文本通知栏样式, 当 style = 1 时可用，内容会被通知栏以大文本的形式展示出来
	Inbox             interface{} `json:"inbox,omitempty"`              // 文本条目通知栏样式, 当 style = 2 时可用， json 的每个 key 对应的 value 会被当作文本条目逐条展示
	BigPicPath        string      `json:"big_pic_path,omitempty"`       // 大图片通知栏样式, 当 style = 3 时可用，可以是网络图片 url，或本地图片的 path
	Extras            interface{} `json:"extras,omitempty"`             // 扩展字段
	LargeIcon         string      `json:"large_icon,omitempty"`         // 通知栏大图标
	SmallIconUri      string      `json:"small_icon_uri,omitempty"`     // 通知栏小图标
	Intent            interface{} `json:"intent,omitempty"`             // 指定跳转页面
	UriActivity       string      `json:"uri_activity,omitempty"`       // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 activity 节点的 “android:name”属性值; 适配华为、小米、vivo厂商通道跳转；
	UriAction         string      `json:"uri_action,omitempty"`         // 指定跳转页面, 该字段用于指定开发者想要打开的 activity，值为 "activity"-"intent-filter"-"action" 节点的 "android:name" 属性值; 适配 oppo、fcm跳转；
	BadgeAddNum       int         `json:"badge_add_num,omitempty"`      // 角标数字，取值范围1-99
	BadgeClass        string      `json:"badge_class,omitempty"`        // 桌面图标对应的应用入口Activity类， 配合badge_add_num使用，二者需要共存，缺少其一不可；
	Sound             string      `json:"sound,omitempty"`              // 填写Android工程中/res/raw/路径下铃声文件名称，无需文件名后缀
	ShowBeginTime     string      `json:"show_begin_time,omitempty"`    //定时展示开始时间（yyyy-MM-dd HH:mm:ss）
	ShowEndTime       string      `json:"show_end_time,omitempty"`      //定时展示结束时间（yyyy-MM-dd HH:mm:ss）
	DisplayForeground string      `json:"display_foreground,omitempty"` //APP在前台，通知是否展示, 值为 "1" 时，APP 在前台会弹出通知栏消息；值为 "0" 时，APP 在前台不会弹出通知栏消息。
}