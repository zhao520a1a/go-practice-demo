package main

import (
	"context"
	"fmt"
	"github.com/sasha-s/go-deadlock"
	"gitlab.pri.ibanyu.com/middleware/seaweed/xlog"
	"time"
)

const reloadPeriodInMinutes = 5

type Manager struct {
	mu deadlock.RWMutex
}

func NewManager() *Manager {
	return &Manager{}
}

//使用 Tick/Sleep 每隔100毫秒打印“Hello TigerwolfC”
func (m *Manager) StartByTimePeriod1(ctx context.Context) {
	for range time.Tick(time.Millisecond * 100) {
		fmt.Println("Hello TigerwolfC")
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	for range ticker.C {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}

	for {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("Hello TigerwolfC")
	}
}

//定时：时间段处理一次
func (m *Manager) StartByTimePeriod(ctx context.Context) {
	fun := "Manager.StartSync -->"

TimedLoop:
	for {
		err := m.Reload(ctx)
		if err != nil {
			xlog.Errorf(ctx, "%s err:%v", fun, err)
		}

		select {
		case <-time.After(reloadPeriodInMinutes * time.Minute):
			xlog.Infof(ctx, "%s start next round", fun)
		case <-ctx.Done():
			xlog.Infof(ctx, "%s about to exit", fun)
			break TimedLoop
		}
	}
}

func (m *Manager) Reload(ctx context.Context) (err error) {

	return nil
}

//定时：每天固定时间点处理一次
func (m *Manager) StartByTimeDot(ctx context.Context) {
	fun := "Manager.StartSync -->"

TimedLoop:
	for {
		err := m.Reload(ctx)
		if err != nil {
			xlog.Errorf(ctx, "%s err:%v", fun, err)
		}

		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 10, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))

		select {
		case <-t.C:
			xlog.Infof(ctx, "%s start next round", fun)
		case <-ctx.Done():
			xlog.Infof(ctx, "%s about to exit", fun)
			break TimedLoop
		}
	}
}

func main() {
	ctx := context.Background()
	defaultManage := NewManager()

	var ch chan int
	//定时任务
	go func() {
		go defaultManage.StartByTimeDot(ctx)
		//go defaultManage.StartByTimePeriod(ctx)
		time.Sleep(10 * time.Minute)
		ch <- 1
	}()
	<-ch

}
