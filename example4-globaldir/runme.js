
// NOTE require path must be always relative to this file
// NOTE paths not required for "mango" because WithGlobalFolders() has "libs" folder added
const m = require("../libs/apple")

m.apple()

const stuff = {
    foo: "baz",
    fruit: m.apple(),
}



const result = JSON.stringify(stuff)
console.log('(js) getData should return: ', result)

function getData(){
    return result
}
