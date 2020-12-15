//	Otsimo Summer Talent Camp Project
/*
At the beginning of the process, the 4th year university students from different departments
will apply to the camp. After the selection progress, the interns will be given a task to work
on during their internship. Among these interns, some will be hired as full-time team members
at the end of the summer. To make the selection process simpler for the team, we would like to
develop a platform that provides a way to manage candidates, their applications, and appointments
relating to them easily and quickly. On this platform, the assigned team member will be able
to quickly assess the candidates and see if there is anything that needs their attention.
They will be able to accept or reject the application, find out if they have any appointments
with the applicants, etc.
*/

//install mongodb for linux
//sudo apt install mongodb-clients

//use mongod in shell to install
//sudo apt install mongodb-server-core

//connect to mongodb shell firstly "mongodb" on shell and then open the other shell
//and write "mongo" then enter.

//mongodump
//mongodump --host localhost:27017 --archive=dump.gz --gzip --db Otsimo//

//mangorestore
//in "data.gz" directory
//mongorestore --gzip --archive=data.gz --nsFrom "$Otsimo.*" --nsTo "$data.*" //

//package clause
package main

//import statement
import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//my code
func main() {

	// _id mongodb tarafından atanıyor dikkat et
	/*aday := Candidate{
		first_name: "okan",
		last_name:  "özşahin",
		email:      "okan@gmail.com",
		department: "Design",
		university: "ODTU",
		experience: true,
	}
	sonuc, err := CreateCandidate(aday)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sonuc.first_name, " ", sonuc.last_name, " ", sonuc.assignee)
	}*/

	/*adayım, err := ReadCandidate("5b758c7d51d9590001def631")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(adayım.first_name, " ", adayım.last_name, " ", adayım.assignee)
	}*/

	//fmt.Println(completeMeeting("5b758c7d51d9590001def631"))
	//DeleteCandidate("5fd87b4c06fbdac6461d6cc1")
	//FindAssigneeIDByName("Zafeer")
	//fmt.Println(FindAssigneeIDByName("Mehmet"))
	//fmt.Println(AcceptCandidate("5b7587a351d9590001def628"))
	/*time := time.Now()
	ArrangeMeeting("5fd89ef18e7ad22fb1c03438", &time)*/

	//fmt.Println(sonuc.first_name)
}

//Candidates collection stores some required information of Candidates.
type Candidate struct {
	_id        string `json:"_id"`        // Unique hash that identifies candidate.
	first_name string `json:"first_name"` // First name of the candidate.
	last_name  string `json:"last_name"`  // Last name of the candidate.
	email      string `json:"email"`      // Contact email of candidate.
	department string `json:"department"` // Department that candidate applied.
	/*Available values are
	-Marketing
	-Design
	-Development*/
	university string `json:"university"` // University of the candidate.
	experience bool   `json:"experience"` // Candidate has previous working experience or not.
	status     string `json:"status"`     // Status of the candidate. status := "Pending"
	/*Available values are
	-Pending
	-In Progress
	-Denied
	-Accepted*/
	meeting_count int       `json:"meeting_count"` // The order of the next meeting. The maximum meeting count is 4.
	next_meeting  time.Time `json:"next_meeting"`  // Timestamp of the next meeting between the Otsimo team and the candidate.
	assignee      string    `json:"assignee"`      // The id of the Otsimo team member who is responsible for this candidate.
}

// Collection Assignees stores some required information relating to the Assignees.
// Assignee means a team member of Otsimo who is responsible for their own team's summer interns.
type Assignee struct {
	_id        string //Unique hash that identifies an assignee.
	name       string //Name of the assignee.
	department string //Department that assignee works in the Otsimo.
	/*Available values are
	-Marketing
	-Design
	-Development*/
}

