package video

import (
	"movie-sync-server/conf"
	"testing"
)

func TestGetUrl(t *testing.T) {
	conf.Init()

	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestGetUrl",
			args: args{
				url: "https://www.iqiyi.com/v_1bbf0eo65vg.html?r_area=pcw_rec_like&r_source=56&bkt=tpfsfallrerank_10%3Btp_fsfall_rule_05%3Btpfsfallrank_09%3Btp_fsfall_prerank_05&e=e1bc43f656d43bb97d40d85be075748e&stype=2&vfrm=pcw_home&vfrmblk=712211_cainizaizhui&vfrmrst=712211_cainizaizhui_float_video_area2",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUrl(tt.args.url); got != tt.want {
				t.Errorf("GetUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
