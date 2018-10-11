package main

import (
	"github.com/miekg/pkcs11"
	"fmt"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"math/big"
	"github.com/btcsuite/btcd/btcec"
	"encoding/hex"
	"bytes"
	"github.com/btcsuite/btcutil"
)

func main() {
	p := pkcs11.New("/usr/local/Cellar/softhsm/2.4.0/lib/softhsm/libsofthsm2.so")
	err := p.Initialize()
	if err != nil {
		panic(err)
	}

	defer p.Destroy()
	defer p.Finalize()

	slots, err := p.GetSlotList(true)

	fmt.Println("number of slots ",slots)

	if err != nil {
		panic(err)
	}

	session, err := p.OpenSession(slots[0], pkcs11.CKF_SERIAL_SESSION|pkcs11.CKF_RW_SESSION)
	if err != nil {
		panic(err)
	}
	defer p.CloseSession(session)

	err = p.Login(session, pkcs11.CKU_USER, "5678")
	if err != nil {
		panic(err)
	}
	defer p.Logout(session)

	sourceTestNet, err := GetTestNetAddressPubKeyFromHsm("my ec key new", p, session)
	destinationTestNet, err := GetTestNetAddressPubKeyFromHsm("my ec key new 2", p, session)
	fmt.Println("--- destination address - ",destinationTestNet)

	if err != nil {
		fmt.Println("new address pub key err ", err)
	}
	fmt.Println("--- source address - ", sourceTestNet.EncodeAddress())

	sourceUtxoHash, _ := chainhash.NewHashFromStr("3d4f5905b1889342a99c7e7f855412ba812e426ecee33f0e0b759b24e4d262a1")

	if err != nil {
		fmt.Println("destination decode err ", err)
	}
	if err != nil {
		fmt.Println("source decode err ", err)
	}
	//destinationPkScript, _ := txscript.PayToAddrScript(destinationAddress)


	scriptHex := "76a914abeef2c797a4ade55c42c38edcd6b111f98c723688ac"
	script, err := hex.DecodeString(scriptHex)
	if err != nil {
		fmt.Println(err)
		return
	}
	redeemTx := wire.NewMsgTx(wire.TxVersion)
	prevOut := wire.NewOutPoint(sourceUtxoHash, 0)
	txIn := wire.NewTxIn(prevOut, nil, nil)
	redeemTx.AddTxIn(txIn)

	fmt.Printf("type is - %T\n",destinationTestNet)
	script1, err := txscript.PayToAddrScript(destinationTestNet.AddressPubKeyHash())
	if err != nil {
		fmt.Println("error creating script1- ",err)
	}
	txOut := wire.NewTxOut(5003730, script1)

	redeemTx.AddTxOut(txOut)
	script2, err := txscript.PayToAddrScript(sourceTestNet.AddressPubKeyHash())
	if err != nil {
		fmt.Println("error creating script2- ",err)
	}
	txOut = wire.NewTxOut(4003730, script2)
	redeemTx.AddTxOut(txOut)

	hash, err := txscript.CalcSignatureHash(script, txscript.SigHashAll, redeemTx, 0)
	sigScript := SignByPrivateKey("my ec key new", p, session, hash)

	finalScript, err := txscript.NewScriptBuilder().AddData(sigScript).AddData(sourceTestNet.ScriptAddress()).Script()
	redeemTx.TxIn[0].SignatureScript = finalScript

	flags := txscript.ScriptBip16 | txscript.ScriptVerifyDERSignatures |
		txscript.ScriptStrictMultiSig |
		txscript.ScriptDiscourageUpgradableNops
	vm, err := txscript.NewEngine(script, redeemTx, 0,
		flags, nil, nil, -1)
	if err != nil {
		fmt.Println("-------------------122--", err)
		return
	}
	if err := vm.Execute(); err != nil {
		fmt.Println("-------------------1223--", err)
		return
	}
	var b bytes.Buffer
	err = redeemTx.Serialize(&b)
	if err != nil {
		fmt.Println("===========", err)
	}
	fmt.Printf("--- %x\n", b.Bytes())
	fmt.Println("Transaction successfully signed")

}

