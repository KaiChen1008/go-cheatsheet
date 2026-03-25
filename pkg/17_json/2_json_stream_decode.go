package json

import (
	"encoding/json"
	"fmt"
	"strings"
)

/*
Token 流：
{ → "metadata" → { → ... } → "users" → [ → { → "id" → 1 → "name" → "Alice" ... } → ... ] → }

	   ↑                     ↑
	略過整塊              觸發 parseUsers()
*/
type User struct {
	ID    int
	Name  string
	Score float64
}

func main() {
	raw := `{
		"metadata": { "version": "1.0", "total": 3 },
		"users": [
			{ "id": 1, "name": "Alice", "secret": "xxx", "score": 95.5 },
			{ "id": 2, "name": "Bob",   "secret": "yyy", "score": 87.0 },
			{ "id": 3, "name": "Carol", "secret": "zzz", "score": 92.3 }
		]
	}`

	decoder := json.NewDecoder(strings.NewReader(raw))
	var users []User

	for {
		token, err := decoder.Token()
		if err != nil { // EOF
			break
		}

		// 找到 key = "users" 時，開始解析陣列
		if key, ok := token.(string); ok && key == "users" {
			users = parseUsers(decoder)
		}
		// 其他 token（metadata 等）直接略過，不做任何事
	}

	for _, u := range users {
		fmt.Printf("ID: %d | Name: %-6s | Score: %.1f\n", u.ID, u.Name, u.Score)
	}
}

// parseUsers 在確認進入 "users" array 後，逐筆解析每個 user object
func parseUsers(decoder *json.Decoder) []User {
	var users []User

	// 確認下一個 token 是 '['
	token, err := decoder.Token()
	if err != nil || token.(json.Delim) != '[' {
		fmt.Println("expected '['")
		return nil
	}

	// 逐筆讀取 array 裡的每個 object
	for decoder.More() {
		user, ok := parseUser(decoder)
		if !ok {
			continue
		}
		users = append(users, user)
	}

	decoder.Token() // 消耗 ']'
	return users
}

// parseUser 解析單一 user object，只取需要的欄位
func parseUser(decoder *json.Decoder) (User, bool) {
	// 確認下一個 token 是 '{'
	token, err := decoder.Token()
	if err != nil || token.(json.Delim) != '{' {
		return User{}, false
	}

	var user User

	for decoder.More() {
		// 讀 key
		keyToken, err := decoder.Token()
		if err != nil {
			return User{}, false
		}
		key := keyToken.(string)

		// 依 key 決定怎麼讀 value
		switch key {
		case "id":
			var v float64
			decoder.Decode(&v) // JSON number 預設是 float64
			user.ID = int(v)

		case "name":
			var v string
			decoder.Decode(&v)
			user.Name = v

		case "score":
			var v float64
			decoder.Decode(&v)
			user.Score = v

		default:
			// 不需要的欄位（secret 等），跳過 value
			var discard json.RawMessage
			decoder.Decode(&discard)
		}
	}

	decoder.Token() // 消耗 '}'
	return user, true
}
