package api

import (
	"errors"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"./service"
	"github.com/gin-gonic/gin"
	"github.com/nfnt/resize"
)

// Upload 图片上传
// @Tags upload
// @Summary 图片上传
// @Description 图片上传
// @Accept  json
// @Produce  json
// @Param	Authorization query	string	flase	"JWT"
// @Success 200 {string} json "{"Error":"Success","Data": object}"
// @Failure  500 {string} json "{"Error":"error","Data": null}"
// @Failure  301 {string} json "{"Error":"Re-login","Data": object}"
// @Router /upload  [post]
func Upload(c *gin.Context) {
	jwt := c.Query("Authorization")
	hmacSample, err := service.ParseJWT(jwt)
	if err != nil {
		retError(c, 19, err)
		return
	}
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		retError(c, 7, err)
		return
	}
	if header.Size > 1048576 {
		retError(c, 29, nil)
		return
	}
	filename := strings.Split(header.Filename, ".")
	if len(filename) != 2 {
		retError(c, 7, nil)
		return
	}
	if filename[1] != "jpg" && filename[1] != "png" {
		retError(c, 30, nil)
		return
	}
	//产生随机的字符串文件名
	newname := "/image/temp/" + strconv.FormatInt(time.Now().Unix(), 10) + getRandomString(4) + strconv.Itoa(hmacSample[0]) + "." + filename[1]
	fi, err := os.Create("./static" + newname)
	if err != nil {
		retError(c, 7, err)
		return
	}
	_, err = io.Copy(fi, file)
	if err != nil {
		retError(c, 7, err)
		return
	}
	//关闭文件
	fi.Close()

	retSuccess(c, newname)
}

func getRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func checkPhoto(dir string) bool {
	_, err := os.Stat("./static" + dir)
	if err != nil {
		log.Println("checkPhoto error! [error]", err)
		return false
	}
	return true
}

func movefile(dir, newdir string) error {
	if !checkPhoto(dir) {
		return errors.New("can not find dirfile")
	}
	body, err := ioutil.ReadFile("./static" + dir)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("./static"+newdir, body, 0644)
	if err != nil {
		return err
	}
	return nil
}

func removefile(dir string) {
	os.Remove("./static" + dir)
}

//PrictureSize 改变传入的图片地址的图片的尺寸
func prictureSize(pricturePath, savepath string, pWidth, pHight uint) error {
	savepath = "./static" + savepath
	pricturePath = "./static" + pricturePath
	//打开图片
	file, err := os.Open(pricturePath)
	if err != nil {
		return err
	}

	//解码图片
	img, err := jpeg.Decode(file)
	if err != nil {
		return err
	}
	file.Close()

	// 改变图片尺寸
	m := resize.Resize(pWidth, pHight, img, resize.Lanczos3)

	//创建新图片
	out, err := os.Create(savepath)
	if err != nil {
		return err
	}
	defer out.Close()
	// write new image to file
	err = jpeg.Encode(out, m, nil)
	if err != nil {
		return err
	}
	return nil
}
