package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var (
	username = "johndoe"
	id = "id"
	secret = "secret"
	sub = "sub = \"sub\""
	userpoolid = "userpollid"
	attr = "email"
	attributesToGet = []*string{&attr}
)

func main() {
	// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#specifying-credentials
	creds := credentials.NewStaticCredentials(id,secret,"")
	// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html
	// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html#creating-sesstion
	mySession := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-2"),Credentials:creds}))

	cip := cognitoidentityprovider.New(mySession)
	//agui := cognitoidentityprovider.AdminGetUserInput{UserPoolId:&userpoolid,Username:&username}
	lui := cognitoidentityprovider.ListUsersInput{UserPoolId:&userpoolid, Filter: &sub, AttributesToGet:attributesToGet}

	//user, err := cip.AdminGetUser(&agui)
	user, err:= cip.ListUsers(&lui)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(user)
}
