package main

import "fmt"

// 观察者模式
// 抽象的观察者
type Listener interface {
	OnTeacherComming() // 观察者得到通知后要触发的具体的动作
}

type Notify interface {
	AddListenner(listener Listener)
	RemoveListenner(listener Listener)
	Notify()
}

// 实现层
type StuZhans struct {
	Baddthing string
}

func (s *StuZhans) OnTeacherComming() { fmt.Println("张三 停止了..", s.Baddthing) }

type Lis struct {
	Baddthing string
}

func (l *Lis) OnTeacherComming() { fmt.Println("Lis 停止了..", l.Baddthing) }

type Wangwu struct {
	Baddthing string
}

func (w *Wangwu) OnTeacherComming() { fmt.Println("王五 停止了..", w.Baddthing) }

// 通知者班长
type ClassMonitor struct {
	listennerList []Listener // 需要通知的全部观察者集合
}

func (m *ClassMonitor) AddListenner(listener Listener) {
	m.listennerList = append(m.listennerList, listener)
}

func (m *ClassMonitor) RemoveListenner(listener Listener) {
	for index, l := range m.listennerList {
		// 找到要删除的元素
		if listener == l {
			m.listennerList = append(m.listennerList[:index], m.listennerList[index+1:]...)
			break
		}
	}
}

func (m *ClassMonitor) Notify() {
	for _, listener := range m.listennerList {
		listener.OnTeacherComming() //多态现象
	}
}

// 业务逻辑层
func main() {
	s1 := &StuZhans{
		Baddthing: "抄作业",
	}
	s2 := &Lis{
		Baddthing: "玩王者荣耀",
	}
	s3 := &Wangwu{
		Baddthing: "看漫画",
	}
	mon := new(ClassMonitor)
	mon.AddListenner(s1)
	mon.AddListenner(s2)
	mon.AddListenner(s3)

	mon.Notify()
}
