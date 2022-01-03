package models

type Pet struct {
	ID      string `json:"id" dynamodbav:"id"`
	Name    string `json:"name" dynamodbav:"name"`
	Age     int    `json:"age" dynamodbav:"age"`
	TutorID int    `json:"tutorID" dynamodbav:"tutor_id"`
}