func GetTestNetAddressPubKeyFromHsm(label string, context *pkcs11.Ctx, session pkcs11.SessionHandle)(*btcutil.AddressPubKey, error) {
	keySearch := []*pkcs11.Attribute{pkcs11.NewAttribute(pkcs11.CKA_LABEL, label), pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PUBLIC_KEY)}
	context.FindObjectsInit(session, keySearch)
	obj, b, err := context.FindObjects(session, 100)

	fmt.Println("b-", b)
	fmt.Println("obj lengths - ", len(obj))
	context.FindObjectsFinal(session)
	o := obj[0]
	fmt.Println(" rsa object ", o)
	attribute, err := context.GetAttributeValue(session, o, []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_EC_POINT, nil)})
	if err != nil {
		fmt.Printf("err getting value of rsa:%d", o)
		fmt.Println(" err:", err)
		return nil, err
	} else {
		return btcutil.NewAddressPubKey(attribute[0].Value[2:], &chaincfg.TestNet3Params)
	}
}

func SignByPrivateKey(label string, context *pkcs11.Ctx, session pkcs11.SessionHandle, text []byte) []byte {
	keySearch := []*pkcs11.Attribute{pkcs11.NewAttribute(pkcs11.CKA_LABEL, label),	pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY)}
	context.FindObjectsInit(session, keySearch)
	obj, b, err := context.FindObjects(session, 100)
	context.FindObjectsFinal(session)
	if err != nil {
		fmt.Println("no such private key exists", err)
		return nil
	} else {
		fmt.Println("b-", b)
		fmt.Println("obj lengths - ", len(obj))
		err := context.SignInit(session, []*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_ECDSA, nil)}, obj[0])
		if err != nil {
			fmt.Println("failed to initialize signing operation: ", err)
		}
		signature, err := context.Sign(session, text)
		if err != nil {
			fmt.Println("failed to sign data: ", err)
		}
		fmt.Println("Signature is: ",signature)
		fmt.Println("Signature byte length is: ",len(signature))
		fmt.Printf("Signature (in hex) is: %x\n",signature)

		context.SignFinal(session)
		r := &big.Int{}
		r = r.SetBytes(signature[:32])
		fmt.Println("----------", len(r.Bytes()))
		s := &big.Int{}
		s = s.SetBytes(signature[32:])
		fmt.Println("----------", len(s.Bytes()))
		sig := &btcec.Signature{R:r, S:s}
		sigB := append(sig.Serialize(), byte(txscript.SigHashAll))
		return sigB
	}
}

func GenerateGcmKey(p *pkcs11.Ctx, session *pkcs11.SessionHandle) pkcs11.ObjectHandle {
	template := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, "sample aes key"),
		pkcs11.NewAttribute(pkcs11.CKA_ENCRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_DECRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_SECRET_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_AES),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
		pkcs11.NewAttribute(pkcs11.CKA_WRAP, true),
		pkcs11.NewAttribute(pkcs11.CKA_UNWRAP, true),
		pkcs11.NewAttribute(pkcs11.CKA_PRIVATE, true),
		pkcs11.NewAttribute(pkcs11.CKA_DERIVE, true),
		pkcs11.NewAttribute(pkcs11.CKA_VALUE_LEN, 32),
		pkcs11.NewAttribute(pkcs11.CKA_EXTRACTABLE, false),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN,  true),
	}
	m := []*pkcs11.Mechanism{
		pkcs11.NewMechanism(pkcs11.CKM_AES_KEY_GEN, nil),
	}
	handler, err := p.GenerateKey(*session, m, template)
	if err != nil {
		fmt.Println(err)
	}
	return handler
}

