// Code generated by go2go; DO NOT EDIT.


//line tables.go2:1
package lingo

//line tables.go2:1
import (
//line tables.go2:1
 "fmt"
//line tables.go2:1
 "reflect"
//line tables.go2:1
 "regexp"
//line tables.go2:1
 "strconv"
//line tables.go2:1
 "strings"
//line tables.go2:1
 "sync/atomic"
//line tables.go2:1
 "testing"
//line tables.go2:1
)

//line tables.go2:5
type Actor struct {
	ActorID  int16
	IsActive bool
}

func NewTActor() TActor {
	return NewTActorForVariable("actor")
}
func NewTActorForVariable(variable string) TActor {
	pm := NewPathMetadataForVariable(variable)
	tablePath := instantiate୦୦NewSimplePath୦lingo୮aActor(Actor{}, pm)
	return NewTActorForPath(tablePath)
}
func NewTActorForPath(p instantiate୦୦SimplePath୦lingo୮aActor,) TActor {
	var t TActor = TActor{
		mixin: p,
	}
	t.ActorID = instantiate୦୦NewNumberPath୦int16(instantiate୦୦NewSimplePathForProperty୦int16(0, "actor_id", p))
	t.IsActive = NewBoolPath(instantiate୦୦NewSimplePathForProperty୦bool(false, "is_active", p))
	return t
}

type TActor struct {
	mixin instantiate୦୦SimplePath୦lingo୮aActor

	ActorID  instantiate୦୦NumberPath୦int16
	IsActive BoolPath
}

//line tables.go2:33
func (t TActor) Type() reflect.Type {
	return t.mixin.Type()
}
func (t TActor) Metadata() PathMetadata {
	return t.mixin.Metadata()
}
func (t TActor) Root() Path {
	return t.mixin.Root()
}
func (t TActor) String() string {
	return instantiate୦୦VisitWithBuilder୦lingo୮aTemplates୦string(ToStringBuilder{}, DefaultTemplates(), t)
}

//line paths.go2:88
func instantiate୦୦NewSimplePath୦lingo୮aActor(value Actor,

//line paths.go2:88
 pm PathMetadata) instantiate୦୦SimplePath୦lingo୮aActor {
	return instantiate୦୦SimplePath୦lingo୮aActor{
		value: value,
		pm:    pm,
	}
}

//line paths.go2:93
type instantiate୦୦SimplePath୦lingo୮aActor struct {
//line paths.go2:95
 value Actor

//line paths.go2:96
 pm PathMetadata
}

//line paths.go2:98
func (p instantiate୦୦SimplePath୦lingo୮aActor,) Type() reflect.Type { return p.pm.Type() }
func (p instantiate୦୦SimplePath୦lingo୮aActor,) String() string {
	return instantiate୦୦VisitWithBuilder୦lingo୮aTemplates୦string(ToStringBuilder{}, DefaultTemplates(), p)
}
func (p instantiate୦୦SimplePath୦lingo୮aActor,) Metadata() PathMetadata { return p.pm }
func (p instantiate୦୦SimplePath୦lingo୮aActor,) Root() Path             { return p.pm.Root() }
//line paths.go2:82
func instantiate୦୦NewSimplePathForProperty୦int16(value int16,

//line paths.go2:82
 propName string, parent Path) instantiate୦୦SimplePath୦int16 {
	return instantiate୦୦NewSimplePath୦int16(value, NewPathMetadataForProperty(propName, parent))
}

//line paths.go2:131
func instantiate୦୦NewNumberPath୦int16(p instantiate୦୦SimplePath୦int16,) instantiate୦୦NumberPath୦int16 {
	return instantiate୦୦NumberPath୦int16{
		mixin: p,
	}
}
//line paths.go2:82
func instantiate୦୦NewSimplePathForProperty୦bool(value bool,

//line paths.go2:82
 propName string, parent Path) instantiate୦୦SimplePath୦bool {
	return instantiate୦୦NewSimplePath୦bool(value, NewPathMetadataForProperty(propName, parent))
}

