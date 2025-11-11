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
    8: i64 status
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
struct Consult {
    string	UserId
    i64	ConsultId
    string	BudgetRange
    string	PreferredType
    string	UseCase
    string	FuelType
    string	BrandPreference
}
struct PointList{
    1:required list<Point> item
    2:required i64 num,
    3:required i64 sum, //总积分
}
struct Point{
    1:required i64 point_id
    2:required string user_id
    3:required i64 points
    4:required string reason
            5: required string created_at
            6: required string updated_at
}
struct Consultation{
    1:required Consult consult
    2: required ConsultResult consult_result
}

struct ConsultationList{
    1:required list<Consultation> item,
    2: required i64 total,
}
struct Gift{
    1:required i64 gift_id
    2:required string gift_name
    3:required i64 required_points
    4:required i64 stock_quantity
    5:required string cover_image_url
    6:required i64 is_online,
                 7: required string created_at
                 8: required string updated_at
}
struct GiftList{
        1:required list<Gift> item
        2:required i64 total,
}
struct Order{
    1: required  i64 Id,
    2: required string user_id,
    3: required string gift_name,
    4: required i64 need_points,
    5: required string orderTime
    6: required i64 status,
}
struct Frequency{
    1: required string frequency_name
    2: required i64 value
}
struct Scene{
        1: required string scene_name
        2: required i64 value
}
struct Budget{
    1: required string budget_name
      2: required i64 value
}
struct FrequencyList{
    1: required list<Frequency> item,
    2:required i64 total
}
struct BudgetList{
    1: required list<Budget> item,
        2:required i64 total
}
struct SceneList{
    1: required list<Scene> item,
        2:required i64 total
}