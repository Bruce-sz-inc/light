// DO NOT EDIT THIS FILE!
// It is generated by `light` tool by source `model.go` at 2017-03-22 11:57:58.

package mapper

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/arstd/light/example/enum"
	"github.com/arstd/light/example/model"
	"github.com/arstd/log"
	"github.com/lib/pq"
)

type ModelMapperImpl struct{}

func (*ModelMapperImpl) Insert(tx *sql.Tx, m *model.Model) (err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` insert into models(name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32) values (%s, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s) returning id`)
	args = append(args, m.Name)
	args = append(args, m.Flag)
	args = append(args, m.Score)
	xmMap, _ := json.Marshal(m.Map)
	args = append(args, xmMap)
	args = append(args, m.Time)
	args = append(args, pq.Array(m.Array))
	args = append(args, pq.Array(m.Slice))
	args = append(args, m.Status)
	xmPointer, _ := json.Marshal(m.Pointer)
	args = append(args, xmPointer)
	xmStructSlice, _ := json.Marshal(m.StructSlice)
	args = append(args, xmStructSlice)
	args = append(args, m.Uint32)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	dest := make([]interface{}, 1)
	dest[0] = &m.Id
	err = db.QueryRow(query, args...).Scan(dest...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	return
}
func (*ModelMapperImpl) BatchInsert(tx *sql.Tx, ms []*model.Model) (i int64, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` insert into models(uint32, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice) values`)
	for i, m := range ms {
		if i != 0 {
			buf.WriteString(",")
		}
		buf.WriteString(` (%s+888, %s, %s, %s, %s, %s, %s, %s, %s, %s, %s)`)
		args = append(args, i)
		args = append(args, m.Name)
		args = append(args, m.Flag)
		args = append(args, m.Score)
		xmMap, _ := json.Marshal(m.Map)
		args = append(args, xmMap)
		args = append(args, m.Time)
		args = append(args, pq.Array(m.Array))
		args = append(args, pq.Array(m.Slice))
		args = append(args, m.Status)
		xmPointer, _ := json.Marshal(m.Pointer)
		args = append(args, xmPointer)
		xmStructSlice, _ := json.Marshal(m.StructSlice)
		args = append(args, xmStructSlice)
	}

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
	}
	return res.RowsAffected()
}
func (*ModelMapperImpl) Get(tx *sql.Tx, id int) (m *model.Model, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where id=%s`)
	args = append(args, id)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	m = &model.Model{}
	dest := make([]interface{}, 12)
	dest[0] = &m.Id
	dest[1] = &m.Name
	dest[2] = &m.Flag
	dest[3] = &m.Score
	var mxMap []byte
	dest[4] = &mxMap
	var mxTime pq.NullTime
	dest[5] = &mxTime
	dest[6] = pq.Array(&m.Array)
	dest[7] = pq.Array(&m.Slice)
	dest[8] = &m.Status
	var mxPointer []byte
	dest[9] = &mxPointer
	var mxStructSlice []byte
	dest[10] = &mxStructSlice
	dest[11] = &m.Uint32
	err = db.QueryRow(query, args...).Scan(dest...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	m.Map = map[string]interface{}{}
	json.Unmarshal(mxMap, m.Map)
	m.Time = mxTime.Time
	m.Pointer = &model.Model{}
	json.Unmarshal(mxPointer, &m.Pointer)
	m.StructSlice = []*model.Model{}
	json.Unmarshal(mxStructSlice, &m.StructSlice)
	return
}
func (*ModelMapperImpl) Update(tx *sql.Tx, m *model.Model) (i int64, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` update models set name=%s, flag=%s, score=%s, map=%s, time=%s, slice=%s, status=%s, pointer=%s, struct_slice=%s, uint32=%s where id=%s`)
	args = append(args, m.Name)
	args = append(args, m.Flag)
	args = append(args, m.Score)
	xmMap, _ := json.Marshal(m.Map)
	args = append(args, xmMap)
	args = append(args, m.Time)
	args = append(args, pq.Array(m.Slice))
	args = append(args, m.Status)
	xmPointer, _ := json.Marshal(m.Pointer)
	args = append(args, xmPointer)
	xmStructSlice, _ := json.Marshal(m.StructSlice)
	args = append(args, xmStructSlice)
	args = append(args, m.Uint32)
	args = append(args, m.Id)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
	}
	return res.RowsAffected()
}
func (*ModelMapperImpl) Delete(tx *sql.Tx, id int) (i int64, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` delete from models where id=%s`)
	args = append(args, id)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Error(query)
		log.Error(args...)
		log.Error(err)
	}
	return res.RowsAffected()
}
func (*ModelMapperImpl) Count(tx *sql.Tx, m *model.Model, ss []enum.Status) (i int64, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` select count(*) from models where name like %s`)
	args = append(args, m.Name)
	if m.Flag != false {
		buf.WriteString(` and flag=%s`)
		args = append(args, m.Flag)
	}
	if len(ss) != 0 {
		buf.WriteString(` and status in (`)
		for i, v := range ss {
			if i != 0 {
				buf.WriteString(",")
			}
			buf.WriteString(` %s`)
			args = append(args, v)
		}
		buf.WriteString(` )`)
	}
	if len(m.Slice) != 0 {
		buf.WriteString(` and slice && %s`)
		args = append(args, pq.Array(m.Slice))
	}

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	err = db.QueryRow(query, args...).Scan(&i)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
	}
	return
}
func (*ModelMapperImpl) List(tx *sql.Tx, m *model.Model, ss []enum.Status, offset int, limit int) (ms []*model.Model, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` select id, name, flag, score, map, time, xarray, slice, status, pointer, struct_slice, uint32 from models where name like %s`)
	args = append(args, m.Name)
	if m.Flag != false {
		if len(ss) != 0 {
			buf.WriteString(` and status in (`)
			for i, v := range ss {
				if i != 0 {
					buf.WriteString(",")
				}
				buf.WriteString(` %s`)
				args = append(args, v)
			}
			buf.WriteString(` )`)
		}
		buf.WriteString(` and flag=%s`)
		args = append(args, m.Flag)
	}
	if len(m.Array) != 0 {
		buf.WriteString(` and xarray && array[`)
		for i, v := range m.Array {
			if i != 0 {
				buf.WriteString(",")
			}
			buf.WriteString(` %s`)
			xv, _ := json.Marshal(v)
			args = append(args, xv)
		}
		buf.WriteString(` ]`)
	}
	if len(m.Slice) != 0 {
		buf.WriteString(` and slice && %s`)
		args = append(args, pq.Array(m.Slice))
	}
	buf.WriteString(` order by id offset %s limit %s`)
	args = append(args, offset)
	args = append(args, limit)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	log.Debug(query)
	log.Debug(args...)
	var rows *sql.Rows
	rows, err = db.Query(query, args...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	defer rows.Close()

	ms = []*model.Model{}
	for rows.Next() {
		elem := &model.Model{}
		ms = append(ms, elem)
		dest := make([]interface{}, 12)
		dest[0] = &elem.Id
		dest[1] = &elem.Name
		dest[2] = &elem.Flag
		dest[3] = &elem.Score
		var elemxMap []byte
		dest[4] = &elemxMap
		var elemxTime pq.NullTime
		dest[5] = &elemxTime
		dest[6] = pq.Array(&elem.Array)
		dest[7] = pq.Array(&elem.Slice)
		dest[8] = &elem.Status
		var elemxPointer []byte
		dest[9] = &elemxPointer
		var elemxStructSlice []byte
		dest[10] = &elemxStructSlice
		dest[11] = &elem.Uint32
		err = rows.Scan(dest...)
		if err != nil {
			log.Error(err)
			return
		}
		elem.Map = map[string]interface{}{}
		json.Unmarshal(elemxMap, elem.Map)
		elem.Time = elemxTime.Time
		elem.Pointer = &model.Model{}
		json.Unmarshal(elemxPointer, &elem.Pointer)
		elem.StructSlice = []*model.Model{}
		json.Unmarshal(elemxStructSlice, &elem.StructSlice)
	}
	if err = rows.Err(); err != nil {
		log.Error(err)
		return
	}
	return
}
func (*ModelMapperImpl) Paging(tx *sql.Tx, m *model.Model, ss []enum.Status, offset int, limit int) (i int64, ms []*model.Model, err error) {

	var (
		buf  bytes.Buffer
		args []interface{}
	)
	buf.WriteString(` select id, name, flag, score, map, time, slice, status, pointer, struct_slice from models where name like %s`)
	args = append(args, m.Name)
	if m.Flag != false {
		if len(ss) != 0 {
			buf.WriteString(` and status in (`)
			for i, v := range ss {
				if i != 0 {
					buf.WriteString(",")
				}
				buf.WriteString(` %s`)
				args = append(args, v)
			}
			buf.WriteString(` )`)
		}
		buf.WriteString(` and flag=%s`)
		args = append(args, m.Flag)
	}
	if len(m.Slice) != 0 {
		buf.WriteString(` and slice && %s`)
		args = append(args, pq.Array(m.Slice))
	}
	buf.WriteString(` order by id offset %s limit %s`)
	args = append(args, offset)
	args = append(args, limit)

	ph := make([]interface{}, len(args))
	for i := range args {
		ph[i] = fmt.Sprintf("$%d", i+1)
	}
	query := fmt.Sprintf(buf.String(), ph...)
	fidx := strings.LastIndex(query, " from ")
	obidx := strings.LastIndex(query, "order by")
	tQuery := `select count(*)` + query[fidx:obidx]
	dNum := strings.Count(query[obidx:], "$")
	tArgs := args[:len(args)-dNum]
	log.Debug(tQuery)
	log.Debug(tArgs...)
	err = db.QueryRow(tQuery, tArgs...).Scan(&i)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
	}
	if i == 0 {
		return
	}
	log.Debug(query)
	log.Debug(args...)
	var rows *sql.Rows
	rows, err = db.Query(query, args...)
	if err != nil {
		log.Error(err)
		log.Error(query)
		log.Error(args...)
		return
	}
	defer rows.Close()

	ms = []*model.Model{}
	for rows.Next() {
		elem := &model.Model{}
		ms = append(ms, elem)
		dest := make([]interface{}, 10)
		dest[0] = &elem.Id
		dest[1] = &elem.Name
		dest[2] = &elem.Flag
		dest[3] = &elem.Score
		var elemxMap []byte
		dest[4] = &elemxMap
		var elemxTime pq.NullTime
		dest[5] = &elemxTime
		dest[6] = pq.Array(&elem.Slice)
		dest[7] = &elem.Status
		var elemxPointer []byte
		dest[8] = &elemxPointer
		var elemxStructSlice []byte
		dest[9] = &elemxStructSlice
		err = rows.Scan(dest...)
		if err != nil {
			log.Error(err)
			return
		}
		elem.Map = map[string]interface{}{}
		json.Unmarshal(elemxMap, elem.Map)
		elem.Time = elemxTime.Time
		elem.Pointer = &model.Model{}
		json.Unmarshal(elemxPointer, &elem.Pointer)
		elem.StructSlice = []*model.Model{}
		json.Unmarshal(elemxStructSlice, &elem.StructSlice)
	}
	if err = rows.Err(); err != nil {
		log.Error(err)
		return
	}
	return
}
