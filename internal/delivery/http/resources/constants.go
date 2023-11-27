package v1

const MsgSuccess = "Success"
const MsgError = "Error"

var SuccessMetaResponse = map[string]interface{}{
	"status": 200,
	"msg":    MsgSuccess,
}
