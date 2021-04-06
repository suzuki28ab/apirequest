package main

import (
	"fmt"
	"sync"

	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/joho/godotenv/autoload"

	"github.com/suzukix/apireq/api_request"
	"github.com/suzukix/apireq/db"
	"github.com/suzukix/apireq/discord"
	"github.com/suzukix/apireq/model"
)

func apiRequest() (string, error) {
	bcasters := []model.Bcaster{}
	db, err := db.GetDbConnect()
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	db.Set("gorm:auto_preload", true).Find(&bcasters)
	nicoUserSession := api_request.GetUserSeesion()
	twitchToken := api_request.GetTwitchToken()
	discordSession := discord.GetDiscordGo()
	defer discordSession.Close()
	var wg sync.WaitGroup

	for _, bcaster := range bcasters {
		wg.Add(1)
		go func(b model.Bcaster) {
			defer wg.Done()
			startFlag := b.RequestBcasterLive(db, nicoUserSession, twitchToken)
			if startFlag == 1 {
				discord.SendMessage(discordSession, b.CreateDiscordNotice())
			}
		}(bcaster)
	}
	wg.Wait()
	model.SetApiTime(db)
	return "success", nil
}

func main() {
	lambda.Start(apiRequest)
	// db, err := db.GetDbConnect()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// sqlDB, _ := db.DB()
	// defer sqlDB.Close()

	// bcaster := model.Bcaster{
	// 	Name: "hoge",
	// 	Twitch: model.Twitch{
	// 		Account: "fuga",
	// 		OffURL:  "https://www.twitch.tv/fuga",
	// 	},
	// }

	// db.Create(&bcaster)
	// db.Save(&bcaster)
}
