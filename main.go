package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	maks := e.Group("/users")
	maksBot := e.Group("/bot")
	maksBot.GET("/commands/", func(c echo.Context) error {
		var answer string
		command := c.QueryParam("command")
		commandType := c.QueryParam("type")
		if command == "start" {
			if commandType == "first" {
				answer = "hello"
			}
			if commandType == "second" {
				answer = "hi you've been gone for a long time"
			}
		}
		return c.JSON(http.StatusOK, answer)
	})
	maks.GET("/:id", func(c echo.Context) error {
		if c.Param("id") != "1" {
			return c.JSON(http.StatusOK, "no id")
		}
		return c.JSON(http.StatusOK, Response{Result: []Result{{Message: Message{From: User{ID: 1}}}}})
	})
	maks.POST("/", func(c echo.Context) error {
		var response Response

		if err := c.Bind(&response); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})
	maks.PATCH("/:id", func(c echo.Context) error {
		var response Response

		if err := c.Bind(&response); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":8088"))
}

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/url"
// )

// func main() {

// 	var update = false
// 	resp, err := GetUserById("http://localhost:8088", "users", 1)
// 	if err != nil {
// 		fmt.Printf("error sending GETuserById request: %v\n", err)
// 		fmt.Printf("\n%+v\n", resp)
// 		resp, err := PostUser("http://localhost:8088", "users", data)
// 		if err != nil {
// 			fmt.Printf("error sending POSTuser request: %v\n", err)
// 		}
// 		update = true
// 		fmt.Printf("\n%+v\nPOST\n", resp)
// 	}
// 	if !update {
// 		resp, err = PatchUserById("http://localhost:8088", "users", data)
// 		if err != nil {
// 			fmt.Printf("error sending PATCHuserById request: %v", err)
// 		}
// 		fmt.Printf("\n%+v\nPATCH\n", resp)
// 	}
// 	result, err := GetAnswer("http://localhost:8088", "bot/commands/", "start", "second")
// 	if err != nil {
// 		fmt.Printf("error sending GETanswer request: %v\n", err)
// 	}
// 	fmt.Printf("\n%s\nGET\n", result)
// }

// func GetAnswer(baseURL, route, command, commandType string) (string, error) {
// 	fullurl := fmt.Sprintf("%s/%s", baseURL, route)

// 	queryParams := url.Values{}
// 	queryParams.Add("command", command)
// 	queryParams.Add("type", commandType)

// 	resp, err := http.Get(fullurl + "?" + queryParams.Encode())
// 	if err != nil {
// 		return "", fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return "", fmt.Errorf("error reading response body: %v", err)
// 	}
// 	return string(respBody), nil
// }

// func PatchUserById(baseURL, route string, updateData Response) (Response, error) {

// 	url := fmt.Sprintf("%s/%s/%d", baseURL, route, updateData.Result[0].Message.From.ID)

// 	jsonData, err := json.Marshal(updateData)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error marshalling request body: %v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error creating request: %v", err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error reading response body: %v", err)
// 	}

// 	if err := json.Unmarshal(respBody, &updateData); err != nil {
// 		return Response{}, fmt.Errorf("error unmarshalling response body: %v", err)
// 	}
// 	return updateData, nil
// }

// func PostUser(baseURL, route string, userData Response) (Response, error) {

// 	url := fmt.Sprintf("%s/%s", baseURL, route)

// 	jsonData, err := json.Marshal(userData)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error marshalling request body: %v", err)
// 	}

// 	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error creating request: %v", err)
// 	}

// 	client := &http.Client{}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()

// 	respBody, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error reading response body: %v", err)
// 	}

// 	if err := json.Unmarshal(respBody, &userData); err != nil {
// 		return Response{}, fmt.Errorf("error unmarshalling response body: %v", err)
// 	}
// 	return userData, nil
// }

// func GetUserById(baseURL, route string, id int) (Response, error) {

// 	url := fmt.Sprintf("%s/%s/%d", baseURL, route, id)

// 	resp, err := http.Get(url)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error creating request: %v", err)
// 	}
// 	defer resp.Body.Close()
// 	if resp.StatusCode != http.StatusOK {
// 		return Response{}, fmt.Errorf("error sending request: %v", err)
// 	}
// 	body, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return Response{}, fmt.Errorf("error reading response body: %v", err)
// 	}

// 	var result Response
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return Response{}, fmt.Errorf("error unmarshalling response body: %v", err)
// 	}
// 	return result, nil

// }

type Response struct {
	Ok     bool     `json:"ok"`
	Result []Result `json:"result"`
}

type Result struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

type Message struct {
	MessageID int      `json:"message_id"`
	From      User     `json:"from"`
	Chat      Chat     `json:"chat"`
	Date      int64    `json:"date"`
	Text      string   `json:"text"`
	Entities  []Entity `json:"entities"`
}

type User struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
	IsPremium    bool   `json:"is_premium"`
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Username  string `json:"username"`
	Type      string `json:"type"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// var data = Response{
// 	Ok: true,
// 	Result: []Result{
// 		{
// 			UpdateID: 64523697,
// 			Message: Message{
// 				MessageID: 2510,
// 				From: User{
// 					ID:           5789357767,
// 					IsBot:        false,
// 					FirstName:    "Тех.специалист Офсет",
// 					Username:     "tehsupport_offset",
// 					LanguageCode: "ru",
// 					IsPremium:    true,
// 				},
// 				Chat: Chat{
// 					ID:        5789357767,
// 					FirstName: "Тех.специалист Офсет",
// 					Username:  "tehsupport_offset",
// 					Type:      "private",
// 				},
// 				Date: 1733419052,
// 				Text: "/start",
// 				Entities: []Entity{
// 					{
// 						Offset: 0,
// 						Length: 6,
// 						Type:   "bot_command",
// 					},
// 				},
// 			},
// 		},
// 	},
// }
