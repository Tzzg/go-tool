package main

import (
	"context"
	"github.com/Tzzg/go-tool/data"
	"github.com/Tzzg/go-tool/worker_pool"
	"github.com/bits-and-blooms/bloom/v3"
	"github.com/blevesearch/bleve/v2"
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
	"github.com/golang-module/carbon/v2"
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/spf13/viper"
	"log"
	"os"
	"time"
)

func main() {
	// test Data
	a := data.StrToMd5("asdsa", 10)
	log.Printf(a)

	// 【work pool】
	wp := worker_pool.NewWorkerPool(context.TODO(), 16, 10000)

	add := wp.TryAdd(func(ctx context.Context) {
		time.Sleep(3)
		log.Println("echo 1")
	})
	if add == false {
		log.Printf("wp.TryAdd false")
	}
	log.Println("WaitAndClose")
	wp.WaitAndClose()

	// 【时间处理】
	//  go get -u github.com/golang-module/carbon/v2
	log.Println(carbon.Now())

	// 【config 配置】
	// go get github.com/spf13/viper
	viper.Set("KEY", "scsds")
	log.Println(viper.GetString("KEY"))

	// 【布隆过滤器】
	// go get -u github.com/bits-and-blooms/bloom/v3
	filter := bloom.NewWithEstimates(1000000, 0.01) // 构建能够接收 100 万个元素且误报率为 1% 的布隆过滤器
	filter.Add([]byte("Love"))
	log.Println(filter.Test([]byte("Love")))
	log.Println(filter.Test([]byte("Lov1e")))

	// 布谷鸟过滤器 提供了动态添加和删除项目的灵活性
	// go get -u github.com/seiflotfy/cuckoofilter
	cf := cuckoo.NewFilter(1000)
	cf.InsertUnique([]byte("geeky ogre"))

	// Lookup a string (and it a miss) if it exists in the cuckoofilter
	log.Println(cf.Lookup([]byte("hello")))
	log.Println(cf.Lookup([]byte("geeky ogre")))

	count := cf.Count()
	log.Println(count) // count == 1

	// Delete a string (and it a miss)
	cf.Delete([]byte("hello"))

	count = cf.Count()
	log.Println(count) // count == 1

	// Delete a string (a hit)
	cf.Delete([]byte("geeky ogre"))

	count = cf.Count()
	log.Println(count) // count == 0

	cf.Reset() // reset

	// 【数据结构】
	// go get -u github.com/emirpasic/gods
	list := arraylist.New()
	list.Add("a")                     // ["a"]
	list.Add("c", "b")                // ["a","c","b"]
	list.Sort(utils.StringComparator) // ["a","b","c"]
	a1, _ := list.Get(0)              // "a",true
	log.Println(a1)
	_, _ = list.Get(100)                  // nil,false
	_ = list.Contains("a", "b", "c")      // true
	_ = list.Contains("a", "b", "c", "d") // false
	list.Swap(0, 1)                       // ["b","a",c"]
	list.Remove(2)                        // ["b","a"]
	list.Remove(1)                        // ["b"]
	list.Remove(0)                        // []
	list.Remove(0)                        // [] (ignored)
	_ = list.Empty()                      // true
	_ = list.Size()                       // 0
	list.Add("a")                         // ["a"]
	list.Clear()                          // []
	list.Insert(0, "b")                   // ["b"]
	list.Insert(0, "a")                   // ["a","b"]

	// 文本分析
	// go get -u github.com/blevesearch/bleve/v2
	// open a new index
	mapping := bleve.NewIndexMapping()
	index, _ := bleve.New("example-bleve", mapping)

	// index some data
	_ = index.Index("2", "THIS IS SOME TEXT ok")

	// search for some text
	query := bleve.NewMatchQuery("text")
	search := bleve.NewSearchRequest(query)
	searchResults, _ := index.Search(search)
	log.Println(searchResults)
	index.Close()

	// 删掉分析持久化数据
	os.RemoveAll("example-bleve")

	//

}
