class Checkout {
    execute() {
        window.location.href = '/checkout';
        // const json_cart = localStorage.getItem("cart");
        // const cart = JSON.parse(json_cart) || [];
        // console.log({ cart })
    }

    goToCheckout() {
        const dataForm = this.getDataForm();
        this.createOrder(dataForm);
    }

    getDataForm() {
        const formData = new FormData(document.querySelector('#checkout-form'))
        const clientName = formData.get("clientName")
        const deliverId = formData.get("deliveryId")
        const address = formData.get("address")
        const homeNumber = formData.get("homeNumber")
        const observation = formData.get("observation")
        const billingType = formData.get("billingType")
        return {
            clientName,
            deliverId,
            address,
            homeNumber,
            observation,
            billingType,
        }
    }

    createOrder(dataForm) {
        const json_cart = localStorage.getItem("cart");
        const cart = JSON.parse(json_cart) || [];
        const cartFormatter = cart.map((c) => {
            return {
                price: Number(c.price),
                productId: c.productId,
                qtd: c.quantity
            }
        })
        var myHeaders = new Headers();
        myHeaders.append("Content-Type", "application/json")
        var myInit = {
            method: "POST",
            body: JSON.stringify({
                cart: cartFormatter,
                delivery: {
                    id: dataForm.deliverId,
                    address: dataForm.address,
                    homeNumber: dataForm.homeNumber,
                    observation: dataForm.observation,
                },
                client: {
                    clientName: dataForm.clientName,
                },
                billingType: dataForm.billingType,
            }),
            headers: myHeaders,
            mode: "cors",
            cache: "default"
        }
        fetch("/api/v1/order", myInit).then((res) => {
            if (res.status != 201) {
                console.log({error: res.status})
            } else {
                res.json().then((body) => {
                    console.log({ body })
                })
            }
        }).catch(err => console.log({ err }))
    }
}