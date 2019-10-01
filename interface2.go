package istorage

import "context"

// PutEvents batch of events
// Transaction is not needed (so only some records can be written)
var PutEvents func(ctx context.Context, WSID int64, batch []Event) (err error)

// GetMaxOffset1 returns max Offset1 for given WSID
var GetMaxOffset1 func(ctx context.Context, WSID int64) (err error)

// GetEvents starting from given Offset1 value
// Must analyze ctx.Err AFTER each write to channel
var GetEvents func(ctx context.Context, WSID int64, offset1 int64) (res <-chan Event, perr *error)
