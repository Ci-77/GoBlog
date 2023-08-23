package qqlogin

import (
	"encoding/json"
	"fmt"
	"gvb/global"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// 写qq登录的
type QQinfo struct {
	Nickname string `json:"nickname"`     //昵称
	Gender   string `json:"gender"`       //性别
	Avater   string `json:"figureurl_qq"` //头像大图
	OpenID   string `json:"open_id"`
}
type QQLogin struct {
	appID     string
	appKey    string
	redirect  string
	code      string
	accessTok string
	openID    string
}

func NewQQLogin(code string) (qqInfo QQinfo, err error) {
	qqlogin := &QQLogin{
		appID:    global.Config.QQ.AppID,
		appKey:   global.Config.QQ.Key,
		redirect: global.Config.QQ.Redirect,
		code:     code,
	}
	err=qqlogin.GetToken()
	if err!=nil{
log.Fatal(err)
	}
	err=qqlogin.GetOpenId()
	if err!=nil {
		log.Fatal(err)
	}
	qqInfo,err=qqlogin.GetUserInfo()
	if err!=nil {
		log.Fatal(err)
	}
	qqInfo.OpenID=qqlogin.openID
	return qqInfo,nil
}
//获取code

//获取token
func (q *QQLogin) GetToken() error {
	params := url.Values{}
	params.Add("grant_type", "authorization_code")
	params.Add("client_id", q.appID)
	params.Add("client_secret", q.appKey)
	params.Add("code", q.code)
	params.Add("redirect_url", q.redirect)
u,err := url.Parse("https://graph.qq.com/oauth2.0/token")
if err!=nil {
	return err
}
u.RawQuery=params.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	qs,err:= parseQS(res.Body)
	if err!=nil{
		return err
	}
	q.accessTok=qs[`access_token`][0]
	return nil
}

func parseQS(r io.Reader)(val map[string][]string,err error) {
	val,err=url.ParseQuery(readAll(r))
	if err!=nil {
		log.Fatalln("出错",err)
	}
	return val,nil

}
//读取所有的数据，并将其转化为字符串
func readAll(r io.Reader) string{
b,err:=io.ReadAll(r)
if err!=nil {
	log.Fatal(err)
}
return string(b)
}
//获取openid
func (q *QQLogin)GetOpenId()error{
	u, err := url.Parse(fmt.Sprintf( "https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessTok))
	if err != nil {
		return err
	}
	res,err:=http.Get(u.String())
	if err!=nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
openID,err:=getOpenID(res.Body)
if err!=nil {
	log.Fatal(err)
}
q.openID=openID
return nil
}
//从http响应的正文中解析出openid
func getOpenID(r io.Reader)(string, error) {
	body:=readAll(r)
	start:=strings.Index(body,`"openid":"`)+len(`"openid":"`)
if start==-1 {
	return "",fmt.Errorf("openid not found")
}
end:=strings.Index(body[start:],`"`)
if end==-1 {
	return "",fmt.Errorf("openid not found")
}
return body[start:start+end],nil
}
func (q *QQLogin)GetUserInfo()(qqInfo QQinfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.accessTok)
	params.Add("openid", q.openID)
	params.Add("oauth_consumer_key",q.appID )
	u,err:=url.Parse("https://graph.qq.com/user/get_user_info?")
	if err!=nil {
		return qqInfo,err
	}
	u.RawQuery=params.Encode()
	resp, err := http.Get(u.String())
	if err!=nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	err=json.Unmarshal(data, &qqInfo)
	if err!=nil {
		return qqInfo,err
	}
	return qqInfo,nil
}
