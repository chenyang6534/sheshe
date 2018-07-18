package game5g

import (
	"dq/timer"
	"dq/utils"
	"sync"
	"time"
)

const (
	PlayerState_GoIn    = 1 //玩家加入中
	PlayerState_Gameing = 2 //游戏中
	PlayerState_Die     = 3 //死亡
	PlayerState_GoOut   = 4 //退出
)

//玩家
type Game5GPlayer struct {
	//基本数据
	Uid         int
	ConnectId   int
	Name        string
	Gold        int64
	WinCount    int
	LoseCount   int
	SeasonScore int
	RankNum     int
	AvatarUrl   string
	SkinId      int

	//游戏
	Game *Game5GLogic

	//蛇
	MySnake *Snake
	State   int //玩家状态
}

//游戏逻辑

const (
	Game5GState_Wait   = 1 //等待玩家加入中
	Game5GState_Gaming = 2 //游戏中
	Game5GState_Result = 3 //结算中
	Game5GState_Over   = 4 //解散
)

//游戏模式
const (
	Game5GMode_CreateRoom     = 1 //自己建房
	Game5GMode_AutoMatching   = 2 //自动匹配
	Game5GMode_SeasonMatching = 3 //赛季天梯匹配
)

type Game5GLogic struct {

	//games
	GameAgent *Game5GAgent

	//游戏ID
	GameId int
	//将要玩游戏的玩家ID
	//WillPlayGamePlayerUid [2]int
	//玩家
	Player []*Game5GPlayer

	//游戏状态
	State int

	//锁
	Lock *sync.Mutex

	//时间到 倒计时
	LogicTimer *timer.Timer

	//游戏模式
	GameMode int

	//创建者UID -1表示服务器自动创建
	CreateId int

	//游戏创建时间戳
	CreateGameTime int64

	//帧率
	GameFrame int
	//当前帧
	CurFrameNum int
}

func (game *Game5GLogic) Init() {
	game.State = Game5GState_Gaming
	game.GameFrame = 20
	game.CurFrameNum = 0

	game.Lock = new(sync.Mutex)
	game.CreateId = -1
	game.CreateGameTime = utils.Milliseconde()
	game.LogicTimer = timer.AddCallback(time.Millisecond*(1000/20), game.Update)

	game.Player = make([]*Game5GPlayer, 0)

}

//玩家进入
func (game *Game5GLogic) GoIn(player *Game5GPlayer) (*Game5GPlayer, error) {
	game.Lock.Lock()
	defer game.Lock.Unlock()
	player.State = PlayerState_GoIn
	game.Player = append(game.Player, player)

	//同步当前游戏给此玩家

	return player, nil

}

//玩家退出
func (game *Game5GLogic) GoOut(player *Game5GPlayer) bool {
	game.Lock.Lock()
	defer game.Lock.Unlock()

	return true
}

//玩家掉线
func (game *Game5GLogic) Disconnect(player *Game5GPlayer) bool {
	game.Lock.Lock()
	defer game.Lock.Unlock()

	return true

}

//
func (game *Game5GLogic) Update() {
	game.Lock.Lock()
	defer game.Lock.Unlock()

	if game.State == Game5GState_Gaming {
		game.CurFrameNum++

		//遍历玩家列表
		for k, v := range game.Player {

		}

	}

}

var g_GameId = 10000
var g_GameId_lock = new(sync.Mutex)

//创建一个新的游戏ID
func GetNewGameId() int {
	g_GameId_lock.Lock()
	defer g_GameId_lock.Unlock()

	g_GameId++
	return g_GameId
}

func NewGame5GLogic_SeasonMatching(ga *Game5GAgent) *Game5GLogic {
	ng := &Game5GLogic{}
	ng.GameId = GetNewGameId()
	ng.GameAgent = ga

	ng.GameMode = Game5GMode_SeasonMatching
	ng.CreateId = -1
	ng.Init()

	return ng
}
