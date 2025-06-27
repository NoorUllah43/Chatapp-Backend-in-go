import { createContext, useState, useRef } from "react";


export const AppContext = createContext()

export const AppProvider = ({ children }) => {
    const [User, setUser] = useState({})
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    const [allUsers, setAllUsers] = useState([])
    const [receiverChat, setReceiverChat] = useState()
    const ws = useRef(null)

    const backendURL = "http://localhost:5000"

    if (isLoggedIn && !ws.current) {
        ws.current = new WebSocket(`ws://localhost:8080/ws?userId=${User.ID}`);
        ws.current.onopen = function () {
            console.log("Connected to WebSocket server");
        };
    }




    const value = {
        User,
        setUser,
        isLoggedIn,
        setIsLoggedIn,
        allUsers,
        setAllUsers,
        receiverChat,
        setReceiverChat,
        backendURL,
        ws,
    }


    return <AppContext.Provider value={value}>
        {children}
    </AppContext.Provider>
}





