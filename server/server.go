package server 


func Init() {
	r := Router()
	
	r.Run(":9000")
}