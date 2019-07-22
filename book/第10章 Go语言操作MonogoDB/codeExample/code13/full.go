package main

import (
	"fmt"
	"log"

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
	fmt.Println("This is a test to use mgo for go.")

	//connect server
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect success.")
	}
	defer session.Close()

	//connect db
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("hykj001")

	//switch collection
	c := db.C("people")

	//insert
	err = c.Insert(&User{
		Id_:       bson.NewObjectId(),
		Name:      "Jimmy Kuu",
		Age:       33,
		JonedAt:   time.Now(),
		Interests: []string{"Develop", "Movie"},
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("insert success.")
	}

	//select no condition
	var users []User
	c.Find(nil).All(&users)
	fmt.Println(users)

	//select one condition
	c.Find(bson.M{"name": "Jimmy Kuu"}).All(&users) //name not Name
	fmt.Println(users)

	//update alter
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911d109c44bc1a30c9472d")},
		bson.M{"$set": bson.M{
			"name": "Jimmy Gu",
			"age":  34,
		}})
	id := "5a911f559c44bc07a4fc612a"
	objectId := bson.ObjectIdHex(id)
	user := new(User)
	c.Find(bson.M{"_id": objectId}).One(&user)
	fmt.Println(user)

	//update add
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$inc": bson.M{
			"age": -1,
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//add ele
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$push": bson.M{
			"interests": "Golang",
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//del ele
	c.Update(bson.M{"_id": bson.ObjectIdHex("5a911f559c44bc07a4fc612a")},
		bson.M{"$pull": bson.M{
			"interests": "Golang",
		}})
	objectId = bson.ObjectIdHex(id)
	c.FindId(objectId).One(&user)
	fmt.Println(user)

	//remove
	c.Remove(bson.M{"name": "Jimmy Kuu"})
	c.Find(nil).All(&users)
	fmt.Println(users)

}
