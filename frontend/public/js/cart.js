function addToCart(productId, productName, price, image) {
    new Toast().execute(productName)
    const json_cart = localStorage.getItem("cart");
    const cart = JSON.parse(json_cart) || [];
    const alreadyExist = cart?.filter(e => e?.productId == productId) || []
    if (alreadyExist.length == 1) {
        cart[alreadyExist[0].index].quantity += 1
        localStorage.setItem("cart", JSON.stringify(cart))
        return
    }
    const index = (cart?.length ?? 0) != 0 ? cart.length : 0
    localStorage.setItem("cart", JSON.stringify([...cart, { index, productId, productName, price, image, quantity: 1 }]));
    return
}

function calculateTotalCart(carts) {
    const setCarts = carts || []
    if (setCarts.length == 0) return 0;
    let total = 0;
    setCarts.forEach((prod) => {
        total += prod?.price * prod?.quantity;
    });
    return total?.toFixed(2);
}

function getCart() {
    const json_cart = localStorage.getItem("cart");
    const cart = JSON.parse(json_cart) || [];
    return cart;
}

function setCart(carts) {
    const card_cart = document.querySelector('#cart-basket')
    card_cart.innerHTML = null;
    carts.forEach((cart) => {
        card_cart.innerHTML += `
        <div class="card mb-3" style="max-width: 540px; max-height: 200px;">
            <div class="row g-0">
                <div class="col-md-4">
                    <img src="${cart.image}" class="img-fluid rounded-start" width="146px" alt="${cart.image}">
                </div>
                <div class="col-md-8">
                    <div class="card-body">
                        <h6 class="card-title"><strong>${cart.productName}</strong></h6>
                        <hr>
                        <p class="card-title">Preço Uni: R$ ${cart.price} - ${cart.quantity}x</p>
                    </div>
                </div>
            </div>
        </div>
        `
    })
}

function setTotalPriceCart(total) {
    const card_total_cart = document.querySelector('#total_cart');
    card_total_cart.innerHTML = `
    <h5 class="text-right">Total: R$ ${total}</h5>
    `;
}

function getAndSetCart() {
    const carts = getCart();
    setCart(carts);
    const total = calculateTotalCart(carts);
    setTotalPriceCart(total)

}
