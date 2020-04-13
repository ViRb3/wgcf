/* SPDX-License-Identifier: MIT
 *
 * Copyright (C) 2019 WireGuard LLC. All Rights Reserved.
 */
// https://github.com/WireGuard/wireguard-windows/blob/master/conf/config.go

package wireguard

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

const KeyLength = 32

type Key [KeyLength]byte

func (k *Key) String() string {
	return base64.StdEncoding.EncodeToString(k[:])
}

func (k *Key) IsZero() bool {
	var zeros Key
	return subtle.ConstantTimeCompare(zeros[:], k[:]) == 1
}

func (k *Key) Public() *Key {
	var p [KeyLength]byte
	curve25519.ScalarBaseMult(&p, (*[KeyLength]byte)(k))
	return (*Key)(&p)
}

func NewPresharedKey() (*Key, error) {
	var k [KeyLength]byte
	_, err := rand.Read(k[:])
	if err != nil {
		return nil, err
	}
	return (*Key)(&k), nil
}

func NewPrivateKey() (*Key, error) {
	k, err := NewPresharedKey()
	if err != nil {
		return nil, err
	}
	k[0] &= 248
	k[31] = (k[31] & 127) | 64
	return k, nil
}

func NewKey(base64Key string) (*Key, error) {
	k, err := base64.StdEncoding.DecodeString(base64Key)
	if err != nil {
		return nil, err
	}
	var key Key
	copy(key[:], k)
	return &key, nil
}
