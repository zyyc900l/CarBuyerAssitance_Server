namespace go admin
include "./model.thrift"

struct AddUserRequest{
    1: required string userId
    2: required string username,
    3: required string phone
    4:required double budget_min
    5: required double budget_max
    6:required string preferred_type
    7:required  string preferred_brand
    8: required i64 status
    9: required string address
    10:required  string password,
}
struct AddUserResponse{
        1: required model.BaseResp base,
        2: required string user_id,
}
struct DeleteUserRequest{
    1:string userId,
}
struct DeleteUserResponse{
            1: required model.BaseResp base,
}
struct QueryAllConsultRequest{
    1 :required i64 page_size,
    2:required i64 page_num,
}
struct QueryAllConsultResponse{
            1: required model.BaseResp base,
            2: required model.ConsultationList data,
}
struct QueryARequest
{

}
struct QueryAResponse
{
     1: required model.BaseResp base,
     2: required model.FrequencyList frequency,
     3: required model.BudgetList budget,
     4: required model.SceneList scene,
}

service AdminService{
    AddUserResponse AddUser(1:AddUserRequest req)(api.post="/api/admin/user/add"),
    DeleteUserResponse DeleteUser(1:DeleteUserRequest req)(api.delete="/api/admin/user/delete"),
   QueryAllConsultResponse QueryAllConsult(1:QueryAllConsultRequest req)(api.get="/api/admin/consult/query"),
   QueryAResponse QueryA(1:QueryARequest req)(api.get ="/api/admin/query")
}