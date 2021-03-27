package model

import "time"

type Article struct {
	title		string
	id 			string
	link 		string
	poster		string
	vote		int
	publishTime time.Time
}