package main

import ("fmt"
		"os")


func main(){
	if len(os.Args) != 3 && len(os.Args) !=4 {
		fmt.Println("ERROR: insufficient number arguments")
		return
	}


	//multiline terminal input (encode decode)
/* 	var input string */
	if len(os.Args) == 3{
		if os.Args[1] == "-start" && os.Args[2] == "-d"{
			input, err:= GetInputMulti()
			if err != nil{
				fmt.Println(err)
				return
			}
			result, err := Decode(input)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		} else if os.Args[1] == "-start" && os.Args[2] == "-e"{
			input, err:= GetInputMulti()
			result,err:= Encode(input)
			if err != nil{
				fmt.Println(err)
				return
			}
			fmt.Println(result)
		}else{
			fmt.Println("ERROR:arguments faulty")
		}

	//single line terminal input OR .txt input (encode decode)
	}else if len(os.Args) == 4{

		//txt input (encode decode)
		if os.Args[2] == "-f"{
			input, err:= GetInputFile()
			if err != nil{
				fmt.Println(err)
				return
			}

			if os.Args[3] == "-d"{
					result,err:=Decode(input)
					fmt.Println(result)
				if err != nil{
					fmt.Println(err)
					return
				}
			}else if os.Args[3] == "-e"{
				result,err:= Encode(input)
				if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(result)
			}else{
				fmt.Println("ERROR:invalid arguments")
			}

		//terminal input (encode decode)
		}else if os.Args[2] == "-t"{
			if os.Args[3] == "-d"{
					result,err:=Decode(os.Args[1])
					fmt.Println(result)
				if err != nil{
					fmt.Println(err)
					return
				}
			}else if os.Args[3] == "-e"{
				result,err:= Encode(os.Args[1])
				if err != nil{
					fmt.Println(err)
					return
				}
			fmt.Println(result)
			}else{
				fmt.Println("ERROR:invalid arguments")
			}
		}
	}

}



	