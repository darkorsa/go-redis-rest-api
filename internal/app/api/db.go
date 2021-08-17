package api

import (
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (s *Server) GetKey(c *gin.Context) {
	item, queryErr := s.queryService.Get(c.Param("id"))
	if queryErr != nil {
		s.internalServerError(queryErr.Error(), c)
		return
	}

	if item == nil {
		s.notFoundError("key not found", c)
		return
	}

	c.JSON(200, item)
}

func (s *Server) GetKeys(c *gin.Context) {
	items, err := s.queryService.List()
	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	c.JSON(200, items)
}

func (s *Server) DelKey(c *gin.Context) {
	res, err := s.queryService.Del(c.Param("id"))

	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	if res == 0 {
		s.notFoundError("key not found", c)
		return
	}

	ok := map[string]string{
		"result":   "OK",
		"affected": strconv.FormatInt(res, 10),
	}

	c.JSON(200, ok)
}

func (s *Server) ListGet(c *gin.Context) {
	key := c.Param("id")

	var start int64 = 0
	var stop int64 = -1
	var err error

	if c.Query("start") != "" {
		start, err = strconv.ParseInt(c.Query("start"), 10, 64)
		if err != nil {
			s.badRequestError("start parameter is invalid", c)
			return
		}
	}

	if c.Query("stop") != "" {
		stop, err = strconv.ParseInt(c.Query("stop"), 10, 64)
		if err != nil {
			s.badRequestError("stop parameter is invalid", c)
			return
		}
	}

	items, err := s.queryService.LRange(key, start, stop)
	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	c.JSON(200, items)
}

func (s *Server) ListLPush(c *gin.Context) {
	s.listPush("LPush", c)
}

func (s *Server) ListRPush(c *gin.Context) {
	s.listPush("RPush", c)
}

func (s *Server) listPush(method string, c *gin.Context) {
	if c.PostForm("value") == "" {
		s.badRequestError("value is empty", c)
		return
	}

	args := []reflect.Value{
		reflect.ValueOf(c.Param("id")),
		reflect.ValueOf(c.PostForm("value")),
	}

	val := reflect.ValueOf(s.queryService).MethodByName(method).Call(args)

	var err error

	if val[1].Interface() != nil {
		err = val[0].Interface().(error)
		s.internalServerError(err.Error(), c)
		return
	}

	ok := map[string]string{
		"result": "OK",
	}

	c.JSON(200, ok)
}

func (s *Server) ListDel(c *gin.Context) {
	var count int64 = 0
	var err error

	if c.PostForm("value") == "" {
		s.badRequestError("value is empty", c)
		return
	}

	if c.Query("count") != "" {
		count, err = strconv.ParseInt(c.Query("count"), 10, 64)
		if err != nil {
			s.badRequestError("count parameter is invalid", c)
			return
		}
	}

	res, err := s.queryService.LRem(c.Param("id"), count, c.PostForm("value"))

	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	ok := map[string]string{
		"result":   "OK",
		"affected": strconv.FormatInt(res, 10),
	}

	c.JSON(200, ok)
}
