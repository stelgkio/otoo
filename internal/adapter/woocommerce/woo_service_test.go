package woocommerce

import "testing"

func TestWoocommerceService_WoocommerceCreateAllWebHook(t *testing.T) {
	type args struct {
		customerKey    string
		customerSecret string
		domainUrl      string
		projectId      string
	}
	tests := []struct {
		name    string
		s       *WoocommerceService
		args    args
		wantErr bool
	}{
		{
			name: "Valid Inputs",
			s:    &WoocommerceService{},
			args: args{
				customerKey:    "valid_customer_key",
				customerSecret: "valid_customer_secret",
				domainUrl:      "https://example.com",
				projectId:      "valid_project_id",
			},
			wantErr: false,
		},
		{
			name: "Empty Domain URL",
			s:    &WoocommerceService{},
			args: args{
				customerKey:    "valid_customer_key",
				customerSecret: "valid_customer_secret",
				domainUrl:      "",
				projectId:      "valid_project_id",
			},
			wantErr: true,
		},
		{
			name: "Invalid HTTPS URL",
			s:    &WoocommerceService{},
			args: args{
				customerKey:    "valid_customer_key",
				customerSecret: "valid_customer_secret",
				domainUrl:      "http://example.com", // Not HTTPS
				projectId:      "valid_project_id",
			},
			wantErr: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.s.WoocommerceCreateAllWebHook(tt.args.customerKey, tt.args.customerSecret, tt.args.domainUrl, tt.args.projectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("WoocommerceService.WoocommerceCreateAllWebHook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
