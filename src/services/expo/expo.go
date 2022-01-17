package services

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/himitery/slack-expo-bot/src/module/slack"
	"net/http"
	"os"
)

type ExpoResult struct {
	Status         string `json:"status"`
	Id             string `json:"id"`
	ArtifactUrl    string `json:"artifactUrl"`
	ReleaseChannel string `json:"releaseChannel"`
	Platform       string `json:"platform"`
}

func Expo(ctx *gin.Context) {
	var expoResult ExpoResult
	if err := ctx.ShouldBindJSON(&expoResult); err != nil {
		ctx.String(http.StatusBadRequest, "Insufficient data")
		return
	}

	expoSignature := ctx.Request.Header.Get("expo-signature")
	if expoSignature != getHashCode(expoResult) {
		ctx.String(http.StatusBadRequest, "Signatures didn't match")
		return
	}

	slack.SendMessage(slack.Message{
		Title: "Expo Build Complete!",
		Content: fmt.Sprintf("%-16s : %s\n", "Id", expoResult.Id) +
			fmt.Sprintf("%-16s : %s\n", "Platform", expoResult.Platform) +
			fmt.Sprintf("%-16s : %s", "ArtifcatUrl", expoResult.ArtifactUrl),
	})

	ctx.String(http.StatusOK, "success")
}

func getHashCode(expoResult ExpoResult) string {
	hash := hmac.New(sha1.New, []byte(os.Getenv("EXPO_KEY")))
	jsonBytes, _ := json.Marshal(expoResult)
	hash.Write(jsonBytes)
	return "sha1=" + hex.EncodeToString(hash.Sum(nil))
}
