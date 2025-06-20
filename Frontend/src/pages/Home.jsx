import { useContext, useState, useEffect } from 'react'
import { AppContext } from '../context/AppContext'
import { MdSearch } from "react-icons/md"
import axios from 'axios'
import { IoImageOutline, IoSendOutline, IoChatbubbleEllipsesOutline } from "react-icons/io5"
import { motion } from 'motion/react'
import SendMessage from '../components/SendMessage'
import ReceiveMessage from '../components/ReceiveMessage'
import { IoIosLogOut } from "react-icons/io";
import { toast } from 'react-toastify'
import { useNavigate } from 'react-router-dom'
import { MdKeyboardArrowRight, MdKeyboardArrowLeft } from "react-icons/md";




const Home = () => {

  const [searchChat, setSearchChat] = useState('')
  const [Message, setMessage] = useState('')
  const [showSidebar, setShowSidebar] = useState(false)
  var ws
  const navigate = useNavigate()


  const { User, setAllUsers, allUsers, receiverChat, setReceiverChat, setIsLoggedIn } = useContext(AppContext)

  useEffect(() => {
    const getAllUsers = async () => {
      try {
        const { data } = await axios.get('http://localhost:5000/user/all', { withCredentials: true })
        if (data.success) {
          setAllUsers(data.users)
        } else {
          console.error('Failed to fetch users:', data.message)
        }
      } catch (error) {
        console.error("Error fetching users:", error);

      }
    }
    getAllUsers()

  }, [])





  const handleChatClick = (user) => {
    setReceiverChat(user)
  }

  const logout = async () => {
    try {

      const { data } = await axios.get('http://localhost:5000/auth/logout', { withCredentials: true })
      if (data.Success) {
        toast.success(data.message)
        navigate('/')
        setIsLoggedIn(false)
      }

    } catch (error) {
      console.error("Error in logout:", error);
    }
  }



  useEffect(() => {
    ws = new WebSocket("ws://localhost:8080/ws");

    ws.onopen = function () {
      console.log("Connected to WebSocket server");
    };

  }, [])

  const sendMessage = () => {
    
    
  }
  





  return (
    <main className='flex items-center w-full h-screen'>


      {/* sidebar  */}
      <div className={`md:w-[260px] h-full flex flex-col items-center bg-[#F8F9FA] border-e-2 border-[#E5E7EB] ${!showSidebar ? 'w-20' : 'w-[260px]'} transition-all duration-300 ease-in-out`}>
        {/* user profile  */}
        <div className='w-full h-20 border-b-2 border-[#E5E7EB] flex items-center justify-between gap-3 px-4'>

          <div className='flex items-center gap-2'>
            <div className={`w-10 h-10 cursor-pointer rounded-full overflow-hidden bg-purple-400 text-gray-200 flex justify-center items-center font-bold text-xl  `}>
              {User.name[0]}
            </div>
            <div className={`md:flex flex-col items-start justify-center ${!showSidebar ? 'hidden' : 'flex'}`}>
              <h1 className='font-bold'>{User.name}</h1>
              <p className='text-[10px] text-green-500 '>Online</p>
            </div>
          </div>

          <div className={`text-2xl text-gray-500 cursor-pointer ${!showSidebar ? 'hidden' : 'block'}`}>
            <span className='p-2' onClick={logout} >
              <IoIosLogOut />
            </span>
          </div>

        </div>


        <div className='w-full h-[calc(100%-80px)] flex flex-col justify-between items-center p-5 gap-4 '>
          {/* search input  */}
          <div className={`md:flex w-full items-center bg-white text-[#9CA3AF] h-[30px] gap-2 rounded-md border-[1px] px-2 border-[#9CA3AF] ${!showSidebar ? 'hidden' : 'flex'}`}>
            <label htmlFor="search" className='w-[20px] text-xl h-full flex items-center justify-center' ><MdSearch /></label>
            <input type="search" name="search" id="search" placeholder='Search Chat' value={searchChat} onChange={(e) => setSearchChat(e.target.value)} className='h-full w-[calc(100%-20px)] outline-none text-sm ' />
          </div>

          {/* users */}
          <ul className='w-full h-[calc(100%-50px)] flex flex-col gap-2 overflow-auto '>
            {allUsers.filter((user) => (
              user.ID !== User.ID && user.name.toLowerCase().includes(searchChat)
            )).map((user) => (
              <li key={user.ID} onClick={() => handleChatClick(user)} className='flex gap-2 w-full hover:bg-gray-200  rounded-xl cursor-pointer' >
                <div className={`w-8 h-8 rounded-full overflow-hidden bg-purple-400 text-gray-200 flex justify-center items-center font-bold text-xl`}>
                  {user.name[0]}
                </div>
                <div className={`w-[130px] md:flex flex-col items-start justify-center ${!showSidebar ? 'hidden' : 'flex'}`}>
                  <h2 className='font-semibold text-md'>{user.name}</h2>
                  <p className='text-[10px] text-gray-400 '>Thank you so much!</p>
                </div>
              </li>
            ))}
          </ul>

        </div>

      </div>


      {/* message chat  */}
      {receiverChat ? (
        <div className={`h-full md:w-[calc(100%-260px)] flex flex-col items-center justify-center ${!showSidebar ? 'w-full' : 'w-[calc(100%-260px)]'} transition-all duration-300 ease-in-out`}>
          <div className='w-full h-20 flex items-center justify-between px-5 border-b-2 border-[#E5E7EB]'>
            <div className='flex items-center gap-2'>
              <div className={`w-8 h-8 rounded-full overflow-hidden cursor-pointer bg-purple-400 text-gray-200 flex justify-center items-center font-bold text-xl`}>
                {receiverChat.name[0]}
              </div>
              <div className='flex flex-col items-start justify-center'>
                <h1 className='font-bold'>{receiverChat.name}</h1>
                <p className='text-[10px] text-green-500 '>Online</p>
              </div>
            </div>
          </div>

          {/* Messages  */}
          <div id="chatMessages" className={`w-full h-[calc(100%-160px)] flex  `}>
            {/* toggle button for mobile view */}

            <div onClick={() => setShowSidebar(!showSidebar)} className='w-6 h-full rounded-md flex items-center justify-center bg-white border-r-2 border-[#E5E7EB] md:hidden'>
              <button className='w-full h-10 flex items-center justify-center'>{showSidebar ? <MdKeyboardArrowLeft /> : <MdKeyboardArrowRight />}</button>
            </div>

            <div className='md:w-full w-[calc(100%-24px)] h-full flex flex-col overflow-auto md:p-3 '>
              <SendMessage message={"hi"} />
              <ReceiveMessage message={"how are you?"} />
            </div>


          </div>

          {/* message input  */}
          <div className='w-full h-20 flex items-center justify-between px-3 border-t-2 border-[#E5E7EB]'>

            <div className='w-[calc(100%-50px)] bg-gray-200 px-2 rounded-full flex items-center gap-2'>
              <input type="text" placeholder='Type a message...' value={Message} onChange={(e) => setMessage(e.target.value)} className='flex-1 border-none outline-none p-2' />
            </div>
            <div className='flex items-center justify-center w-10 h-10 rounded-full bg-blue-500 text-white cursor-pointer'>
              <button onClick={sendMessage} className='bg-blue-500 flex justify-center items-center w-10 h-10 rounded-full text-white cursor-pointer'><IoSendOutline /></button>
            </div>
          </div>
        </div>
      ) : (
        <div className="h-full w-[calc(100%-320px)] flex flex-col items-center justify-center" >
          <motion.div animate={{
            y: [0, -40, 0],
            rotate: [0, -10, 0],
          }}
            transition={{
              duration: 5,
              repeat: Infinity,
              repeatType: "loop",
              ease: "easeInOut",
            }} className='flex items-center justify-center text-9xl font-bold text-purple-500 '>
            <IoChatbubbleEllipsesOutline />
          </motion.div>
        </div>
      )}




    </main>
  )
}

export default Home
