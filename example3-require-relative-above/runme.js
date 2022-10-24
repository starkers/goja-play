
// NOTE require path must be always relative to this file
// NOTE "./libs" is made accessible in this case because WithGlobalFolders() supports ".."
const m = require("../libs/banana.js")

m.banana()

const stuff = {
    foo: "bar",
    fruit: m.banana(),
}



const result = JSON.stringify(stuff)
console.log('(js) getData should return: ', result)

function getData(){
    return result
}
