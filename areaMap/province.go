// Package areaMap 该文件是由go generate自动生成的，请勿直接修改代码！！！
// 如需更新请更新/data文件的数据源，然后在/generate下执行 make all
package areaMap

type ProvinceId struct {
	Name string `json:"name"`
	Pid  int    `json:"pid"`
}

type ProvincePid struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type ProvinceName struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
	Pid  int    `json:"pid"`
}

var ProvinceById = map[int]ProvinceId{
	1:  {Name: "北京", Pid: 0},
	2:  {Name: "天津", Pid: 0},
	3:  {Name: "河北省", Pid: 0},
	4:  {Name: "山西省", Pid: 0},
	5:  {Name: "内蒙古自治区", Pid: 0},
	6:  {Name: "辽宁省", Pid: 0},
	7:  {Name: "吉林省", Pid: 0},
	8:  {Name: "黑龙江省", Pid: 0},
	9:  {Name: "上海", Pid: 0},
	10: {Name: "江苏省", Pid: 0},
	11: {Name: "浙江省", Pid: 0},
	12: {Name: "安徽省", Pid: 0},
	13: {Name: "福建省", Pid: 0},
	14: {Name: "江西省", Pid: 0},
	15: {Name: "山东省", Pid: 0},
	16: {Name: "河南省", Pid: 0},
	17: {Name: "湖北省", Pid: 0},
	18: {Name: "湖南省", Pid: 0},
	19: {Name: "广东省", Pid: 0},
	20: {Name: "广西壮族自治区", Pid: 0},
	21: {Name: "海南省", Pid: 0},
	22: {Name: "重庆", Pid: 0},
	23: {Name: "四川省", Pid: 0},
	24: {Name: "贵州省", Pid: 0},
	25: {Name: "云南省", Pid: 0},
	26: {Name: "西藏自治区", Pid: 0},
	27: {Name: "陕西省", Pid: 0},
	28: {Name: "甘肃省", Pid: 0},
	29: {Name: "青海省", Pid: 0},
	30: {Name: "宁夏回族自治区", Pid: 0},
	31: {Name: "新疆维吾尔自治区", Pid: 0},
	32: {Name: "台湾省", Pid: 0},
	33: {Name: "香港特别行政区", Pid: 0},
	34: {Name: "澳门特别行政区", Pid: 0},
	35: {Name: "海外", Pid: 0},
}

var ProvinceByPid = map[int][]ProvincePid{
	0: {{Name: "北京", Id: 1}, {Name: "天津", Id: 2}, {Name: "河北省", Id: 3}, {Name: "山西省", Id: 4}, {Name: "内蒙古自治区", Id: 5}, {Name: "辽宁省", Id: 6}, {Name: "吉林省", Id: 7}, {Name: "黑龙江省", Id: 8}, {Name: "上海", Id: 9}, {Name: "江苏省", Id: 10}, {Name: "浙江省", Id: 11}, {Name: "安徽省", Id: 12}, {Name: "福建省", Id: 13}, {Name: "江西省", Id: 14}, {Name: "山东省", Id: 15}, {Name: "河南省", Id: 16}, {Name: "湖北省", Id: 17}, {Name: "湖南省", Id: 18}, {Name: "广东省", Id: 19}, {Name: "广西壮族自治区", Id: 20}, {Name: "海南省", Id: 21}, {Name: "重庆", Id: 22}, {Name: "四川省", Id: 23}, {Name: "贵州省", Id: 24}, {Name: "云南省", Id: 25}, {Name: "西藏自治区", Id: 26}, {Name: "陕西省", Id: 27}, {Name: "甘肃省", Id: 28}, {Name: "青海省", Id: 29}, {Name: "宁夏回族自治区", Id: 30}, {Name: "新疆维吾尔自治区", Id: 31}, {Name: "台湾省", Id: 32}, {Name: "香港特别行政区", Id: 33}, {Name: "澳门特别行政区", Id: 34}, {Name: "海外", Id: 35}},
}

var ProvinceByName = map[string][]ProvinceName{
	"福建省":      {{Name: "福建省", Id: 13, Pid: 0}},
	"湖南省":      {{Name: "湖南省", Id: 18, Pid: 0}},
	"吉林省":      {{Name: "吉林省", Id: 7, Pid: 0}},
	"黑龙江省":     {{Name: "黑龙江省", Id: 8, Pid: 0}},
	"江苏省":      {{Name: "江苏省", Id: 10, Pid: 0}},
	"西藏自治区":    {{Name: "西藏自治区", Id: 26, Pid: 0}},
	"陕西省":      {{Name: "陕西省", Id: 27, Pid: 0}},
	"甘肃省":      {{Name: "甘肃省", Id: 28, Pid: 0}},
	"北京":       {{Name: "北京", Id: 1, Pid: 0}},
	"天津":       {{Name: "天津", Id: 2, Pid: 0}},
	"河北省":      {{Name: "河北省", Id: 3, Pid: 0}},
	"江西省":      {{Name: "江西省", Id: 14, Pid: 0}},
	"山东省":      {{Name: "山东省", Id: 15, Pid: 0}},
	"云南省":      {{Name: "云南省", Id: 25, Pid: 0}},
	"贵州省":      {{Name: "贵州省", Id: 24, Pid: 0}},
	"台湾省":      {{Name: "台湾省", Id: 32, Pid: 0}},
	"香港特别行政区":  {{Name: "香港特别行政区", Id: 33, Pid: 0}},
	"浙江省":      {{Name: "浙江省", Id: 11, Pid: 0}},
	"重庆":       {{Name: "重庆", Id: 22, Pid: 0}},
	"四川省":      {{Name: "四川省", Id: 23, Pid: 0}},
	"广西壮族自治区":  {{Name: "广西壮族自治区", Id: 20, Pid: 0}},
	"海南省":      {{Name: "海南省", Id: 21, Pid: 0}},
	"澳门特别行政区":  {{Name: "澳门特别行政区", Id: 34, Pid: 0}},
	"山西省":      {{Name: "山西省", Id: 4, Pid: 0}},
	"辽宁省":      {{Name: "辽宁省", Id: 6, Pid: 0}},
	"河南省":      {{Name: "河南省", Id: 16, Pid: 0}},
	"内蒙古自治区":   {{Name: "内蒙古自治区", Id: 5, Pid: 0}},
	"广东省":      {{Name: "广东省", Id: 19, Pid: 0}},
	"海外":       {{Name: "海外", Id: 35, Pid: 0}},
	"新疆维吾尔自治区": {{Name: "新疆维吾尔自治区", Id: 31, Pid: 0}},
	"湖北省":      {{Name: "湖北省", Id: 17, Pid: 0}},
	"青海省":      {{Name: "青海省", Id: 29, Pid: 0}},
	"宁夏回族自治区":  {{Name: "宁夏回族自治区", Id: 30, Pid: 0}},
	"上海":       {{Name: "上海", Id: 9, Pid: 0}},
	"安徽省":      {{Name: "安徽省", Id: 12, Pid: 0}},
}
