package main

import (
	"fmt"

	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id_       bson.ObjectId `bson:"_id"`
	Username  string        `bson:"username"`
	Friends   int           `bson:"friends"`
	LastLogin time.Time     `bson:"loastLogin"`
	Interests []string      `bson:"interests"`
}

func main() {
	//connect server
	session, err := mgo.Dial("")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("connect success.")
	defer session.Close()

	//connect db
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("gostudy")
	fmt.Println(db.Name)
	//switch collection
	c := db.C("user")
	fmt.Println(c.FullName)
}
