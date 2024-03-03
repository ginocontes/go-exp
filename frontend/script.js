
let userButton = document.getElementById("user-button")
let createUserButton = document.getElementById("create-user")

let url = "http://localhost:8080"

let createList = (elements) => {
    let divUsers = document.getElementById('div-users')
    for(let c of divUsers.children) {
        divUsers.removeChild(c) //TODO: find a better way
    }
    let l = document.createElement('ul')
    for(let el of elements) {
        let i = document.createElement("li")
        i.textContent = el.Name
        l.appendChild(i)
    }

    divUsers.appendChild(l)
}

let createUserReq = async (username, password, email ) => {
    let userBody = {username, password, email}
    console.log(userBody)
    let res = await fetch(url + "/createUser", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        redirect: "follow",
        mode: "no-cors",
        body: JSON.stringify(userBody)
    })
    if( res.status != 200) {
        console.error("Error creating User")
    }
}

let createUser = (e) => {
    createUserReq("ginocontes", "pass", "contes@gmail.com")
}

let getUsers = async (e) => {
    let req = await fetch("http://localhost:8080/getUsers")
    if (req.status == 200) {
        let users = await req.json()
        console.log(users)
        for(let user of users) {
            console.log(user.Name)
        }
        createList(users)
    } else {
        console.error("Get Users failed!")
    }
}

userButton.addEventListener('click',  getUsers)

createUserButton.addEventListener('click', createUser)


