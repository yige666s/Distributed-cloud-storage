package microhandler

import (
	"Distributed-cloud-storage/common"
	"Distributed-cloud-storage/config"
	"Distributed-cloud-storage/db"
	"Distributed-cloud-storage/service/account/proto/user"
	"Distributed-cloud-storage/util"
	"context"
)

type User struct {
}

// Server,处理用户注册请求
func SignUp(ctx context.Context, req *user.ReqSignup, resp *user.RespSignup) error {
	username := req.Username
	passwd := req.Password
	if len(username) < 3 || len(passwd) < 5 {
		resp.Code = common.StatusParamInvalid
		resp.Message = "注册参数无效"
		return nil
	}

	enc_passwd := util.Sha1([]byte(passwd + config.PasswordSalt)) // 对密码进行加盐处理
	suc := db.UserSignup(username, enc_passwd)
	if suc {
		resp.Code = common.StatusOK
		resp.Message = "注册成功"
	} else {
		resp.Code = common.StatusRegisterFailed
		resp.Message = "注册失败"
	}
	return nil
}
