// Copyright 2022 zenpk
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package eh

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/pkg/zap"
	"net/http"
	"reflect"
	"strconv"
)

type JSONHandler struct {
	C *gin.Context
	V interface{}
}

// Handle return the given interface with tagged values and error messages
func (h *JSONHandler) Handle(err error, httpCode ...int) {
	ref := reflect.ValueOf(&h.V).Elem()
	refCopy := reflect.New(ref.Elem().Type()).Elem()
	for i := 0; i < ref.Elem().NumField(); i++ {
		tag := ref.Elem().Type().Field(i).Tag.Get("eh")
		if len(tag) <= 0 {
			continue
		}
		if tag == "err" {
			refCopy.Field(i).SetString(err.Error())
			continue
		}
		if len(tag) > 4 && tag[0:4] == "pre:" {
			field := tag[4:]
			presetVal := reflect.ValueOf(Preset).FieldByName(field)
			refCopy.Field(i).Set(presetVal)
			continue
		}
		switch ref.Elem().Field(i).Interface().(type) {
		case string:
			refCopy.Field(i).SetString(tag)
		case int64, int32, int:
			val, _ := strconv.ParseInt(tag, 10, 64)
			refCopy.Field(i).SetInt(val)
		case uint64, uint32, uint:
			val, _ := strconv.ParseUint(tag, 10, 64)
			refCopy.Field(i).SetUint(val)
		case bool:
			val, _ := strconv.ParseBool(tag)
			refCopy.Field(i).SetBool(val)
		case float64, float32:
			val, _ := strconv.ParseFloat(tag, 64)
			refCopy.Field(i).SetFloat(val)
		default:
			// other implementations...
		}
	}
	zap.Logger.Warn(err.Error())
	ref.Set(refCopy)
	if len(httpCode) > 0 {
		h.C.JSON(httpCode[0], h.V)
		return
	}
	h.C.JSON(http.StatusOK, h.V)
}
