package error_test

import (
	"code.byted.org/im_cloud/common_tiktok/pkg/errs"
	"fmt"
)

var globalError error

func TestingError() {

	err := errs.FromCodeMsg(101000, "test err")

	newErr := errs.New(errs.DownstreamErr, errs.Msg("errmsg"), errs.Wrap(err))

	fmt.Println(errs.Code(newErr))
	fmt.Println(newErr.Error())

}
