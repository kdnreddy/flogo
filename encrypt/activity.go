package encrypt

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
	svField1 = "consumerKey"
	svField2 = "consumerSecret"
	svField3 = "accessToken"
	svField4 = "accessSecret"
	ovResult = "result"
)

type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

var activityLog = logger.GetLogger("tibco-activity-encrypt")

type EncryptActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EncryptActivity{metadata: metadata}
}

func (a *EncryptActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *EncryptActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing encrypt activity")
	//Read Inputs
	if context.GetInput(ivField1) == nil {
		// Input string is not configured
		// return error to the engine
		return false, activity.NewError("Input string is not configured", "ENCRYPT-4001", nil)
	}
	field1v := context.GetInput(ivField1).(string)

	if context.GetInput(svField1) == nil {
		// Salt is not configured
		// return error to the engine
		return false, activity.NewError("Salt is not present", "ENCRYPT-4002", nil)
	}
	field1s := context.GetInput(svField1).(string)

	if context.GetInput(svField2) == nil {
		// Salt is not configured
		// return error to the engine
		return false, activity.NewError("Salt is not present", "ENCRYPT-4002", nil)
	}
	field2s := context.GetInput(svField2).(string)

	if context.GetInput(svField3) == nil {
		// Salt is not configured
		// return error to the engine
		return false, activity.NewError("Salt is not present", "ENCRYPT-4002", nil)
	}
	field3s := context.GetInput(svField3).(string)

	if context.GetInput(svField4) == nil {
		// Salt is not configured
		// return error to the engine
		return false, activity.NewError("Salt is not present", "ENCRYPT-4002", nil)
	}
	field4s := context.GetInput(svField4).(string)

	creds := Credentials{
		ConsumerKey:       field1s,
		ConsumerSecret:    field2s,
		AccessToken:       field3s,
		AccessTokenSecret: field4s,
	}

	fmt.Printf("%+v\n", creds)
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)

	log.Printf("User's ACCOUNT:\n%+v\n", user)
	/* client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Println(err)
	} */

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
