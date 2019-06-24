package api

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Resqonse return request
type Resqonse struct {
	Result interface{}
}

func verifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func verifyMobile(mobile string) bool {
	reg := `^1([38][0-9]|14[57]|5[^4])\d{8}$`
	rgx := regexp.MustCompile(reg)
	return rgx.MatchString(mobile)
}

func retError(c *gin.Context, errs int, e interface{}) {
	if e != nil {
		log.Printf("[error] web fail  [Msg]: %v", e)
	} else {
		log.Printf("[error] web fail  [Msg]: %v", errs)
	}
	c.JSON(http.StatusInternalServerError, Resqonse{
		Result: errs,
	})
}

func retSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Resqonse{
		Result: data,
	})
}

func getOffset(c *gin.Context) (int, error) {
	offsets := c.Query("offset")
	if offsets == "" {
		return 0, nil
	}
	offset, err := strconv.Atoi(offsets)
	return offset, err
}

func getLimit(c *gin.Context) (int, error) {
	limits := c.Query("limit")
	if limits == "" {
		return 30, nil
	}
	limit, err := strconv.Atoi(limits)
	return limit, err
}

func getStatus(c *gin.Context) (int, error) {
	statuss := c.Query("status")
	if statuss == "" {
		return 0, nil
	}
	status, err := strconv.Atoi(statuss)
	return status, err
}

func getStarttime(c *gin.Context) (int, error) {
	starttimes := c.Query("starttime")
	if starttimes == "" {
		return 0, nil
	}
	starttime, err := strconv.Atoi(starttimes)
	return starttime, err
}

func getEndtime(c *gin.Context) (int, error) {
	endtimes := c.Query("endtime")
	if endtimes == "" {
		return 0, nil
	}
	endtime, err := strconv.Atoi(endtimes)
	return endtime, err
}

func getName(c *gin.Context) string {
	name := c.Query("name")
	return name
}

func getClassid(c *gin.Context) (int, error) {
	classids := c.Query("classid")
	if classids == "" {
		return 0, nil
	}
	classid, err := strconv.Atoi(classids)
	return classid, err
}

func getOwnid(c *gin.Context) (int, error) {
	ownids := c.Query("ownid")
	if ownids == "" {
		return 0, nil
	}
	ownid, err := strconv.Atoi(ownids)
	return ownid, err
}

func getUID(c *gin.Context) (int, error) {
	uids := c.Query("uid")
	if uids == "" {
		return 0, nil
	}
	uid, err := strconv.Atoi(uids)
	return uid, err
}

func getID(c *gin.Context) (int, error) {
	ids := c.Query("id")
	if ids == "" {
		return 0, nil
	}
	id, err := strconv.Atoi(ids)
	return id, err
}
