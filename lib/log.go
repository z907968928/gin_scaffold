package lib

import (
	"fmt"
	logger "github.com/e421083458/gin_scaffold/core/logger"
	"strings"
)

// 通用DLTag常量定义
const (
	DLTagUndefind     = "_undef"
	DLTagMySqlSuccess = "_com_mysql_success"
	DLTagMySqlFailed  = "_com_mysql_failure"

	DLTagRedisSuccess = "_com_redis_success"
	DLTagRedisFailed  = "_com_redis_failure"

	DLTagThriftFailed  = "_com_thrift_failure"
	DLTagThriftSuccess = "_com_thrift_success"

	DLTagHTTPSuccess = "_com_http_success"
	DLTagHTTPFailed  = "_com_http_failure"

	DLTagTCPFailed = "_com_tcp_failure"

	DLTagRequestIn  = "_com_request_in"
	DLTagRequestOut = "_com_request_out"
)

const (
	_dlTag          = "dltag"
	_traceId        = "traceid"
	_spanId         = "spanid"
	_childSpanId    = "cspanid"
	_dlTagBizPrefix = "_com_"
	_dlTagBizUndef  = "_com_undef"
)

var Log *Logger

type Logger struct{}

func (l *Logger) TagInfo(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	logger.Info(parseParams(m))
}

func (l *Logger) TagWarn(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	logger.Warn(parseParams(m))
}

func (l *Logger) TagError(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	logger.Error(parseParams(m))
}

func (l *Logger) TagTrace(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	logger.Trace(parseParams(m))
}

func (l *Logger) TagDebug(trace *TraceContext, dltag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dltag)
	m[_traceId] = trace.TraceId
	m[_childSpanId] = trace.CSpanId
	m[_spanId] = trace.SpanId
	logger.Debug(parseParams(m))
}

// 校验dltag合法性
func checkDLTag(dltag string) string {
	if strings.HasPrefix(dltag, _dlTagBizPrefix) {
		return dltag
	}

	if dltag == DLTagUndefind {
		return dltag
	}
	return dltag
}

//map格式化为string
func parseParams(m map[string]interface{}) string {
	var dltag = "_undef"
	if _dltag, _have := m["dltag"]; _have {
		if __val, __ok := _dltag.(string); __ok {
			dltag = __val
			delete(m, "dltag")
		}
	}
	for _key, _val := range m {
		dltag = dltag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	dltag = strings.Trim(fmt.Sprintf("%q", dltag), "\"")
	return dltag
}
