package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"userweb/forms"
	"userweb/global"
	"userweb/global/response"
	"userweb/middlewares"
	"userweb/models"
	"userweb/proto"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func HandleGrpcErrorToHttp(err error, c *gin.Context) {
	//将grpc code 转换成http状态码
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, gin.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "内部错误",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, gin.H{
					"msg": "参数错误",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": "用户服务不可用",
				})

			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"msg": e.Code(),
				})
			}
		}
	}
	return
}
func HandleValidatorErr(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}

	fmt.Println(err.Error())
	ctx.JSON(http.StatusBadRequest, gin.H{
		"error": removeTopStruct(errs.Translate(global.Trans)),
	})
}
func GetUserList(ctx *gin.Context) {
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.MysqlInfo.Host,
		global.ServerConfig.MysqlInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接用户服务失败", "msg", err)
	}

	//打印当前用户
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)
	zap.S().Infof("访问用户:%d", currentUser.ID)

	//grpc client
	userSrvClient := proto.NewUserClient(userConn)
	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	pageInfo := &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	}
	userinfo, err := userSrvClient.GetUserList(ctx, pageInfo)
	if err != nil {
		zap.S().Error("获取用户列表页失败：", err)
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	result := make([]interface{}, 0)
	for _, value := range userinfo.Data {
		// birth, _ := time.Parse("2006-01-02 15:04:05", value.Birthday)
		user := response.UserResponse{
			ID:       value.Id,
			NickName: value.NickName,
			BirthDay: value.Birthday, //todo 去除多余时间信息
			Gender:   int(value.Gender),
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	ctx.JSON(http.StatusOK, result)
	return
}

func PasswordLogin(ctx *gin.Context) {
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := ctx.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorErr(ctx, err)
		return
	}
	//登录逻辑
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.MysqlInfo.Host,
		global.ServerConfig.MysqlInfo.Port), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("连接用户服务失败", "msg", err)
	}
	userSrvClient := proto.NewUserClient(userConn)
	if rsp, err := userSrvClient.GetUserByMobile(context.Background(), &proto.MobileRequest{
		Moblie: passwordLoginForm.Mobile,
	}); err != nil { //是否存在
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusBadRequest, gin.H{
					"mobile": "用户不存在",
				})
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else { //检查密码 todo 加密加盐
		if rsp.Password == passwordLoginForm.Password {
			//生成token
			j := middlewares.NewJWT()
			claims := models.CustomClaims{
				ID:          uint(rsp.Id),
				NickName:    rsp.NickName,
				AuthorityId: uint(rsp.Role),
				StandardClaims: jwt.StandardClaims{
					NotBefore: time.Now().Unix(), //签名的生效时间
					ExpiresAt: time.Now().Unix() + 60*60*24*30,
					Issuer:    "wei",
				},
			}
			token, err := j.CreateToken(claims)
			if err != nil {
				zap.S().Error("生成token失败：", err)
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"msg": "生成token失败",
				})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{
				"id":         rsp.Id,
				"nickname":   rsp.NickName,
				"token":      token,
				"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"msg": "密码错误",
			})
		}
	}

}

func removeTopStruct(fields map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fields {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}
