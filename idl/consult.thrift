namespace go consult
include "./model.thrift"

struct ConsultRequest{
    required string budget_range,
    required string preferred_type,
    required string use_casecase,
    required string fuel_type,
    required string brand_preference,
}
struct ConsultResponse{
    1: required model.BaseResp base,
    2: required model.ConsultResult data,
}
service ConsultService{
    ConsultResponse Consult(1:ConsultRequest req)(api.get ="/api/consult/purchase"),
}