package gosms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// MapStrings 一纬map
type MapStrings map[string]string

// MapStringAny 二维以上map
type MapStringAny map[string]any

// HashMap hashmao
type HashMap struct {
	HashMap sync.Map
}

// MergeMapsString 多个map合并
func MergeMapsString(maps ...MapStrings) MapStrings {
	result := make(MapStrings)
	for _, m := range maps {
		for k, v := range m {
			result[k] = v
		}
	}
	return result
}

// ReplaceMapStringsForStringIndex
// 您手机注册的验证码为："您手机注册的验证码为：【变量】，如有问题请拨打客服电话：【变量】"
// new = MapStrings{"0": "111", "1": "222"}
// 您手机注册的验证码为：【111】，如有问题请拨打客服电话：【222】
func ReplaceMapStringsForStringIndex(s, old string, news MapStrings) string {
	var e []string
	if len(news) <= 0 {
		return s
	}
	ss := strings.Split(s, old)
	for i, v := range ss {
		if val, ok := news[strconv.Itoa(i)]; ok {
			if i < len(ss)-1 {
				e = append(e, fmt.Sprintf("%v%v", v, val))
			}
		} else {
			if len(v) > 0 {
				e = append(e, fmt.Sprintf("%v", v))
			}
		}
	}
	return strings.Join(e, "")
}

// ReplaceMapStrings 您手机注册的验证码为：【变量1】，如有问题请拨打客服电话：【变量2】
// new = MapStrings{"变量1"："111"，"变量2"::"222"}
// 您手机注册的验证码为：111，如有问题请拨打客服电话：222
func ReplaceMapStrings(s string, repl MapStrings) string {
	var r []string
	for k, v := range repl {
		r = append(r, k, v)
	}
	return strings.NewReplacer(r...).Replace(s)
}

// Set 添加值 force=true 覆盖原值
func (maps MapStrings) Set(k, v string) {
	maps.SetForce(k, v, true)
}

// SetForce Set 添加值 force=true 覆盖原值
func (maps MapStrings) SetForce(k, v string, force bool) {
	if maps.Exist(k) {
		if force {
			maps[k] = v
		}
		return
	}
	maps[k] = v
}

// Exist 键值是否存在
func (maps MapStrings) Exist(k string) bool {
	_, ok := maps[k]
	return ok
}

// Delete 删除指定到key值
func (maps MapStrings) Delete(k string) {
	delete(maps, k)
}

// GetDefault 获取值并携带默认值
func (maps MapStrings) GetDefault(k, d string) string {
	if v, ok := maps[k]; ok {
		return v
	}
	return d
}

// Keys keys
func (maps MapStrings) Keys() []string {
	var keys []string
	for k := range maps {
		keys = append(keys, k)
	}
	return keys
}

// Values values
func (maps MapStrings) Values() []string {
	var values []string
	for _, v := range maps {
		values = append(values, v)
	}
	return values
}

// Verify 验证出参数
func (maps MapStrings) Verify(mustKey []string) MapStrings {
	errs := MapStrings{}
	for _, key := range mustKey {
		if _, found := maps[key]; !found {
			errs.Set(key, key+" parameter is missing")
		}
	}
	return errs
}

// Implode values 转字符串
func (maps MapStrings) Implode(sep string) string {
	if len(maps) <= 0 {
		return ""
	}
	buf := new(bytes.Buffer)
	for _, s := range maps {
		buf.Write([]byte(s))
		buf.Write([]byte(sep))
	}
	str := buf.String()
	return str[:len(str)-1]
}

// ToUrlQuery 转url请求参数
func (maps MapStrings) ToUrlQuery(kSort bool) string {
	l := len(maps)
	if l <= 0 {
		return ""
	}
	var query string
	if kSort {
		var keys []string
		for k := range maps {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			query += k
			query += "="
			query += url.QueryEscape(maps[k])
			query += "&"
		}
	} else {
		for k, v := range maps {
			query += k
			query += "="
			query += url.QueryEscape(v)
			query += "&"
		}
	}
	return query[:len(query)-1]
}

// ToJson 转json字符串
func (maps MapStrings) ToJson() (string, error) {
	marshal, err := json.Marshal(maps)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

// JsonUnmarshal 解析json到实体上
func (maps MapStringAny) JsonUnmarshal(v any) error {
	jsonStr, err := maps.ToJson()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(jsonStr), v)
}

// ToJson  转json 字节流
func (maps MapStringAny) ToJson() (string, error) {
	marshal, err := json.Marshal(maps)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}

// Set 添加值 force=true 覆盖原值
func (maps MapStringAny) Set(k string, v any) {
	maps.SetForce(k, v, true)
}

// SetForce Set 添加值 force=true 覆盖原值
func (maps MapStringAny) SetForce(k string, v any, force bool) {
	if maps.Exist(k) {
		if force {
			maps[k] = v
		}
		return
	}
	maps[k] = v
}

// Exist 键值是否存在
func (maps MapStringAny) Exist(k string) bool {
	_, ok := maps[k]
	return ok
}

// Get 获取value 要设置默认值
func (maps MapStringAny) Get(k string, dv any) any {
	if vv, ok := maps[k]; ok {
		return vv
	}
	return dv
}

// Verify 验证出参数
func (maps MapStringAny) Verify(mustKey []string) MapStrings {
	errMaps := MapStrings{}
	for _, key := range mustKey {
		if _, found := maps[key]; !found {
			errMaps.Set(key, key+" parameter is missing")
		}
	}
	return errMaps
}

// ToIMessage 获取value 要设置默认值
func (maps MapStringAny) ToIMessage() IMessage {
	var message Message
	if typ, ok := maps.Get("type", 0).(int); ok {
		message.Type = typ
	}
	message.SignName = maps.Get("sign_name", "")
	message.Content = maps.Get("content", "")
	message.Template = maps.Get("template", "")
	message.Data = maps.Get("data", MapStrings{})
	if gateways, ok := maps.Get("gateways", nil).([]IGateway); ok {
		message.Gateways = gateways
	}
	return message.I()
}

// Set 设置
func (g *HashMap) Set(key string, value any) {
	g.HashMap.Store(key, value)
}

// Delete 删除
func (g *HashMap) Delete(key string) {
	g.HashMap.Delete(key)
}

// Get 获取
func (g *HashMap) Get(key string) (any, bool) {
	return g.HashMap.Load(key)
}

// Keys 键值
func (g *HashMap) Keys() []string {
	var s []string
	maps := g.ToMaps()
	for k := range maps {
		s = append(s, k)
	}
	return s
}

// Values value值
func (g *HashMap) Values() []any {
	var s []any
	maps := g.ToMaps()
	for _, v := range maps {
		s = append(s, v)
	}
	return s
}

// ToMaps 转maps
func (g *HashMap) ToMaps() map[string]any {
	maps := map[string]any{}
	g.HashMap.Range(func(key, value any) bool {
		if k, ok := key.(string); ok {
			maps[k] = value
		}
		return true
	})
	return maps
}
