
const m = require("../libs/testmod.js")

m.test()

const stuff = {
    foo: "bar",
}



const result = JSON.stringify(stuff)
console.log('(js) getData should return: ', result)

function getData(){
    return result
}
