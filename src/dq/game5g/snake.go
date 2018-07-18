package game5g

import (
	"dq/vec2d"
	"math"
	"math/rand"
)

type SnakeNode struct {
	Position  *vec2d.Vector //位置
	Direction *vec2d.Vector //方向

}

//玩家
type Snake struct {
	Player *Game5GPlayer

	//基本数据
	Uid    int
	SkinId int

	//初始属性
	InitEnergy float32 //初始能量

	Speed            float32 //移动速度  (像素每秒)
	DoubleSpeed      float32 //加速状态速度倍率(2.0)
	RadioSpeed       float32 //转身速度
	DieJiChengEnergy float32 //死亡时继承 吃到的能量百分比
	DieDropEnergy    float32 //死亡时掉落能量 百分比
	BodyRadius       int     //身体碰撞半径
	AddEnergySpeed   float32 //加分速率

	//基本属性
	Energy float32 //能量

	//蛇头
	Head SnakeNode
	//蛇身体
	Body []SnakeNode
}

//节点数
func (snake *Snake) GetNodeNum() int {
	return int(snake.Energy / 20)
}

//创建蛇节点
func (snake *Snake) CreateSnakeNode() {
	addenergy := (snake.Energy - snake.InitEnergy) * snake.DieJiChengEnergy
	if addenergy < 0 {
		addenergy = 0
	}
	snake.Energy = snake.InitEnergy + addenergy

	snake.Head = SnakeNode{Position: vec2d.New(float64(rand.Intn(1000)+500), 1000), Direction: vec2d.New(0, 1)}

	//创建身体
	bodynum := snake.GetNodeNum()
	snake.Body = make([]SnakeNode, 0)
	frame := snake.Player.Game.GameFrame
	for k := 0; k < bodynum; k++ {
		body := SnakeNode{Position: vec2d.New(snake.Head.Position.X, snake.Head.Position.Y-float64(snake.Speed)/float64(frame)), Direction: vec2d.New(0, 1)}

		snake.Body = append(snake.Body, body)
	}
}

func (snake *Snake) Init(skinid int) {
	snake.SkinId = skinid

	//初始化属性()
	snake.InitEnergy = 100
	snake.Speed = 240
	snake.DoubleSpeed = 2.0
	snake.RadioSpeed = math.Pi / 180 * 15 //转向速度 弧度每帧（15摄氏度）
	snake.DieJiChengEnergy = 0.2
	snake.DieDropEnergy = 0.5
	snake.BodyRadius = 50
	snake.AddEnergySpeed = 1.0

	//snake.Energy = snake.InitEnergy

	snake.CreateSnakeNode()

}

func CreateSnake(skinid int, player *Game5GPlayer) {
	snake := &Snake{}
	snake.Player = player
	snake.Init(skinid)
}
