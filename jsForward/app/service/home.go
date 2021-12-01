package service

import (
	"jsForward/library/logger"
)

var Home = new(serviceHome)

type serviceHome struct{}

// RequestData 请求
func (s *serviceHome) RequestData(data string)string {
	logger.WebLog.Debugf("[+] 请求: %s", data)
	return data
}

// ResponseData 返回
func (s *serviceHome) ResponseData(data string)string {
	logger.WebLog.Debugf("[+] 返回: %s", data)
	return data
}