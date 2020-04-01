package encode

import (
	"io/ioutil"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}
	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())
	//setup attrs
	tc.SetInput("inputString", "HelloWorld")
	tc.SetInput("salt", "mysecret")
	_, err := act.Eval(tc)
	assert.Nil(t, err)
	result := tc.GetOutput("result")
	assert.Equal(t, result, "fcf8f68a842757346a83f867de9863dd8864fd5e7cf5e333cd19193c2411dc49fcfd081d3ca2a1804c85358bde4119252ad722c79e40ae3bb950ad986d16f8ca")
}
