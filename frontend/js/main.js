
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