package utils

import (
	"crypto/md5"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
	"path"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// 随机生成指定位数的大写字母和数字的组合
func GetRandomString(l int) string {
	str := "0123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// MergeArray 合并数组
func MergeArray(dest []interface{}, src []interface{}) (result []interface{}) {
	result = make([]interface{}, len(dest)+len(src))
	//将第一个数组传入result
	copy(result, dest)
	//将第二个数组接在尾部，也就是 len(dest):
	copy(result[len(dest):], src)
	return
}

//判断元素是否存在
func IsRepeatEleArr(a []int, e int) bool {
	for _, aele := range a {
		if aele == e {
			return true
		}
	}
	return false
}

//判断元素是否存在
func IsRepeatInt64Arr(a []int64, e int64) bool {
	for _, aele := range a {
		if aele == e {
			return true
		}
	}
	return false
}

//判断元素是否存在
func IsRepeatStringArr(a []string, e string) bool {
	for _, aele := range a {
		if aele == e {
			return true
		}
	}
	return false
}

// 某值是否存在切片里
func IsValueInSlice(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

// []int 去重
func RemoveRepeatedIntElement(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// []int64 去重
func RemoveRepeatedInt64Element(arr []int64) (newArr []int64) {
	newArr = make([]int64, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// []string 去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// []int 去重
func RemoveRepeatedElementInt(arr []int) (newArr []int) {
	newArr = make([]int, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 传递数据
func PassData(left interface{}, right interface{}) (err error) {
	rightType := reflect.TypeOf(right)
	if rightType.Kind() != reflect.Ptr {
		return errors.New("the right value is not a pointer type. ")
	}
	leftByteSlice, err := json.Marshal(left)
	if err != nil {
		return err
	}
	err = json.Unmarshal(leftByteSlice, right)
	return
}

//取交集并且去除重复数据
func RemoveRepeatedIntersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times >= 1 {
			nn = append(nn, v)
			delete(m, v)
		}
	}
	return nn
}

//如果参数含重复元素,求交集不去除重复数据
func IntersectInt(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//如果参数含重复元素,求交集不去除重复数据
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//int slice求差集 slice1(在slice1中没在slice2中)
func DifferenceInt(slice1, slice2 []int) []int {
	slice1 = RemoveRepeatedElementInt(slice1)
	slice2 = RemoveRepeatedElementInt(slice2)
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := IntersectInt(slice1, slice2) //取交集并去重
	for _, v := range inter {
		m[v]++
	}
	//取差集
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//string slice求差集 slice1(在slice1中没在slice2中)
func Difference(slice1, slice2 []string) []string {
	slice1 = RemoveRepeatedElement(slice1)
	slice2 = RemoveRepeatedElement(slice2)
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//逗号分隔字符串，遍历元素转切片
func StringToSlice(param string) []string {
	var externalUserid []string
	for _, v := range strings.Split(param, ",") {
		v = strings.TrimSpace(v)
		externalUserid = append(externalUserid, v)
	}
	return externalUserid
}

//翻转int slice
func reverseInt(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// SliceColumnJoin 将slice 中struct 的任何一列 取出并拼接.
func SliceColumnJoin(data interface{}, column string, sep string) (str string) {
	var s strings.Builder
	va := reflect.ValueOf(data)
	if va.Kind() == reflect.Ptr {
		va = va.Elem()
	}
	if va.Kind() == reflect.Slice {
		for i := 0; i < va.Len(); i++ {
			elem := va.Index(i)
			if elem.Kind() == reflect.Ptr {
				elem = elem.Elem()
			}
			if elem.Kind() == reflect.Struct {
				elemTy := elem.Type()
				num := elem.NumField()
				for j := 0; j < num; j++ {
					name := elemTy.Field(j).Name
					if name == column {
						if elem.Field(j).Kind() != reflect.String {
							return
						}
						s.WriteString(elem.Field(j).String())
						s.WriteString(sep)
					}
				}
			}
		}
	}
	str = strings.Trim(s.String(), sep)
	return
}

//生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func UniqueId() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(cryptoRand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

//切片查找
func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func Int64ToInt(s int64) int {
	intString := strconv.FormatInt(s, 10)
	ind10, _ := strconv.Atoi(intString)
	return ind10
}

//返回输入切片中某个单一列的值。
func SliceColumn(structSlice []interface{}, key string) []interface{} {
	rt := reflect.TypeOf(structSlice)
	rv := reflect.ValueOf(structSlice)
	if rt.Kind() == reflect.Slice { //切片类型
		var sliceColumn []interface{}
		elemt := rt.Elem() //获取切片元素类型
		for i := 0; i < rv.Len(); i++ {
			inxv := rv.Index(i)
			if elemt.Kind() == reflect.Struct {
				for i := 0; i < elemt.NumField(); i++ {
					if elemt.Field(i).Name == key {
						strf := inxv.Field(i)
						switch strf.Kind() {
						case reflect.String:
							sliceColumn = append(sliceColumn, strf.String())
						case reflect.Float64:
							sliceColumn = append(sliceColumn, strf.Float())
						case reflect.Int, reflect.Int64:
							sliceColumn = append(sliceColumn, strf.Int())
						default:
							//do nothing
						}
					}
				}
			}
		}
		return sliceColumn
	}
	return nil
}

//通过合并两个切片来创建一个新切片，其中的一个切片元素为键名，另一个切片元素为键值：
func SliceCombine(s1, s2 []interface{}) map[interface{}]interface{} {
	if len(s1) != len(s2) {
		panic("the number of elements for each slice isn't equal")
	}
	m := make(map[interface{}]interface{}, len(s1))
	for i, v := range s1 {
		m[v] = s2[i]
	}
	return m
}

/**
 * 根据 user agent string 判断用户的平台、浏览器
 * 参考资料
 * **************************************************************************************************************************************************
 *
 * 台式机
 *
 * Linux Ubuntu
 * Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.2.2pre) Gecko/20100225 Ubuntu/9.10 (karmic) Namoroka/3.6.2pre
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Linux Mandriva 2008.1
 * Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.1) Gecko/2008072403 Mandriva/3.0.1-1mdv2008.1 (2008.1) Firefox/3.0.1
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Linux suSE 10.1
 * Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.8.0.3) Gecko/20060425 SUSE/1.5.0.3-7 Firefox/1.5.0.31
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Windows XP SP3
 * Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.9.1) Gecko/20090624 Firefox/3.5 (.NET CLR 3.5.30729)
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Windows Vista
 * Mozilla/5.0 (Windows; U; Windows NT 6.1; nl; rv:1.9.2.13) Gecko/20101203 Firefox/3.6.13
 * Mozilla/5.0 (Windows; U; Windows NT 6.0; en-US; rv:1.9.2.6) Gecko/20100625 Firefox/3.6.6 (.NET CLR 3.5.30729)
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * windows 2000
 * Mozilla/5.0 (Windows; U; Windows NT 5.0; en-GB; rv:1.8.1b2) Gecko/20060821 Firefox/2.0b2
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Windows 7
 * Mozilla/5.0 (Windows NT 6.1; WOW64; rv:14.0) Gecko/20100101 Firefox/14.0.1
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Windows Server 2008
 * Mozilla/5.0 (Windows; U; Windows NT 6.1; en-US; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5 (.NET CLR 3.5.30729)
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * iMac OSX 10.7.4
 * Mozilla/5.0 (Macintosh; Intel Mac OS X 10.7; rv:13.0) Gecko/20100101 Firefox/13.0.1
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Mac OS X
 * Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.9) Gecko/20100824 Firefox/3.6.9
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 *
 * 手持设备
 *
 * iPad
 * Mozilla/5.0 (iPad; U; CPU OS 3_2 like Mac OS X; en-us) AppleWebKit/531.21.10 (KHTML, like Gecko) Version/4.0.4 Mobile/7B334b Safari/531.21.10
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * iPad 2
 * Mozilla/5.0 (iPad; CPU OS 5_1 like Mac OS X; en-us) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9B176 Safari/7534.48.3
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * iPhone 4
 * Mozilla/5.0 (iPhone; U; CPU iPhone OS 4_0 like Mac OS X; en-us) AppleWebKit/532.9 (KHTML, like Gecko) Version/4.0.5 Mobile/8A293 Safari/6531.22.7
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * iPhone 5
 * Mozilla/5.0 (iPhone; CPU iPhone OS 5_0 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3
 * --------------------------------------------------------------------------------------------------------------------------------------------------
 * Android
 * Mozilla/5.0 (Linux; U; Android 2.2; en-us; Nexus One Build/FRF91) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1
 * **************************************************************************************************************************************************
 * @author YH
 */
type HeaderGetSystemInfo struct {
	/** 浏览器类型 */
	BrowserType  string `json:"browser_type" form:"browser_type" comment:"浏览器类型" `
	PlatformType string `json:"platform_type" form:"platform_type" comment:"操作系统类型" `
}

/**
 * 根据请求头来判断浏览器类型及操作系统类型
 * @param userAgent 请求头
 * @return
 */
func GetUserAgent(userAgent string) (data HeaderGetSystemInfo) {
	if userAgent == "" {
		userAgent = ""
	}
	if strings.Index(userAgent, "Windows") != -1 { //主流应用靠前
		if strings.Index(userAgent, "Windows NT 10.0") != -1 { //Windows 10
			return judgeBrowser(userAgent, "Windows 10") //判断浏览器
		} else if strings.Index(userAgent, "Windows NT 6.2") != -1 { //Windows 8
			return judgeBrowser(userAgent, "Windows 8") //判断浏览器
		} else if strings.Index(userAgent, "Windows NT 6.1") != -1 { //Windows 7
			return judgeBrowser(userAgent, "Windows 7")
		} else if strings.Index(userAgent, "Windows NT 6.0") != -1 { //Windows Vista
			return judgeBrowser(userAgent, "Windows Vista")
		} else if strings.Index(userAgent, "Windows NT 5.2") != -1 { //Windows XP x64 Edition
			return judgeBrowser(userAgent, "Windows XP")
		} else if strings.Index(userAgent, "Windows NT 5.1") != -1 { //Windows XP
			return judgeBrowser(userAgent, "Windows XP")
		} else if strings.Index(userAgent, "Windows NT 5.01") != -1 { //Windows 2000, Service Pack 1 (SP1)
			return judgeBrowser(userAgent, "Windows 2000")
		} else if strings.Index(userAgent, "Windows NT 5.0") != 1 { //Windows 2000
			return judgeBrowser(userAgent, "Windows 2000")
		} else if strings.Index(userAgent, "Windows NT 4.0") != -1 { //Microsoft Windows NT 4.0
			return judgeBrowser(userAgent, "Windows NT 4.0")
		} else if strings.Index(userAgent, "Windows 98; Win 9x 4.90") != -1 { //Windows Millennium Edition (Windows Me)
			return judgeBrowser(userAgent, "Windows ME")
		} else if strings.Index(userAgent, "Windows 98") != -1 { //Windows 98
			return judgeBrowser(userAgent, "Windows 98")
		} else if strings.Index(userAgent, "Windows 95") != -1 { //Windows 95
			return judgeBrowser(userAgent, "Windows 95")
		} else if strings.Index(userAgent, "Windows CE") != -1 { //Windows CE
			return judgeBrowser(userAgent, "Windows CE")
		}
	} else if strings.Index(userAgent, "Mac OS X") != -1 {
		if strings.Index(userAgent, "iPhone") != -1 {
			return judgeBrowser(userAgent, "iPhone")
		} else if strings.Index(userAgent, "iPad") != -1 {
			return judgeBrowser(userAgent, "iPad") //判断系统
		} else {
			return judgeBrowser(userAgent, "Mac") //判断系统
		}
	} else if strings.Index(userAgent, "Android") != -1 {
		return judgeBrowser(userAgent, "Android") //判断系统
	} else if strings.Index(userAgent, "Linux") != -1 {
		return judgeBrowser(userAgent, "Linux") //判断系统
	} else if strings.Index(userAgent, "FreeBSD") != -1 {
		return judgeBrowser(userAgent, "FreeBSD") //判断系统
	} else if strings.Index(userAgent, "Solaris") != -1 {
		return judgeBrowser(userAgent, "Solaris") //判断系统
	}
	return judgeBrowser(userAgent, "其他")
}

func judgeBrowser(userAgent string, platformType string) (data HeaderGetSystemInfo) {
	//fmt.Printf("userAgent=%v\n	platformType=%v\n", userAgent, platformType)
	data.PlatformType = platformType
	if strings.Index(userAgent, "Edge") != -1 {
		data.BrowserType = "Microsoft Edge"
		return data
	} else if strings.Index(userAgent, "QQBrowser") != -1 {
		data.BrowserType = "腾讯浏览器"
		return data
	} else if strings.Index(userAgent, "Chrome") != -1 && strings.Index(userAgent, "Safari") != -1 {
		data.BrowserType = "Chrome"
		return data
	} else if strings.Index(userAgent, "Firefox") != -1 {
		data.BrowserType = "Firefox"
		return data
	} else if strings.Index(userAgent, "360") != -1 { //Internet Explorer 6
		data.BrowserType = "360浏览器"
		return data
	} else if strings.Index(userAgent, "Opera") != -1 { //Internet Explorer 6
		data.BrowserType = "Opera"
		return data
	} else if strings.Index(userAgent, "Safari") != -1 && strings.Index(userAgent, "Chrome") == -1 { //Internet Explorer 6
		data.BrowserType = "Safari"
		return data
	} else if strings.Index(userAgent, "MetaSr1.0") != -1 { //Internet Explorer 6
		data.BrowserType = "搜狗浏览器"
		return data
	} else if strings.Index(userAgent, "TencentTraveler") != -1 { //Internet Explorer 6
		data.BrowserType = "腾讯浏览器"
		return data
	} else if strings.Index(userAgent, "UC") != -1 { //Internet Explorer 6
		data.BrowserType = "UC浏览器"
		return data
	} else if strings.Index(userAgent, "MSIE") != -1 {
		if strings.Index(userAgent, "MSIE 10.0") != -1 { //Internet Explorer 10
			data.BrowserType = "IE 10"
			return data
		} else if strings.Index(userAgent, "MSIE 9.0") != -1 { //Internet Explorer 9
			data.BrowserType = "IE 9"
			return data
		} else if strings.Index(userAgent, "MSIE 8.0") != -1 { //Internet Explorer 8
			data.BrowserType = "IE 8"
			return data
		} else if strings.Index(userAgent, "MSIE 7.0") != -1 { //Internet Explorer 7
			data.BrowserType = "IE 7"
			return data
		} else if strings.Index(userAgent, "MSIE 6.0") != -1 { //Internet Explorer 6
			data.BrowserType = "IE 6"
			return data
		}
	} else { //暂时支持以上三个主流.其它浏览器,待续...
		data.BrowserType = "其他"
		return data
	}
	data.BrowserType = "其他"
	return data
}

// HasLocalIPAddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPAddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		if ip = strings.TrimSpace(ip); ip != "" && !HasLocalIPAddr(ip) {
			return ip
		}
	}

	if ip = strings.TrimSpace(r.Header.Get("X-Real-Ip")); ip != "" && !HasLocalIPAddr(ip) {
		return ip
	}

	if ip = RemoteIP(r); !HasLocalIPAddr(ip) {
		return ip
	}

	return ""
}

// RemoteIP 通过 RemoteAddr 获取 IP 地址， 只是一个快速解析方法。
func RemoteIP(r *http.Request) string {
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

//获取函数调用者信息
func GetCallerInfo(skip int) (info string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		info = "runtime.Caller() failed"
		return
	}
	funcName := runtime.FuncForPC(pc).Name()
	fileName := path.Base(file) // Base函数返回路径的最后一个元素
	return fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, fileName, lineNo)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//生成随机id
func CreateRandomId() string {
	h := md5.New()
	h.Write([]byte(fmt.Sprint(time.Now().UnixNano() + rand.Int63n(1000000))))
	return hex.EncodeToString(h.Sum(nil))
}