//line paths.go2:84
type instantiate୦୦NumberPath୦int16 struct {
//line paths.go2:137
 mixin instantiate୦୦SimplePath୦int16
}

//line paths.go2:139
func (p instantiate୦୦NumberPath୦int16,) Type() reflect.Type { return p.mixin.Type() }
func (p instantiate୦୦NumberPath୦int16,) String() string {
	return instantiate୦୦VisitWithBuilder୦lingo୮aTemplates୦string(ToStringBuilder{}, DefaultTemplates(), p)
}
func (p instantiate୦୦NumberPath୦int16,) Metadata() PathMetadata { return p.mixin.Metadata() }
func (p instantiate୦୦NumberPath୦int16,) Root() Path             { return p.mixin.Root() }

func (p instantiate୦୦NumberPath୦int16,) EqValue(v int16,

//line paths.go2:146
) BooleanExpression {
	return p.Eq(instantiate୦୦NewSimpleConstant୦int16(v))
}
func (p instantiate୦୦NumberPath୦int16,) Eq(v instantiate୦୦TypedExpression୦int16,) BooleanExpression {
	return NewOpEqual(p, v)
}
func (p instantiate୦୦NumberPath୦int16,) GTValue(v int16,

//line paths.go2:152
) BooleanExpression {
	return p.GT(instantiate୦୦NewSimpleConstant୦int16(v))
}
func (p instantiate୦୦NumberPath୦int16,) GT(e instantiate୦୦TypedExpression୦int16,) BooleanExpression {
	return NewOpGreaterThan(p, e)
}

//line paths.go2:157
type instantiate୦୦SimplePath୦int16 struct {
//line paths.go2:95
 value int16

//line paths.go2:96
 pm PathMetadata
}

//line paths.go2:98
func (p instantiate୦୦SimplePath୦int16,) Type() reflect.Type { return p.pm.Type() }
func (p instantiate୦୦SimplePath୦int16,) String() string {
	return instantiate୦୦VisitWithBuilder୦lingo୮aTemplates୦string(ToStringBuilder{}, DefaultTemplates(), p)
}
func (p instantiate୦୦SimplePath୦int16,) Metadata() PathMetadata { return p.pm }
func (p instantiate୦୦SimplePath୦int16,) Root() Path             { return p.pm.Root() }
//line paths.go2:88
func instantiate୦୦NewSimplePath୦int16(value int16,

//line paths.go2:88
 pm PathMetadata) instantiate୦୦SimplePath୦int16 {
	return instantiate୦୦SimplePath୦int16{
		value: value,
		pm:    pm,
	}
}

//line expression.go2:168
func instantiate୦୦NewSimpleConstant୦int16(v int16,

//line expression.go2:168
) instantiate୦୦SimpleConstant୦int16 {
	return instantiate୦୦SimpleConstant୦int16{
		value: v,
	}
}

//line expression.go2:172
type instantiate୦୦TypedExpression୦int16 interface {
//line expression.go2:123
 Expression
}
//line expression.go2:124
type instantiate୦୦SimpleConstant୦int16 struct {
//line expression.go2:173
 instantiate୦୦TypedExpression୦int16

	value int16
}

//line expression.go2:177
func (e instantiate୦୦SimpleConstant୦int16,) Type() reflect.Type {
//line expression.go2:177
 return reflect.TypeOf(int16(e.value))
//line expression.go2:177
}
func (e instantiate୦୦SimpleConstant୦int16,) String() string {
	return instantiate୦୦VisitWithBuilder୦lingo୮aTemplates୦string(ToStringBuilder{}, DefaultTemplates(), e)
}
func (e instantiate୦୦SimpleConstant୦int16,) Interface() interface{} { return e.value }

//line expression.go2:181
var _ = fmt.Errorf
//line expression.go2:181
var _ = reflect.Append
//line expression.go2:181
var _ = regexp.Compile
//line expression.go2:181
var _ = strconv.AppendBool

//line expression.go2:181
type _ strings.Builder

//line expression.go2:181
var _ = atomic.AddInt32
//line expression.go2:181
var _ = testing.AllocsPerRun
