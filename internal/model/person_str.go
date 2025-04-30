package main

import "time"

type Person struct {
	id         int
	name       string
	surname    string
	patronymic string
	sex        string
	national   string
	creation   time.Time
	update     time.Time
}
