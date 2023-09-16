function addToCart(productId, productName, price, image) {
    new Toast().execute(productName, image)
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

function removeToCart(productId) {
    const json_cart = localStorage.getItem("cart");
    const cart = JSON.parse(json_cart) || [];
    const removeIfExist = cart?.filter(e => e?.productId != productId) || []
    localStorage.setItem("cart", JSON.stringify(removeIfExist))
    setCart(removeIfExist)
    setTotalPriceCart(calculateTotalCart(removeIfExist))
}

function incrementToCart(productId) {
    const json_cart = localStorage.getItem("cart");
    const cart = JSON.parse(json_cart) || [];
    const index = cart.findIndex(e => e.productId == productId)
    if (index >= 0) {
        cart[index].quantity += 1
        localStorage.setItem("cart", JSON.stringify(cart))
        setCart(cart)
        setTotalPriceCart(calculateTotalCart(cart))
        return true
    }
    return false
}

function decrementToCart(productId) {
    const json_cart = localStorage.getItem("cart");
    const cart = JSON.parse(json_cart) || [];
    const index = cart.findIndex(e => e.productId == productId)
    if (index < 0) return false
    if(cart[index].quantity > 1){
        cart[index].quantity -= 1
        localStorage.setItem("cart", JSON.stringify(cart))
        setCart(cart)
        setTotalPriceCart(calculateTotalCart(cart))
        return true
    }
    removeToCart(productId)
    return true
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
        <div class="card mb-3" style="max-width: 540px;">
            <div class="row g-0">
                <div class="col-md-4">
                    <img src="${cart.image}" class="img-fluid rounded-start" width="146px" alt="${cart.image}">
                </div>
                <div class="col-md-8">
                    <div class="card-body">
                        <h6 class="card-title"><strong>${cart.productName}</strong></h6>
                        <hr>
                        <p class="card-title">Preço Uni: R$ ${cart.price} - ${cart.quantity}x</p>
                        <div class="row">
                            <div class="btn-group" role="group" aria-label="Basic example">
                              <button type="button" class="btn btn-dark" onclick="incrementToCart('${cart.productId}')">+</button>
                              <button type="button" class="btn btn-dark" onclick="decrementToCart('${cart.productId}')">-</button>
                              <button type="button" class="btn btn-dark" onclick="removeToCart('${cart.productId}')">x</button>
                            </div>
                        </div>
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
    <h5 class="text-center"><strong>Total: </strong> R$ ${total}</h5>
    `;
}

function getAndSetCart() {
    const carts = getCart();
    setCart(carts);
    const total = calculateTotalCart(carts);
    setTotalPriceCart(total)

}
