package network

import (
	"RMazeE-server/service"
	"RMazeE-server/types"
	"RMazeE-server/types/errors"
	"github.com/gin-gonic/gin"
	"strconv"
	"sync"
)

var (
	rankRouterInit     sync.Once
	rankRouterInstance *rankRouter
)

type rankRouter struct {
	router *Network
	//service
	rankService *service.Rank
}

func newRankRouter(router *Network, rankService *service.Rank) *rankRouter {
	rankRouterInit.Do(func() {
		rankRouterInstance = &rankRouter{
			router:      router,
			rankService: rankService,
		}

	})
	router.registerGET("/", rankRouterInstance.get)
	router.registerPOST("/", rankRouterInstance.create)
	router.registerUPDATE("/", rankRouterInstance.update)
	router.registerDELETE("/", rankRouterInstance.delete)

	return rankRouterInstance
}

func (u *rankRouter) create(c *gin.Context) {
	var req types.CreateRankRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.CreateRankResponse{
			ApiResponse: types.NewApiResponse("바인딩 에러입니다.", -1, err.Error()),
		})
	} else if err = u.rankService.Create(req.ToRank()); err != nil {
		u.router.failedResponse(c, &types.CreateRankResponse{
			ApiResponse: types.NewApiResponse("Create 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.CreateRankResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *rankRouter) get(c *gin.Context) {
	if params, err := getParams(c); err != nil {
		u.router.failedResponse(c, &types.GetRankResponse{
			ApiResponse: types.NewApiResponse("Invalid params", -1, err.Error()),
		})
	} else if foundRanking, err := u.rankService.Get(params.Algorithm, params.MazeLevel); err != nil {
		u.router.failedResponse(c, &types.GetRankResponse{
			ApiResponse: types.NewApiResponse("Get 요청에 실패하였습니다", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.GetRankResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
			Ranking:     foundRanking,
		})
	}
}

func getParams(c *gin.Context) (*types.GetRankParams, error) {
	params := new(types.GetRankParams)
	if str := c.Query("algorithm"); str == "" {
		return nil, errors.Errorf(errors.InvalidParams, nil)
	} else {
		params.Algorithm = str
	}

	if str := c.Query("mazeLevel"); str == "" {
		return nil, errors.Errorf(errors.InvalidParams, nil)
	} else if num, err := strconv.ParseInt(str, 10, 64); err != nil {
		return nil, errors.Errorf(errors.InvalidParams, nil)
	} else {
		params.MazeLevel = num
	}

	return params, nil
}

func (u *rankRouter) update(c *gin.Context) {
	var req types.UpdateRankRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.UpdateRankResponse{
			ApiResponse: types.NewApiResponse("바인딩 에러입니다.", -1, err.Error()),
		})
	} else if err = u.rankService.Update(req.ToRank()); err != nil {
		u.router.failedResponse(c, &types.UpdateRankResponse{
			ApiResponse: types.NewApiResponse("update 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.UpdateRankResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}
}

func (u *rankRouter) delete(c *gin.Context) {
	var req types.DeleteRankRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.failedResponse(c, &types.DeleteRankResponse{
			ApiResponse: types.NewApiResponse("바인딩 에러입니다.", -1, err.Error()),
		})
	} else if err = u.rankService.Delete(req.ToRank()); err != nil {
		u.router.failedResponse(c, &types.DeleteRankResponse{
			ApiResponse: types.NewApiResponse("delete 에러입니다.", -1, err.Error()),
		})
	} else {
		u.router.okResponse(c, &types.DeleteRankResponse{
			ApiResponse: types.NewApiResponse("성공입니다.", 1, nil),
		})
	}

}
