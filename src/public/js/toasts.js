class Toast {
    execute(productName, image) {
        this.injectToast(productName, image)
        this.showToast()
    }

    showToast() {
        const toastLiveExample = document.getElementById('liveToast')
        const toastBootstrap = bootstrap.Toast.getOrCreateInstance(toastLiveExample)
        toastBootstrap.show()
    }

    injectToast(productName, image) {
        const date = new Date()
        const injectToast = document.querySelector('#injected-toast')
        injectToast.innerHTML = `
        <div class="toast-header">
            <img src="${image}" class="rounded me-2" width="40" alt="...">
            <strong class="me-auto">${productName}</strong>
            <small>${date.toISOString().substring(0,10)}</small>
            <button type="button" class="btn-close" data-bs-dismiss="toast" aria-label="Close"></button>
        </div>
        <div class="toast-body">
            Produto ${productName} adicionado no carrinho de compra!
        </div>
        `
    }
}



