package main

import (
	"bytes"
	"fmt"
	"net/http"
    "io"
    "os/user"
    "os"
)

func main() {
    currentUser, err := user.Current()
    if err != nil {
        //fmt.Println("error getting username")
    }
    username := currentUser.Username
    //fmt.Printf("Username is: %s\n", username)

    hostname, err := os.Hostname()
    if err != nil {
        //fmt.Println("error getting hostname")
    }
    //fmt.Printf("Hostname: %s\n", hostname)
    

    url := "https://genoscloudstorage.blob.core.windows.net/results/" + hostname + "?sp=acw&st=2024-09-17T20:27:13Z&se=2024-09-18T04:27:13Z&spr=https&sv=2022-11-02&sr=c&sig=1AJSimTZxVdk3sAJfyU9nE6Ykie9LGfpo%2BwO1Utaatg%3D"
    data := []byte(username + " , " + hostname)
    //resp, err := http.NewRequest(http.MethodGet, url, nil)
    //generate the request
    req, err := http.NewRequest("PUT", url, bytes.NewBuffer(data))
    req.Header.Set("Content-Type", "text/plain")
	req.Header.Set("x-ms-blob-type", "BlockBlob")
    if err != nil {
        //fmt.Println("error creating http request")
    }
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        //fmt.Println("http response error")
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        // handle error
		//fmt.Println("error in reading response")
    }
    fmt.Println(string(body))
    defer resp.Body.Close()
}