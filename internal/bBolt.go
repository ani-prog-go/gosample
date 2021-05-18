package internal

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"time"

	bolt "go.etcd.io/bbolt"
)

func init() {

}
func getBoltDB() (*bolt.DB, error) {
	if db, err := bolt.Open("my.db", 0600, &bolt.Options{Timeout: 2 * time.Second}); err != nil {
		return nil, err
	} else {
		return db, nil
	}
}
func Bolt() {
	// Откройте файл данных my.db в вашем текущем каталоге..
	//Он будет создан, если его не существует.
	//db, err := bolt.Open("my.db", 0600, nil)
	// timeout если по какимто причинам файл базы данных занят
	var db *bolt.DB

	if dbx, err := getBoltDB(); err != nil {
		log.Fatal(err.Error())
	} else {
		db = dbx
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket
		bucket, err1 := tx.CreateBucketIfNotExists([]byte("configs")) //CreateBucketIfNotExists
		if err1 != nil {
			return fmt.Errorf("create bucket: %s", err1)
		}
		err2 := bucket.Put([]byte("answer"), []byte("42"))
		if err2 != nil {
			return fmt.Errorf("error programm: %s", err2)

		}

		return nil
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("configs"))
		v := b.Get([]byte("answer"))
		PrintColor(ColorGreen, fmt.Sprintf("Прочитали из базы: %s\n", v))
		return nil
	})
}

type User struct {
	Id   int
	Name string
	Age  int
}

////
func BoltJson() {

	user := User{
		Id:   20,
		Name: "Arkadii",
		Age:  50,
	}
	// Откройте файл данных my.db в вашем текущем каталоге..
	//Он будет создан, если его не существует.
	//db, err := bolt.Open("my.db", 0600, nil)
	// timeout если по какимто причинам файл базы данных занят
	var db *bolt.DB
	if dbx, err := getBoltDB(); err != nil {
		log.Fatal(err.Error())
	} else {
		db = dbx
	}
	defer db.Close()

	if err := saveDb(db, 25, user); err != nil {
		fmt.Printf("create bucket: %s", err.Error())
	} else {
		PrintColor(ColorGreen, fmt.Sprintf("Записали в базу: %v", user))
	}

	if user, err := findDb(db, 20); err == nil {
		if user == (User{}) {
			PrintColor(ColorRed, fmt.Sprintf("Не найдено: ид %d", 26))
		} else {
			PrintColor(ColorGreen, fmt.Sprintf("Прочитали из базы: %s лет %d ид %d", user.Name, user.Age, user.Id))
		}
	} else {
		fmt.Println(err)
	}
}

// itob returns an 8-byte big endian representation of v.
func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func findDb(db *bolt.DB, id int) (User, error) {
	var u User
	var v []byte
	//var b *bolt.Tx
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("users"))
		v = b.Get(itob(id))
		// прокрутить все ключи в ведре id мы храним в []byte оно uint64 8 байт
		b.ForEach(func(k, val []byte) error {
			fmt.Printf("key=%d, value=%s\n", binary.BigEndian.Uint64(k), val)
			return nil
		})
		return nil
	})
	if err != nil {
		return u, err
	}
	if v == nil {
		return u, nil
	}

	if err := json.Unmarshal(v, &u); err != nil {
		return u, err
	}
	return u, nil
}
func saveDb(db *bolt.DB, id int, user User) error {
	if err := db.Update(func(tx *bolt.Tx) error {
		var bucket *bolt.Bucket
		var err error
		if bucket, err = tx.CreateBucketIfNotExists([]byte("users")); err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		// Generate ID for the user.
		// This returns an error only if the Tx is closed or not writeable.
		// That can't happen in an Update() call so I ignore the error check.
		// id, _ := bucket.NextSequence()
		// user.ID = int(id)

		// Marshal user data into bytes.
		if buf, err := json.Marshal(user); err != nil {
			return err
		} else {
			if err := bucket.Put(itob(id), buf); err != nil {
				return fmt.Errorf("error programm: %s", err)

			}
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