// this function creates new candidates and control some statement about
// candidate to correctly insert database.
func CreateCandidate(candidate Candidate) (Candidate, error) {

	// define candidate to insert database
	cand := Candidate{
		_id:           primitive.NewObjectID().Hex(),
		first_name:    candidate.first_name,
		last_name:     candidate.last_name,
		email:         candidate.email,
		department:    candidate.department,
		university:    candidate.university,
		experience:    candidate.experience,
		status:        "Pending",
		meeting_count: 0,
		assignee:      candidate.assignee,
	}

	/* ASSIGNEES'S DATA */
	//{ "_id" : "5bb6360e55c98300013a087b", "name" : "Sercan", "department" : "Development" }
	//{ "_id" : "5bb6368f55c98300013a087d", "name" : "Can", "department" : "Development" }
	//{ "_id" : "5c052d4af410a50001d0c76b", "name" : "Duygu", "department" : "Development" }
	//{ "_id" : "5bf92a19f410a50001d0adb3", "name" : "Elif", "department" : "Marketing" }
	//{ "_id" : "5bfc1a59f410a50001d0b332", "name" : "Ali", "department" : "Marketing" }
	//{ "_id" : "5c18ae31a7948900011168b9", "name" : "Mehmet", "department" : "Design" }
	//{ "_id" : "5c18ad7ea7948900011168b7", "name" : "Murat", "department" : "Design" }
	//{ "_id" : "5c191acea7948900011168d4", "name" : "Zafer", "department" : "CEO" }

	//this is not good a way to dynamic programming but because of easy applicability I did it that
	if cand.department == "Development" {
		// if random num 0 then assigne Sercan or 1 assgine Can else assignee duygu
		random := rand.Intn(3 - 0)
		if random == 0 {
			cand.assignee = "5bb6360e55c98300013a087b"
		} else if random == 1 {
			cand.assignee = "5bb6368f55c98300013a087d"
		} else {
			cand.assignee = "5c052d4af410a50001d0c76b"
		}
	} else if cand.department == "Marketing" {
		if rand.Intn(2-0) == 1 {
			cand.assignee = "5bf92a19f410a50001d0adb3" // if random num 1 then assigne elif else Ali
		} else {
			cand.assignee = "5bfc1a59f410a50001d0b332"
		}
	} else if cand.department == "Design" {
		if rand.Intn(2-0) == 1 {
			cand.assignee = "5c18ae31a7948900011168b9" // if random num 1 then assigne mehmet else murat
		} else {
			cand.assignee = "5c18ad7ea7948900011168b7"
		}
	}

	departments := [3]string{"Marketing", "Design", "Development"}

	//check departments for correctly
	for _, b := range departments {
		if b == cand.department {
			collection, ctx := connectDbToCollection("Candidates")
			candidateResult, err := collection.InsertOne(ctx, bson.D{
				{"_id", cand._id},
				{"first_name", cand.first_name},
				{"last_name", cand.last_name},
				{"email", cand.email},
				{"department", cand.department},
				{"university", cand.university},
				{"experience", cand.experience},
				{"status", cand.status},
				{"meeting_count", cand.meeting_count},
				{"assignee", cand.assignee},
			})
			fmt.Println(candidateResult.InsertedID)
			return cand, err
		}
	}
	return cand, errors.New("invalid department")

}

// find candidate with candidate's id and return candidate's info.
// and check if there is any error.
func ReadCandidate(_id string) (Candidate, error) {

	collection, ctx := connectDbToCollection("Candidates")
	//fmt.Println(collection)

	// find candidate with _id to return
	filterCursor, err := collection.Find(ctx, bson.M{"_id": _id})
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(err)

	// assign candidate's value in bson.M format to easily read.
	var candidateFiltered []bson.M
	if err = filterCursor.All(ctx, &candidateFiltered); err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(candidateFiltered))
	cand := Candidate{}

	// check if database dont have any candidate with given id then throw error "not found cadidate".
	if len(candidateFiltered) == 0 {
		return cand, errors.New("not found Candidate")
	}

	// I did this two variable because of GO throws panic.
	meetingCount, _ := candidateFiltered[0]["meeting_count"].(int)
	nextMeeting, _ := candidateFiltered[0]["next_meeting"].(time.Time)

	candidate := Candidate{
		_id:           candidateFiltered[0]["_id"].(string),
		first_name:    candidateFiltered[0]["first_name"].(string),
		last_name:     candidateFiltered[0]["last_name"].(string),
		email:         candidateFiltered[0]["email"].(string),
		department:    candidateFiltered[0]["department"].(string),
		university:    candidateFiltered[0]["university"].(string),
		experience:    candidateFiltered[0]["experience"].(bool),
		status:        candidateFiltered[0]["status"].(string),
		meeting_count: meetingCount,
		next_meeting:  nextMeeting,
		assignee:      candidateFiltered[0]["assignee"].(string),
	}
	return candidate, err

}

