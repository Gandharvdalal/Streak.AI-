package main

import "fmt"
import "encoding/json"
import "net/http"

func main() {
	// routing for /findPairs
	https.HandleFunc("/find-pairs",findPairsHandler)

	//start server

	fmt.Println("Server is listening to the port 8080 !!")

	if err:=http.ListenAndServe(":8080",nil);
	err!=nil{
		fmt.Println("Server failed to start", err)
	}
}

type Request struct{
	Numbers[]int json:"numbers"  
	Target int json:"target"

}
type Response struct {
	Solutions [][]int json:"solutions"
}

func findPairs(numbers[] int , target int){
	var solutions[][]int
	for i:=0;i<len(numbers);i++{
		for j:i+1;j<len(numbers);j++{
			if numbers[i]+numbers[j]==target{
				solutions=append(solution,[]int{i,j})
			}
		}
	}
	return solutions

}

// API handler for finding Pairs
func findPairsHandler(w http.ResponseWriter, r *http.Request){
	var req Request

	// decode the request body

	if err:=json.NewDecoder(r.body).Decode(&req);
	err!=nil{
		http.Error(w,"This is a Really Bad Request Body !!!!")
		return
	}
		// Now we will find the Pairs

	solutions:=findPairs(req.Numbers,req.Target)


	// Creating Response

	resp:=Response{
		Solutions:solutions,
	}

	// Setting Response header

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)



	// Sending the Resonse

	if error:=json.NewEncoder(w).Encode(resp);
	err!=nil{
		http.Error(w,"Error in encoding Responseeee !!!", http.StatusInternalServerError)
	}
}


