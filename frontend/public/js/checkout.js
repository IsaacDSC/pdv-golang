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
        const CEP = formData.get("CEP")
        const city = formData.get("city")
        const neighborhood = formData.get("neighborhood")
        const address = formData.get("address")
        const homeNumber = formData.get("homeNumber")
        const observation = formData.get("observation")
        return {
            CEP,
            city,
            neighborhood,
            address,
            homeNumber,
            observation,
        }
    }

    sendOrder(dataForm){

    }
}