/**
 * @Project: mipush
 * @Author: qun
 * @Description:
 * @File: topic
 * @Version: 1.0.0
 * @Date: 2020/1/31 12:27
 *
 */
package topic

import (
    "github.com/houseme/mipush/builder"
    "github.com/houseme/mipush/miconst"
    "github.com/houseme/mipush/result"
    "github.com/houseme/mipush/util/http"
    "github.com/houseme/mipush/util/json"
)

//注册topic
func SubscribeTopicByRegId(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//批量注册主题
func BatchSubscribeTopicByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//注册topic
func SubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//批量注册主题
func BatchSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//通过regId取消注册topic
func UnSubscribeTopicByRegId(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//通过regId批量取消注册主题
func BatchUnSubscribeTopicByRegIds(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//通过Alias取消注册topic
func UnSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}

//通过Alias批量取消主题
func BatchUnSubscribeTopicByAlias(builder *builder.Builder, params builder.Params) (*result.Result, error) {
    params.MiUrl = miconst.MessageAllURL
    res, err := http.DoPost(builder, params)
    if err != nil {
        return nil, err
    }
    var result = &result.Result{}
    return json.ToJson(res, result)
}
