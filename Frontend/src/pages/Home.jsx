import { useContext, useState } from 'react'
import { AppContext } from '../context/AppContext'
import { MdSearch } from "react-icons/md"

const Home = () => {

  const [searchChat, setSearchChat] = useState('')


  const { user } = useContext(AppContext)


  return (
    <main className='flex items-center justify-between w-full h-screen'>
      <div className='w-[320px] h-full flex flex-col items-center bg-[#F8F9FA] border-e-2 border-[#E5E7EB] '>
        {/* user profile  */}
        <div className='w-full h-20 border-b-2 border-[#E5E7EB] flex items-center justify-start gap-3 px-4'>
          <div className={`w-10 h-10 rounded-full overflow-hidden bg-purple-400 flex justify-center items-center font-bold text-xl  `}>
            {user.name[0]}
          </div>
          <div className='flex flex-col items-start justify-center'>
            <h1 className='font-bold'>{user.name}</h1>
            <p className='text-[10px] text-green-500 '>Online</p>
          </div>

        </div>


        <div className='w-full h-[calc(100%-80px)] flex flex-col justify-between items-center p-5 gap-4 '>
          {/* search input  */}
          <div className='flex w-full items-center bg-white text-[#9CA3AF] h-[30px] gap-2 rounded-md border-[1px] px-2 border-[#9CA3AF] '>
            <label htmlFor="search" className='w-[20px] text-xl h-full flex items-center justify-center' ><MdSearch /></label>
            <input type="search" name="search" id="search" placeholder='Search Chat' value={searchChat} onChange={(e) => setSearchChat(e.target.value)} className='h-full w-[calc(100%-20px)] outline-none text-sm ' />
          </div>

          {/* Chats */}
          <ul className='w-full h-[calc(100%-50px)] flex flex-col gap-2 overflow-auto '>

          </ul>

        </div>

      </div>



    </main>
  )
}

export default Home
