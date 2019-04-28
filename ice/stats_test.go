package ice

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

func TestAdd(t *testing.T) {
	name := "xname"
	for i := 0; i < 10; i++ {
		method := "xmethod-" + strconv.Itoa(i)
		for j := 0; j < 10; j++ {
			var x = rand.Int63n(1000)
			//log.Info(name, "  ", method, "  ", x)
			Add(name, method, x)
		}
	}
}

func TestGetAll(t *testing.T) {
	all := GetAll(false)
	for _,stat := range all{
		fmt.Println(stat.String())
	}

	//db.Close()
	//db, err = bolt.Open("stats.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	//db.Update(func(tx *bolt.Tx) error {
	//	root, err := tx.CreateBucketIfNotExists([]byte("goice"))
	//	if err != nil {
	//		log.Println(err)
	//		return err
	//	}
	//	log.Info("准备开始")
	//	for i := 0; i < 10; i++ {
	//		userTableName := "user-" + strconv.Itoa(i)
	//		userTable, err := root.CreateBucketIfNotExists([]byte(userTableName))
	//		if err != nil {
	//			return err
	//		}
	//		for j := 0; j < 10; j++ {
	//			userTable.Put([]byte("name-"+strconv.Itoa(j)), []byte("aloxc-"+strconv.Itoa(j)))
	//		}
	//	}
	//
	//	tx.ForEach(func(name []byte, bucket *bolt.Bucket) error {
	//		fmt.Println("bucket : " ,string(name))
	//		bucket.ForEach(func(k, v []byte) error {
	//			inb := bucket.Bucket(k)
	//			fmt.Printf("key = %s,value=%s\n",k,v)
	//			inb.ForEach(func(k1, v1 []byte) error {
	//				fmt.Printf("key = %s,value=%s\n",k1,v1)
	//				return nil
	//			})
	//			return nil
	//		})
	//		return nil
	//
	//
	//	})
	//	//rootCursor := root.Cursor()
	//	//for rootKey, rootValue := rootCursor.First(); rootKey != nil; rootKey, rootValue = rootCursor.Next() {
	//	//	//userCursor := rootCursor.Bucket().Cursor()
	//	//	fmt.Printf("rootKey=%s rootValue = %s\n", rootKey, rootValue)
	//	//	userBucket.ForEach(func(userKey, userValue []byte) error {
	//	//			fmt.Printf("userKey=%s userValue = %s\n", userKey, userValue)
	//	//		return nil
	//	//	})
	//		//for userKey, userValue := userCursor.First(); userKey != nil; userKey, userValue = userCursor.Next() {
	//		//	fmt.Printf("userKey=%s userValue = %s\n", userKey, userValue)
	//		//
	//		//}
	//	//}
	//	return nil
	//})
}
