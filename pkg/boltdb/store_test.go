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
	"fmt"
	"log"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type TestUser struct {
	Id   int
	Name string
}

func (t *TestUser) SetId(id int) {
	t.Id = id
}

func (t *TestUser) SetName(name string) {
	t.Name = name
}

func (t *TestUser) SetUpdateAt() {}

var _ = Describe("Store", func() {

	var (
		db       *BoltStore
		err      error
		testUser TestUser
	)

	var (
		users1 = []*TestUser{
			{
				Id:   0,
				Name: "test",
			}, {
				Id:   1,
				Name: "test1",
			},
		}
	)

	BeforeEach(func() {
		db, err = NewBoltStore()
		err = db.CreateBucket("user")
		for i, v := range users1 {
			if err := db.Put("user", "user-"+fmt.Sprint(i), v); err != nil {
				log.Fatal(err)
			}
		}
		err = db.Get("user", "user-1", &testUser)
		defer db.Close()
	})

	Context("Test BoltDB", func() {
		It("should get test data correctly", func() {
			Expect(testUser.Name).To(Equal("test1"))
		})

		It("shouldn't error", func() {
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
