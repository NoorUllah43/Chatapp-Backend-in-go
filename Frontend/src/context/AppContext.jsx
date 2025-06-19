import { createContext, useState } from "react";


export const AppContext = createContext()

export const AppProvider = ({ children }) => {
    const [User, setUser] = useState({})
    const [isLoggedIn, setIsLoggedIn] = useState(false)
    const [allUsers, setAllUsers] = useState([])
const [receiverChat, setReceiverChat] = useState()
    const value = {
        User,
        setUser,
        isLoggedIn,
        setIsLoggedIn,
        allUsers,
        setAllUsers,
        receiverChat,
        setReceiverChat,
    }


    return <AppContext.Provider value={value}>
        {children}
    </AppContext.Provider>
}





