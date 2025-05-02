import { animal } from "./class"; /* herança */

class cachorro extends animal {
    constructor (especie, vivo, nome, cor) {
        super(especie, vivo, nome);
        this.cor = cor;
    }
    latir() {
        console.log("Latindo");
    }
}

let cachorrao = new cachorro(24 ,true, "Bonitão", "amarelo");

console.log(cachorrao.cor);

let animais = [];

animais.push(cachorro);
console.log(animais);