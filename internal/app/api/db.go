package api

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

type actionResponse struct {
	Result   string `json:"result"`
	Affected int64  `json:"affected"`
}

// @Summary Get key
// @Description Get value for key ID
// @Tags keys
// @Produce json
// @Param id path string true "Key ID"
// @Success 200 {object} domain.Item
// @Failure 404 {object} apiErrors.apiError
// @Failure 500 {object} apiErrors.apiError
// @Router /keys/{id} [get]
func (s *Server) GetKey(c *gin.Context) {
	key, queryErr := s.queryService.Get(c.Param("id"))
	if queryErr != nil {
		s.internalServerError(queryErr.Error(), c)
		return
	}

	if key == nil {
		s.notFoundError("key not found", c)
		return
	}

	c.JSON(200, key)
}

// @Summary Get keys
// @Description Get list of all keys
// @Tags keys
// @Produce json
// @Success 200 {object} domain.Keys
// @Failure 500 {object} apiErrors.apiError
// @Router /keys [get]
func (s *Server) GetKeys(c *gin.Context) {
	keys, err := s.queryService.List()
	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	c.JSON(200, keys)
}

// @Summary Find keys
// @Description Get list of keys matching pattern
// @Tags keys
// @Produce json
// @Param p query string true "Pattern"
// @Success 200 {object} domain.Keys
// @Failure 400,500 {object} apiErrors.apiError
// @Router /keys/find [get]
func (s *Server) FindKeys(c *gin.Context) {
	if len(c.Query("p")) == 0 {
		s.badRequestError("pattern parameter required", c)
		return
	}

	items, err := s.queryService.Find(c.Query("p"))
	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	c.JSON(200, items)
}

// @Summary Delete key
// @Description Delete key by ID
// @Tags keys
// @Produce json
// @Param id path string true "Key ID"
// @Success 200 {object} actionResponse
// @Failure 404 {object} apiErrors.apiError
// @Failure 500 {object} apiErrors.apiError
// @Router /keys/{id} [delete]
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

	ok := actionResponse{
		Result:   "OK",
		Affected: res,
	}

	c.JSON(200, ok)
}

// @Summary Delete multiple keys
// @Description Delete multiple keys by IDs
// @Tags keys
// @Accept json
// @Produce json
// @Param payload body []string true "Keys IDs"
// @Success 200 {object} actionResponse
// @Failure 404 {object} apiErrors.apiError
// @Failure 500 {object} apiErrors.apiError
// @Router /keys/delete [post]
func (s *Server) DelKeys(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	var ids []string
	json.Unmarshal(jsonData, &ids)

	var i int64 = 0
	for _, id := range ids {
		res, err := s.queryService.Del(id)
		if err != nil || res == 0 {
			continue
		}
		i++
	}

	ok := actionResponse{
		Result:   "OK",
		Affected: i,
	}

	c.JSON(200, ok)
}

// @Summary Get list
// @Description Get elements of the list
// @Tags lists
// @Produce json
// @Param id path string true "Key ID"
// @Param start query int false "Offset"
// @Param stop query int false "Limit"
// @Success 200 {object} domain.Item
// @Failure 400,500 {object} apiErrors.apiError
// @Router /list/key/{id} [get]
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

// @Summary List LPUSH
// @Description Insert all the specified values at the head of the list stored at key.
// @Tags lists
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Key ID"
// @Param value formData string true "Value"
// @Success 200 {object} actionResponse
// @Failure 400,500 {object} apiErrors.apiError
// @Router /list/lpush/key/{id} [post]
func (s *Server) ListLPush(c *gin.Context) {
	s.listPush("LPush", c)
}

// @Summary List RPUSH
// @Description Insert all the specified values at the tail of the list stored at key.
// @Tags lists
// @Accept x-www-form-urlencoded
// @Produce json
// @Param id path string true "Key ID"
// @Param value formData string true "Value"
// @Success 200 {object} actionResponse
// @Failure 400,500 {object} apiErrors.apiError
// @Router /list/rpush/key/{id} [post]
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

	ok := actionResponse{
		Result:   "OK",
		Affected: val[0].Interface().(int64),
	}

	c.JSON(200, ok)
}

// @Summary List delete
// @Description Removes the first count occurrences of elements equal to element from the list stored at key.
// @Tags lists
// @Accept json
// @Produce json
// @Param id path string true "Key ID"
// @Param payload body map[string]int true "List elements to delete with corresponding count argument"
// @Success 200 {object} actionResponse
// @Failure 400,500 {object} apiErrors.apiError
// @Router /list/key/{id} [DELETE]
func (s *Server) ListDel(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		s.internalServerError(err.Error(), c)
		return
	}

	var m map[string]int
	json.Unmarshal(jsonData, &m)

	var i int64 = 0
	var id string = c.Param("id")
	for el, count := range m {
		res, err := s.queryService.LRem(id, int64(count), el)
		if err != nil || res == 0 {
			continue
		}
		i++
	}

	ok := actionResponse{
		Result:   "OK",
		Affected: i,
	}

	c.JSON(200, ok)
}
