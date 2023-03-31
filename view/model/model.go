package model

type Service struct {
	Index func(req *IndexReq) (rsp *IndexRsp) `method:"get" path:"/api/v1/hello"`
}

type IndexReq struct {
}

type IndexRsp struct {
}
