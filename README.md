################  
jmap is a tool used to store million to 100 millions k-v ojects in local server and costs less memory. 

first of all,you should estimate your requirements.    

cap = cap(your k-v objects)  
if cap < 10k {  
    use go map  
 } else {  
    use jmap instead  
}   
################  

#benchMark test

j-map ready! len: 10000000  
goMap ready! len: 10000000  
goos: linux  
goarch: amd64  
pkg: jmap  
cpu: Intel(R) Xeon(R) CPU E5-2630 v4 @ 2.20GHz  
BenchmarkJMapSet-40      	1000000000	         0.0000019 ns/op  
BenchmarkJMapGet-40      	1000000000	         0.0000017 ns/op  
BenchmarkJMapDel-40      	1000000000	         0.0000015 ns/op  
BenchmarkMapSet-40       	1000000000	         0.0000019 ns/op  
BenchmarkMapGet-40       	1000000000	         0.0000020 ns/op  
BenchmarkMapDelete-40    	1000000000	         0.0000027 ns/op  
PASS  
ok  	jmap	5.919s
