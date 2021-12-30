package usecase

import (
	"testing"
)

func TestDSGApiService_Login(t *testing.T) {
	//util.AssertEqual(t, getMD5Key(2, 1606445062), "9db3d172c2ea12ac9aebddbba257469a", "md5 doesn't match the result")

	//param := getLoginParam("dsgdev_ssp_624762318", 10)
	//util.AssertEqual(t, param, "Qylp4ZZuc1Olg1TOiduCZjf/LUHLcGd0Z85KptzSvy1SYPoL0BOEn0fHDcZQW5xm4O3AJXitGPlF2ioOD7lb4TnhAmDjGUNQ72eMPDfQXtg=", "param doesn't match the result")

	//postURL := getPostURL("103.103.81.3", 8200, 2, param, 1606451021)
	//postURLResult := "http://103.103.81.3:8200/channelHandle?agent=2&timestamp=1606451021&param=Qylp4ZZuc1Olg1TOiduCZjf%2FLUHLcGd0Z85KptzSvy1SYPoL0BOEn0fHDcZQW5xm4O3AJXitGPlF2ioOD7lb4TnhAmDjGUNQ72eMPDfQXtg%3D&key=c40a3e44e6c724ed999b5551d31c30fe"
	//util.AssertEqual(t, postURL, postURLResult, "post doesn't match the result")
	type args struct {
		loginDomain string
		agentID     int
		account     string
		gameID      int32
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test_login",
			args: args{
				loginDomain: "103.103.81.3",
				agentID:  2,
				account: "dsgdev_ssp_624762318",
				gameID:  10,
			},
			want: "",
			wantErr: false,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DSGApiService{}
			got, err := ds.Login(tt.args.loginDomain, tt.args.agentID, tt.args.account, tt.args.gameID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Login() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Login() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDSGApiService_StoreMoney(t *testing.T) {
	type args struct {
		loginDomain string
		agentID     int
		account     string
		money       int32
	}
	tests := []struct {
		name    string
		args    args
		want    int32
		wantErr bool
	}{
		{
			name: "test_login",
			args: args{
				loginDomain: "103.103.81.3",
				agentID:  2,
				account: "dsgdev_ssp_624762318",
				money:  10,
			},
			want: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ds := &DSGApiService{}
			got, err := ds.StoreMoney(tt.args.loginDomain, tt.args.agentID, tt.args.account, tt.args.money)
			if (err != nil) != tt.wantErr {
				t.Errorf("StoreMoney() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("StoreMoney() got = %v, want %v", got, tt.want)
			}
		})
	}
}
