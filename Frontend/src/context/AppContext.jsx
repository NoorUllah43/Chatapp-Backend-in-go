import { createContext, useState } from "react";


export const AppContext = createContext()

export const AppProvider = ({ children }) => {
    const [user, setUser] = useState({})
    const [isLoggedIn, setIsLoggedIn] = useState(false)

    const value = {
        user,
        setUser,
        isLoggedIn,
        setIsLoggedIn
    }


    return <AppContext.Provider value={value}>
        {children}
    </AppContext.Provider>
}





