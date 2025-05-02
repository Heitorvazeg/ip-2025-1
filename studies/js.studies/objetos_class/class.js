export class animal {
    constructor(especie, vivo, nome) {
        this.especie = especie;
        this.vivo = vivo;
        this.nome = nome;
    }
    set especie(novaespecie) {
        if (typeof novaespecie == "string") {
            this.__especie = novaespecie;
        }
        else if (typeof novaespecie != "string") {
            console.error("Deve ser string");
        }
    }
    get especie() {
        return this.__especie;
    }
    static sobreviver() {
        console.log("sobrevivendo")
    }

    reproduzir(){
        console.log("se reproduziu")
    }
}