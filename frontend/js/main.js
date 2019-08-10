
import {fecthData} from './common.js'
import { getScrollTop, getScrollHeight, getWindowHeight } from './scrollBottom.js'

console.log('hello world')

// 定义全局变量 data为空数组
var gData = []
var index = 1
var content = document.getElementsByClassName('content')[0]

// 定义一个页码的修改器回调函数
function add() {
    index++
}

// 绑定 onready 事件
document.addEventListener('readystatechange', async function (event) {
    // 为什么此处除了 complete 其他的事件都没有进入，是因为 module 导入异步加载导致的吗？
    if (document.readyState === "complete") {
        var data = await fecthData(add, index)
        gData = data
        gData.forEach(goods => {
            content.appendChild(newComic(goods))
        })
        // console.log('是否修改成功', index)
    }
})

// 此处如果不加 true 无法捕获事件！！！
window.addEventListener('scroll', async function(e) {
    // console.log('查看 gdata 数据', gData)
    // console.log('打印参数', '窗口高度为', content.scrollTop, '元素高度为', content.clientHeight, '元素可视化高度', content.scrollHeight, '窗口可视化高度为：', document.documentElement.clientHeight, document.body.clientHeight)
    if (content.scrollHeight - content.scrollTop === content.clientHeight) {
        var data = await fecthData(add, index)
        gData.push(...data)  // 这个全局变量好像没用
        data.forEach(goods => {
            // 此时改变了 content 的数据
            content.appendChild(newComic(goods))
        })
    }
    // if(getScrollTop() + getWindowHeight() == getScrollHeight()){
    //     console.log(getScrollTop(), getWindowHeight(), getScrollHeight())
    //     alert("you are in the bottom!")
    // }
}, true)

// document.addEventListener('readystatechange', function () {
//     if (document.readyState === "DOMContentLoaded") {
//         console.log(2)
//     }
// })

function newComic(goods) {
    var para = document.createElement('div')
    var img = document.createElement('img')
    var h1 = document.createElement('h1')
    var text = document.createTextNode(goods.goods_short_title)

    img.setAttribute('src', goods.goods_pic)
    // h1.appendChild(text)

    para.className = 'comic'
    para.appendChild(img)
    para.appendChild(h1)
    return para
}
// content.appendChild(newComic())