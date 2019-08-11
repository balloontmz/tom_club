
var close = document.getElementById('close')
var bg = document.getElementById('pop-bg')
var popup = document.getElementById('popup')

export function Init() {    
    close.onclick = function (e) {
        e.preventDefault()
        popup.classList.toggle('show')
    }
    bg.onclick = function (e) {
        e.preventDefault()
        popup.classList.toggle('show')
    }
}

export function bind() {
    console.log('在此处')
    console.log(popup)
    popup.classList.toggle('show')
}