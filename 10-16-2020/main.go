package main
import("encoding/csv"
"fmt"
"os"
"strconv")
func main(){c,_:=os.Open(".csv")
r,_:=csv.NewReader(c).ReadAll()
n:=make([][]int,len(r))
for i:=0;i<len(r);i++{for j:=0;j<len(r[i]);j++{if r[i][j]!=""{t,_:=strconv.Atoi(r[i][j])
n[i]=append(n[i],t)}}}
for i:=0;i<len(n);i++{x:=n[i]
for range x{for k:=range x{if k<len(x)-1&&x[k]>x[k+1]{n[i][k],n[i][k+1]=x[k+1],x[k]}}}}
fmt.Printf("%v\n",n)}