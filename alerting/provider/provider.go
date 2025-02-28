package provider

import (
	"github.com/pschlump/gatus/alerting/alert"
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
	"github.com/pschlump/gatus/core"
)

// AlertProvider is the interface that each providers should implement
type AlertProvider interface {
	// IsValid returns whether the provider's configuration is valid
	IsValid() bool

	// GetDefaultAlert returns the provider's default alert configuration
	GetDefaultAlert() *alert.Alert

	// Send an alert using the provider
	Send(endpoint *core.Endpoint, alert *alert.Alert, result *core.Result, resolved bool) error
}

// ParseWithDefaultAlert parses an Endpoint alert by using the provider's default alert as a baseline
func ParseWithDefaultAlert(providerDefaultAlert, endpointAlert *alert.Alert) {
	if providerDefaultAlert == nil || endpointAlert == nil {
		return
	}
	if endpointAlert.Enabled == nil {
		endpointAlert.Enabled = providerDefaultAlert.Enabled
	}
	if endpointAlert.SendOnResolved == nil {
		endpointAlert.SendOnResolved = providerDefaultAlert.SendOnResolved
	}
	if endpointAlert.Description == nil {
		endpointAlert.Description = providerDefaultAlert.Description
	}
	if endpointAlert.FailureThreshold == 0 {
		endpointAlert.FailureThreshold = providerDefaultAlert.FailureThreshold
	}
	if endpointAlert.SuccessThreshold == 0 {
		endpointAlert.SuccessThreshold = providerDefaultAlert.SuccessThreshold
	}
}

var (
	// Validate interface implementation on compile
	_ AlertProvider = (*custom.AlertProvider)(nil)
	_ AlertProvider = (*discord.AlertProvider)(nil)
	_ AlertProvider = (*email.AlertProvider)(nil)
	_ AlertProvider = (*github.AlertProvider)(nil)
	_ AlertProvider = (*googlechat.AlertProvider)(nil)
	_ AlertProvider = (*matrix.AlertProvider)(nil)
	_ AlertProvider = (*mattermost.AlertProvider)(nil)
	_ AlertProvider = (*messagebird.AlertProvider)(nil)
	_ AlertProvider = (*ntfy.AlertProvider)(nil)
	_ AlertProvider = (*opsgenie.AlertProvider)(nil)
	_ AlertProvider = (*pagerduty.AlertProvider)(nil)
	_ AlertProvider = (*slack.AlertProvider)(nil)
	_ AlertProvider = (*teams.AlertProvider)(nil)
	_ AlertProvider = (*telegram.AlertProvider)(nil)
	_ AlertProvider = (*twilio.AlertProvider)(nil)
)
