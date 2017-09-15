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

func getGroups(board string){
    api:="https://api.trello.com/1/boards/"+board+"?fields=id,name,idOrganization,dateLastActivity&lists=open&list_fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    http.Get(api)
    // resp,_ := http.Get(api)
    // defer resp.Body.Close()
    // body,_:=ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body: ",string(body))
}

func getBoards() []BoardResp {
    api:="https://api.trello.com/1/members/"+API_USER+"/boards?fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    resp,err:=http.Get(api)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    body,_ := ioutil.ReadAll(resp.Body)
    boardResps := make{[]BoardResp}
    json.Unmarshal(body,&boardResps)
    return boardResps
}

func getCards(board string){
    api:="https://api.trello.com/1/boards/"+board+"/cards/?fields=id,name,desc,idList&key="+API_KEY+"&token="+API_TOKEN
    http.Get(api)
}

func createCard(card string,desc string,idList string){
    api:="https://api.trello.com/1/cards"

    jsonStr:= []byte("{\"name\":\""+card+"\",\"desc\":\""+desc+"\",\"idList\":\""+idList+"\",\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("POST",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    client.Do(req)
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

func main(){
    fmt.Println(getBoard())
}


type BoardResp struct {
    Id string
    Name string
    Desc string 
}

type CardResp struct {
    Id string
    Name string
    Desc string
    IdList string
}