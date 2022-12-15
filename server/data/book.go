// server/data/book.go
package data

import bookpd "github.com/SeoEunkyo/gRPC-golang/protos/v1"

var BookData = map[int32]*bookpd.BookMessage{
	1: {
		Id:          1,
		Title:       "book1",
		Writer:      "Henry",
		Description: "01012341234",
	},
	2: {
		Id:          2,
		Title:       "book2",
		Writer:      "new jin",
		Description: "make cookie",
	},
	3: {
		Id:          3,
		Title:       "book3",
		Writer:      "grammy",
		Description: "hi Grammy",
	},
}
