package main

import (
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// ユーザー定義構造体
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	jwt.StandardClaims
}

// シークレットキー(仮)
const secretKey = "foobar"

// jwt(トークン文字列)を生成する関数
func createTokenString() string {
	// User情報をtokenに込める
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &User{
			Name: "otiai10",
			Age:  30,
	})
	// Secretで文字列にする. このSecretはサーバだけが知っている
	tokenstring, err := token.SignedString([]byte(secretKey))
	if err != nil {
			log.Fatalln(err)
	}
	return tokenstring
}

// jwtを生成してクッキーに保存
func setJwt(w http.ResponseWriter, r *http.Request) {
	// jwt tokenをを生成
	tokenstring := createTokenString()
	// tokenの表示
	log.Println(tokenstring)

	// cookieに格納
	c := http.Cookie{
		Name: "jwt-test",
		Value: tokenstring,
		HttpOnly: true,
	}

	http.SetCookie(w, &c)
}

// jwtの認証をしてjwtの中身を解読してレスポンス
// jwtはcookieの中から取り出す
func getJwt(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("jwt-test")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first_cookie")
		return
	}

	token, err := jwt.Parse(c.Value, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	fmt.Fprintln(w, token.Claims)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_jwt", setJwt)
	http.HandleFunc("/get_jwt", getJwt)
	server.ListenAndServe()
}