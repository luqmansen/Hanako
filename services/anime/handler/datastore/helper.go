package datastore

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"unicode"
)

func queryFormater(request map[string]interface{}) bson.M {
	var query bson.M

	if request["title"] != "" && request["type"] != "" {
		query = bson.M{
			"$and": []interface{}{
				bson.M{
					"$text": bson.M{"$search": request["title"]},
				},
				bson.M{"type": request["type"]},
			},
		}
	} else if request["title"] != "" {
		// Format string if it is contain more than one word to be queried as title
		s := fmt.Sprintf("%v", request["title"])
		for _, v := range s {
			if unicode.IsSpace(v) {
				s = "\"" + s + "\""
				break
			}
		}
		query = bson.M{
			"$text": bson.M{
				"$search": s,
			},
		}
	} else {
		query = bson.M{
			"type": request["type"],
		}
	}
	return query
}
