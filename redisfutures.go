//   Copyright 2009-2012 Joubin Houshyar
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package goredis

import "time"

// FutureKeys
//
type FutureKeys interface {
	Get() ([]string, Error)
	TryGet(timeoutnano time.Duration) (keys []string, error Error, timedout bool)
}
type _futurekeys struct {
	future FutureBytesArray
}

func newFutureKeys(future FutureBytesArray) FutureKeys {
	return _futurekeys{future}
}
func (fvc _futurekeys) Get() (v []string, error Error) {
	gv, err := fvc.future.Get()
	if err != nil {
		return nil, err
	}
	v = make([]string, len(gv))
	for i, key_bytes := range gv {
		v[i] = string(key_bytes)
	}
	return
}
func (fvc _futurekeys) TryGet(ns time.Duration) ([]string, Error, bool) {
	gv, err, timedout := fvc.future.TryGet(ns)
	if timedout || err != nil {
		return nil, err, timedout
	}

	v := make([]string, len(gv))
	for i, key_bytes := range gv {
		v[i] = string(key_bytes)
	}
	return v, nil, timedout
}

// FutureInfo
//
type FutureKVs interface {
	Get() (map[string][]byte, Error)
	TryGet(timeoutnano time.Duration) (keys map[string][]byte, error Error, timedout bool)
}
type _futurekvs struct {
	future FutureBytesArray
}

func newFutureKVs(future FutureBytesArray) FutureKVs {
	return _futurekvs{future}
}

func (fvc _futurekvs) Get() (v map[string][]byte, error Error) {
	gv, err := fvc.future.Get()
	if err != nil {
		return nil, err
	}
	v = parseKVs(gv)
	return v, nil
}
func (fvc _futurekvs) TryGet(ns time.Duration) (map[string][]byte, Error, bool) {
	gv, err, timedout := fvc.future.TryGet(ns)
	if timedout || err != nil {
		return nil, err, timedout
	}
	return parseKVs(gv), nil, timedout
}

func parseKVs(v [][]byte) map[string][]byte{
	if v == nil {
		return nil
	}
	l := len(v)
	result := map[string][]byte{}
	for i := 0; i < l - 1; i = i + 2  {
		result[string(v[i])] = v[i+1]
	}
	return result
}
// FutureInfo
//
type FutureInfo interface {
	Get() (map[string]string, Error)
	TryGet(timeoutnano time.Duration) (keys map[string]string, error Error, timedout bool)
}
type _futureinfo struct {
	future FutureBytes
}

func newFutureInfo(future FutureBytes) FutureInfo {
	return _futureinfo{future}
}
func (fvc _futureinfo) Get() (v map[string]string, error Error) {
	gv, err := fvc.future.Get()
	if err != nil {
		return nil, err
	}
	v = parseInfo(gv)
	return v, nil
}
func (fvc _futureinfo) TryGet(ns time.Duration) (map[string]string, Error, bool) {
	gv, err, timedout := fvc.future.TryGet(ns)
	if timedout || err != nil {
		return nil, err, timedout
	}
	return parseInfo(gv), nil, timedout
}

// FutureKeyType
//
type FutureKeyType interface {
	Get() (KeyType, Error)
	TryGet(timeoutnano time.Duration) (keys KeyType, error Error, timedout bool)
}
type _futurekeytype struct {
	future FutureString
}

func newFutureKeyType(future FutureString) FutureKeyType {
	return _futurekeytype{future}
}
func (fvc _futurekeytype) Get() (v KeyType, error Error) {
	gv, err := fvc.future.Get()
	if err != nil {
		return RT_NONE, err
	}
	v = GetKeyType(gv)
	return v, nil
}
func (fvc _futurekeytype) TryGet(ns time.Duration) (KeyType, Error, bool) {
	gv, err, timedout := fvc.future.TryGet(ns)
	if timedout || err != nil {
		var defv KeyType
		return defv, err, timedout
	}
	return GetKeyType(gv), nil, timedout
}
