package main

import (
	"github.com/miekg/pkcs11"
	"fmt"
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

	/*aesGcm := GenerateGcmKey(p, &session)

	fmt.Println("aes key", aesGcm)
*/
	keySearch := []*pkcs11.Attribute{pkcs11.NewAttribute(pkcs11.CKA_LABEL, "sample aes key")}
	p.FindObjectsInit(session, keySearch)
	obj, b, err:= p.FindObjects(session, 100)

	fmt.Println("b-", b)
	fmt.Println("obj lengths - ", len(obj))
	p.FindObjectsFinal(session)

	if err != nil {
		panic(err)
	}
	for _,o := range obj {
		fmt.Println(" aes obj handle ",o)
		ed := EncryptByGcm(p, session, []byte("this is test data"), o)
		fmt.Println("gcm encrypted data - ",ed)

		dec := DecryptByGcm(p, session, ed, o)
		fmt.Println("gcm decrypted data - ",string(dec))

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
