package db

import (
	"encoding/json"
	"time"

	"github.com/boltdb/bolt"
)

type Work struct {
	ID          int
	Name        string
	Completed   bool
	CompletedAt time.Time
}

var bucket_name = []byte("work")
var bucket *bolt.DB


func InitBucket(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return err
	}
	
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(bucket_name)
		return err
	})
}

func CreateWork(task string) (int, error) {
	var id int
	work := Work{
		Name:      task,
		Completed: false,
	}
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket_name)
		id64, _ := bucket.NextSequence()
		id = int(id64)
		key := itob(id)

		workBytes, err := json.Marshal(work)
		if err != nil {
			return err
		}
		return bucket.Put(key, workBytes)
	})
	if err != nil {
		return -1, err
	}
	return id, nil
}

func ListWork() ([]Work, error) {
	var works []Work
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket_name)
		c := bucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			var work Work
			err := json.Unmarshal(v, &work)
			if err != nil {
				return err
			}
			works = append(works, work)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return works, nil
}

func DeleteWork(key int) error {
	return db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket_name)
		return bucket.Delete(itob(key))
	})
}

func UpdateWork(key int, status bool) error {
	var work *Work
	var err error

	work, err = GetWork(key)

	if err != nil {
		return err
	}

	err = db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket_name)

		// work.ID = key
		work.Completed = status
		work.CompletedAt = time.Now()

		workBytes, err := json.Marshal(work)
		if err != nil {
			return err
		}
		return bucket.Put(itob(key), workBytes)
	})

	if err != nil {
		return err
	}
	return nil
}

func GetWork(key int) (*Work, error) {
	var work *Work
	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(bucket_name)
		v := bucket.Get(itob(key))
		err := json.Unmarshal(v, &work)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return work, nil
}
