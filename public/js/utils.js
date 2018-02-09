function $(select) {
    return document.querySelector(select);
}

function requestAjax(method, url, obj) {
    return new Promise(function(resolver, rechazar){
        let xhr = new XMLHttpRequest();

        xhr.open(method, url, true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.addEventListener("load", e => {
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            };
            resolver(respuesta)
        })
        xhr.addEventListener("error", e => {
            let self = e.target;
    
            rechazar(selft);
        });
    
        xhr.send(obj);
    });

}