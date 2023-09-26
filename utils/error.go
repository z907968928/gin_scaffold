package utils

import "fmt"

type E int

const (
	// 参数错误
	ErrorCodeInvalidParam E = 1000
	ErrorCodeEmptyParam   E = 1001

	// 数据库操作错误
	ErrorCodeDdError           E = 2000
	ErrorCodeSQLConnectFailure E = 2001
	ErrorCodeSQLFailure        E = 2002
	ErrorCodePackSQLFailure    E = 2003
	ErrorCodeSQLQueryFailure   E = 2004
	ErrorCodeSQLQueryEmpty     E = 2005
	ErrorCodeSQLDataNotExist   E = 2006
	ErrorCodeSQLInsertFailure  E = 2007
	ErrorCodeSQLDeleteFailure  E = 2008
	ErrorCodeSQLUpdateFailure  E = 2009

	ErrorCodeRedisConnectFailure E = 2020
	ErrorCodeRedisFailure        E = 2021

	// 用户权限
	ErrorCodeNotLogin E = 4000
	ErrorCodeNotAuth  E = 4001

	// 内部错误
	ErrorCodeINNERERR          E = 5000
	ErrorCodeUNKNOW            E = 5001
	ErrorCodeRELYSERVERFAILURE E = 5002
)

type ErrCatch struct {
	ErrCode E
	ErrMsg  string
}

// Error 方法实现了 error 接口，它返回错误消息
func (e *ErrCatch) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.ErrCode, e.ErrMsg)
}

func (e *ErrCatch) ErrCatch(errNum E, errMsg string) ErrCatch {
	return ErrCatch{
		ErrCode: errNum,
		ErrMsg:  errMsg,
	}
}
