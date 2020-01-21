package twitterbot

import (
	"fmt"
	"log"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	ivField1 = "inputString"
	ovResult = "result"
	svField1 = "consumerKey"
	svField2 = "consumerSecret"
	svField3 = "accessToken"
	svField4 = "accessTokenSecret"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var activityLog = logger.GetLogger("tibco-activity-twitterbot")

type TwitterActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &TwitterActivity{metadata: metadata}
}

func (a *TwitterActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	// Pass in your consumer key (API Key) and your Consumer Secret (API Secret)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	// Pass in your Access Token and your Access Token Secret
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	// Verify Credentials
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

func (a *TwitterActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing encrypt activity")
	//Read Inputs
	if context.GetInput(ivField1) == nil {
		// Input string is not configured
		// return error to the engine
		return false, activity.NewError("Input string is not configured", "ENCRYPT-4001", nil)
	}
	field1v := context.GetInput(ivField1).(string)
	sField1 := context.GetInput(svField1).(string)
	sField2 := context.GetInput(svField2).(string)
	sField3 := context.GetInput(svField3).(string)
	sField4 := context.GetInput(svField4).(string)

	//Set output
	// Create a new HMAC by defining the hash type and the key (as byte array)

	creds := Credentials{
		AccessToken:       sField1,
		AccessTokenSecret: sField2,
		ConsumerKey:       sField3,
		ConsumerSecret:    sField4,
	}

	fmt.Printf("%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	}

	// Print out the pointer to our client
	// for now so it doesn't throw errors
	tweet, resp, err := client.Statuses.Update(field1v, nil)
	if err != nil {
		log.Println(err)
	}
	log.Printf("%+v\n", resp)
	log.Printf("%+v\n", tweet)

	/* h := hmac.New(sha512.New, []byte(field2v))
	h.Write([]byte(field1v))
	sha := hex.EncodeToString(h.Sum(nil)) */
	context.SetOutput(ovResult, resp)
	return true, nil
}
