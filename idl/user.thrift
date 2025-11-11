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
service UserService {
    RegisterResponse Register(1: RegisterRequest req)(api.post = "/api/user/register"),
    LoginResponse Login(1: LoginRequest req)(api.post = "/api/user/login"),
}
