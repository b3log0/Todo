package main
 
import (
    "net/http"
    "io/ioutil"
    "bytes"
    "fmt"
    "encoding/json"
)

const(
    API_USER="zephyrcheung"
    API_KEY="31238c876dc00f9e858fe9f70b939697"
    API_TOKEN="79ba20868621208c142d3c8afcf69489bfd29dc73e590c4d9a83ff1a9a37f663"
)

func createBoard(board string) string{
    api:="https://api.trello.com/1/boards/"

    jsonStr:= []byte("{\"name\":\""+board+"\",\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("POST",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    boardResp:=BoardResp{}
    json.Unmarshal(body,&boardResp)
    return boardResp.Id
}

func getGroups(board string) []GroupResp{
    api:="https://api.trello.com/1/boards/"+board+"?fields=id,name&lists=open&list_fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    resp,err:=http.Get(api)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body,_ := ioutil.ReadAll(resp.Body)
    boardResp := BoardResp{}
    json.Unmarshal(body,&boardResp)
    return boardResp.GroupList
}

func getBoards() []BoardResp {
    api:="https://api.trello.com/1/members/"+API_USER+"/boards?fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    resp,err:=http.Get(api)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body,_ := ioutil.ReadAll(resp.Body)
    boardResps := []BoardResp{}
    json.Unmarshal(body,&boardResps)
    return boardResps
}

func getCards(board string) []CardResp {
    api:="https://api.trello.com/1/boards/"+board+"/cards/?fields=id,name,desc,idList&key="+API_KEY+"&token="+API_TOKEN
    resp,err:=http.Get(api)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body,_ := ioutil.ReadAll(resp.Body)
    cardResps := []CardResp{}
    json.Unmarshal(body,&cardResps)
    return cardResps
}

func createCard(card string,desc string,idList string) CardResp {
    api:="https://api.trello.com/1/cards"

    jsonStr:= []byte("{\"name\":\""+card+"\",\"desc\":\""+desc+"\",\"idList\":\""+idList+"\",\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("POST",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body, _ := ioutil.ReadAll(resp.Body)
    cardResp:=CardResp{}
    json.Unmarshal(body,&cardResp)
    return cardResp
}

func updateCard(idCard string,card string,desc string,idList string){
    api:="https://api.trello.com/1/cards/"+idCard
    jsonStr:= []byte("{\"name\":\""+card+"\",\"desc\":\""+desc+"\",\"idList\":\""+idList+"\",\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("PUT",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client:=&http.Client{}
    client.Do(req) 
}

func closeCard(idCard string){
    api:="https://api.trello.com/1/cards/"+idCard
    jsonStr:= []byte("{\"closed\":true,\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("PUT",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client:=&http.Client{}
    client.Do(req) 
}

func openCard(idCard string){
    api:="https://api.trello.com/1/cards/"+idCard
    jsonStr:= []byte("{\"closed\":false,\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("PUT",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client:=&http.Client{}
    client.Do(req) 
}

func closeBoard(idBoard string){
    api:="https://api.trello.com/1/boards/"+idBoard
    jsonStr:= []byte("{\"closed\":true,\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("PUT",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client:=&http.Client{}
    client.Do(req)
}

func openBoard(idBoard string){
    api:="https://api.trello.com/1/boards/"+idBoard
    jsonStr:= []byte("{\"closed\":false,\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("PUT",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client:=&http.Client{}
    client.Do(req) 
}

func main2(){
    fmt.Println(getGroups("59b644d161104b3e83978b19"))
}


