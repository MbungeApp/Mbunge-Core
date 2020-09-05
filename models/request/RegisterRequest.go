/*
 * Copyright (c) 2020.
 * MbungeApp Inc all rights reserved
 */

package request

import (
	"encoding/json"
	"github.com/MbungeApp/mbunge-core/models/db"
)

func UnmarshalRegisterRequest(data []byte) (RegisterRequest, error) {
	var r RegisterRequest
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *RegisterRequest) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type RegisterRequest struct {
	User     db.User `json:"user"`
	Type     string  `json:"type" form:"type"`
	FcmToken string  `json:"fcm_token" form:"fcm_token"`
}

/*
{
   "user":{
      "first_name":"patrick",
      "last_name":"waweru",
      "email_address":"example@gmail.com",
      "password":"123456789",
      "phone_number":"123456789",
      "county":"nakuru",
      "date_birth":"iufidfdjfjd",
      "gender":"1",
      "profile_pic":"df"
   },
   "type":"mobile",
   "fcm_token":"1234567dv"
}
*/
