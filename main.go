package main
import (
  "bufio"
  "fmt"
  "os"
  "regexp"
  "sort"
  "strings"
)
// IPStat helps sort our IP map
type IPSat struct {
  IP string
  Count int
}

func main() {
  if len(os.Args) < 2 {
    fmt.Println("Usage: go run main.go <log_file>")
    return
  }

  fileName := os.Args[1]

  //1. Open the file
  file, err := os.Open(fileName)
  if err != nil {
    fmt.Printf("Error opening file: %v\n", err)
    return
  }
  defer file.Close()

  //2. Compile Regex : Captures IP(Group 1) and Status Code (Group 2)
  pattern := regexp.MustCompile(`^(\S+) \S+ \S+ \[.*?\] ".*?" (\d{3})`)

  //3. Intialize Counters and map
  totalRequests := 0
  error404 := 0
  error5xx := 0
  ipCounts := make(map[string]int)

  //4. Memory-safe file scanner
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    line := scannerText()

    // skip empty lines
    if strings.TrimSpace(line) == "" {
      continue
      }

    totalRequests++


    //5. Extract data via Regex
    matches := pattern.FindStringSubmatch(line)
    if len(matches) == 3 {
      ip := matches[1]
      status := matches[2]


  //Track IP visits
   ipCounts[ip]++

 //Track errors
  if status == "404" {
    error404++
  } else if strings.HasPrefix(status, "5") {
    error5xx++
   }     
  }      
}

// Check for scanner errors
if err := scanner.Err(); err != nil {
  fmt.Printf("Error scanning file: %v\n", err)
  return
}

// 6.Sort IPs by visit count with a secondary tie-breaker
var topIPs []IPStat

//Convert map to slice for sorting
for ip, count := range ipCounts {
  newStat := IPStat{
    IP: ip,
    Count: count,
  }
  topIPs = append(topIPs, newStat)

//Sort descending using a custom comparison function
sort.Slice(topIPs, func(i, int, j int) bool {
  //Tie-breaker: if counts are equal, sort alphabetically by IP
  if topIPs[i].Count == topIPs[j].Count {
        return topIPs[i].IP < topIPs[j].IP
    }

  //Default : Sort by Count (descending order)
  return topIPs[i].Count > topIPs[j].Count
  })

  //7. Print Audit Report 
  fmt.Println("=== LogLens Audit Report ===")
  fmt.Println()
  fmt.Printf("Total Requests: %d\n", totalRequests)
  fmt.Printf("Total 404 Not Found: %d\n", error404)
  fmt.Printf("Total 5xx Server Errors: %d\n", error5xx)


  fmt.Println("\nTop 3 Visiting IPs:")

  //Print up to top 3 IPs safely
  limit := 3
  if len(topIPs) < 3 {
    limit = len(topIPs)
  }

  for i := 0; i<limit; i++ {
    fmt.Printf("%d. %s (%d requests)\n", i+1, topIPs[i].IP, topIPs[i].Count)
    }
}

  

  

  




  

  

    



      
      
      

















  














  
