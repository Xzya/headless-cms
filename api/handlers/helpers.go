package handlers

import (
	apimodels "github.com/Xzya/headless-cms/gen/models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"fmt"
	"github.com/go-openapi/swag"
)

func StringToUint(s string) uint {
	n, _ := strconv.ParseUint(s, 10, 64)
	return uint(n)
}

func UintToString(n uint) string {
	return fmt.Sprintf("%v", n)
}

func StringPtrToInt32Ptr(s *string) *int32 {
	if n, err := swag.ConvertInt32(swag.StringValue(s)); err == nil {
		return &n
	}
	return nil
}

func Int32PtrToStringPtr(n *int32) *string {
	if n != nil {
		return swag.String(fmt.Sprintf("%v", *n))
	}
	return nil
}

func StringValueIfNotEmpty(s *string) *string {
	if s != nil {
		if len(*s) != 0 {
			return s
		}
	}
	return nil
}

func APIError(err error) *apimodels.Error {
	if err != nil {
		s, ok := status.FromError(err)
		if ok {
			return &apimodels.Error{
				Code:  int32(s.Code()),
				Error: s.Message(),
			}
		}
	}
	return &apimodels.Error{
		Code:  int32(codes.Unknown),
		Error: err.Error(),
	}
}
