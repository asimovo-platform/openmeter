// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// BalanceSnapshot is the predicate function for balancesnapshot builders.
type BalanceSnapshot func(*sql.Selector)

// Entitlement is the predicate function for entitlement builders.
type Entitlement func(*sql.Selector)

// Feature is the predicate function for feature builders.
type Feature func(*sql.Selector)

// Grant is the predicate function for dbgrant builders.
type Grant func(*sql.Selector)

// NotificationChannel is the predicate function for notificationchannel builders.
type NotificationChannel func(*sql.Selector)

// NotificationChannelOrErr calls the predicate only if the error is not nit.
func NotificationChannelOrErr(p NotificationChannel, err error) NotificationChannel {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// NotificationEvent is the predicate function for notificationevent builders.
type NotificationEvent func(*sql.Selector)

// NotificationEventOrErr calls the predicate only if the error is not nit.
func NotificationEventOrErr(p NotificationEvent, err error) NotificationEvent {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// NotificationEventDeliveryStatus is the predicate function for notificationeventdeliverystatus builders.
type NotificationEventDeliveryStatus func(*sql.Selector)

// NotificationRule is the predicate function for notificationrule builders.
type NotificationRule func(*sql.Selector)

// NotificationRuleOrErr calls the predicate only if the error is not nit.
func NotificationRuleOrErr(p NotificationRule, err error) NotificationRule {
	return func(s *sql.Selector) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// UsageReset is the predicate function for usagereset builders.
type UsageReset func(*sql.Selector)
