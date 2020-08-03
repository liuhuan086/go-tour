package timer

import "time"

func GetNowTime() time.Time {
	// set time zone
	location, _ := time.LoadLocation("Asia/shanghai")
	return time.Now().In(location)
}

func GetCalculateTime( currentTimer time.Time, d string) (time.Time, error) {
	// 因为不清楚传入的d到底是什么格式，因此使用ParseDuration来预处理
	// 如果预先知道准确的d的且不需要适配，则可以直接调用Add方法
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}

	return currentTimer.Add(duration), err
}