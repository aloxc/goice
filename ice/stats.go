package ice

import (
	"github.com/aloxc/goice/utils"
	"github.com/boltdb/bolt"
	_ "github.com/sirupsen/logrus"
	"strconv"
	"sync/atomic"
	"time"
)

//统计相关代码
var (
	db, err = bolt.Open("stats.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	curCacheCount int32 = 0
	lastTime int64 = time.Now().UnixNano()

)

const (
	maxCacheCount int32 =  1000
	maxCacheTime = 50
)
type Stat struct {
	Name   []byte `json:"服务",type:"string"`
	Method []byte `json:"方法",type:"string"`
	Time   []byte `json:"统计时间",type:"int64"`
	Spend  []byte `json:"花费",type:"int64"`
}

func (this *Stat) String() string {
	var desc = ""
	desc += "{"
	desc += "\""
	desc += "服务"
	desc += ":"
	desc += "\""
	desc += string(this.Name)
	desc += "\","

	desc += "\""
	desc += "方法"
	desc += ":"
	desc += "\""
	desc += string(this.Method)
	desc += "\","

	desc += "\""
	desc += "统计时间"
	desc += ":"
	desc += "\""
	desc += strconv.FormatInt(utils.BytesToInt64(this.Time),10)
	desc += "\","

	desc += "\""
	desc += "花费"
	desc += ":"
	desc += "\""
	desc += strconv.FormatInt(utils.BytesToInt64(this.Spend),10)
	desc += "\"}"
	return desc
}
func Add(name, method string, spendTime int64) {
	db.Update(func(tx *bolt.Tx) error {
		nbyte := []byte(name)
		mbyte := []byte(method)
		nameB := tx.Bucket(nbyte)
		if nameB == nil {
			if nameB, err = tx.CreateBucket(nbyte); err != nil {
				return err
			}
		}
		methodB := nameB.Bucket(mbyte)
		if methodB == nil {
			if methodB, err = nameB.CreateBucket(mbyte); err != nil {
				return err
			}
		}
		return methodB.Put(utils.Int64ToBytes(time.Now().UnixNano()), utils.Int64ToBytes(spendTime))
	})
	atomic.AddInt32(&curCacheCount,1)
	//nano := time.Now().UnixNano()
	//if (curCacheCount> maxCacheCount) || (nano- maxCacheTime * 1000000 > lastTime){
	//	stats := GetAll(true)
	//	log.Infof("开始执行统计，缓存中有[%d]数据.[%t][%t]", len(stats),(curCacheCount> maxCacheCount),(nano- maxCacheTime * 1000000 > lastTime))
	//	//log.Infof("开始执行统计，缓存中有[%d]数据.count=[%d],maxCount=[%d],nano=[%d],time=[%d]", len(stats),curCacheCount,maxCacheCount,nano,( maxCacheTime * 1000000 + lastTime))
	//	lastTime = nano
	//	curCacheCount = 0
	//	//fmt.Println(stats)
	//}
}

func GetAll(delete bool) []Stat {
	var stats = []Stat{}
	db.Update(func(tx *bolt.Tx) error {
		tx.ForEach(func(name []byte, rootBucket *bolt.Bucket) error {
			rootBucket.ForEach(func(method, _ []byte) error {
				inb := rootBucket.Bucket(method)
				inb.ForEach(func(time, spend []byte) error {
					stats = append(stats, Stat{
						Name:   name,
						Method: method,
						Time:   time,
						Spend:  spend,
					})
					if delete {
						inb.Delete(time)
					}
					return nil
				})
				return nil
			})
			return nil
		})
		return nil
	})
	return stats
}
