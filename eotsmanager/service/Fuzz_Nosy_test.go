package service

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/finality-provider/eotsmanager/proto"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"google.golang.org/grpc"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Server_RunUntilShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Server
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RunUntilShutdown(ctx)
	})
}

// skipping Fuzz_Nosy_Server_startGrpcListen__ because parameters include func, chan, or unsupported interface: []net.Listener

func Fuzz_Nosy_rpcServer_CreateKey__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.CreateKeyRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.CreateKey(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_CreateRandomnessPairList__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.CreateRandomnessPairListRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.CreateRandomnessPairList(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_KeyRecord__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.KeyRecordRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.KeyRecord(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var _x3 *proto.PingRequest
		fill_err = tp.Fill(&_x3)
		if fill_err != nil {
			return
		}
		if r == nil || _x3 == nil {
			return
		}

		r.Ping(_x2, _x3)
	})
}

func Fuzz_Nosy_rpcServer_RegisterWithGrpcServer__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var grpcServer *grpc.Server
		fill_err = tp.Fill(&grpcServer)
		if fill_err != nil {
			return
		}
		if r == nil || grpcServer == nil {
			return
		}

		r.RegisterWithGrpcServer(grpcServer)
	})
}

func Fuzz_Nosy_rpcServer_SaveEOTSKeyName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.SaveEOTSKeyNameRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.SaveEOTSKeyName(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_SignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.SignEOTSRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.SignEOTS(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_SignSchnorrSig__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.SignSchnorrSigRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.SignSchnorrSig(_x2, req)
	})
}

func Fuzz_Nosy_rpcServer_UnsafeSignEOTS__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var r *rpcServer
		fill_err = tp.Fill(&r)
		if fill_err != nil {
			return
		}
		var _x2 context.Context
		fill_err = tp.Fill(&_x2)
		if fill_err != nil {
			return
		}
		var req *proto.SignEOTSRequest
		fill_err = tp.Fill(&req)
		if fill_err != nil {
			return
		}
		if r == nil || req == nil {
			return
		}

		r.UnsafeSignEOTS(_x2, req)
	})
}
