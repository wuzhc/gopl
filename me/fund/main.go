// 基金管理
// 定时获取最新数据,实时显示收益
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/wuzhc/go-logs"
)

const (
	FUNDGZ_HOST = "http://fundgz.1234567.com.cn/js/"
)

var (
	ErrHasRun  = errors.New("Finance can not run twice")
	ErrHasStop = errors.New("Finance has stop")
)

type fundData struct {
	Fundcode string
	Name     string
	Jzrq     string // 截止日期
	Dwjz     string
	Gsz      string
	Gszzl    string
	Gztime   string
}

type Finance struct {
	sync.Mutex
	funds     []*Fund
	polling   time.Duration
	closeChan chan struct{}
	running   int
	logger    *logs.Dispatcher
}

type Fund struct {
	sync.Mutex
	finance   *Finance
	capital   float64
	profit    float64
	code      string
	lastGszzl float64
}

func NewFinance(poll int) *Finance {
	logger := logs.NewDispatcher()
	logger.SetTarget(logs.TARGET_FILE, `{"filename":"log.log","level":10,"max_size":50000000,"rotate":true}`)
	return &Finance{
		polling:   time.Duration(poll) * time.Second,
		closeChan: make(chan struct{}),
		logger:    logger,
	}
}

func (f *Finance) Run() error {
	defer func() {
		if err := recover(); err != nil {
			f.running = 0
		}
	}()
	if f.running == 1 {
		return ErrHasRun
	}

	f.running = 1 // optimistically
	tc := time.NewTicker(f.polling)
	defer tc.Stop()

	for {
		select {
		case <-tc.C:
			var wg sync.WaitGroup
			for _, fund := range f.funds {
				wg.Add(1)
				go fund.UpdateMoney(&wg)
			}
			wg.Wait()
			f.PrintProfit()
		case <-f.closeChan:
			f.running = 0
			break
		}
	}
}

func (f *Finance) Close() {
	f.closeChan <- struct{}{}
}

func (f *Finance) AddFund(money float64, code string) {
	if f.running == 1 {
		return
	}

	f.Lock()
	fund := &Fund{
		finance: f,
		capital: money,
		code:    code,
	}
	f.funds = append(f.funds, fund)
	f.Unlock()
}

func (f *Finance) RemoveFund(code string) {
	if f.running == 1 {
		return
	}

	for i, fund := range f.funds {
		if fund.code == code {
			f.Lock()
			if len(f.funds) == 1 {
				f.funds = f.funds[:0]
			} else {
				f.funds = append(f.funds[:i], f.funds[i+1:]...)
			}
			f.Unlock()
			break
		}
	}
}

func (f *Finance) PrintProfit() {
	if f.running == 0 {
		fmt.Println(ErrHasStop)
		return
	}

	var profit, total float64
	for _, rund := range f.funds {
		profit += rund.profit
		total += rund.capital
	}

	now := time.Now().Format("2006-01-02 15:04:05")
	if profit < 0 {
		fmt.Printf("当前时间:%v\n总收益\x1b[0;32m%v\x1b[0m\n\n", now, profit)
	} else {
		fmt.Printf("当前时间:%v\n总收益\x1b[0;31m%v\x1b[0m\n\n", now, profit)
	}

	f.logger.Info(fmt.Sprintf("总收益为%v\n", profit))
}

func (fund *Fund) UpdateMoney(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
	}()

	url := FUNDGZ_HOST + fund.code + ".js?rt=" + string(time.Now().Unix())
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	nbyte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	s := string(nbyte)
	s = strings.Replace(s, "jsonpgz(", "", -1)
	s = strings.Replace(s, ");", "", -1)

	var data fundData
	if err := json.Unmarshal([]byte(s), &data); err != nil {
		panic(err)
	}
	gszzl, err := strconv.ParseFloat(data.Gszzl, 64)
	if err != nil {
		panic(err)
	}

	fund.PrintFundProfit(data.Name, gszzl)
	if fund.lastGszzl == gszzl {
		return
	}

	fund.Lock()
	var capital = fund.capital
	fund.capital = capital * (float64(100) + gszzl) / 100
	fund.profit += capital * gszzl / 100
	fund.lastGszzl = gszzl
	fund.Unlock()
}

func (fund *Fund) PrintFundProfit(name string, gszzl float64) {
	var format string
	profit := fund.capital * gszzl / 100
	if profit > 0 {
		format = "[%s] 上次日涨幅为%v, 现在为%v, 收益为\x1b[0;31m%v\x1b[0m\n"
	} else {
		format = "[%s] 上次日涨幅为%v, 现在为%v, 收益为\x1b[0;32m%v\x1b[0m\n"
	}

	fmt.Printf(format, name, fund.lastGszzl, gszzl, profit)
	fund.finance.logger.Info(fmt.Sprintf("[%s] 收益为%v", name, profit))
}

func main() {
	finance := NewFinance(1800)
	finance.AddFund(1001.56, "090010")
	finance.AddFund(306.44, "110003")
	finance.AddFund(200, "001548")
	finance.Run()
}
