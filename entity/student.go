package entity

type Student struct {
	FirstName  string `valid:"alpha~must be alphabet,required~First name required"`
	LastName   string `valid:"alpha~must be alphabet,required~Last name required"`
	Email      string `valid:"email~must be email,required~Email required"`
	Age        int    `valid:"int~must be integer,required~Age required"`
	Student_id string `valid:"alphanum~must be alphanum,required~Student_id required"`
}
