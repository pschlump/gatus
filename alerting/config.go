package alerting

import (
	"log"
	"reflect"
	"strings"

	"github.com/pschlump/gatus/alerting/alert"
	"github.com/pschlump/gatus/alerting/provider"
	"github.com/pschlump/gatus/alerting/provider/custom"
	"github.com/pschlump/gatus/alerting/provider/discord"
	"github.com/pschlump/gatus/alerting/provider/email"
	"github.com/pschlump/gatus/alerting/provider/github"
	"github.com/pschlump/gatus/alerting/provider/googlechat"
	"github.com/pschlump/gatus/alerting/provider/matrix"
	"github.com/pschlump/gatus/alerting/provider/mattermost"
	"github.com/pschlump/gatus/alerting/provider/messagebird"
	"github.com/pschlump/gatus/alerting/provider/ntfy"
	"github.com/pschlump/gatus/alerting/provider/opsgenie"
	"github.com/pschlump/gatus/alerting/provider/pagerduty"
	"github.com/pschlump/gatus/alerting/provider/slack"
	"github.com/pschlump/gatus/alerting/provider/teams"
	"github.com/pschlump/gatus/alerting/provider/telegram"
	"github.com/pschlump/gatus/alerting/provider/twilio"
)

// Config is the configuration for alerting providers
type Config struct {
	// Custom is the configuration for the custom alerting provider
	Custom *custom.AlertProvider `yaml:"custom,omitempty"`

	// Discord is the configuration for the discord alerting provider
	Discord *discord.AlertProvider `yaml:"discord,omitempty"`

	// Email is the configuration for the email alerting provider
	Email *email.AlertProvider `yaml:"email,omitempty"`

	// GitHub is the configuration for the github alerting provider
	GitHub *github.AlertProvider `yaml:"github,omitempty"`

	// GoogleChat is the configuration for the googlechat alerting provider
	GoogleChat *googlechat.AlertProvider `yaml:"googlechat,omitempty"`

	// Matrix is the configuration for the matrix alerting provider
	Matrix *matrix.AlertProvider `yaml:"matrix,omitempty"`

	// Mattermost is the configuration for the mattermost alerting provider
	Mattermost *mattermost.AlertProvider `yaml:"mattermost,omitempty"`

	// Messagebird is the configuration for the messagebird alerting provider
	Messagebird *messagebird.AlertProvider `yaml:"messagebird,omitempty"`

	// Ntfy is the configuration for the ntfy alerting provider
	Ntfy *ntfy.AlertProvider `yaml:"ntfy,omitempty"`

	// Opsgenie is the configuration for the opsgenie alerting provider
	Opsgenie *opsgenie.AlertProvider `yaml:"opsgenie,omitempty"`

	// PagerDuty is the configuration for the pagerduty alerting provider
	PagerDuty *pagerduty.AlertProvider `yaml:"pagerduty,omitempty"`

	// Slack is the configuration for the slack alerting provider
	Slack *slack.AlertProvider `yaml:"slack,omitempty"`

	// Teams is the configuration for the teams alerting provider
	Teams *teams.AlertProvider `yaml:"teams,omitempty"`

	// Telegram is the configuration for the telegram alerting provider
	Telegram *telegram.AlertProvider `yaml:"telegram,omitempty"`

	// Twilio is the configuration for the twilio alerting provider
	Twilio *twilio.AlertProvider `yaml:"twilio,omitempty"`
}

// GetAlertingProviderByAlertType returns an provider.AlertProvider by its corresponding alert.Type
func (config *Config) GetAlertingProviderByAlertType(alertType alert.Type) provider.AlertProvider {
	entityType := reflect.TypeOf(config).Elem()
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		if strings.ToLower(field.Name) == string(alertType) {
			fieldValue := reflect.ValueOf(config).Elem().Field(i)
			if fieldValue.IsNil() {
				return nil
			}
			return fieldValue.Interface().(provider.AlertProvider)
		}
	}
	log.Printf("[alerting][GetAlertingProviderByAlertType] No alerting provider found for alert type %s", alertType)
	return nil
}

// SetAlertingProviderToNil Sets an alerting provider to nil to avoid having to revalidate it every time an
// alert of its corresponding type is sent.
func (config *Config) SetAlertingProviderToNil(p provider.AlertProvider) {
	entityType := reflect.TypeOf(config).Elem()
	for i := 0; i < entityType.NumField(); i++ {
		field := entityType.Field(i)
		if field.Type == reflect.TypeOf(p) {
			reflect.ValueOf(config).Elem().Field(i).Set(reflect.Zero(field.Type))
		}
	}
}
