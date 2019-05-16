package main

import (
	"fmt"
	"time"
)

func main() {
	//-------------------时间点time.Time-----------------------
	// 格式化当前时间为Y-m-d H:i:s
	fmt.Println("格式化当前时间为Y-m-d H:i:s ", time.Now().Format("2006-01-02 15:04:05"))
	// 当时时间戳
	fmt.Println("当前时间戳:", time.Now().Unix())
	// 时间字符串格式转时间戳(先将字符串时间转为time.Time类型,再变为时间戳)
	dt, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 12:00:00")
	fmt.Println("时间字符串格式转为时间戳:", dt.Unix())
	// 时间戳转时间字符串格式(先将时间戳转为time.Time类型,再格式化为时间字符串)
	dt2 := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("时间戳转时间字符串格式:", dt2)
	// 获取当前时间的年份
	y := time.Now().Year()
	fmt.Println("获取当前时间的年份:", y)

	//-------------------时间段time.Duration--------------------------
	// 1h3m15s表示1小时3分钟15秒,对应转换为多少个hours,minutes,seconds
	tp, _ := time.ParseDuration("1h3m15s")
	fmt.Println(tp.Hours(), tp.Minutes(), tp.Seconds(), tp.String(), tp.Truncate(1000))

	//-------------------------时间运算--------------------------------
	// go语言时间类型为time.Duration
	sec := time.Duration(1) * time.Second
	fmt.Println(sec)
	// 睡眠1秒
	time.Sleep(sec)
	// 延迟1秒
	time.After(sec)
	// 两个时间点的间隔计算,记录从start时间点开始后消耗了多少时间,等价与time() - start
	start := time.Now()
	time.Sleep(sec)
	fmt.Println("耗时为:", time.Since(start))
	// 时间点加上时间段,等价与php的time() + time.Duration
	add := time.Now().Add(time.Duration(10) * time.Minute)
	fmt.Println("当前时间加上10分钟为:", add.Format("2006-01-02 15:04:05"))
	// 两个时间点相差,t1 - t2
	t1, _ := time.Parse("2006-01-02 15:04:05", "2019-05-05 09:00:00")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2019-05-05 12:00:00")
	fmt.Println("t1 - t2 = ", t1.Sub(t2))

	//--------------------------定时器------------------------------------
	//	for {
	//		select {
	//		case <-time.After(time.Duration(1) * time.Second):
	//			fmt.Println("time out.......")
	//		}
	//	}
	a := time.After(2 * time.Second)
	<-a
	fmt.Println("timer receive")

	time.AfterFunc(2*time.Second, func() {
		fmt.Println("timer receive")
	})
}
