# jmap: use jmap instead of map
# 
    jmap is a tool used to store million to 100 millions k-v objects in local server and costs less memory. 
    
    first of all,you should estimate your requirements.    
    
    cap = cap(your k-v objects)  
    if cap < 10k {  
        use go map  
     } else {  
        use jmap instead  
    }   
#  

# benchMark

    j-map ready! len: 0
    goMap ready! len: 0
    goos: linux
    goarch: amd64
    pkg: github.com/junjiefly/jmap
    cpu: Intel(R) Core(TM) i3-8100T CPU @ 3.10GHz
    BenchmarkJMapSet-4      	14406464	        85.08 ns/op	      34 B/op	       0 allocs/op
    BenchmarkJMapGet-4      	28466576	        47.05 ns/op	       0 B/op	       0 allocs/op
    BenchmarkJMapDelete-4   	17903854	        62.91 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMapSet-4       	15370224	       218.7 ns/op	      86 B/op	       0 allocs/op
    BenchmarkMapGet-4       	28448533	        61.85 ns/op	       0 B/op	       0 allocs/op
    BenchmarkMapDelete-4    	18742928	        70.59 ns/op	       0 B/op	       0 allocs/op
    PASS
    ok  	github.com/junjiefly/jmap	11.371s
