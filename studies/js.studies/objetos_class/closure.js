function closure() { /* faz toda a função inner e as variaveis ficarem privada */
    function inner() {
        let contagem = setInterval(() => {
            let data = new Date();
            console.log(data);
        }, 1000);  /* fica dando console repetidamente a cada 1 seg */
    }
    let cont = inner();
} 

let start = closure();
start.inner();
