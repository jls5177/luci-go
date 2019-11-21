// Code generated by svcdec; DO NOT EDIT.

package rpcpb

import (
	"context"

	proto "github.com/golang/protobuf/proto"

	empty "github.com/golang/protobuf/ptypes/empty"
)

type DecoratedRecorder struct {
	// Service is the service to decorate.
	Service RecorderServer
	// Prelude is called for each method before forwarding the call to Service.
	// If Prelude returns an error, then the call is skipped and the error is
	// processed via the Postlude (if one is defined), or it is returned directly.
	Prelude func(ctx context.Context, methodName string, req proto.Message) (context.Context, error)
	// Postlude is called for each method after Service has processed the call, or
	// after the Prelude has returned an error. This takes the the Service's
	// response proto (which may be nil) and/or any error. The decorated
	// service will return the response (possibly mutated) and error that Postlude
	// returns.
	Postlude func(ctx context.Context, methodName string, rsp proto.Message, err error) error
}

func (s *DecoratedRecorder) CreateInvocation(ctx context.Context, req *CreateInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateInvocation", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) UpdateInvocation(ctx context.Context, req *UpdateInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "UpdateInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.UpdateInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "UpdateInvocation", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) FinalizeInvocation(ctx context.Context, req *FinalizeInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "FinalizeInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.FinalizeInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "FinalizeInvocation", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) Include(ctx context.Context, req *IncludeRequest) (rsp *empty.Empty, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "Include", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.Include(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "Include", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) CreateTestResult(ctx context.Context, req *CreateTestResultRequest) (rsp *TestResult, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateTestResult", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateTestResult(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateTestResult", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) BatchCreateTestResults(ctx context.Context, req *BatchCreateTestResultsRequest) (rsp *BatchCreateTestResultsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "BatchCreateTestResults", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.BatchCreateTestResults(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "BatchCreateTestResults", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) CreateTestExoneration(ctx context.Context, req *CreateTestExonerationRequest) (rsp *TestExoneration, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "CreateTestExoneration", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.CreateTestExoneration(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "CreateTestExoneration", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) BatchCreateTestExonerations(ctx context.Context, req *BatchCreateTestExonerationsRequest) (rsp *BatchCreateTestExonerationsResponse, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "BatchCreateTestExonerations", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.BatchCreateTestExonerations(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "BatchCreateTestExonerations", rsp, err)
	}
	return
}

func (s *DecoratedRecorder) DeriveInvocation(ctx context.Context, req *DeriveInvocationRequest) (rsp *Invocation, err error) {
	if s.Prelude != nil {
		var newCtx context.Context
		newCtx, err = s.Prelude(ctx, "DeriveInvocation", req)
		if err == nil {
			ctx = newCtx
		}
	}
	if err == nil {
		rsp, err = s.Service.DeriveInvocation(ctx, req)
	}
	if s.Postlude != nil {
		err = s.Postlude(ctx, "DeriveInvocation", rsp, err)
	}
	return
}