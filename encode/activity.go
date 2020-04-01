package encode

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivField1 = "inputString"
	ivField2 = "salt"
	ovResult = "result"
)

var activityLog = logger.GetLogger("tibco-activity-encode")

type EncodeActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EncodeActivity{metadata: metadata}
}

func (a *EncodeActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *EncodeActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing encode activity")
	//Read Inputs
	if context.GetInput(ivField1) == nil {
		// Input string is not configured
		// return error to the engine
		return false, activity.NewError("Input string is not configured", "ENCODE-4001", nil)
	}
	field1v := context.GetInput(ivField1).(string)

	if context.GetInput(ivField2) == nil {
		// Salt is not configured
		// return error to the engine
		return false, activity.NewError("Salt is not present", "ENCODE-4002", nil)
	}
	field2v := context.GetInput(ivField2).(string)

	//Set output
	// Create a new HMAC by defining the hash type and the key (as byte array)

	h := hmac.New(sha512.New, []byte(field2v))
	h.Write([]byte(field1v))
	sha := hex.EncodeToString(h.Sum(nil))
	context.SetOutput(ovResult, sha)
	return true, nil
}
