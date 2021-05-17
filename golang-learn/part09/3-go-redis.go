package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init()  {
	pool = &redis.Pool{
		MaxIdle: 16,
		MaxActive: 0,
		IdleTimeout: 300,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379")
		},
	}

}

func set()  {
	c:= pool.Get()
	defer c.Close()

	_, err := c.Do("set", "abc", 100)
	if err != nil{
		fmt.Println(err)
		return
	}

	// get
	r, err := redis.Int(c.Do("get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :",err)
		return
	}
	fmt.Println(r)
	pool.Close()
}


// string批量操作
func StringPool()  {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("MSet", "abc", 100, "def", 300)
	if err != nil{
		fmt.Println(err)
		return
	}
	r, err := redis.Ints(c.Do("MGet", "abc", "def"))
	if err != nil{
		fmt.Println(err)
		return
	}
	for _, v := range r{
		fmt.Println(v)
	}

	pool.Close()
}

// 过期时间 expire

func Expire()  {
	c := pool.Get()
	defer c.Close()


	//_, err := c.Do("set", "abc", 100)
	//if err != nil{
	//	fmt.Println(err)
	//	return
	//}

	_, err := c.Do("expire", "abc", 10)
	if err != nil{
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :",err)
		return
	}
	fmt.Println(r)

}


// list
func TestList()  {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("lpush", "book_list", "abc", "def", 300)
	if err != nil{
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil{
		fmt.Println("get value failed,", err)
		return
	}
	fmt.Println(r)

}

// hash
func TestHash()  {
	c := pool.Get()
	defer c.Close()

	_, err := c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}




func main() {
	//set()
	//StringPool()
	//Expire()
	//TestList()
	TestHash()
}
