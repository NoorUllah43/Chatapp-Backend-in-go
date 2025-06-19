import { createBrowserRouter, RouterProvider, Navigate } from 'react-router-dom'
import Login from './pages/Login'
import Home from './pages/Home'
import { useContext } from 'react'
import { AppContext } from './context/AppContext'



// protect routes 
const ProtectRoute = ({ children }) => {
  const { isLoggedIn } = useContext(AppContext);

  return isLoggedIn ? children : <Navigate to="/" />;
};



function App() {

  const router = createBrowserRouter([
    {
      path: '/',
      element: <Login />
    },
    {
      path: '/home',
      element: <ProtectRoute>
        <Home />
      </ProtectRoute>
    }
  ])
  return (
    <>
      <RouterProvider router={router} />
    </>
  )
}

export default App
