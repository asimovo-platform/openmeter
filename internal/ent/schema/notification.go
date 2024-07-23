package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/openmeterio/openmeter/internal/notification"
	notificationpostgres "github.com/openmeterio/openmeter/internal/notification/repository/postgres"
	"github.com/openmeterio/openmeter/pkg/framework/entutils"
)

type NotificationChannel struct {
	ent.Schema
}

func (NotificationChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.TimeMixin{},
	}
}

func (NotificationChannel) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			GoType(notification.ChannelType("")).
			Immutable(),
		field.String("name").
			NotEmpty(),
		field.Bool("disabled").
			Default(false).
			Optional(),
		field.String("config").
			GoType(notification.ChannelConfig{}).
			ValueScanner(notificationpostgres.ChannelConfigValueScanner).
			SchemaType(map[string]string{
				dialect.Postgres: "jsonb",
			}),
	}
}

func (NotificationChannel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("rules", NotificationRule.Type),
	}
}

func (NotificationChannel) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("namespace", "id"),
		index.Fields("namespace", "type"),
		index.Fields("namespace", "id", "type"),
	}
}

type NotificationRule struct {
	ent.Schema
}

func (NotificationRule) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entutils.IDMixin{},
		entutils.NamespaceMixin{},
		entutils.TimeMixin{},
	}
}

func (NotificationRule) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			GoType(notification.RuleType("")).
			Immutable().
			Comment("The event type the rule associated with"),
		field.String("name").
			NotEmpty().
			Comment("The name of the rule"),
		field.Bool("disabled").Default(false).Optional().
			Comment("Whether the rule is disabled or not"),
		field.String("config").
			GoType(notification.RuleConfig{}).
			ValueScanner(notificationpostgres.RuleConfigValueScanner).
			SchemaType(map[string]string{
				dialect.Postgres: "jsonb",
			}),
	}
}

func (NotificationRule) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("channels", NotificationChannel.Type).
			Ref("rules"),
	}
}
