package types

import "errors"

// ErrUidAlreadyUsed is returned when the server rejects a deposit-address
// generation request because the given uid is already registered.
// Detect it with errors.Is(err, types.ErrUidAlreadyUsed).
var ErrUidAlreadyUsed = errors.New("uid already used")
