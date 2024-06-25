package db

import (
	"encoding/binary"
	"time"

	"github.com/boltdb/bolt"
)

var taskBucket = []byte("tasks")
var db *bolt.DB

type Task struct {
	Key   int
	Value string
}

func Init(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(taskBucket)
		return err
	})
}

func CreateTask(task string) (int, error) {
	var id int
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)
		return bucket.Put(key, []byte(task))
		// return bucket.Put(itob(id64), []byte(task))
	})
	if err != nil {
		return -1, err
	}

	return id, nil
}

func itob(id int) []byte {
	b := make([]byte, 8)
	// binary.BigEndian.PutUint64(b, id)
	binary.BigEndian.PutUint64(b, uint64(id))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}