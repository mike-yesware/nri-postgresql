// Package args contains the argument list, defined as a struct, along with a method that validates passed-in args
package args

import (
	"encoding/json"
	"errors"
	"fmt"

	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
)

// ArgumentList struct that holds all PostgreSQL arguments
type ArgumentList struct {
	sdkArgs.DefaultArgumentList
	Name                   string `default:"" help:"The name used to identify this databse within New Relic"`
	Username               string `default:"" help:"The username for the PostgreSQL database"`
	Password               string `default:"" help:"The password for the specified username"`
	Hostname               string `default:"localhost" help:"The PostgreSQL hostname to connect to"`
	Database               string `default:"postgres" help:"The PostgreSQL database name to connect to"`
	Port                   string `default:"5432" help:"The port to connect to the PostgreSQL database"`
	CollectionList         string `default:"{}" help:"A JSON object which defines the databases, schemas, tables, and indexes to collect. Collects nothing by default."`
	EnableSSL              bool   `default:"false" help:"If true will use SSL encryption, false will not use encryption"`
	TrustServerCertificate bool   `default:"false" help:"If true server certificate is not verified for SSL. If false certificate will be verified against supplied certificate"`
	Pgbouncer              bool   `default:"false" help:"Collects metrics from PgBouncer instance. Assumes connection is through PgBouncer."`
	SSLRootCertLocation    string `default:"" help:"Absolute path to PEM encoded root certificate file"`
	SSLCertLocation        string `default:"" help:"Absolute path to PEM encoded client cert file"`
	SSLKeyLocation         string `default:"" help:"Absolute path to PEM encoded client key file"`
	Timeout                string `default:"10" help:"Maximum wait for connection, in seconds. Set 0 for no timeout"`
}

// DatabaseList is a map from database name to SchemaLists to collect
type DatabaseList map[string]SchemaList

// SchemaList is a map from schema name to TableList to collect
type SchemaList map[string]TableList

// TableList is a map from table name to an array of indexes to collect
type TableList map[string][]string

// Validate validates PostgreSQl arguments
func (al ArgumentList) Validate() error {
	if al.Name == "" {
		return errors.New("invalid configuration: must specify a database name for identification")
	}

	if al.Username == "" || al.Password == "" {
		return errors.New("invalid configuration: must specify a username and password")
	}

	if err := al.validateSSL(); err != nil {
		return err
	}

	var dl DatabaseList
	err := json.Unmarshal([]byte(al.CollectionList), &dl)
	if err != nil {
		return fmt.Errorf("invalid configuration: failed to unmarshal CollectionList JSON: %s", err.Error())
	}

	return nil
}

func (al ArgumentList) validateSSL() error {
	if al.EnableSSL {
		if !al.TrustServerCertificate && al.SSLRootCertLocation == "" {
			return errors.New("invalid configuration: must specify a certificate file when using SSL and not trusting server certificate")
		}

		if al.SSLCertLocation == "" || al.SSLKeyLocation == "" {
			return errors.New("invalid configuration: must specify both a client cert and key file when enabling SSL")
		}
	}

	return nil
}

// GetCollectionList unmarshals the argument collection list into a DatabaseList
func (al ArgumentList) GetCollectionList() DatabaseList {
	var dl DatabaseList
	_ = json.Unmarshal([]byte(al.CollectionList), &dl) // already checked in Validate()

	return dl
}
