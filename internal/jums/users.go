package jums

// SetID 设置用户id
func (u *Users) SetID(userID uint) *Users {
	u.UserID = userID
	return u
}

// Add 添加数据
func (u *Users) Add(types ...*userType) *Users {
	add := map[string]any{}
	for _, data := range types {
		for k, v := range *data {
			add[k] = v
		}
	}
	u.Data["add"] = add
	return u
}

// Set 设置数据
func (u *Users) Set(types ...*userType) *Users {
	set := map[string]any{}
	for _, data := range types {
		for k, v := range *data {
			u.Data[k] = v
		}
	}
	u.Data["set"] = set
	return u
}

// Del 删除数据
func (u *Users) Del(types ...*userType) *Users {
	del := map[string]any{}
	for _, data := range types {
		for k, v := range *data {
			u.Data[k] = v
		}
	}
	u.Data["del"] = del
	return u
}

// Send 发送消息内容
func (u *Users) Send() error {
	return Request("v1/sent/user/opt", u.Key, u.Secret, u.Data)
}

type userType map[string]any

// UserTag 渠道标签
func UserTag(channelKey string, data ...string) *userType {
	return &userType{
		"tag": map[string]any{
			channelKey: data,
		},
	}
}

// UserPhone 渠道手机号
func UserPhone(channelKey string, data ...string) *userType {
	return &userType{
		"phone": map[string]any{
			channelKey: data,
		},
	}
}

// UserEmail 渠道邮箱
func UserEmail(channelKey string, data ...string) *userType {
	return &userType{
		"email": map[string]any{
			channelKey: data,
		},
	}
}

// UserApp app类型
func UserApp(instance string, data ...string) *userType {
	return &userType{
		"app": map[string]any{
			instance: data,
		},
	}
}

// UserWechatOa 微信公众号
func UserWechatOa(instance string, data string) *userType {
	return &userType{
		"wechatoa": map[string]any{
			instance: data,
		},
	}
}

// UserWechatMp 微信小程序
func UserWechatMp(instance string, data string) *userType {
	return &userType{
		"wechatmp": map[string]any{
			instance: data,
		},
	}
}

// UserAlipayLife 支付宝生活号
func UserAlipayLife(instance string, data string) *userType {
	return &userType{
		"alipaylife": map[string]any{
			instance: data,
		},
	}
}

// UserDingtalkCC 钉钉
func UserDingtalkCC(data string) *userType {
	return &userType{
		"dingtalkcc": data,
	}
}

// UserWechatWk 企业微星
func UserWechatWk(data string) *userType {
	return &userType{
		"wechatwk": data,
	}
}
