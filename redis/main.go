package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6380",
		Password: "123456",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("link redis failed ,err :", err)
	}
	fmt.Println("link redis success")
}

func redisString() {
	//判断k1是否存在
	// 当然了，如果存在，set也会重新进行赋值
	if r, _ := rdb.Exists(ctx, "k1").Result(); r != 0 {
		fmt.Println("k1  exists!")
		return
	} else if r == 0 {
		goto setK1
	}

setK1:
	//set设置
	// set k1 100,expire k1 60 设置新的字符串，过期时间为60；不过期为0
	//err := rdb.Set(ctx, "k1", 100, 60*time.Second).Err()
	err := rdb.Set(ctx, "k1", 100, 0).Err()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("set k1 success")
	}

	//get获取
	v, err := rdb.Get(ctx, "k1").Result()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("k1的值为", v)
	}

	//其他命令
	//Do() 方法返回 Cmd 类型，你可以使用它获取你想要的类型
	v, _ = rdb.Do(ctx, "get", "k1").Text()
	fmt.Println("Do()：k1的值为", v)

	n, _ := rdb.Do(ctx, "append", "k1", "kkk").Int()
	fmt.Println("k1新增后的元素数量", n)
}

func redisNil() {
	//rdb.Do(ctx, "flushall")
	v, err := rdb.Get(ctx, "k1").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key不存在")
	case err != nil:
		fmt.Println("get k1 failed ,err:", err)
	case v == "":
		fmt.Println("值是空的")
	default:
		fmt.Println("k1的值为：", v)
	}
}

func main() {
	initRedis()
	//redisString()
	redisNil()

}
