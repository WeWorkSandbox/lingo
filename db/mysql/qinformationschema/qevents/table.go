// Code generated by Lingo for table information_schema.EVENTS - DO NOT EDIT

package qevents

import (
	"github.com/weworksandbox/lingo/pkg/core"
	"github.com/weworksandbox/lingo/pkg/core/path"
)

func As(alias string) QEvents {
	return newQEvents(alias)
}

func New() QEvents {
	return newQEvents("")
}

func newQEvents(alias string) QEvents {
	q := QEvents{_alias: alias}
	q.eventCatalog = path.NewStringPath(q, "EVENT_CATALOG")
	q.eventSchema = path.NewStringPath(q, "EVENT_SCHEMA")
	q.eventName = path.NewStringPath(q, "EVENT_NAME")
	q.definer = path.NewStringPath(q, "DEFINER")
	q.timeZone = path.NewStringPath(q, "TIME_ZONE")
	q.eventBody = path.NewStringPath(q, "EVENT_BODY")
	q.eventDefinition = path.NewStringPath(q, "EVENT_DEFINITION")
	q.eventType = path.NewStringPath(q, "EVENT_TYPE")
	q.executeAt = path.NewTimePath(q, "EXECUTE_AT")
	q.intervalValue = path.NewStringPath(q, "INTERVAL_VALUE")
	q.intervalField = path.NewStringPath(q, "INTERVAL_FIELD")
	q.sqlMode = path.NewStringPath(q, "SQL_MODE")
	q.starts = path.NewTimePath(q, "STARTS")
	q.ends = path.NewTimePath(q, "ENDS")
	q.status = path.NewStringPath(q, "STATUS")
	q.onCompletion = path.NewStringPath(q, "ON_COMPLETION")
	q.created = path.NewTimePath(q, "CREATED")
	q.lastAltered = path.NewTimePath(q, "LAST_ALTERED")
	q.lastExecuted = path.NewTimePath(q, "LAST_EXECUTED")
	q.eventComment = path.NewStringPath(q, "EVENT_COMMENT")
	q.originator = path.NewInt64Path(q, "ORIGINATOR")
	q.characterSetClient = path.NewStringPath(q, "CHARACTER_SET_CLIENT")
	q.collationConnection = path.NewStringPath(q, "COLLATION_CONNECTION")
	q.databaseCollation = path.NewStringPath(q, "DATABASE_COLLATION")
	return q
}

type QEvents struct {
	_alias              string
	eventCatalog        path.StringPath
	eventSchema         path.StringPath
	eventName           path.StringPath
	definer             path.StringPath
	timeZone            path.StringPath
	eventBody           path.StringPath
	eventDefinition     path.StringPath
	eventType           path.StringPath
	executeAt           path.TimePath
	intervalValue       path.StringPath
	intervalField       path.StringPath
	sqlMode             path.StringPath
	starts              path.TimePath
	ends                path.TimePath
	status              path.StringPath
	onCompletion        path.StringPath
	created             path.TimePath
	lastAltered         path.TimePath
	lastExecuted        path.TimePath
	eventComment        path.StringPath
	originator          path.Int64Path
	characterSetClient  path.StringPath
	collationConnection path.StringPath
	databaseCollation   path.StringPath
}

// core.Table Functions

func (q QEvents) GetColumns() []core.Column {
	return []core.Column{
		q.eventCatalog,
		q.eventSchema,
		q.eventName,
		q.definer,
		q.timeZone,
		q.eventBody,
		q.eventDefinition,
		q.eventType,
		q.executeAt,
		q.intervalValue,
		q.intervalField,
		q.sqlMode,
		q.starts,
		q.ends,
		q.status,
		q.onCompletion,
		q.created,
		q.lastAltered,
		q.lastExecuted,
		q.eventComment,
		q.originator,
		q.characterSetClient,
		q.collationConnection,
		q.databaseCollation,
	}
}

func (q QEvents) GetSQL(d core.Dialect) (core.SQL, error) {
	return path.ExpandTableWithDialect(d, q)
}

func (q QEvents) GetAlias() string {
	return q._alias
}

func (q QEvents) GetName() string {
	return "EVENTS"
}

func (q QEvents) GetParent() string {
	return "information_schema"
}

// Column Functions

func (q QEvents) EventCatalog() path.StringPath {
	return q.eventCatalog
}

func (q QEvents) EventSchema() path.StringPath {
	return q.eventSchema
}

func (q QEvents) EventName() path.StringPath {
	return q.eventName
}

func (q QEvents) Definer() path.StringPath {
	return q.definer
}

func (q QEvents) TimeZone() path.StringPath {
	return q.timeZone
}

func (q QEvents) EventBody() path.StringPath {
	return q.eventBody
}

func (q QEvents) EventDefinition() path.StringPath {
	return q.eventDefinition
}

func (q QEvents) EventType() path.StringPath {
	return q.eventType
}

func (q QEvents) ExecuteAt() path.TimePath {
	return q.executeAt
}

func (q QEvents) IntervalValue() path.StringPath {
	return q.intervalValue
}

func (q QEvents) IntervalField() path.StringPath {
	return q.intervalField
}

func (q QEvents) SqlMode() path.StringPath {
	return q.sqlMode
}

func (q QEvents) Starts() path.TimePath {
	return q.starts
}

func (q QEvents) Ends() path.TimePath {
	return q.ends
}

func (q QEvents) Status() path.StringPath {
	return q.status
}

func (q QEvents) OnCompletion() path.StringPath {
	return q.onCompletion
}

func (q QEvents) Created() path.TimePath {
	return q.created
}

func (q QEvents) LastAltered() path.TimePath {
	return q.lastAltered
}

func (q QEvents) LastExecuted() path.TimePath {
	return q.lastExecuted
}

func (q QEvents) EventComment() path.StringPath {
	return q.eventComment
}

func (q QEvents) Originator() path.Int64Path {
	return q.originator
}

func (q QEvents) CharacterSetClient() path.StringPath {
	return q.characterSetClient
}

func (q QEvents) CollationConnection() path.StringPath {
	return q.collationConnection
}

func (q QEvents) DatabaseCollation() path.StringPath {
	return q.databaseCollation
}
