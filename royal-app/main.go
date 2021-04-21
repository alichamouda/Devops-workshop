package main

var RoyalMembers []RoyalMember

func initRoyals() {
	RoyalMembers = []RoyalMember{
		RoyalMember{Id: "1", Name: "Elizabeth Alexandra Mary Windsor", Reign: "6 February 1952 – present", Birthdate: "21 April 1926"},
		RoyalMember{Id: "2", Name: "Albert Frederick Arthur George", Reign: "11 December 1936 – 6 February 1952", Birthdate: "14 December 1895"},
		RoyalMember{Id: "3", Name: "George Frederick Ernest Albert", Reign: "6 May 1910 – 20 January 1936", Birthdate: "3 June 1865"},
	}
}

func main() {
	initRoyals()
	handleRequests()
}
