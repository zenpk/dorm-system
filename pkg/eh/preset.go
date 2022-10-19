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

type PresetType struct {
	CodeOK              int64
	CodeInputError      int64
	CodeTokenError      int64
	CodeMiddlewareError int64
	CodeDatabaseError   int64
	CodeLogicalError    int64
	CodeUncaughtError   int64
}

// Error Code
// ABC
// A: [2: success, 4: client error, 5: server error]
// B: [0: success, 1: input error, 2: middleware error, 3: database error, 4: logical error, 5: uncaught error]
// C: [0: success, 1: failure]

var Preset = PresetType{
	CodeOK:              200,
	CodeInputError:      411,
	CodeTokenError:      421,
	CodeMiddlewareError: 521,
	CodeDatabaseError:   531,
	CodeLogicalError:    541,
	CodeUncaughtError:   551,
}
