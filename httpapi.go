package httpapi

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/binary"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"time"
)

type SearchSPONSOR_Message struct {
	Appid []byte `json:"appid"`
	Time  []byte `json:"emit"`
	Data  []byte `json:"data"`
	Addr  []byte `json:"addr"`
}

type auth struct {
	Username string `json:"username"`
	Pwd      string `json:"password"`
}
type AdminCreateToken_Message struct {
	Appid     []byte `json:"appid"`
	Time      []byte `json:"emit"`
	Data      []byte `json:"data"`
	AdminAddr []byte `json:"adminaddr"`
	Number    []byte `json:"number"`
	Creater   []byte `json:"creater"`
}
type GetNftIndex_Message struct {
	Appid []byte `json:"appid"`
	Time  []byte `json:"emit"`
	Data  []byte `json:"data"`
}
type UserRegite_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Data     []byte `json:"data"`
	Password []byte `json:"password"`
}
type AdminTransferTokenBatch_Message struct {
	Appid      []byte `json:"appid"`
	Time       []byte `json:"emit"`
	Data       []byte `json:"data"`
	AdminAddr  []byte `json:"adminaddr"`
	Creaters   []byte `json:"creaters"`
	Ids        []byte `json:"ids"`
	AddrNumber []byte `json:"addrnumber"`
}
type UserUpdataPassword_Message struct {
	Appid       []byte `json:"appid"`
	Time        []byte `json:"emit"`
	Data        []byte `json:"data"`
	Addr        []byte `json:"password"`
	OldPassword []byte `json:"oldpassword"`
	NewPassword []byte `json:"newpassword"`
}
type SafeTransferFrom_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Data     []byte `json:"data"`
	Form     []byte `json:"from"`
	To       []byte `json:"to"`
	Id       []byte `json:"id"`
	Number   []byte `json:"number"`
	Password []byte `json:"password"`
}
type SafeBatchTransferFrom_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Data     []byte `json:"data"`
	Form     []byte `json:"from"`
	To       []byte `json:"to"`
	Ids      []byte `json:"ids"`
	Numbers  []byte `json:"numbers"`
	Password []byte `json:"password"`
}
type BanlaceOf_Message struct {
	Appid []byte `json:"appid"`
	Time  []byte `json:"emit"`
	Data  []byte `json:"data"`
	Addr  []byte `json:"addr"`
	Id    []byte `json:"id"`
}
type BanlaceOfBatch_Message struct {
	Appid      []byte `json:"appid"`
	Time       []byte `json:"emit"`
	Data       []byte `json:"data"`
	Addr       []byte `json:"addr"`
	Id         []byte `json:"id"`
	AddrNumber []byte `json:"addrnumber"`
}

//Event res
type safeTransferFromEvent struct {
	Type     string `json:"type"`
	Txhash   string `json:"txhash"`
	Operator string `json:"operator"`
	From     string `json:"from"`
	To       string `json:"to"`
	Id       string `json:"id"`
	Value    string `json:"value"`
}
type safeBatchTransferFromEvent struct {
	Type     string    `json:"type"`
	Txhash   string    `json:"txhash"`
	Operator string    `json:"operator"`
	From     string    `json:"from"`
	To       string    `json:"to"`
	Ids      []big.Int `json:"ids"`
	Values   []big.Int `json:"values"`
}
type adminCreateNFT struct {
	Type     string `json:"type"`
	Txhash   string `json:"txhash"`
	Operator string `json:"operator"`
	Minter   string `json:"minter"`
	Amount   string `json:"amount"`
	NFTID    string `json:"NFTID"`
}
type RespbalanceOfBatch struct {
	Code    string    `json:"code"`
	Msg     []big.Int `json:"msg"`
	TimeSub int64     `json:"timesub"`
}
type RespNor struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	TimeSub int64  `json:"timesub"`
}
type Resp struct {
	Code    string         `json:"code"`
	Msg     adminCreateNFT `json:"msg"`
	TimeSub int64          `json:"timesub"`
}
type Resp_trs struct {
	Code    string                `json:"code"`
	Msg     safeTransferFromEvent `json:"msg"`
	TimeSub int64                 `json:"timesub"`
}
type Resp_trsBatch struct {
	Code    string                     `json:"code"`
	Msg     safeBatchTransferFromEvent `json:"msg"`
	TimeSub int64                      `json:"timesub"`
}
type GetNFTCreator_Message struct {
	Appid []byte `json:"appid"`
	Time  []byte `json:"emit"`
	Data  []byte `json:"data"`
	Id    []byte `json:"id"`
}
type AdminCreateTokenBatch_Message struct {
	Appid      []byte `json:"appid"`
	Time       []byte `json:"emit"`
	Data       []byte `json:"data"`
	AdminAddr  []byte `json:"adminaddr"`
	Creaters   []byte `json:"creaters"`
	AddrNumber []byte `json:"addrnumber"`
}
type BurnFrom_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Data     []byte `json:"data"`
	Form     []byte `json:"from"`
	Id       []byte `json:"id"`
	Number   []byte `json:"number"`
	Password []byte `json:"password"`
}
type BurnFromBatch_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Data     []byte `json:"data"`
	Form     []byte `json:"from"`
	Ids      []byte `json:"ids"`
	Password []byte `json:"password"`
}
type Uri_Message struct {
	Appid []byte `json:"appid"`
	Time  []byte `json:"emit"`
	Data  []byte `json:"data"`
	ID    []byte `json:"ID"`
}

