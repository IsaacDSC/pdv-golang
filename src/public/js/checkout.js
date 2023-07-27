class Checkout {
    execute() {
        window.location.href = '/checkout';
    }

    sendCheckout() {
        const dataForm = this.getDataForm();
        this.sendOrder(dataForm);
    }

    getDataForm() {
        const formData = new FormData(document.querySelector('#checkout-form'))
        const deliverId = formData.get("delivery")
        const address = formData.get("address")
        const homeNumber = formData.get("homeNumber")
        const observation = formData.get("observation")
        return {
            deliverId,
            address,
            homeNumber,
            observation,
        }
    }

    sendOrder(dataForm) {
        console.log(dataForm)
        var myInit = {
            method: "POST",
            body: JSON.stringify(dataForm),
            mode: "cors",
            cache: "default"
        }
        fetch("/checkout", myInit).then((res) => {
            console.log(res)
        }).catch(err => console.log({ err }))
    }
}