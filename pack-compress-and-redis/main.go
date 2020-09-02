package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"gopkg.in/redis.v5"
	"io"
	"time"

	"github.com/vmihailenco/msgpack/v5"
)

type Item struct {
	Foo  string
	List []ArrayElement
}

type ArrayElement struct {
	Timestamp time.Time
	Bar       string
}

func main() {
	rc := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	rc.FlushAll()

	// pack and compress
	item := Item{Foo: "foo foo", List: []ArrayElement{{
		Timestamp: time.Now().Add(5 * time.Minute),
		Bar:       "bar bar",
	}}}
	packed, err := pack(item)
	if err != nil {
		panic(err)
	}
	compressedAndPacked := compress(packed)

	// push to redis
	rc.RPush("queue", compressedAndPacked)

	// pop from redis
	var redisBytes []byte
	redisStr, err := rc.LPop("queue").Result()
	if err != nil {
		panic(err)
	}
	redisBytes = []byte(redisStr)

	// uncompress and unpack
	uncompressed := uncompress(redisBytes)
	res, err := unpack(uncompressed)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

func compress(input []byte) []byte {
	var in bytes.Buffer
	w := zlib.NewWriter(&in)
	w.Write(input)
	w.Close()
	return in.Bytes()
}
func uncompress(in []byte) []byte {
	var out bytes.Buffer
	r, _ := zlib.NewReader(bytes.NewBuffer(in))
	io.Copy(&out, r)
	return out.Bytes()
}

func pack(i Item) ([]byte, error) {
	return msgpack.Marshal(i)
}

func unpack(input []byte) (Item, error) {
	var i Item
	err := msgpack.Unmarshal(input, &i)
	if err != nil {
		return Item{}, err
	}
	return i, nil
}
