
import {fecthData} from './common.js'
import { getScrollTop, getScrollHeight, getWindowHeight } from './scrollBottom.js'
import { Init, bind } from './popup.js'

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
        // document.documentElement.appendChild(popup)
        Init()
        var data = await fecthData(add, index)
        gData = data
        // for (let i = 0; i < data.length; i++) {
        //     const goods = data[i];
        //     var e = newComic(goods)
        //     content.appendChild(e)
        // }
        // var col = document.getElementsByClassName('comic')
        // console.log('here', col)
        // for (let i = 0; i < col.length; i++) {
        //     const element = col[i]
        //     console.log('here')
        //     element.addEventListener('click', clickFunc, true)
            
        // }
        gData.forEach(goods => {
            var e = newComic(goods)
            content.appendChild(e)
        })
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

// 同时应该添加点击事件 
function newComic(goods) {
    var para = document.createElement('div')
    var img = document.createElement('img')
    var h1 = document.createElement('h1')
    var text = document.createTextNode(goods.goods_short_title)

    img.setAttribute('src', goods.goods_pic)
    h1.appendChild(text)

    var priceArea = document.createElement('div')
    var priceSpan = document.createElement('span')
    var couponImg = document.createElement('img')
    // var spanText = document.createElement('<span>' + '<em>¥</em>' + goods.buy_price + '</span>')
    couponImg.setAttribute('src', 'data:image/png;base64,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')

    priceSpan.innerHTML = '<em>¥</em>' + goods.buy_price
    priceArea.className = 'price-area'

    priceArea.appendChild(priceSpan)
    priceArea.appendChild(couponImg)

    para.className = 'comic'
    para.appendChild(img)
    para.appendChild(h1)
    para.appendChild(priceArea)
    para.addEventListener('click', clickFunc(goods.ID), false) // 事件监听记得传入一个回调函数，而不是函数的调用。如果传入调用记得返回闭包
    return para
}
// content.appendChild(newComic())

function clickFunc(id){
    var callback = function () {
        console.log(id)
        bind()
    }
    return callback
}