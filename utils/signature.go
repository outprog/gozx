package utils

import (
	"crypto/ecdsa"
	"errors"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core"
)

func Signature(key *ecdsa.PrivateKey, orderHash string, signType int) (signHex string, err error) {
	var sign []byte

	switch signType {
	case SIGNTYPE_Illegal:
	case SIGNTYPE_Invalid:
	case SIGNTYPE_EIP712:
	case SIGNTYPE_EthSign:
		sign, err = ethSign(key, orderHash)
	case SIGNTYPE_Wallet:
	case SIGNTYPE_Validator:
	case SIGNTYPE_PreSigned:
	default:
		return signHex, errors.New("Not support signatrue")
	}

	signHex = common.ToHex(sign)
	return
}

func IsValidSignture(orderHash string, makerAddr common.Address, signature string) (isValid bool) {
	sign := common.FromHex(signature)
	if len(sign) < 1 {
		return
	}

	signType, err := strconv.Atoi(common.Bytes2Hex(sign[len(sign)-1:]))
	if err != nil {
		return
	}

	switch signType {
	case SIGNTYPE_Illegal:
	case SIGNTYPE_Invalid:
	case SIGNTYPE_EIP712:
	case SIGNTYPE_EthSign:
		isValid = ethSignValidator(orderHash, makerAddr, sign[:len(sign)-1])
	case SIGNTYPE_Wallet:
	case SIGNTYPE_Validator:
	case SIGNTYPE_PreSigned:
	}

	return
}

func ethSign(key *ecdsa.PrivateKey, orderHash string) (sign []byte, err error) {
	hash, _ := core.SignHash(common.FromHex(orderHash))
	sign, err = crypto.Sign(hash, key)
	if err != nil {
		return
	}

	sign[64] += 27
	sign = append(sign[len(sign)-1:], sign[:len(sign)-1]...)
	sign = append(sign, common.FromHex(strconv.Itoa(SIGNTYPE_EthSign))...)

	return
}

func ethSignValidator(orderHash string, address common.Address, sign []byte) (isValid bool) {
	if len(sign) != 65 {
		return
	}

	if sign[0] != 27 && sign[0] != 28 {
		return
	}
	sign[0] -= 27

	sign = append(sign[1:], sign[:1]...)
	hash, _ := core.SignHash(common.FromHex(orderHash))

	recoverPub, err := crypto.Ecrecover(hash, sign)
	if err != nil {
		return
	}
	pubKey, err := crypto.UnmarshalPubkey(recoverPub)
	if err != nil {
		return
	}
	addr := crypto.PubkeyToAddress(*pubKey)

	isValid = addr == address
	return
}
