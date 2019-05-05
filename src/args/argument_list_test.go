package args

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		name      string
		arg       *ArgumentList
		wantError bool
	}{
		{
			"No Errors",
			&ArgumentList{
				Name:           "name",
				Username:       "user",
				Password:       "password",
				Hostname:       "localhost",
				Port:           "90",
				CollectionList: "{}",
			},
			false,
		},
		{
			"No Name",
			&ArgumentList{
				Name:           "",
				Username:       "user",
				Password:       "password",
				Hostname:       "localhost",
				Port:           "90",
				CollectionList: "{}",
			},
			true,
		},
		{
			"No Username",
			&ArgumentList{
				Name:           "name",
				Username:       "",
				Password:       "password",
				Hostname:       "localhost",
				Port:           "90",
				CollectionList: "{}",
			},
			true,
		},
		{
			"No Password",
			&ArgumentList{
				Name:           "name",
				Username:       "user",
				Hostname:       "localhost",
				Port:           "90",
				CollectionList: "{}",
			},
			true,
		},
		{
			"SSL and No Server Certificate",
			&ArgumentList{
				Name:                   "name",
				Username:               "user",
				Password:               "password",
				Hostname:               "localhost",
				Port:                   "90",
				EnableSSL:              true,
				TrustServerCertificate: false,
				SSLRootCertLocation:    "",
				CollectionList:         "{}",
			},
			true,
		},
		{
			"Missing Key file with Cert file",
			&ArgumentList{
				Name:                   "name",
				Username:               "user",
				Password:               "password",
				Hostname:               "localhost",
				Port:                   "90",
				EnableSSL:              true,
				TrustServerCertificate: true,
				SSLKeyLocation:         "",
				SSLCertLocation:        "my.crt",
				CollectionList:         "{}",
			},
			true,
		},
		{
			"Missing Cert file with Key file",
			&ArgumentList{
				Name:                   "name",
				Username:               "user",
				Password:               "password",
				Hostname:               "localhost",
				Port:                   "90",
				EnableSSL:              true,
				TrustServerCertificate: true,
				SSLKeyLocation:         "my.key",
				SSLCertLocation:        "",
				CollectionList:         "{}",
			},
			true,
		},
	}

	for _, tc := range testCases {
		err := tc.arg.Validate()
		if tc.wantError && err == nil {
			t.Errorf("Test Case %s Failed: Expected error", tc.name)
		} else if !tc.wantError && err != nil {
			t.Errorf("Test Case %s Failed: Unexpected error: %v", tc.name, err)
		}
	}

}
