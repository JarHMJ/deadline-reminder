package feishu

import "testing"

func TestFeishu_Notify(t *testing.T) {
	type fields struct {
		BaseUrl string
	}
	type args struct {
		msg string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{name: "test", fields: fields{BaseUrl: "https://open.feishu.cn/open-apis/bot/v2/hook/xxxxxxxx"}, args: args{msg: "aa"}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Feishu{
				baseUrl: tt.fields.BaseUrl,
			}
			if err := f.Notify(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Notify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
