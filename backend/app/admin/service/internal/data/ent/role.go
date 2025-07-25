// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"kratos-admin/app/admin/service/internal/data/ent/role"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 角色表
type Role struct {
	config `json:"-"`
	// ID of the ent.
	// id
	ID uint32 `json:"id,omitempty"`
	// 创建时间
	CreateTime *time.Time `json:"create_time,omitempty"`
	// 更新时间
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// 删除时间
	DeleteTime *time.Time `json:"delete_time,omitempty"`
	// 状态
	Status *role.Status `json:"status,omitempty"`
	// 创建者ID
	CreateBy *uint32 `json:"create_by,omitempty"`
	// 更新者ID
	UpdateBy *uint32 `json:"update_by,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty"`
	// 租户ID
	TenantID *uint32 `json:"tenant_id,omitempty"`
	// 角色名称
	Name *string `json:"name,omitempty"`
	// 角色标识
	Code *string `json:"code,omitempty"`
	// 上一层角色ID
	ParentID *uint32 `json:"parent_id,omitempty"`
	// 排序ID
	SortID *int32 `json:"sort_id,omitempty"`
	// 分配的菜单列表
	Menus []uint32 `json:"menus,omitempty"`
	// 分配的API列表
	Apis []uint32 `json:"apis,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleQuery when eager-loading is set.
	Edges        RoleEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleEdges holds the relations/edges for other nodes in the graph.
type RoleEdges struct {
	// Parent holds the value of the parent edge.
	Parent *Role `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Role `json:"children,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleEdges) ParentOrErr() (*Role, error) {
	if e.Parent != nil {
		return e.Parent, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: role.Label}
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e RoleEdges) ChildrenOrErr() ([]*Role, error) {
	if e.loadedTypes[1] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Role) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case role.FieldMenus, role.FieldApis:
			values[i] = new([]byte)
		case role.FieldID, role.FieldCreateBy, role.FieldUpdateBy, role.FieldTenantID, role.FieldParentID, role.FieldSortID:
			values[i] = new(sql.NullInt64)
		case role.FieldStatus, role.FieldRemark, role.FieldName, role.FieldCode:
			values[i] = new(sql.NullString)
		case role.FieldCreateTime, role.FieldUpdateTime, role.FieldDeleteTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Role fields.
func (r *Role) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case role.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = uint32(value.Int64)
		case role.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				r.CreateTime = new(time.Time)
				*r.CreateTime = value.Time
			}
		case role.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				r.UpdateTime = new(time.Time)
				*r.UpdateTime = value.Time
			}
		case role.FieldDeleteTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delete_time", values[i])
			} else if value.Valid {
				r.DeleteTime = new(time.Time)
				*r.DeleteTime = value.Time
			}
		case role.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = new(role.Status)
				*r.Status = role.Status(value.String)
			}
		case role.FieldCreateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by", values[i])
			} else if value.Valid {
				r.CreateBy = new(uint32)
				*r.CreateBy = uint32(value.Int64)
			}
		case role.FieldUpdateBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by", values[i])
			} else if value.Valid {
				r.UpdateBy = new(uint32)
				*r.UpdateBy = uint32(value.Int64)
			}
		case role.FieldRemark:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field remark", values[i])
			} else if value.Valid {
				r.Remark = new(string)
				*r.Remark = value.String
			}
		case role.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				r.TenantID = new(uint32)
				*r.TenantID = uint32(value.Int64)
			}
		case role.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				r.Name = new(string)
				*r.Name = value.String
			}
		case role.FieldCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field code", values[i])
			} else if value.Valid {
				r.Code = new(string)
				*r.Code = value.String
			}
		case role.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				r.ParentID = new(uint32)
				*r.ParentID = uint32(value.Int64)
			}
		case role.FieldSortID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort_id", values[i])
			} else if value.Valid {
				r.SortID = new(int32)
				*r.SortID = int32(value.Int64)
			}
		case role.FieldMenus:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field menus", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Menus); err != nil {
					return fmt.Errorf("unmarshal field menus: %w", err)
				}
			}
		case role.FieldApis:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field apis", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.Apis); err != nil {
					return fmt.Errorf("unmarshal field apis: %w", err)
				}
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Role.
// This includes values selected through modifiers, order, etc.
func (r *Role) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryParent queries the "parent" edge of the Role entity.
func (r *Role) QueryParent() *RoleQuery {
	return NewRoleClient(r.config).QueryParent(r)
}

// QueryChildren queries the "children" edge of the Role entity.
func (r *Role) QueryChildren() *RoleQuery {
	return NewRoleClient(r.config).QueryChildren(r)
}

// Update returns a builder for updating this Role.
// Note that you need to call Role.Unwrap() before calling this method if this Role
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Role) Update() *RoleUpdateOne {
	return NewRoleClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Role entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Role) Unwrap() *Role {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Role is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Role) String() string {
	var builder strings.Builder
	builder.WriteString("Role(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	if v := r.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.DeleteTime; v != nil {
		builder.WriteString("delete_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := r.Status; v != nil {
		builder.WriteString("status=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.CreateBy; v != nil {
		builder.WriteString("create_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.UpdateBy; v != nil {
		builder.WriteString("update_by=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.Remark; v != nil {
		builder.WriteString("remark=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.TenantID; v != nil {
		builder.WriteString("tenant_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.Name; v != nil {
		builder.WriteString("name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.Code; v != nil {
		builder.WriteString("code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := r.ParentID; v != nil {
		builder.WriteString("parent_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := r.SortID; v != nil {
		builder.WriteString("sort_id=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("menus=")
	builder.WriteString(fmt.Sprintf("%v", r.Menus))
	builder.WriteString(", ")
	builder.WriteString("apis=")
	builder.WriteString(fmt.Sprintf("%v", r.Apis))
	builder.WriteByte(')')
	return builder.String()
}

// Roles is a parsable slice of Role.
type Roles []*Role
