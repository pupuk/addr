package addr

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/pupuk/addr/areaMap"
)

type Address struct {
	IdNumber string `json:"id_number"`
	Mobile   string `json:"mobile"`
	PostCode string `json:"post_code"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Province string `json:"province"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Street   string `json:"street"`
}

// FilterWord 需要过滤掉收货地址中的常用说明字符，排除干扰词
var FilterWord = []string{"身份证号", "地址", "收货人", "收件人", "收货", "邮编", "电话", "手机", "手机号", "手机号码", "身份证号码", "身份证号", "身份证", "：", ":", "；", ";", "，", ",", "。", "."}

// Decompose 分离手机号(座机)，身份证号，姓名，地址等信息
func Decompose(info *Address, str string) *Address {
	//1. 过滤掉收货地址中的常用说明字符，排除干扰词
	for _, value := range FilterWord {
		str = strings.Replace(str, value, " ", -1)
	}

	//2. 多个空白字符(包括空格\r\n\t)换成一个空格
	reg := regexp.MustCompile(`\s+`)
	str = strings.TrimSpace(reg.ReplaceAllString(str, " "))

	//3. 去除手机号码中的短横线 如0136-3333-6666 主要针对苹果手机
	reg = regexp.MustCompile(`0?(\d{3})-(\d{4})-(\d{4})([-_]\d{2,})`)
	str = reg.ReplaceAllString(str, "$1$2$3$4")

	//4. 提取中国境内身份证号码
	reg = regexp.MustCompile(`(?i)\d{18}|\d{17}X`)
	IdNumber := reg.FindString(str)
	str = strings.Replace(str, IdNumber, "", -1)
	info.IdNumber = strings.ToUpper(IdNumber)

	//5. 提取11位手机号码或者7位以上座机号，支持虚拟号的提取
	reg = regexp.MustCompile(`\d{7,11}[\-_]\d{2,6}|\d{7,11}|\d{3,4}-\d{6,8}`)
	mobile := reg.FindString(str)
	str = strings.Replace(str, mobile, "", -1)
	info.Mobile = mobile

	//6. 提取6位邮编 邮编也可用后面解析出的省市区地址从数据库匹配出
	reg = regexp.MustCompile(`\d{6}`)
	postcode := reg.FindString(str)
	str = strings.Replace(str, postcode, "", -1)
	info.PostCode = postcode

	//再次把2个及其以上的空格合并成一个，并首位TRIM
	reg = regexp.MustCompile(` {2,}`)
	str = strings.TrimSpace(reg.ReplaceAllString(str, " "))

	//7. 按照空格切分 长度长的为地址 短的为姓名 因为不是基于自然语言分析，所以采取统计学上高概率的方案
	r := strings.Split(str, " ")

	name := r[0]
	for _, v := range r {
		if len(v) < len(name) {
			name = v
		}
	}

	if len(r) <= 1 {
		info.Address = r[0]
		return info
	}

	info.Name = name
	address := strings.TrimSpace(strings.Replace(str, name, "", -1))
	info.Address = address

	return info
}

// Smart 智能解析
func Smart(str string) *Address {
	var info Address
	info = *Decompose(&info, str)
	Parse(&info)
	if info.Region != "" && strings.Contains(info.Address, info.Region) {
		info.Street = info.Address[strings.LastIndex(info.Address, info.Region)+len(info.Region):]
	}
	if info.City != "" && strings.Contains(info.Address, info.City) {
		info.Street = info.Address[strings.LastIndex(info.Address, info.City)+len(info.City):]
	}

	return &info
}

// Parse 智能解析出省市区+街道地址
func Parse(address *Address) *Address {
	// 匹配所有省级
	pReg := regexp.MustCompile(`.+?(省|市|自治区|特别行政区|区)`)
	pArr := pReg.FindAllString(address.Address, -1)
	// 匹配所有市级
	// 由于该匹配可能会遗漏部分，所以合并省级匹配
	cReg := regexp.MustCompile(`.+?(省|市|自治州|州|地区|盟|县|自治县|区|林区)`)
	cArr := append(cReg.FindAllString(address.Address, -1), pArr...)
	// 匹配所有区县级
	// 由于该匹配可能会遗漏部分(如：东乡区)所以合并市级匹配
	rReg := regexp.MustCompile(`.+?(市|县|自治县|旗|自治旗|区|林区|特区|街道|镇|乡)`)
	rArr := append(rReg.FindAllString(address.Address, -1), cArr...)

	// 处理区县级
I:
	for _, r := range rArr {
		if r1, ok := areaMap.RegionByName[r]; ok && len(r1) == 1 {
			address.Region = r1[0].Name
			address.PostCode = strconv.Itoa(r1[0].Zipcode)
			getAddressById(address, r1[0].Pid, city)
			break
		} else if ok {
			for _, r2 := range r1 {
				address.Region = r2.Name
				address.PostCode = strconv.Itoa(r1[0].Zipcode)
				getAddressById(address, r2.Pid, city)
				for _, v := range cArr {
					if address.City == v {
						break I
					}
				}
			}
		}
	}
	// 处理市级
	if address.City == "" {
		for _, c := range cArr {
			if r1, ok := areaMap.CityByName[c]; ok {
				address.City = r1[0].Name
				address.PostCode = strconv.Itoa(r1[0].Zipcode)
				getAddressById(address, r1[0].Pid, province)
				getAddressByPid(address, r1[0].Id, region, rArr)
				break
			}
		}
	}

	// 处理省级
	if address.Province == "" {
		for _, p := range pArr {
			if r1, ok := areaMap.ProvinceByName[p]; ok {
				address.Province = r1[0].Name
				getAddressByPid(address, r1[0].Id, city, cArr)
				getAddressByPid(address, r1[0].Id, region, rArr)
				break
			}
		}
	}
	return address
}

const (
	// 定义map等级常量
	province = "province"
	city     = "city"
	region   = "region"
)

// 根据id获取地址信息
func getAddressById(address *Address, id int, rank string) *Address {
	if rank == province {
		info := areaMap.ProvinceById[id]
		address.Province = info.Name
	}
	if rank == city {
		info := areaMap.CityById[id]
		address.City = info.Name
		getAddressById(address, info.Pid, province)
	}
	if rank == region {
		info := areaMap.RegionById[id]
		address.Region = info.Name
		getAddressById(address, info.Pid, city)
	}
	return address
}

// 根据pid获取下一级行政地址信息
func getAddressByPid(address *Address, pid int, rank string, arr []string) *Address {
	if rank == city && address.City == "" {
		for _, addr := range arr {
			for _, info := range areaMap.CityByPid[pid] {
				if strings.Contains(info.Name, addr) {
					address.City = info.Name
					address.PostCode = strconv.Itoa(info.Zipcode)
					return address
				}
			}
		}
	}
	if rank == region && address.Region == "" {
		for _, addr := range arr {
			for _, info := range areaMap.RegionByPid[pid] {
				if strings.Contains(info.Name, addr) {
					address.Region = info.Name
					address.PostCode = strconv.Itoa(info.Zipcode)
					return address
				}
			}
		}
	}
	return address
}
