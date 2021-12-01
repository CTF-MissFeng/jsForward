package api

import (
	"jsForward/app/service"
	"jsForward/library/logger"

	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/net/ghttp"
)

var Home = new(apiHome)

type apiHome struct{}

// RequestData 请求
func (a *apiHome) RequestData(r *ghttp.Request){
	json1,err := gjson.DecodeToJson(r.GetBodyString())
	if err != nil { // 非json数据
		respData := service.Home.RequestData(r.GetBodyString())
		r.Response.Write(respData)
	}else{
		jsonStr,err1 := json1.ToJsonString()
		if err1 != nil {
			logger.WebLog.Warningf("[-] %s", err1.Error())
			r.Response.WriteJson(r.GetBodyString())
		}
		respData := service.Home.RequestData(jsonStr)
		r.Response.WriteJson(respData)
	}
}

// ResponseData 返回
func (a *apiHome) ResponseData(r *ghttp.Request){
	_,err := gjson.DecodeToJson(r.GetBodyString())
	if err != nil {
		respData := service.Home.ResponseData(r.GetBodyString())
		r.Response.Write(respData)
	}else{
		respData := service.Home.ResponseData(r.GetBodyString())
		r.Response.WriteJson(respData)
	}
}