package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/midoks/go-p2p-server/internal/announce"
	// "github.com/midoks/go-p2p-server/internal/hub"
	"github.com/midoks/go-p2p-server/internal/tools/uniqid"
)

func initAnnounce() {
	announce.Init()
}

type AnPeer struct {
	Id string `json:"id"`
}

//接收announce信息
func p2pChannel(c *gin.Context) {
	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.BindJSON(&postJson)

	uniqidId := uniqid.New(uniqid.Params{"", false})
	gPeers := []AnPeer{}
	if channel, ok := postJson["channel"]; ok {
		key := channel.(string)
		if peer, ok := announce.Get(key); ok {
			for _, p := range peer {

				gPeers = append(gPeers, AnPeer{Id: p})
			}
			announce.Set(key, uniqidId)
		} else {
			announce.Set(key, uniqidId)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              uniqidId,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               uniqidId,
		},
	})
}

//接收announce信息
func p2pChannelPeers(c *gin.Context) {
	postJson := make(map[string]interface{}) //注意该结构接受的内容
	c.ShouldBind(&postJson)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	gPeers := []AnPeer{}
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	announce.SetPeerHeartbeat(peers, 60*time.Second)

	c.JSON(http.StatusOK, gin.H{
		"ret": 0,
		"data": gin.H{
			"id":              peers,
			"peers":           gPeers,
			"report_interval": 3,
			"v":               "scadasd",
		},
	})
}

//接收announce信息
func p2pChannelStats(c *gin.Context) {
	json := make(map[string]interface{})
	c.ShouldBind(&json)
	fmt.Println(json)

	channel_id := c.Param("channel_id")
	peers := c.Param("peer")

	announce.SetPeerHeartbeat(peers, 60*time.Second)

	gPeers := []AnPeer{}
	key := channel_id
	if peer, ok := announce.Get(key); ok {
		for _, p := range peer {
			gPeers = append(gPeers, AnPeer{Id: p})
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"ret":   0,
		"id":    peers,
		"name":  "stats",
		"peers": gPeers,
		"data":  gin.H{},
	})
}
