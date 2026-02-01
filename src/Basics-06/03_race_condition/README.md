```bash
go@linux:~/GolandProjects/rebrain-go$ go run src/Basics-06/03_race_condition/7dollars.go 
-4
go@linux:~/GolandProjects/rebrain-go$ go run -race src/Basics-06/03_race_condition/7dollars.go 
==================
WARNING: DATA RACE
Write at 0x00c0000be018 by goroutine 7:
  main.(*BankCell).SubBalance()
      /home/go/GolandProjects/rebrain-go/src/Basics-06/03_race_condition/7dollars.go:19 +0x84
  main.main.gowrap1()
      /home/go/GolandProjects/rebrain-go/src/Basics-06/03_race_condition/7dollars.go:32 +0x1f

Previous read at 0x00c0000be018 by main goroutine:
  main.(*BankCell).SubBalance()
      /home/go/GolandProjects/rebrain-go/src/Basics-06/03_race_condition/7dollars.go:19 +0xe8
  main.main()
      /home/go/GolandProjects/rebrain-go/src/Basics-06/03_race_condition/7dollars.go:33 +0xac

Goroutine 7 (running) created at:
  main.main()
      /home/go/GolandProjects/rebrain-go/src/Basics-06/03_race_condition/7dollars.go:32 +0xab
==================
3
Found 1 data race(s)
exit status 66
go@linux:~/GolandProjects/rebrain-go$ 
```

_"-race"_ key is Go **race detector**

```bash
go@linux:~/GolandProjects/rebrain-go$ go run -race src/Basics-06/03_race_condition/mutex7dollars.go 
3
```

# Links:
* https://go.dev/blog/race-detector
* https://en.wikipedia.org/wiki/Race_condition