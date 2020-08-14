//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import "time"

// MockDB used to simulate a database model
type MockDB struct{}

// Get only returns an empty string, as this is only for demonstration purposes
func (*MockDB) Get(key string) (string, error) {
	d, _ := time.ParseDuration("20ms")
	time.Sleep(d)
	return "", nil
}

// GetMockDB returns an instance of MockDB
func GetMockDB() *MockDB {
	return &MockDB{}
}
