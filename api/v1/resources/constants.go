package v1

const MsgSuccess = "Success"

var SuccessMetaResponse map[string]interface{} = map[string]interface{}{
	"status": 200,
	"msg":    MsgSuccess,
}
