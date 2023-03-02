// Copyright 2023 SphereEx Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"sync"

	bolt "go.etcd.io/bbolt"
)

type Config struct {
	Path string
}

type BoltStore struct {
	Config
	db   *bolt.DB
	lock *sync.Mutex
}

func NewBoltStore() (*BoltStore, error) {
	s := &BoltStore{
		lock: &sync.Mutex{},
	}

	if err := s.init(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *BoltStore) init() error {
	var err error

	if s.Path == "" {
		s.Path = "./bolt.db"
	}
	s.db, err = bolt.Open(s.Path, 0600, nil)

	if err != nil {
		return err
	}
	return nil
}

func (s *BoltStore) Close() error {
	return s.db.Close()
}

func (s *BoltStore) GetDB() *bolt.DB {
	return s.db
}

func (s *BoltStore) CreateBucket(name string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return fmt.Errorf("create bucket err: %w", err)
		}
		return nil
	})
}

type BoltStoreValue interface {
	SetUpdateAt()
}

func (s *BoltStore) Put(name string, key string, value interface{}) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		s.lock.Lock()
		defer s.lock.Unlock()

		b, err := tx.CreateBucketIfNotExists([]byte(name))
		if err != nil {
			return err
		}

		value.(BoltStoreValue).SetUpdateAt()
		valueBytes, err := json.Marshal(value)
		if err != nil {
			return err
		}

		return b.Put([]byte(key), valueBytes)
	})
}

func (s *BoltStore) Get(name string, key string, value interface{}) error {
	return s.db.View(func(tx *bolt.Tx) error {
		s.lock.Lock()
		defer s.lock.Unlock()

		bucket := tx.Bucket([]byte(name))
		if bucket == nil {
			return bolt.ErrBucketNotFound
		}
		data := bucket.Get([]byte(key))

		if data == nil {
			return nil
		}

		return json.Unmarshal(data, value)
	})
}

func (s *BoltStore) Delete(name string, key string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))
		return b.Delete([]byte(key))
	})
}

func (s *BoltStore) GetAllPrefix(name string, prefix string, vtype interface{}) (map[string]interface{}, error) {
	var res = make(map[string]interface{})
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))
		c := b.Cursor()
		prefix := []byte(prefix)
		for k, v := c.Seek(prefix); k != nil && bytes.HasPrefix(k, prefix); k, v = c.Next() {
			elem := reflect.New(reflect.TypeOf(vtype).Elem())
			value := elem.Interface()

			if err := json.Unmarshal(v, value); err != nil {
				return err
			}
			res[string(k)] = value
		}
		return nil
	})

	return res, err
}

func (s *BoltStore) GetAll(name string, vtype interface{}) (map[string]interface{}, error) {
	var res = make(map[string]interface{})
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(name))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			elem := reflect.New(reflect.TypeOf(vtype).Elem())
			value := elem.Interface()

			if err := json.Unmarshal(v, value); err != nil {
				return err
			}
			res[string(k)] = value
		}
		return nil
	})

	return res, err
}
