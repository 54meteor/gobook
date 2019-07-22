package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/golang/groupcache"
)

var (
	peers_addrs = []string{"http://127.0.0.1:8001", "http://127.0.0.1:8002", "http://127.0.0.1:8003"}
)

func main() {
	//	验证启动参数
	if len(os.Args) <= 1 {
		fmt.Println("invalid hosts")
		os.Exit(1)
	}
	local_addr := os.Args[1]
	//	创建HTTP池
	peers := groupcache.NewHTTPPool("http://" + local_addr)
	peers.Set(peers_addrs...)
	//	创建Group
	var image_cache = groupcache.NewGroup("cache", 8<<30, groupcache.GetterFunc(
		//	无法命中缓存的处理方式
		func(ctx groupcache.Context, key string, dest groupcache.Sink) error {
			//		从文件中读取内容
			result, err := ioutil.ReadFile(key)
			if err != nil {
				fmt.Printf("read file error %s.\n", err.Error())
				message := "{'status':'0','msg':'" + err.Error() + "'}"
				dest.SetBytes([]byte(message))
				return nil
			}
			fmt.Printf("asking for %s from local file system\n", key)
			//		设置缓存内容
			dest.SetBytes([]byte(result))
			return nil
		}))
	//	创建请求接口
	http.HandleFunc("/cache", func(rw http.ResponseWriter, r *http.Request) {
		//		响应cache请求
		var data []byte
		//		获取参数
		k := r.URL.Query().Get("id")
		fmt.Printf("user get %s from groupcache\n", k)
		//		根据key获取缓存信息
		image_cache.Get(nil, k, groupcache.AllocatingByteSliceSink(&data))
		//		提供返回值
		rw.Write([]byte(data))
	})
	log.Fatal(http.ListenAndServe(local_addr, nil))
}