// select a candidate with id and meeting_count added up 1.
// if candidate's meeting_count is 4 then it's status be "In Progress" and assignee is Zafer who is CEO of Otsimo.
func completeMeeting(_id string) error {

	collection, ctx := connectDbToCollection("Candidates")

	var meet bson.M
	if err := collection.FindOne(ctx, bson.M{"_id": _id}).Decode(&meet); err != nil {
		//log.Fatal(err)
		return errors.New("not found Candidate")
	}

	fmt.Println(meet["meeting_count"])

	mCount, _ := meet["meeting_count"].(int32)
	//fmt.Println(mCount)

	//when meeting completed add +1 to meeting_count for candidate
	mCount = mCount + 1
	//fmt.Println(mCount)

	//update candidate's meeting_count
	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": _id},
		bson.D{
			{"$set", bson.D{{"meeting_count", mCount}}},
		},
	)

	// mCount == 4 that's mean status in progress and candidate's assignee be Zafer
	if mCount >= 4 {
		resultStatus, err := collection.UpdateOne(
			ctx,
			bson.M{"_id": _id},
			bson.D{
				{"$set", bson.D{{"status", "In Progress"}, {"assignee", "5c191acea7948900011168d4"}}},
			},
		)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Updated %v Documents!\n", resultStatus.ModifiedCount)
	}
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return err

}

// delete candidate from database with id
func DeleteCandidate(_id string) error {

	collection, ctx := connectDbToCollection("Candidates")

	//check candidate with id then delete it from database
	result, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
	return err

}

// with candidate's id arrange meeting to candidate and assign a time to meeting.
// because of in task no function for connect to database, I have written all connection coding stuff again and again.
func ArrangeMeeting(_id string, nextMeetingTime *time.Time) error {

	var next_meeting time.Time

	next_meeting = *nextMeetingTime

	collection, ctx := connectDbToCollection("Candidates")

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": _id},
		bson.D{
			{"$set", bson.D{{"next_meeting", next_meeting}}},
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return err

}

// candidate's status be "Denied" with candidate's id
func DenyCandidate(_id string) error {

	collection, ctx := connectDbToCollection("Candidates")

	result, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": _id},
		bson.D{
			{"$set", bson.D{{"status", "Denied"}}},
		},
	)
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return err

}

// candidate's status be "Accepted" with candidate's id
func AcceptCandidate(_id string) error {

	collection, ctx := connectDbToCollection("Candidates")

	filterCursor, err := collection.Find(ctx, bson.M{"_id": _id})
	if err != nil {
		log.Fatal(err)
	}

	var candidateFiltered []bson.M
	if err = filterCursor.All(ctx, &candidateFiltered); err != nil {
		log.Fatal(err)
	}

	// check candidate's meeting_cound and if value is equal or greater than four update candidate's
	// status to "Accepted"
	if candidateFiltered[0]["meeting_count"].(int32) >= 4 {
		result, err := collection.UpdateOne(
			ctx,
			bson.M{"_id": _id},
			bson.D{
				{"$set", bson.D{{"status", "Accepted"}}},
			},
		)
		fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
		return err
	} else {
		err = errors.New("not accepted because of meeting count")
	}
	return err

}

// find assignee id with its name.
func FindAssigneeIDByName(name string) string {

	collection, ctx := connectDbToCollection("Assignees")

	// find assignee with name
	filterCursor, err := collection.Find(ctx, bson.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}

	// return datas and assign it in candidateFiltered variable in []bson.M format.
	var candidateFiltered []bson.M
	if err = filterCursor.All(ctx, &candidateFiltered); err != nil {
		log.Fatal(err)
	}
	fmt.Println(err)

	if candidateFiltered == nil {
		return "not found assignee with this name"
	}
	fmt.Println(candidateFiltered[0]["_id"])
	return candidateFiltered[0]["_id"].(string)

}

// function takes collection name with type string and return collection.
func connectDbToCollection(data string) (*mongo.Collection, context.Context) {

	// connect localhost to fetch data
	// we're configuring our client to use the correct URI, but we're not yet connecting to it.
	client, err := mongo.NewClient(options.Client().ApplyURI(
		"mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// we can define a timeout duration that we want to use when trying to connect.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// to fetch data from the requested database then return collection to use data for our needs.
	collection := client.Database("Otsimo").Collection(data)
	return collection, ctx
}
