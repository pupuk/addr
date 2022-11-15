# 收货地址智能解析（Go语言版）

## 1. 基本介绍

### 1.1 功能

> 项目内含2个功能
>
> - 把字符串解析成姓名、电话、邮编、身份证号、收货地址
> - 把收货地址再解析成省、市、区县、街道地址
> - 支持虚拟手机号（美团、饿了吗等）更新日期：2022-11-15

该项目仍然采用的是统计特征分析，以最大的概率来匹配，得出大概率的解。是根据PHP的addrss_prob版本改写优化而来，解析成功率在98%以上。

### 1.2 性能
Go语言版本采用hash map索引检索模式，对关键词进行匹配，将性能提升到了最高。实测解析10000个地址时间消耗在600ms左右（一条耗时约：0.06ms/条，即1s可解析约1.67w条），性能非常不错。相比纯PHP版本，性能提升近80倍。

### 1.3 特别感谢
 因为工作繁忙，没有时间重头到尾来实现go语言的版本，在关键点的解析省、市、区的部分以及最后的梳理校订，[cxy-chenxuanyu](https://github.com/cxy-chenxuanyu)同学贡献了智慧和代码，再次特别感谢，欢迎围观。

## 2. 使用说明

```
 golang版本 >= v1.11
```

- Install

  - ```git
    go get github.com/pupuk/addr
    ```

- 使用git克隆本项目

  - ```git
    git clone https://github.com/pupuk/addr.git
    ```

- 索引树自动生成

  - ```shell-script
    cd /generate
    make all
    如需更新最新的行政信息，请自行修改/data下的json数据然后重新执行一遍以上命令
    ```
  
- 使用Demo
```go
package main

import (
	"fmt"

	"github.com/pupuk/addr"
)

func main() {
	parse := addr.Smart("张三 13800138000 龙华区龙华街道1980科技文化产业园3栋308 身份证120113196808214821")

	// 输出解析结果
	fmt.Println(parse.Name)     // 张三
	fmt.Println(parse.IdNumber) // 120113196808214821
	fmt.Println(parse.Mobile)   // 13800138000
	fmt.Println(parse.PostCode) // 570100
	fmt.Println(parse.Province) // 广东省
	fmt.Println(parse.City)     // 深圳市
	fmt.Println(parse.Region)   // 龙华区
	fmt.Println(parse.Street)   // 龙华街道1980科技文化产业园3栋317
	fmt.Println(parse.Address)  // 深圳市龙华区龙华街道1980科技文化产业园3栋317
}

```

- 构建HTTP服务端DEMO（使用fiber，gin类似，几行代码即可构建一个解析服务）
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pupuk/addr"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		parse := addr.Smart(c.Query("addr"))
		return c.JSON(parse)
	})

	app.Listen(":3000")
}

```
HTTP请求结果
![image](https://user-images.githubusercontent.com/7934974/184922637-9d909cc2-fa47-45aa-8297-0ea69495f215.png)


### 反馈 &改进
#### Issue
如果有什么问题或建议，或者发现有不能识别，或者识别错误的地址，
提交到[Github Issue](https://github.com/pupuk/addr/issues)

#### 协作            
1. 本版本提供了注释，希望大家能fork，优化，提PR，点个star，大家一起来维护地更好
2. 欢迎改写成其它语言版本，只需注明参考链接即可。

#### 联系我
* Email：pujiexuan@gmail.com
* QQ: 632085136 欢迎一起学习讨论

### 致谢
* [cxy-chenxuanyu](https://github.com/cxy-chenxuanyu)
