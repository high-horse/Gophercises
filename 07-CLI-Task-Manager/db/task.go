package db

import (
	"encoding/binary"

	"github.com/boltdb/bolt"
)



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

func AllTasks() ([]Task, error) {
	var tasks []Task

	err := db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)

		curser := bucket.Cursor()
		for k, v := curser.First(); k != nil; k, v = curser.Next() {
			tasks = append(tasks, Task{
				Key:   btoi(k),
				Value: string(v),
			})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func DeleteTask(key int) error {
	err := db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(taskBucket)
		
		return bucket.Delete(itob(key))
	})

	return err
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