 async function fetchData(){
 let response = await  fetch("http://localhost:3000/cnpj")
   let { data } = await response.json()
   return data 
}

function formatCnpj(value){
    return value.replace(/^(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, "$1.$2.$3/$4-$5")
}

function copy(e){
    const input = e.target;
    input.select();
    input.setSelectionRange(0, 99999);
    input.classList.add("copied")
    document.execCommand("copy")

    iziToast.show({
        title: 'Sucesso!',
        message: 'Cpnj copiado',
        position: "topRight",
        titleColor: "#fff",
        messageColor: "#fff",
        backgroundColor: "#28a745",
        timeout: 2000
    });
}

async function makeList(){
    let element = document.querySelector(".list-box");
    let cnpjs = await fetchData();
    const previousList = document.querySelector(".list");

    if(previousList){
        previousList.parentNode.removeChild(previousList);
    }

    const ul = document.createElement("ul");
    ul.classList.add("list");

    for(let i = 0; i < cnpjs.length; i++){
        let li = document.createElement("li");
        li.classList.add("list-item");

        let input = document.createElement("input");
        input.readOnly = true;
        input.value = formatCnpj(cnpjs[i])
        input.classList.add("value")
        input.addEventListener("click", copy)

        li.appendChild(input)
        ul.appendChild(li)
    }
    element.appendChild(ul)
}