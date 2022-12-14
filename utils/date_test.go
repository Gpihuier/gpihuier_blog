package utils_test

import (
	"testing"
	"time"

	"github.com/Gpihuier/gpihuier_blog/utils"
)

func TestTimeAgo(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"小于1分钟",
			args{time.Now().Add(-30 * time.Second)},
			"约 1 分钟前",
		},
		{
			"约2小时前",
			args{time.Now().Add(-100 * time.Minute)},
			"约 2 小时前",
		},
		{
			"约1天前",
			args{time.Now().Add(-25 * time.Hour)},
			"约 1 天前",
		},
		{
			"2017-11-02",
			args{time.Date(2017, 11, 2, 14, 0, 0, 0, time.Local)},
			"2017-11-02",
		},
		{
			"2016-02-02",
			args{time.Date(2016, 2, 2, 0, 0, 0, 0, time.Local)},
			"2016-02-02",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := utils.TimeAgo(tt.args.t); got != tt.want {
				t.Errorf("TimeAgo() = %v, want %v", got, tt.want)
			}
		})
	}
}
