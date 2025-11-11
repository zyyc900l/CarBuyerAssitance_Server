struct BaseResp{
    1: i64 code,
    2: string msg,
}
struct UserInfo{
    1: string userId
    2: string username,
    3: string phone
    4: double budge_min
    5: double budge_max
    6: string preferred_type
    7: string preferred_brand
    8: string status
    9: string address
        10: required string created_at
        11: required string updated_at
        12: required string deleted_at
}
struct ConsultResult{
    1:required string	Analysis
    2:required string Proposal
    3:required list<Car> Result
}
struct Car{
	 string  ImageUrl
	 string    CarName
	string FuelConsumption
	string Power
	string Seat
	string Drive
	string RecommendedReason
}