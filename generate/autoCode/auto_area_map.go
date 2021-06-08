package autoCode

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type AreaInfoId struct {
	Name    string `json:"name"`
	Pid     int    `json:"pid"`
	Zipcode int    `json:"zipcode"`
}

type AreaInfoPid struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Zipcode int    `json:"zipcode"`
}

type AreaInfoName struct {
	Name    string `json:"name"`
	Id      int    `json:"id"`
	Pid     int    `json:"pid"`
	Zipcode int    `json:"zipcode"`
}

// 自动生成json数据源文件地址
const filePath = "../data/"

// 自动生成代码地址
const packageName = "areaMap"
const codePath = "../areaMap/"

// 定义需要生成map的行政等级
var ranks = []string{"city", "province", "region"}

// AutoAreaMap 自动生成行政区划map
func AutoAreaMap() {
	for _, rank := range ranks {
		data, err := ioutil.ReadFile(filePath + rank)
		if err != nil {
			fmt.Println("读取json文件失败,行政等级为："+rank+"，请检查文件！---", err)
			return
		}

		m := make(map[int]AreaInfoId)
		err = json.Unmarshal(data, &m)
		if err != nil {
			fmt.Println("json序列化数据源失败！请检出数据！path："+filePath+rank+"---", err)
			return
		}

		// 让map数据根据key排序
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)

		str := ""
		// 构建package
		str += "// Package " + packageName + " 该文件是由go generate自动生成的，请勿直接修改代码！！！\n"
		str += "// 如需更新请更新/data文件的数据源，然后在/generate下执行 make all\n"
		str += "package areaMap\n\n"

		// 构建struct
		str += "type " + strings.Title(rank) + "Id struct {\n"
		str += "Name string `json:\"name\"`\n"
		str += "Pid int `json:\"pid\"`\n"
		if rank != "province" {
			str += "Zipcode int `json:\"zipcode\"`\n"
		}
		str += "}\n\n"

		str += "type " + strings.Title(rank) + "Pid struct {\n"
		str += "Name string `json:\"name\"`\n"
		str += "Id int `json:\"id\"`\n"
		if rank != "province" {
			str += "Zipcode int `json:\"zipcode\"`\n"
		}
		str += "}\n\n"

		str += "type " + strings.Title(rank) + "Name struct {\n"
		str += "Name string `json:\"name\"`\n"
		str += "Id int `json:\"id\"`\n"
		str += "Pid int `json:\"pid\"`\n"
		if rank != "province" {
			str += "Zipcode int `json:\"zipcode\"`\n"
		}
		str += "}\n\n"

		// 为构建pid索引树创造条件
		str1Arr := make(map[int][]interface{})
		// 为构建name索引树创造条件
		str2Arr := make(map[string][]interface{})

		// 构建map
		// str是构建根据行政id来生成的索引树
		str += "var " + strings.Title(rank) + "ById = map[int]" + strings.Title(rank) + "Id{\n"
		for _, key := range keys {
			var infoPid AreaInfoPid
			infoPid.Id = key
			infoPid.Name = m[key].Name
			infoPid.Zipcode = m[key].Zipcode
			str1Arr[m[key].Pid] = append(str1Arr[m[key].Pid], infoPid)

			var infoName AreaInfoName
			infoName.Id = key
			infoName.Pid = m[key].Pid
			infoName.Name = m[key].Name
			infoName.Zipcode = m[key].Zipcode
			str2Arr[m[key].Name] = append(str2Arr[m[key].Name], infoName)

			t := reflect.TypeOf(m[key])
			v := reflect.ValueOf(m[key])
			name := t.Field(0).Name
			name1 := v.Field(0).String()
			pid := t.Field(1).Name
			pid1 := v.Field(1).Int()
			if rank == "province" {
				str += strconv.Itoa(key) + ":{" + name + ":\"" + name1 + "\"," + pid + ":" + strconv.Itoa(int(pid1)) + "},\n"
				continue
			}
			zipCode := t.Field(2).Name
			zipCode1 := v.Field(2).Int()
			str += strconv.Itoa(key) + ":{" + name + ":\"" + name1 + "\"," + pid + ":" + strconv.Itoa(int(pid1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},\n"
		}
		str += "}\n\n"

		// str1是构建根据行政父id(pid)来生成的索引树
		str1 := "var " + strings.Title(rank) + "ByPid = map[int][]" + strings.Title(rank) + "Pid{\n"
		for key, value := range str1Arr {
			str1 += strconv.Itoa(key) + ":{"
			for _, value2 := range value {
				t := reflect.TypeOf(value2)
				v := reflect.ValueOf(value2)
				name := t.Field(0).Name
				name1 := v.Field(0).String()
				id := t.Field(1).Name
				id1 := v.Field(1).Int()
				if rank == "province" {
					str1 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "},"
					continue
				}
				zipCode := t.Field(2).Name
				zipCode1 := v.Field(2).Int()
				str1 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},"
			}
			str1 = strings.TrimRight(str1, ",")
			str1 += "},\n"
		}
		str1 += "}\n\n"

		// str2是构建根据地名(name)来生成的索引树
		str2 := "var " + strings.Title(rank) + "ByName = map[string][]" + strings.Title(rank) + "Name{\n"
		for key, value := range str2Arr {
			str2 += "\"" + key + "\":{"
			for _, value2 := range value {
				t := reflect.TypeOf(value2)
				v := reflect.ValueOf(value2)
				name := t.Field(0).Name
				name1 := v.Field(0).String()
				id := t.Field(1).Name
				id1 := v.Field(1).Int()
				pid := t.Field(2).Name
				pid1 := v.Field(2).Int()
				if rank == "province" {
					str2 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + pid + ":" + strconv.Itoa(int(pid1)) + "},"
					continue
				}
				zipCode := t.Field(3).Name
				zipCode1 := v.Field(3).Int()
				str2 += "{" + name + ":\"" + name1 + "\"," + id + ":" + strconv.Itoa(int(id1)) + "," + pid + ":" + strconv.Itoa(int(pid1)) + "," + zipCode + ":" + strconv.Itoa(int(zipCode1)) + "},"
			}
			str2 = strings.TrimRight(str2, ",")
			str2 += "},\n"
		}
		str2 += "}\n"

		str = str + str1 + str2

		// 尝试创建此路径
		uploadDir := codePath
		mkdirErr := os.MkdirAll(uploadDir, os.ModePerm)
		if mkdirErr != nil {
			fmt.Println(err)
		}

		// 打开文件
		file, err := os.OpenFile(codePath+rank+".go", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}

		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		_, err = write.WriteString(str)

		//Flush将缓存的文件真正写入到文件中
		err = write.Flush()

		//及时关闭file句柄
		func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("关闭file句柄失败！，可能回导致内存泄漏，请care一下！---", err)
			}
		}(file)
	}
}
