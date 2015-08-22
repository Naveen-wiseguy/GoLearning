package main

import(
"fmt"
)


//Graph nodes for the list
type Graphnode struct{
  value int;
  next [] int;
}

//Graph adjacency list
type GraphList struct{
  elements map[int] Graphnode;
  nodesNum int; 
}

func initGraph(graph *GraphList) {
  graph.elements=make(map[int] Graphnode);
  graph.nodesNum=0;
}

func addNode(graph *GraphList,elt int) {
  node:= Graphnode{elt,make([] int,0)};
  graph.elements[elt]=node;
}

func addEdge(graph *GraphList,from int,to int) {
  value,ok:=graph.elements[from];
  if(!ok) {
    addNode(graph,from);
    value=graph.elements[from];
  }
  fmt.Printf("Adding an edge from %d to %d",from,to);
  value.next=append(value.next,to);
  graph.elements[from]=value;
}

func printGraph(graph *GraphList) {
  for key,value := range graph.elements {
     fmt.Printf("Details for node %d:\n",key);
     for _,node:= range value.next {
        fmt.Printf("%d to %d\n",key,node); 
     }
  }
}

func main() {
  var graph GraphList;
  initGraph(&graph);
  var num int;
  fmt.Println("Enter the number of nodes in the graph :");
  _,err:=fmt.Scanf("%d",&num);
  if(err!=nil) {
    fmt.Println(err);
  } else{
    for i:=1;i<=num;i++ {
      addNode(&graph,i);
    }
    fmt.Println("Enter the number of edges :");
    var edges int;
    fmt.Scanf("%d",&edges);
    for i:=0;i<edges;i++ {
      fmt.Println("Enter the from and to nodes :");
      var from,to int;
      fmt.Scanf("%d %d",&from,&to);
      addEdge(&graph,from,to);
    }
  }
  printGraph(&graph);
}
