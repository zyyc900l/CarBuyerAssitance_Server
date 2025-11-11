namespace go user
include "./model.thrift"

// register
struct RegisterRequest {
    1: required string username,
    2: required string password,
    3: required string phone_number,
    4: required string Id,
}
struct RegisterResponse {
    1: required model.BaseResp base,
    2: required string user_id,
}

// login
struct LoginRequest{
    1: required string Id,
    2: required string password,
}
struct LoginResponse{
    1: required model.BaseResp base,
    2: required model.UserInfo data,
}
//
struct ProposeFeedbackRequest{
    1: optional string consult_id,
    2: required string feedback,
}
struct ProposeFeedbackResponse{
     1: required model.BaseResp base,
}
struct QueryUserInfoRequest{
    1: required string user_id
}
struct QueryUserInfoResponse{
        1: required model.BaseResp base,
        2: required model.UserInfo data,
}
struct UpdateUserInfoRequest{
    1: required string userId
        4:required double budget_min
        5: required double budget_max
        6:required string preferred_type
        7:required  string preferred_brand
            9: required string address
}
struct UpdateUserInfoResponse{
            1: required model.BaseResp base,
            2: required model.UserInfo data,
}
service UserService {
    RegisterResponse Register(1: RegisterRequest req)(api.post = "/api/user/register"),
    LoginResponse Login(1: LoginRequest req)(api.post = "/api/user/login"),
    QueryUserInfoResponse QueryUserInfo(1:QueryUserInfoRequest req)(api.get="/api/user/query/Info"),
    UpdateUserInfoResponse UpdateUserInfo(1:UpdateUserInfoRequest req)(api.put="/api/user/update/Info"),
}