func EncryptByGcm(p *pkcs11.Ctx, session pkcs11.SessionHandle, data []byte, o pkcs11.ObjectHandle) []byte{
	gcmParams := pkcs11.NewGCMParams(make([]byte, 16), []byte("something"), 128)

	m := []*pkcs11.Mechanism{
		pkcs11.NewMechanism(pkcs11.CKM_AES_GCM, gcmParams),
	}
	err := p.EncryptInit(session, m, o)
	if err != nil {
		panic(err)
	}
	encrypted, err := p.Encrypt(session, data)

	if err != nil {
		panic(err)
	}
	p.EncryptFinal(session)
	return encrypted
}

func DecryptByGcm(p *pkcs11.Ctx, session pkcs11.SessionHandle, data []byte, o pkcs11.ObjectHandle) []byte {
	gcmParams := pkcs11.NewGCMParams(make([]byte, 16), []byte("something"), 128)

	m := []*pkcs11.Mechanism{
		pkcs11.NewMechanism(pkcs11.CKM_AES_GCM, gcmParams),
	}

	err := p.DecryptInit(session, m ,o)
	if err != nil {
		panic(err)
	}
	decrypted, err := p.Decrypt(session, data)

	if err != nil {
		panic(err)
	}
	p.EncryptFinal(session)
	return decrypted
}


func generateRSAKeyPair(p *pkcs11.Ctx, session pkcs11.SessionHandle, tokenLabel string) (pkcs11.ObjectHandle, pkcs11.ObjectHandle) {
	publicKeyTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PUBLIC_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_RSA),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, true),
		pkcs11.NewAttribute(pkcs11.CKA_VERIFY, true),
		pkcs11.NewAttribute(pkcs11.CKA_ENCRYPT, true),
		pkcs11.NewAttribute(pkcs11.CKA_PUBLIC_EXPONENT, []byte{1, 0, 1}),
		pkcs11.NewAttribute(pkcs11.CKA_MODULUS_BITS, 2048),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, tokenLabel),
	}
	privateKeyTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_RSA),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, true),
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, true),
		pkcs11.NewAttribute(pkcs11.CKA_SIGN, true),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, tokenLabel),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
		pkcs11.NewAttribute(pkcs11.CKA_EXTRACTABLE, true),
		pkcs11.NewAttribute(pkcs11.CKA_DECRYPT, true),
	}
	pbk, pvk, e := p.GenerateKeyPair(session,
		[]*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_RSA_PKCS_KEY_PAIR_GEN, nil)},
		publicKeyTemplate, privateKeyTemplate)
	if e != nil {
		fmt.Println(e)
	}

	return pbk, pvk
}


func GenerateEliticalCurveKeyPair(p *pkcs11.Ctx, session pkcs11.SessionHandle, tokenLabel string) (pkcs11.ObjectHandle, pkcs11.ObjectHandle) {
	publicKeyTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PUBLIC_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, true),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, tokenLabel),
		pkcs11.NewAttribute(pkcs11.CKA_EC_PARAMS, []byte("\x06\x05+\x81\x04\x00\n")),
	}

	privateKeyTemplate := []*pkcs11.Attribute{
		pkcs11.NewAttribute(pkcs11.CKA_CLASS, pkcs11.CKO_PRIVATE_KEY),
		pkcs11.NewAttribute(pkcs11.CKA_KEY_TYPE, pkcs11.CKK_EC),
		pkcs11.NewAttribute(pkcs11.CKA_TOKEN, true),
		pkcs11.NewAttribute(pkcs11.CKA_LABEL, tokenLabel),
		pkcs11.NewAttribute(pkcs11.CKA_SENSITIVE, true),
	}
	pbk, pvk, e := p.GenerateKeyPair(session,
	[]*pkcs11.Mechanism{pkcs11.NewMechanism(pkcs11.CKM_EC_KEY_PAIR_GEN, nil)},
	publicKeyTemplate, privateKeyTemplate)
	if e != nil {
	fmt.Println(e)
	}

	return pbk, pvk
}