package gdsb

/*
	Copyright 2018 Rewati Raman rewati.raman@gmail.com https://github.com/rewati/gdsb

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
	limitations under the License.
*/
import (
	"encoding/json"
	"log"
	"sync"
	"time"

	uuid "github.com/satori/go.uuid"
)

//Hold will hold execution
func Hold() {
	var wg = &sync.WaitGroup{}
	wg.Add(1)
	wg.Wait()
}

//UUIDstring to generate UUID
func UUIDstring() string {
	u2, err := uuid.NewV4()
	if err != nil {
		log.Println("Error while generating UUID. Erro: ", err)
		return ""
	}
	return u2.String()
}

//ToJSON takes object o and returns string and error
func ToJSON(o interface{}) (string, error) {
	j, e := json.Marshal(o)
	if e == nil {
		return string(j), e
	}
	return "", e
}

//UTCNow is to generate utc time
func UTCNow() time.Time {
	loc, _ := time.LoadLocation("UTC")
	return time.Now().In(loc)
}

//UTCMilisec is to generate utc time milisecond
func UTCMilisec() int64 {
	v := UTCNow().UnixNano() / 1000000
	return v
}
