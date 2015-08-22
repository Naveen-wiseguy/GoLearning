package main

import(
"fmt"
"net/http"
"io/ioutil"
)


func main() {
  resp,err:=http.Get("http://www.google.co.in");
if (err != nil) {
  fmt.Printf("%s\n",err);
} else{
  defer resp.Body.Close();
  data,err := ioutil.ReadAll(resp.Body);

  if(err != nil) {
    fmt.Printf("Unable to read data");
  } else {
    fmt.Printf("%s\n",data);
  }
}
}
