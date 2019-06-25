{/* <script>
export default {
    data:function () {
        return{
            // 警告图片
            alarm:9900,
            // 弹窗图层
            pop:8800,
            othen:7700,
        }
    }
}
</script> */}


var code = []
var logo = {}
var userinfo = {
}

var color = [
    "blueviolet",
    "aqua",
    "brown",
    "blueviolet",
    "cyan",
    "blue",
    "chartreuse",
    "yellow",
    "violet"
]
export default {
    logo,
    code,
    getcode,
    userinfo,
    color
}

function getcode(id) {
    if (id > this.code.length) {
        return "失败!"
    }
    return this.code[id - 1].error
}