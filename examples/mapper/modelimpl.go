// DO NOT EDIT THIS FILE !
// It is generated by gobatis tool, source from model.go.
package mapper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/arstd/gobatis/examples/domain"
	e "github.com/arstd/gobatis/examples/enums"
	"github.com/wothing/log"
)

type ModelMapperImpl struct{}

func (*ModelMapperImpl) Insert(tx *sql.Tx, m *domain.Model) (err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `insert into models(buildin_bool, buildin_byte, buildin_float32, buildin_float64, buildin_int, buildin_int16, buildin_int32, buildin_int64, buildin_int8, buildin_rune, buildin_string, buildin_uint, buildin_uint16, buildin_uint32, buildin_uint64, buildin_uint8, buildin_map, enum_status, ptr_model, time) values (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) returning id `
	args = append(args, m.BuildinBool)
	args = append(args, m.BuildinByte)
	args = append(args, m.BuildinFloat32)
	args = append(args, m.BuildinFloat64)
	args = append(args, m.BuildinInt)
	args = append(args, m.BuildinInt16)
	args = append(args, m.BuildinInt32)
	args = append(args, m.BuildinInt64)
	args = append(args, m.BuildinInt8)
	args = append(args, m.BuildinRune)
	args = append(args, m.BuildinString)
	args = append(args, m.BuildinUint)
	args = append(args, m.BuildinUint16)
	args = append(args, m.BuildinUint32)
	args = append(args, m.BuildinUint64)
	args = append(args, m.BuildinUint8)
	var x_m_BuildinMap []byte
	x_m_BuildinMap, err = json.Marshal(m.BuildinMap)
	if err != nil {
		log.Error(err)
		return
	}
	args = append(args, x_m_BuildinMap)
	args = append(args, m.EnumStatus)
	var x_m_PtrModel []byte
	x_m_PtrModel, err = json.Marshal(m.PtrModel)
	if err != nil {
		log.Error(err)
		return
	}
	args = append(args, x_m_PtrModel)
	args = append(args, m.Time)
	buf.WriteString(stmt)

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	var dest []interface{}

	dest = append(dest, &m.Id)

	err = db.QueryRow(query, args...).Scan(dest...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}

	return nil

}
func (*ModelMapperImpl) Update(tx *sql.Tx, m *domain.Model) (i int64, err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `update models set buildin_bool=%s, buildin_byte=%s, buildin_float32=%s, buildin_float64=%s, buildin_int=%s, buildin_int16=%s, buildin_int32=%s, buildin_int64=%s, buildin_int8=%s, buildin_rune=%s, buildin_string=%s, buildin_uint=%s, buildin_uint16=%s, buildin_uint32=%s, buildin_uint64=%s, buildin_uint8=%s, buildin_map=%s, enum_status=%s, ptr_model=%s, time=%s where id=%s `
	args = append(args, m.BuildinBool)
	args = append(args, m.BuildinByte)
	args = append(args, m.BuildinFloat32)
	args = append(args, m.BuildinFloat64)
	args = append(args, m.BuildinInt)
	args = append(args, m.BuildinInt16)
	args = append(args, m.BuildinInt32)
	args = append(args, m.BuildinInt64)
	args = append(args, m.BuildinInt8)
	args = append(args, m.BuildinRune)
	args = append(args, m.BuildinString)
	args = append(args, m.BuildinUint)
	args = append(args, m.BuildinUint16)
	args = append(args, m.BuildinUint32)
	args = append(args, m.BuildinUint64)
	args = append(args, m.BuildinUint8)
	var x_m_BuildinMap []byte
	x_m_BuildinMap, err = json.Marshal(m.BuildinMap)
	if err != nil {
		log.Error(err)
		return
	}
	args = append(args, x_m_BuildinMap)
	args = append(args, m.EnumStatus)
	var x_m_PtrModel []byte
	x_m_PtrModel, err = json.Marshal(m.PtrModel)
	if err != nil {
		log.Error(err)
		return
	}
	args = append(args, x_m_PtrModel)
	args = append(args, m.Time)
	args = append(args, m.Id)
	buf.WriteString(stmt)

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return 0, err
	}
	return res.RowsAffected()

}
func (*ModelMapperImpl) Delete(tx *sql.Tx, id int) (i int64, err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `delete from models where id=%s `
	args = append(args, id)
	buf.WriteString(stmt)

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return 0, err
	}
	return res.RowsAffected()
}
func (*ModelMapperImpl) Get(tx *sql.Tx, id int) (m *domain.Model, err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `select id, buildin_bool, buildin_byte, buildin_float32, buildin_float64, buildin_int, buildin_int16, buildin_int32, buildin_int64, buildin_int8, buildin_rune, buildin_string, buildin_uint, buildin_uint16, buildin_uint32, buildin_uint64, buildin_uint8, buildin_map, enum_status, ptr_model, time from models where id=%s `
	args = append(args, id)
	buf.WriteString(stmt)

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	var dest []interface{}
	m = &domain.Model{}
	dest = append(dest, &m.Id)
	dest = append(dest, &m.BuildinBool)
	dest = append(dest, &m.BuildinByte)
	dest = append(dest, &m.BuildinFloat32)
	dest = append(dest, &m.BuildinFloat64)
	dest = append(dest, &m.BuildinInt)
	dest = append(dest, &m.BuildinInt16)
	dest = append(dest, &m.BuildinInt32)
	dest = append(dest, &m.BuildinInt64)
	dest = append(dest, &m.BuildinInt8)
	dest = append(dest, &m.BuildinRune)
	dest = append(dest, &m.BuildinString)
	dest = append(dest, &m.BuildinUint)
	dest = append(dest, &m.BuildinUint16)
	dest = append(dest, &m.BuildinUint32)
	dest = append(dest, &m.BuildinUint64)
	dest = append(dest, &m.BuildinUint8)
	var x_m_BuildinMap []byte
	dest = append(dest, &x_m_BuildinMap)
	dest = append(dest, &m.EnumStatus)
	var x_m_PtrModel []byte
	dest = append(dest, &x_m_PtrModel)
	dest = append(dest, &m.Time)
	err = db.QueryRow(query, args...).Scan(dest...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	m.BuildinMap = map[string]interface{}{}
	err = json.Unmarshal(x_m_BuildinMap, &m.BuildinMap)
	if err != nil {
		log.Error(err)
		return
	}
	m.PtrModel = &domain.Model{}
	err = json.Unmarshal(x_m_PtrModel, &m.PtrModel)
	if err != nil {
		log.Error(err)
		return
	}
	return
}
func (*ModelMapperImpl) Count(tx *sql.Tx, m *domain.Model, ss []e.Status) (i int64, err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `select count(*) from models where buildin_bool=%s `
	args = append(args, m.BuildinBool)
	buf.WriteString(stmt)

	if m.BuildinInt != 0 {
		stmt = `and buildin_int=%s  `
		args = append(args, m.BuildinInt)
		buf.WriteString(stmt)
	}
	if len(ss) != 0 {
		stmt = `and enum_status in (${ss})  `
		stmt = strings.Replace(stmt, "${"+"ss"+"}",
			strings.Repeat(",%s", len(ss))[1:], -1)
		for _, s := range ss {
			args = append(args, int32(s))
		}
		buf.WriteString(stmt)
	}

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	var dest []interface{}
	dest = append(dest, &i)
	err = db.QueryRow(query, args...).Scan(dest...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	return
}
func (*ModelMapperImpl) List(tx *sql.Tx, m *domain.Model, ss []e.Status, offset int, limit int) (ms []*domain.Model, err error) {
	var (
		stmt string
		buf  bytes.Buffer
		args []interface{}
	)

	stmt = `select id, buildin_bool, buildin_byte, buildin_float32, buildin_float64, buildin_int, buildin_int16, buildin_int32, buildin_int64, buildin_int8, buildin_rune, buildin_string, buildin_uint, buildin_uint16, buildin_uint32, buildin_uint64, buildin_uint8, buildin_map, enum_status, ptr_model, time from models where buildin_bool=%s `
	args = append(args, m.BuildinBool)
	buf.WriteString(stmt)

	if m.BuildinInt != 0 {
		stmt = `and buildin_int=%s  `
		args = append(args, m.BuildinInt)
		buf.WriteString(stmt)
	}
	if len(ss) != 0 {
		stmt = `and enum_status in (${ss})  `
		stmt = strings.Replace(stmt, "${"+"ss"+"}",
			strings.Repeat(",%s", len(ss))[1:], -1)
		for _, s := range ss {
			args = append(args, int32(s))
		}
		buf.WriteString(stmt)
	}

	stmt = `order by id offset %s limit %s `
	args = append(args, offset)
	args = append(args, limit)
	buf.WriteString(stmt)

	var ph []interface{}
	for i := range args {
		ph = append(ph, "$"+strconv.Itoa(i+1))
	}

	query := fmt.Sprintf(buf.String(), ph...)

	log.Debug(query)
	log.Debug(args...)

	rows, err := tx.Query(query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()

	var data []*domain.Model
	for rows.Next() {
		x := &domain.Model{}
		data = append(data, x)

		var dest []interface{}
		dest = append(dest, &x.Id)
		dest = append(dest, &x.BuildinBool)
		dest = append(dest, &x.BuildinByte)
		dest = append(dest, &x.BuildinFloat32)
		dest = append(dest, &x.BuildinFloat64)
		dest = append(dest, &x.BuildinInt)
		dest = append(dest, &x.BuildinInt16)
		dest = append(dest, &x.BuildinInt32)
		dest = append(dest, &x.BuildinInt64)
		dest = append(dest, &x.BuildinInt8)
		dest = append(dest, &x.BuildinRune)
		dest = append(dest, &x.BuildinString)
		dest = append(dest, &x.BuildinUint)
		dest = append(dest, &x.BuildinUint16)
		dest = append(dest, &x.BuildinUint32)
		dest = append(dest, &x.BuildinUint64)
		dest = append(dest, &x.BuildinUint8)
		var x_x_BuildinMap []byte
		dest = append(dest, &x_x_BuildinMap)
		dest = append(dest, &x.EnumStatus)
		var x_x_PtrModel []byte
		dest = append(dest, &x_x_PtrModel)
		dest = append(dest, &x.Time)
		err = rows.Scan(dest...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		x.BuildinMap = map[string]interface{}{}
		err = json.Unmarshal(x_x_BuildinMap, &x.BuildinMap)
		if err != nil {
			log.Error(err)
			return
		}
		x.PtrModel = &domain.Model{}
		err = json.Unmarshal(x_x_PtrModel, &x.PtrModel)
		if err != nil {
			log.Error(err)
			return
		}
	}
	if err = rows.Err(); err != nil {
		log.Error(err)
		return nil, err
	}

	return data, nil
}
