const cart = [
    {
        "index": 1,
        "productId": "f751750e-fde6-4bd1-aae0-53203e6a08ac",
        "productName": "Delicious Pizza",
        "price": "54.99",
        "image": "images/f3.png",
        "quantity": 12
    },
    {
        "index": 2,
        "productId": "4b206768-357f-43cf-934a-1ad5fd8fadc5",
        "productName": "French Fries",
        "price": "13.5",
        "image": "images/f5.png",
        "quantity": 8
    }
]

const index = cart.findIndex(e => e.productId == "4b206768-357f-43cf-934a-1ad5fd8fadc5111")

console.log({index})