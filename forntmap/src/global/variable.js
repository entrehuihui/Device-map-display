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
var state  = []

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
    color,
    state,
    sleep
}

function getcode(id) {
    if (id == undefined) {
        return "失败!"
    }
    if (id > this.code.length) {
        return "失败!"
    }
    return this.code[id - 1].error
}

function sleep(ms){
  return new Promise((resolve)=>setTimeout(resolve,ms));
}