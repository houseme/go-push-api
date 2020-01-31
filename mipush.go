package mipush

import (
    "github.com/houseme/mipush/api"
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/result"
)

func sendMessage(builder *builder.Builder, params builder.CommonParams) (*result.Result, error)  {
    return api.SendMessageByRegIds(builder,params);
}