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
