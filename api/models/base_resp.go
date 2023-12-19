/*
Work for: Simpaix.net
Author: Z3NTL3
License: GNU
*/
package models

/*
Telegram API

The response contains a JSON object, which always has a Boolean field 'ok' and may have an optional String field 'description' with a human-readable description of the result. If 'ok' equals True, the request was successful and the result of the query can be found in the 'result' field. In case of an unsuccessful request, 'ok' equals false and the error is explained in the 'description'. An Integer 'error_code' field is also returned, but its contents are subject to change in the future. Some errors may also have an optional field 'parameters' of the type ResponseParameters, which can help to automatically handle the error.
*/
type BaseResp struct {
	Success     bool   `json:"ok"`                    // always present
	Description string `json:"description,omitempty"` // optional
}

type Result struct {
	Result interface{} // type assertion towards
}

type API_Resp[R BaseWrapI] struct {
	BaseResp
	Result R
}

type BaseWrapI interface {
	interface{}
}
