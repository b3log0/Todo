package main
 
import (
    "net/http"
    "io/ioutil"
    "bytes"
)

const(
    API_USER="zephyrcheung"
    API_KEY=""
    API_TOKEN=""
)

func createBoard(board string){
    api:="https://api.trello.com/1/boards/"

    jsonStr:= []byte("{\"name\":\""+board+"\",\"key\":\""+API_KEY+"\",\"token\":\""+API_TOKEN+"\"}")
    req,_:=http.NewRequest("POST",api,bytes.NewBuffer(jsonStr))
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    client.Do(req)
    // resp, err := client.Do(req)
    // if err != nil {
    //     panic(err)
    // }
    // defer resp.Body.Close()

    // fmt.Println("response Status:", resp.Status)
    // fmt.Println("response Headers:", resp.Header)
    // body, _ := ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body:", string(body))    
}

func getGroups(board string){
    api:="https://api.trello.com/1/boards/"+board+"?fields=id,name,idOrganization,dateLastActivity&lists=open&list_fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    http.Get(api)
    // resp,_ := http.Get(api)
    // defer resp.Body.Close()
    // body,_:=ioutil.ReadAll(resp.Body)
    // fmt.Println("response Body: ",string(body))
}

func getBoards(){
    api:="https://api.trello.com/1/members/"+API_USER+"/boards?fields=id,name&key="+API_KEY+"&token="+API_TOKEN
    http.Get(api)
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