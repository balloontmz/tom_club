export function getScrollTop(){
    var scrollTop=0;
    if(document.documentElement&&document.documentElement.scrollTop){
        scrollTop=document.documentElement.scrollTop;
    }else if(document.body){
        scrollTop=document.body.scrollTop;
    }
    return scrollTop;
}

// 取文档内容实际高度
export function getScrollHeight(){
    return Math.max(document.body.scrollHeight,document.documentElement.scrollHeight);
}

// 取窗口可视范围的高度 
export function getWindowHeight(){
    var clientHeight=0;
    if(document.body.clientHeight&&document.documentElement.clientHeight){
        console.log('获取到的窗口高度')
        var clientHeight = (document.body.clientHeight<document.documentElement.clientHeight)?document.body.clientHeight:document.documentElement.clientHeight;        
    }else{
        console.log('获取到的窗口高度 else')
        var clientHeight = (document.body.clientHeight>document.documentElement.clientHeight)?document.body.clientHeight:document.documentElement.clientHeight;    
    }
    return clientHeight;
}

//给Window设置滚动事件
// window.onscroll = function(){
// 　　if(getScrollTop() + getWindowHeight() == getScrollHeight()){
// 　　　　alert("you are in the bottom!");
// 　　}
// };