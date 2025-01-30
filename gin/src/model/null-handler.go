package model

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/go-sql-driver/mysql"
)

// NullInt64 adalah tipe pembungkus untuk sql.NullInt64 agar dapat digunakan dalam JSON.
type NullInt64 sql.NullInt64

// Scan untuk NullInt64 mengimplementasikan interface sql.Scanner.
func (ni *NullInt64) Scan(value interface{}) error {
	var i sql.NullInt64
	if err := i.Scan(value); err != nil {
		return err
	}
	*ni = NullInt64(i)
	return nil
}

// MarshalJSON untuk NullInt64 mengubah nilai menjadi JSON.
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return json.Marshal(nil) // Nilai tidak valid, kembalikan nil
	}
	return json.Marshal(ni.Int64) // Kembalikan nilai int64 sebagai JSON
}

// NullFloat64 adalah tipe pembungkus untuk sql.NullFloat64 agar dapat digunakan dalam JSON.
type NullFloat64 sql.NullFloat64

// Scan untuk NullFloat64 mengimplementasikan interface sql.Scanner.
func (nf *NullFloat64) Scan(value interface{}) error {
	var f sql.NullFloat64
	if err := f.Scan(value); err != nil {
		return err
	}
	*nf = NullFloat64(f)
	return nil
}

// MarshalJSON untuk NullFloat64 mengubah nilai menjadi JSON.
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return json.Marshal(nil) // Nilai tidak valid, kembalikan nil
	}
	return json.Marshal(nf.Float64) // Kembalikan nilai float64 sebagai JSON
}

// NullString adalah tipe pembungkus untuk sql.NullString agar dapat digunakan dalam JSON.
type NullString sql.NullString

// Scan untuk NullString mengimplementasikan interface sql.Scanner.
func (ns *NullString) Scan(value interface{}) error {
	var s sql.NullString
	if err := s.Scan(value); err != nil {
		return err
	}
	*ns = NullString(s)
	return nil
}

// MarshalJSON untuk NullString mengubah nilai menjadi JSON.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return json.Marshal(nil) // Nilai tidak valid, kembalikan nil
	}
	return json.Marshal(ns.String) // Kembalikan nilai string sebagai JSON
}

// NullTime adalah tipe pembungkus untuk mysql.NullTime agar dapat digunakan dalam JSON.
type NullTime mysql.NullTime

// Scan untuk NullTime mengimplementasikan interface sql.Scanner.
func (nt *NullTime) Scan(value interface{}) error {
	var t mysql.NullTime
	if err := t.Scan(value); err != nil {
		return err
	}
	*nt = NullTime(t)
	return nil
}

// MarshalJSON untuk NullTime mengubah nilai waktu menjadi JSON dengan format RFC3339.
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return json.Marshal(nil) // Nilai tidak valid, kembalikan nil
	}
	return json.Marshal(nt.Time.Format(time.RFC3339)) // Kembalikan waktu dalam format RFC3339
}
