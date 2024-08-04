const changeMessage = () => {
    const messageElement = document.getElementById("message")
    const currenMessage = messageElement.textContent

    const alternateMessage = [
        "Hola de nuevo",
        "Bienvenido de nuevo",
        "Que tengas un buen dia"
    ]

    let newMessage

    do {
        newMessage = alternateMessage[Math.floor(Math.random() * alternateMessage.length)]
    } while (newMessage === currenMessage)

    messageElement.textContent = newMessage

}