package main

import (
	"fmt"

	"github.com/google/wire"
)

// wire 依赖注入
// 依赖注入有两个概念
// 注入器和构造器
// 构造器就是创建函数，每个注入器实际就是一个对象创建和初始化函数
// 在这个函数中，只需要告诉wire要创建什么类型的对象，这个类型的依赖，wire
// 会自动生成依赖

// 看一个简单的例子
// 假如黑暗的世界里有一头怪兽
type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "Kitty"}
}

// 假如黑暗的世界里还有一个勇士
type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

// 比方说现在，勇士打败了怪兽
type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

func (m Mission) Start() {
	fmt.Printf("%s defeats %s, world peace!\n", m.Player.Name, m.Monster.Name)
}

// 当我们需要实现的时候
func main() {
	monster := NewMonster()
	player := NewPlayer("dj")
	mission := NewMission(player, monster)
	mission.Start()
}

// 当我们使用wire的时候，我们可以这样做
func InitMission(name string) Mission {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}
}
