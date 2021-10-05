package main

import (
	"fmt"
	"time"
)

type FilterBuilder func(next Filter) Filter

type handlerFunc func(c *Context)

type Filter func(c *Context)

var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		start := time.Now().Nanosecond()
		next(c)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 纳秒", end-start)
	}
}

var _ FilterBuilder = LogFilterBuilder

func LogFilterBuilder(next Filter) Filter {
	return func(c *Context) {
		fmt.Printf("Request.URL : %s", c.R.Host + c.R.URL.Path)
		fmt.Println()
		next(c)
		// todo c.W.status can't get value here
		fmt.Printf("这里我没有取到Response中的StatusCode，c.Resp.StatusCode : %s", c.Resp.StatusCode)
		fmt.Println()
	}
}

//func TransferHeader(next Filter) Filter{
//	return func(c *Context){
//		//fmt.Fprintf(ctx.W, "header is %v\n", ctx.R.Header)
//		for headerKey, values := range c.R.Header {
//			c.W.Header().Add(headerKey, values[0])
//		}
//		verFromOsEnv := os.Getenv("VERSION")
//		fmt.Printf("从环境变量中取得 VERSION的值 ： %s ,并写入Header中",verFromOsEnv)
//		fmt.Println()
//		c.W.Header().Add("VERSION", verFromOsEnv)
//		//c.W.WriteHeader(http.StatusCreated)
//	}
//}