func Testget(thurl string) {
	//get??????
	//http.Get?????????????????????http://??????????????????url,????????????????????????
	resp, err := http.Get(thurl + "/login?username=admin&password=123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error")
	}

	//fmt.Println(string(body))
	fmt.Printf("Get request result: %s\n", string(body))
}

type PopPrk_Message struct {
	Appid    []byte `json:"appid"`
	Time     []byte `json:"emit"`
	Password []byte `json:"password"`
	Addr     []byte `json:"addr"`
}

func PostWithJson_TongBuPrivateKeyToZKVerse(thurl string, actionName string, myappid string, Password string, Addr string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_password := publicEncode([]byte(Password), "public.pem")
	src_addr := publicEncode([]byte(Addr), "public.pem")

	//post????????????json??????
	messages := PopPrk_Message{src_appid, src_mytime, src_password, src_addr}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

///???
func PostWithJson_AdminCreateNFT(thurl string, actionName string, myappid string, Adminaddr string, addr string, nu uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("createnft")       //??????
	NUM := make([]byte, 8)              //????????????
	binary.BigEndian.PutUint64(NUM, nu) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_number := publicEncode(NUM, "public.pem")
	src_creator := publicEncode([]byte(addr), "public.pem")
	src_admin := publicEncode([]byte(Adminaddr), "public.pem")

	//post????????????json??????
	messages := AdminCreateToken_Message{src_appid, src_mytime, src_mydata, src_admin, src_number, src_creator}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//????????????1???NFT
func PostWithJson_AdminCreateNFTBatch(thurl string, actionName string, myappid string, Adminaddr string, addrs []string, addrsNumber uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("AdminCreateNFTBatch")
	//??????
	NUM := make([]byte, 8)                       //????????????
	binary.BigEndian.PutUint64(NUM, addrsNumber) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_number := publicEncode(NUM, "public.pem")

	var addrs_one []byte
	for i := 0; i < len(addrs); i++ {
		ADDR := []byte(addrs[i])
		addrs_one = append(addrs_one, ADDR...)
	}

	src_myaddrs, err := publicEncodeLong(addrs_one, "public.pem")
	if err != nil {
		fmt.Println("RSA??????????????????", err)
		panic(err)
	}
	src_admin := publicEncode([]byte(Adminaddr), "public.pem")
	// fmt.Println(string(src_myaddrs))
	//post????????????json??????
	messages := AdminCreateTokenBatch_Message{src_appid, src_mytime, src_mydata, src_admin, src_myaddrs, src_number}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//????????????NFT
func PostWithJson_AdminTransferNFTBatch(thurl string, actionName string, myappid string, Adminaddr string, addrs []string, ids []uint64, addrsNumber uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("AdminTransferNFTBatch")
	//??????
	NUM := make([]byte, 8)                       //????????????
	binary.BigEndian.PutUint64(NUM, addrsNumber) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_number := publicEncode(NUM, "public.pem")

	var addrs_one []byte
	for i := 0; i < len(addrs); i++ {
		ADDR := []byte(addrs[i])
		addrs_one = append(addrs_one, ADDR...)
	}
	var ids_one []byte
	for i := 0; i < len(ids); i++ {
		ID := make([]byte, 8)                  //????????????
		binary.BigEndian.PutUint64(ID, ids[i]) //uint64?????????
		ids_one = append(ids_one, ID...)
	}
	src_myaddrs, err := publicEncodeLong(addrs_one, "public.pem")
	if err != nil {
		fmt.Println("RSA??????????????????", err)
		panic(err)
	}
	src_ids_one, err := publicEncodeLong(ids_one, "public.pem")
	src_admin := publicEncode([]byte(Adminaddr), "public.pem")
	// fmt.Println(string(src_myaddrs))
	//post????????????json??????
	messages := AdminTransferTokenBatch_Message{src_appid, src_mytime, src_mydata, src_admin, src_myaddrs, src_ids_one, src_number}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//????????????NFT
func PostWithJson_Burn(thurl string, actionName string, myappid string, from string, id uint64, number uint64, password string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("Burn")
	ID := make([]byte, 8)                   //????????????
	binary.BigEndian.PutUint64(ID, id)      //uint64?????????
	NUM := make([]byte, 8)                  //????????????
	binary.BigEndian.PutUint64(NUM, number) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")

	src_from := publicEncode([]byte(from), "public.pem")
	src_id := publicEncode(ID, "public.pem")
	src_number := publicEncode(NUM, "public.pem")
	src_password := publicEncode([]byte(password), "public.pem")
	//post????????????json??????
	messages := BurnFrom_Message{src_appid, src_mytime, src_mydata, src_from, src_id, src_number, src_password}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//????????????NFT??????????????????
func PostWithJson_BurnBatch(thurl string, actionName string, myappid string, addrs string, ids []uint64, password string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????

	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("BurnBatch")

	var ids_one []byte
	for i := 0; i < len(ids); i++ {
		ID := make([]byte, 8)                  //????????????
		binary.BigEndian.PutUint64(ID, ids[i]) //uint64?????????
		ids_one = append(ids_one, ID...)
	}
	ADDR := []byte(addrs)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddrs := publicEncode(ADDR, "public.pem")
	src_myids := publicEncode(ids_one, "public.pem")
	src_password := publicEncode([]byte(password), "public.pem")
	//post????????????json??????
	messages := BurnFromBatch_Message{src_appid, src_mytime, src_mydata, src_myaddrs, src_myids, src_password}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//?????? NFT ???matadata  uri
func PostWithJson_Uri(thurl string, actionName string, myappid string, ids uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("Uri")
	nu := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(nu, ids) //uint64?????????
	myid := []byte(nu)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myid := publicEncode(myid, "public.pem")

	//post????????????json??????
	messages := Uri_Message{src_appid, src_mytime, src_mydata, src_myid}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//?????? ?????????????????????NFT
func PostWithJson_UserNFTs(thurl string, actionName string, myappid string, addr string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("UserNFTs")
	myaddr := []byte(addr)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddr := publicEncode(myaddr, "public.pem")

	//post????????????json??????
	messages := SearchSPONSOR_Message{src_appid, src_mytime, src_mydata, src_myaddr}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//?????? NFT ????????????
func PostWithJson_OwnerOf(thurl string, actionName string, myappid string, ids uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("Uri")
	nu := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(nu, ids) //uint64?????????
	myid := []byte(nu)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myid := publicEncode(myid, "public.pem")

	//post????????????json??????
	messages := Uri_Message{src_appid, src_mytime, src_mydata, src_myid}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}
func PostWithJson_safeTransfer(thurl string, actionName string, myappid string, from string, to string, id uint64, number uint64, password string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("safeTransfer")
	ID := make([]byte, 8)                   //????????????
	binary.BigEndian.PutUint64(ID, id)      //uint64?????????
	NUM := make([]byte, 8)                  //????????????
	binary.BigEndian.PutUint64(NUM, number) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")

	src_from := publicEncode([]byte(from), "public.pem")
	src_to := publicEncode([]byte(to), "public.pem")
	src_id := publicEncode(ID, "public.pem")
	src_number := publicEncode(NUM, "public.pem")
	src_password := publicEncode([]byte(password), "public.pem")
	//post????????????json??????
	messages := SafeTransferFrom_Message{src_appid, src_mytime, src_mydata, src_from, src_to, src_id, src_number, src_password}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal 1error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}
func PostWithJson_userregit(thurl string, actionName string, myappid string, pass string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("userregit")
	var nu uint64 = 1                   //??????
	NUM := make([]byte, 8)              //????????????
	binary.BigEndian.PutUint64(NUM, nu) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_password := publicEncode([]byte(pass), "public.pem")
	var a string
	var b string
	var c string

	json.Unmarshal(src_appid, &a)
	json.Unmarshal(src_mytime, &b)
	json.Unmarshal(src_mydata, &c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//post????????????json??????
	messages := UserRegite_Message{src_appid, src_mytime, src_mydata, src_password}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}
func PostWithJson_userupdatapassword(thurl string, actionName string, myappid string, addr string, oldpassword string, newpassword string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("userregit")
	var nu uint64 = 1                   //??????
	NUM := make([]byte, 8)              //????????????
	binary.BigEndian.PutUint64(NUM, nu) //uint64?????????
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_addr := publicEncode([]byte(addr), "public.pem")
	src_oldpassword := publicEncode([]byte(oldpassword), "public.pem")
	src_newpassword := publicEncode([]byte(newpassword), "public.pem")
	var a string
	var b string
	var c string

	json.Unmarshal(src_appid, &a)
	json.Unmarshal(src_mytime, &b)
	json.Unmarshal(src_mydata, &c)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	//post????????????json??????
	messages := UserUpdataPassword_Message{src_appid, src_mytime, src_mydata, src_addr, src_oldpassword, src_newpassword}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

func PostWithJson_safeBatchTransfer(thurl string, actionName string, myappid string, from string, to string, ids []uint64, numbers []uint64, password string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("safeBatchTransfer")
	var ids_one []byte
	for i := 0; i < len(ids); i++ {
		ID := make([]byte, 8)                  //????????????
		binary.BigEndian.PutUint64(ID, ids[i]) //uint64?????????
		ids_one = append(ids_one, ID...)
	}
	var numbers_one []byte
	for i := 0; i < len(numbers); i++ {
		NUM := make([]byte, 8)                      //????????????
		binary.BigEndian.PutUint64(NUM, numbers[i]) //uint64?????????
		numbers_one = append(numbers_one, NUM...)
	}
	fmt.Printf("%d", ids_one)
	fmt.Printf("%d", numbers_one)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")

	src_from := publicEncode([]byte(from), "public.pem")
	src_to := publicEncode([]byte(to), "public.pem")
	src_ids := publicEncode(ids_one, "public.pem")
	src_numbers := publicEncode(numbers_one, "public.pem")
	src_password := publicEncode([]byte(password), "public.pem")
	//post????????????json??????
	messages := SafeBatchTransferFrom_Message{src_appid, src_mytime, src_mydata, src_from, src_to, src_ids, src_numbers, src_password}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//??????NFT_INDEX
func PostWithJson_getNftIndex(thurl string, actionName string, myappid string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("getNftIndex")
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	//post????????????json??????
	messages := GetNftIndex_Message{src_appid, src_mytime, src_mydata}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//??????
func PostWithJson_balanceOf(thurl string, actionName string, myappid string, addr string, ids uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("balanceOf")
	myaddr := []byte(addr)
	nu := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(nu, ids) //uint64?????????
	myid := []byte(nu)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddr := publicEncode(myaddr, "public.pem")
	src_myid := publicEncode(myid, "public.pem")

	//post????????????json??????
	messages := BanlaceOf_Message{src_appid, src_mytime, src_mydata, src_myaddr, src_myid}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//Batch??????
func PostWithJson_balanceOfBatch(thurl string, actionName string, myappid string, addrs []string, ids []uint64, addrnumber uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????

	addrnum := make([]byte, 8)                      //????????????
	binary.BigEndian.PutUint64(addrnum, addrnumber) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("balanceOfBatch")

	var ids_one []byte
	for i := 0; i < len(ids); i++ {
		ID := make([]byte, 8)                  //????????????
		binary.BigEndian.PutUint64(ID, ids[i]) //uint64?????????
		ids_one = append(ids_one, ID...)
	}
	var addrs_one []byte
	for i := 0; i < len(addrs); i++ {
		ADDR := []byte(addrs[i])
		addrs_one = append(addrs_one, ADDR...)
	}

	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddrs := publicEncode(addrs_one, "public.pem")
	src_myids := publicEncode(ids_one, "public.pem")
	src_addrnum := publicEncode(addrnum, "public.pem")
	//post????????????json??????
	messages := BanlaceOfBatch_Message{src_appid, src_mytime, src_mydata, src_myaddrs, src_myids, src_addrnum}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//??????
func PostWithJson_GetNFTCreator(thurl string, actionName string, myappid string, ids uint64) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("GetNFTCreator")
	nu := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(nu, ids) //uint64?????????
	myid := []byte(nu)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myid := publicEncode(myid, "public.pem")

	//post????????????json??????
	messages := GetNFTCreator_Message{src_appid, src_mytime, src_mydata, src_myid}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//???????????????
func PostWithJson_SearchSPONSOR(thurl string, actionName string, myappid string, addr string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("SearchSPONSOR")
	myaddr := []byte(addr)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddr := publicEncode(myaddr, "public.pem")
	//post????????????json??????
	messages := SearchSPONSOR_Message{src_appid, src_mytime, src_mydata, src_myaddr}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//???????????????
func PostWithJson_SetSPONSOR(thurl string, actionName string, myappid string, addr string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("SetSPONSOR")
	myaddr := []byte(addr)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddr := publicEncode(myaddr, "public.pem")
	//post????????????json??????
	messages := SearchSPONSOR_Message{src_appid, src_mytime, src_mydata, src_myaddr}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//???????????????
func PostWithJson_RemoveSPONSOR(thurl string, actionName string, myappid string, addr string) []byte {
	now := uint64(time.Now().Unix()) //??????????????????
	fmt.Println(now)
	by := make([]byte, 8)               //????????????
	binary.BigEndian.PutUint64(by, now) //uint64?????????
	//????????????
	appid := []byte(myappid)
	mytime := []byte(by)
	mydata := []byte("RemoveSPONSOR")
	myaddr := []byte(addr)
	src_appid := publicEncode(appid, "public.pem")
	src_mytime := publicEncode(mytime, "public.pem")
	src_mydata := publicEncode(mydata, "public.pem")
	src_myaddr := publicEncode(myaddr, "public.pem")
	//post????????????json??????
	messages := SearchSPONSOR_Message{src_appid, src_mytime, src_mydata, src_myaddr}
	ba, err := json.Marshal(messages)
	if err != nil {
		return []byte("json.Marshal error")
	}
	resp, err := http.Post(thurl+"/"+actionName+"", "application/json", bytes.NewBuffer([]byte(ba)))
	if err != nil {
		return []byte("http error")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("ReadAll error")
	}
	return body
}

//??????rsa??????????????????
func publicEncode(plainText []byte, filename string) []byte {
	//1. ?????????????????? ??????data?????????
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stat, _ := file.Stat() //????????????????????????
	data := make([]byte, stat.Size())
	file.Read(data)
	file.Close()
	//2. ?????????????????????pem??????
	block, _ := pem.Decode(data)

	//3. ??????x509????????????????????????????????????
	pubInterface, err2 := x509.ParsePKIXPublicKey(block.Bytes)
	if err2 != nil {
		panic(err2)
	}
	pubKey := pubInterface.(*rsa.PublicKey)

	//4. ??????????????????
	cipherText, err3 := rsa.EncryptPKCS1v15(rand.Reader, pubKey, plainText)
	if err3 != nil {
		panic(err3)
	}
	return cipherText
}

//??????rsa??????????????????
func publicEncodeLong(plainText []byte, filename string) ([]byte, error) {

	//1. ?????????????????? ??????data?????????
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	stat, _ := file.Stat() //????????????????????????
	data := make([]byte, stat.Size())
	file.Read(data)
	file.Close()
	//2. ?????????????????????pem??????
	block, _ := pem.Decode(data)

	//3. ??????x509????????????????????????????????????
	pubInterface, err2 := x509.ParsePKIXPublicKey(block.Bytes)
	if err2 != nil {
		panic(err2)
	}
	pubKey := pubInterface.(*rsa.PublicKey)
	partLen := pubKey.N.BitLen()/8 - 11
	chunks := split(plainText, partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, chunk)
		if err != nil {
			return nil, err
		}
		buffer.Write(bytes)
	}
	return buffer.Bytes(), nil

}

// ???????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:])
	}
	return chunks
}
