package types

import "C"

type Rank struct {
	Name        string  `json:"name"`
	Rank        int64   `json:"rank"`
	ElapsedTime float64 `json:"elapsedTime" bson:"elapsedTime"`
	Algorithm   string  `json:"algorithm"`
	MazeLevel   int64   `json:"mazeLevel" bson:"mazeLevel"`
}

// -----CREATE-----

type CreateRankRequest struct {
	Name        string  `json:"name" binding:"required"`
	ElapsedTime float64 `json:"elapsedTime" binding:"required"`
	Algorithm   string  `json:"algorithm" binding:"required"`
	MazeLevel   int64   `json:"mazeLevel" binding:"required"`
}

func (c *CreateRankRequest) ToRank() *Rank {
	return &Rank{
		Name:        c.Name,
		ElapsedTime: c.ElapsedTime,
		Algorithm:   c.Algorithm,
		MazeLevel:   c.MazeLevel,
	}
}

type CreateRankResponse struct {
	*ApiResponse
	*Rank
}

// -----GET-----

type GetRankRequest struct {
}
type GetRankResponse struct {
	*ApiResponse
	Ranking []*Rank `json:"result"`
}
type GetRankParams struct {
	Algorithm string
	MazeLevel int64
}

// -----UPDATE-----

type UpdateRankRequest struct {
	Name        string  `json:"name" binding:"required"`
	ElapsedTime float64 `json:"elapsedTime" binding:"required"`
	Algorithm   string  `json:"algorithm" binding:"required"`
	MazeLevel   int64   `json:"mazeLevel" binding:"required"`
}
type UpdateRankResponse struct {
	*ApiResponse
	*Rank
}

func (c *UpdateRankRequest) ToRank() *Rank {
	return &Rank{
		Name:        c.Name,
		ElapsedTime: c.ElapsedTime,
		Algorithm:   c.Algorithm,
		MazeLevel:   c.MazeLevel,
	}
}

// -----DELETE-----

type DeleteRankRequest struct {
	Name string `json:"name" binding:"required"`
}

func (c *DeleteRankRequest) ToRank() *Rank {
	return &Rank{
		Name: c.Name,
	}
}

type DeleteRankResponse struct {
	*ApiResponse
}
