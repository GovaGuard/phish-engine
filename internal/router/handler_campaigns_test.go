package router

//
// func TestRouter_AddCampaign(t *testing.T) {
// 	usc := usecase.New(c repository.CampaignRepo, t repository.TargetsRepo, m mail.Sender)
//
// 	type fields struct {
// 		usecase *usecase.Usecase
// 	}
// 	type args struct {
// 		body string
// 	}
// 	tests := []struct {
// 		name       string
// 		fields     fields
// 		args       args
// 		wantStatus int
// 	}{
// 		{
// 			name:   "Simple",
// 			fields: fields{}, args: args{},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			body := strings.NewReader(tt.args.body)
// 			request := httptest.NewRequest(http.MethodPost, "", body)
//
// 			rr := httptest.NewRecorder()
//
// 			router := &Router{
// 				usecase: tt.fields.usecase,
// 			}
//
// 			router.AddCampaign(rr, request)
//
// 			result := rr.Result()
// 			if result.StatusCode != tt.wantStatus {
// 				t.Fail()
// 			}
// 		})
// 	}
// }
