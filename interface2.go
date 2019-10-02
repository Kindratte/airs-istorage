package istorage

import "context"

// PutEvents batch of events
// Transaction is not needed (so only some records can be written)
// Offset1 MUST start from zero
// Offset1 MUST monotonically increase
// If it is not GetEvents may stop on first missed number
var PutEvents func(ctx context.Context, WSID int64, batch []Event) (err error)

// GetEvents between offset1From and offset1To (offset1To INCLUDED into result)
// Must analyze ctx.Err AFTER each write to channel
var GetEvents func(ctx context.Context, WSID int64, offset1From, offset1To int64) (res <-chan Event, perr *error)
