package main
import("encoding/csv"
"fmt"
"os"
"strconv")
func main(){c,_:=os.Open(".csv")
r,_:=csv.NewReader(c).ReadAll()
n:=make([][]int,len(r))
for i:=range r{q:=r[i]
for j:=range q{if q[j]!=""{t,_:=strconv.Atoi(q[j])
n[i]=append(n[i],t)}}}
for i:=range n{x:=n[i]
for range x{for k:=0;k<len(x)-1;k++{if x[k]>x[k+1]{n[i][k],n[i][k+1]=x[k+1],x[k]}}}}
fmt.Printf("%v\n",n)}