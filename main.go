package main

import (
   _ "encoding/xml"
   "net/http"  
   "github.com/mmcdole/gofeed"
   "github.com/gin-gonic/gin"
   "regexp"
)

   type Response struct{
      Title  string `json:"title"`
      Content string `json:"description"`
      URL    string `json:"web_url"`
      PubDate string `json:"publishing_date"` 
      ImgUrl string `json:"picture_url"`
   }

   var result = []Response{}

   func getvalue(c *gin.Context){
      i:= 0
      fp := gofeed.NewParser()
      fp.UserAgent = "MyCustomAgent 1.0"
      feed, _ := fp.ParseURL("http://feeds.feedburner.com/PoorlyDrawnLines")
      data := feed.Items
      for _ ,item := range(data){
         if(i > 10){
            break
         }
         cont := item.Content
         findLinks := regexp.MustCompile("<a.*?href=\"(.*?)\"")
         matches := findLinks.FindStringSubmatch(cont)
         imgsrc := matches[0][9:len(matches[0]) - 1]   
         result = append(result, Response{
            Title: item.Title,
            URL:  item.Link,
            PubDate: item.Published,
            ImgUrl: imgsrc,
            Content: item.Content,
         })
         i += 1
      }
         c.JSON(http.StatusOK, gin.H{"result": result})      
   }


   func SetUpRouter() *gin.Engine{
      router := gin.Default()
      router.GET("/getdata", getvalue)
      return router
   }

   func main(){
      router := SetUpRouter()
      router.Run("localhost:8080")  
   }




