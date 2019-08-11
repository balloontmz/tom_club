export async function fecthData(callback, index) {
    var host = 'https://' + window.location.host
    var url = '/get-goods?page=' + index
    callback()
    var data =  await fetch(url, {
        // body: JSON.stringify(data), // must match 'Content-Type' header
        cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
        credentials: 'same-origin', // include, same-origin, *omit
        headers: {
          'user-agent': 'Mozilla/4.0 MDN Example',
          'content-type': 'application/json'
        },
        method: 'GET', // *GET, POST, PUT, DELETE, etc.
        mode: 'cors', // no-cors, cors, *same-origin
        redirect: 'follow', // manual, *follow, error
        referrer: 'no-referrer', // *client, no-referrer
    })
    data = await data.json()
    if (data.ret == 1) {
        return data.data
    } else {
        return []
    }
    // console.log(await data.json())
    // data.then(function (data) {
    //     console.log(data)
    // })
}