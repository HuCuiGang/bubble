package controller
/*
	控制层

	url --> controller --> logic --> model
	请求 --》控制器--》业务逻辑--》模型层的增删改查
 */
import (
	"github.com/HuCuiGang/bubble/models"
	"github.com/gin-gonic/gin"
	"net/http"
)


func CreateATodo(c *gin.Context) {
	// 前段页面填写代办事项 ，点击提交
	// 1.从请求中把数据拿出来存入数据库
	var todo models.Todo
	c.BindJSON(&todo)
	// 2.存入数据库
	// 3.返回响应
	err := models.CreateATodo(&todo)
	if err != nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}
}

func GetTodoList(c *gin.Context) {
	if todoList, err := models.GetAllTodoList();err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"无效的ID"})
		return
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
		})
		return
	}
	c.BindJSON(&todo)
	if err = models.UpdateATodo(todo);err != nil{
		c.JSON(http.StatusOK,gin.H{"error":err.Error()})
	}else {
		c.JSON(http.StatusOK,todo)
	}

}

func DeleteATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK,gin.H{"error":"无效的ID"})
		return
	}
	if err := models.DeleteATodo(id);err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}else {
		c.JSON(http.StatusOK,gin.H{id:"deleted"})
	}
}